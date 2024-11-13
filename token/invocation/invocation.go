// Package invocation implements the UCAN [invocation] specification with
// an immutable Token type as well as methods to convert the Token to and
// from the [envelope]-enclosed, signed and DAG-CBOR-encoded form that
// should most commonly be used for transport and storage.
//
// # Invocation token validation
//
// Per the specification, invocation Tokens must be validated before the
// command is executed.  This validation can happen in multiple stages:
//
//  1. When the invocation is unsealed from its containing envelope:
//     a. The envelope can be decoded.
//     b. The envelope contains a Signature, VarsigHeader and Payload.
//     c. The Payload contains an iss field that contains a valid did:key.
//     d. The a public key can be extracted from the did:key.
//     e. The public key type is supported by go-ucan.
//     f. The Signature can be decoded per the VarsigHeader.
//     g. The SigPayload can be verified using the Signature and public key.
//     h. The field key of the TokenPayload matches the expected tag.
//  2. When the invocation is created or passes step one:
//     a. The Issuer field is not empty.
//     b. The Subject field is not empty
//     c. The Command field is not empty and the Command is not a wildcard.
//     d. The Policy field is present (but may be empty).
//     e. The Arguments field is present (but may be empty).
//  3. When an unsealed invocation passes steps one and two for execution:
//     a. The invocation can not be expired.
//     b. Invoked at should not be in the future.
//  4. When the proof chain is being validated:
//     a. There must be at least one delegation in the proof chain.
//     b. All referenced delegations must be available.
//     c. The first proof must be issued to the Invoker (audience DID).
//     d. The token must not be expired (expiration in the future or absent).
//     e. The token must be active (nbf in the past or absent).
//     f. The Issuer of each delegation must be the Audience in the next
//     one.
//     g. The last token must be a root delegation.
//     h. The Subject of each delegation must equal the invocation's
//     Audience field.
//     i. The command of each delegation must "allow" the one before it.
//  5. If steps 1-4 pass:
//     a. The policy must "match" the arguments.
//     b. The nonce (if present) is not reused.
//
// [envelope]: https://github.com/ucan-wg/spec#envelope
// [invocation]: https://github.com/ucan-wg/invocation
package invocation

import (
	"errors"
	"fmt"
	"time"

	"github.com/ipfs/go-cid"

	"github.com/ucan-wg/go-ucan/did"
	"github.com/ucan-wg/go-ucan/pkg/args"
	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/pkg/meta"
	"github.com/ucan-wg/go-ucan/token/delegation"
	"github.com/ucan-wg/go-ucan/token/internal/nonce"
	"github.com/ucan-wg/go-ucan/token/internal/parse"
)

// Token is an immutable type that holds the fields of a UCAN invocation.
type Token struct {
	// The DID of the Invoker
	issuer did.DID
	// The DID of Subject being invoked
	subject did.DID
	// The DID of the intended Executor if different from the Subject
	audience did.DID

	// The Command
	command command.Command
	// The Command's arguments
	arguments *args.Args
	// CIDs of the delegation.Token that prove the chain of authority
	// They need to form a strictly linear chain, and being ordered starting from the
	// leaf Delegation (with aud matching the invocation's iss), in a strict sequence
	// where the iss of the previous Delegation matches the aud of the next Delegation.
	proof []cid.Cid
	// Arbitrary Metadata
	meta *meta.Meta

	// A unique, random nonce
	nonce []byte
	// The timestamp at which the Invocation becomes invalid
	expiration *time.Time
	// The timestamp at which the Invocation was created
	invokedAt *time.Time

	// An optional CID of the Receipt that enqueued the Task
	cause *cid.Cid
}

// New creates an invocation Token with the provided options.
//
// If no nonce is provided, a random 12-byte nonce is generated. Use the
// WithNonce or WithEmptyNonce options to specify provide your own nonce
// or to leave the nonce empty respectively.
//
// If no invokedAt is provided, the current time is used. Use the
// WithInvokedAt or WithInvokedAtIn Options to specify a different time
// or the WithoutInvokedAt Option to clear the Token's invokedAt field.
//
// With the exception of the WithMeta option, all others will overwrite
// the previous contents of their target field.
func New(iss, sub did.DID, cmd command.Command, prf []cid.Cid, opts ...Option) (*Token, error) {
	iat := time.Now()

	tkn := Token{
		issuer:    iss,
		subject:   sub,
		command:   cmd,
		arguments: args.New(),
		proof:     prf,
		meta:      meta.NewMeta(),
		nonce:     nil,
		invokedAt: &iat,
	}

	for _, opt := range opts {
		if err := opt(&tkn); err != nil {
			return nil, err
		}
	}

	if len(tkn.nonce) == 0 {
		tkn.nonce, err = nonce.Generate()
		if err != nil {
			return nil, err
		}
	}

	if err := tkn.validate(); err != nil {
		return nil, err
	}

	return &tkn, nil
}

type DelegationLoader interface {
	GetDelegation(cid cid.Cid) (*delegation.Token, error)
}

func (t *Token) ExecutionAllowed(loader DelegationLoader) (bool, error) {
	return t.executionAllowed(loader, t.arguments)
}

func (t *Token) ExecutionAllowedWithArgsHook(loader DelegationLoader, hook func(*args.Args) *args.Args) (bool, error) {
	return t.executionAllowed(loader, hook(t.arguments))
}

func (t *Token) executionAllowed(loader DelegationLoader, arguments *args.Args) (bool, error) {
	// There must be at least one delegation referenced - 4a
	if len(t.proof) < 1 {
		return false, ErrNoProof
	}

	type chainer interface {
		Issuer() did.DID
		Subject() did.DID // TODO: if the invocation token's Audience is nil, copy the subject into it
		Command() command.Command
	}

	// This starts as the invocation token but will be the root delegation
	// after the for loop below completes
	var lastChainer chainer = t

	for i, dlgCid := range t.proof {
		// The token must be present - 4b
		dlg, err := loader.GetDelegation(dlgCid)
		if err != nil {
			return false, fmt.Errorf("%w: need %s", ErrMissingDelegation, dlgCid)
		}

		// No tokens in the proof chain may be expired - 4d
		if dlg.Expiration() != nil && dlg.Expiration().Before(time.Now()) {
			return false, fmt.Errorf("%w: CID is %s", ErrDelegationExpired, dlgCid)
		}

		// No tokens in the proof chain may be inactive - 4e
		if dlg.NotBefore() != nil && dlg.NotBefore().After(time.Now()) {
			return false, fmt.Errorf("%w: CID is %s", ErrDelegationInactive, dlgCid)
		}

		// First proof must have the invoker's Issuer as the Audience - 4c
		if i == 0 && dlg.Audience() != t.Issuer() {
			return false, fmt.Errorf("%w: expected %s, got %s", ErrNotIssuedToInvoker, t.issuer, dlg.Audience())
		}

		// Tokens must form a chain with current issuer equal to the
		// next audience - 4f
		if lastChainer.Issuer() != dlg.Audience() {
			return false, fmt.Errorf("%w: expected %s, got %s", ErrBrokenChain, lastChainer.Issuer(), dlg.Audience())
		}

		// TODO: Checking the subject consistency can happen here - 4h
		// TODO: Checking the command equivalence or attenuation can happen here - 4i

		lastChainer = dlg
	}

	// The last prf value must be a root delegation (have the issuer field
	// match the Subject field) - 4g
	if lastChainer.Issuer() != lastChainer.Subject() {
		return false, fmt.Errorf("%w: expected %s, got %s", ErrLastNotRoot, lastChainer.Subject(), lastChainer.Issuer())
	}

	return true, nil
}

// Issuer returns the did.DID representing the Token's issuer.
func (t *Token) Issuer() did.DID {
	return t.issuer
}

// Subject returns the did.DID representing the Token's subject.
func (t *Token) Subject() did.DID {
	return t.subject
}

// Audience returns the did.DID representing the Token's audience.
func (t *Token) Audience() did.DID {
	return t.audience
}

// Command returns the capability's command.Command.
func (t *Token) Command() command.Command {
	return t.command
}

// Arguments returns the arguments to be used when the command is
// invoked.
func (t *Token) Arguments() *args.Args {
	return t.arguments
}

// Proof() returns the ordered list of cid.Cid which reference the
// delegation Tokens that authorize this invocation.
func (t *Token) Proof() []cid.Cid {
	return t.proof
}

// Meta returns the Token's metadata.
func (t *Token) Meta() meta.ReadOnly {
	return t.meta.ReadOnly()
}

// Nonce returns the random Nonce encapsulated in this Token.
func (t *Token) Nonce() []byte {
	return t.nonce
}

// Expiration returns the time at which the Token expires.
func (t *Token) Expiration() *time.Time {
	return t.expiration
}

// InvokedAt returns the time.Time at which the invocation token was
// created.
func (t *Token) InvokedAt() *time.Time {
	return t.invokedAt
}

// Cause returns the Token's (optional) cause field which may specify
// which describes the Receipt that requested the invocation.
func (t *Token) Cause() *cid.Cid {
	return t.cause
}

// IsValidNow verifies that the token can be used at the current time, based on expiration or "not before" fields.
// This does NOT do any other kind of verifications.
func (t *Token) IsValidNow() bool {
	return t.IsValidAt(time.Now())
}

// IsValidNow verifies that the token can be used at the given time, based on expiration or "not before" fields.
// This does NOT do any other kind of verifications.
func (t *Token) IsValidAt(ti time.Time) bool {
	if t.expiration == nil && ti.After(*t.expiration) {
		return false
	}
	return true
}

func (t *Token) validate() error {
	var errs error

	requiredDID := func(id did.DID, fieldname string) {
		if !id.Defined() {
			errs = errors.Join(errs, fmt.Errorf(`a valid did is required for %s: %s`, fieldname, id.String()))
		}
	}

	requiredDID(t.issuer, "Issuer")
	requiredDID(t.subject, "Subject")

	if len(t.nonce) < 12 {
		errs = errors.Join(errs, fmt.Errorf("token nonce too small"))
	}

	return errs
}

// tokenFromModel build a decoded view of the raw IPLD data.
// This function also serves as validation.
func tokenFromModel(m tokenPayloadModel) (*Token, error) {
	var (
		tkn Token
		err error
	)

	if tkn.issuer, err = did.Parse(m.Iss); err != nil {
		return nil, fmt.Errorf("parse iss: %w", err)
	}

	if tkn.subject, err = did.Parse(m.Sub); err != nil {
		return nil, fmt.Errorf("parse subject: %w", err)
	}

	if tkn.audience, err = parse.OptionalDID(m.Aud); err != nil {
		return nil, fmt.Errorf("parse audience: %w", err)
	}

	if tkn.command, err = command.Parse(m.Cmd); err != nil {
		return nil, fmt.Errorf("parse command: %w", err)
	}

	if len(m.Nonce) == 0 {
		return nil, fmt.Errorf("nonce is required")
	}
	tkn.nonce = m.Nonce

	tkn.arguments = m.Args
	tkn.proof = m.Prf
	tkn.meta = m.Meta

	tkn.expiration = parse.OptionalTimestamp(m.Exp)
	tkn.invokedAt = parse.OptionalTimestamp(m.Iat)

	tkn.cause = m.Cause

	if err := tkn.validate(); err != nil {
		return nil, err
	}

	return &tkn, nil
}
