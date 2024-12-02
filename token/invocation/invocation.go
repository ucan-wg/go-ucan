// Package invocation implements the UCAN [invocation] specification with
// an immutable Token type as well as methods to convert the Token to and
// from the [envelope]-enclosed, signed and DAG-CBOR-encoded form that
// should most commonly be used for transport and storage.
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

	var err error
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

func (t *Token) ExecutionAllowed(loader delegation.Loader) error {
	return t.executionAllowed(loader, t.arguments)
}

func (t *Token) ExecutionAllowedWithArgsHook(loader delegation.Loader, hook func(args args.ReadOnly) (*args.Args, error)) error {
	newArgs, err := hook(t.arguments.ReadOnly())
	if err != nil {
		return err
	}
	return t.executionAllowed(loader, newArgs)
}

func (t *Token) executionAllowed(loader delegation.Loader, arguments *args.Args) error {
	delegations, err := t.loadProofs(loader)
	if err != nil {
		// All referenced delegations must be available - 4b
		return err
	}

	if err := t.verifyProofs(delegations); err != nil {
		return err
	}

	if err := t.verifyTimeBound(delegations); err != nil {
		return err
	}

	if err := t.verifyArgs(delegations, arguments); err != nil {
		return err
	}

	return nil
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
func (t *Token) Arguments() args.ReadOnly {
	return t.arguments.ReadOnly()
}

// Proof() returns the ordered list of cid.Cid which reference the
// delegation Tokens that authorize this invocation.
// Ordering is from the leaf Delegation (with aud matching the invocation's iss)
// to the root delegation.
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
	if t.expiration != nil && ti.After(*t.expiration) {
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

func (t *Token) loadProofs(loader delegation.Loader) (res []*delegation.Token, err error) {
	res = make([]*delegation.Token, len(t.proof))
	for i, c := range t.proof {
		res[i], err = loader.GetDelegation(c)
		if err != nil {
			return nil, fmt.Errorf("%w: need %s", ErrMissingDelegation, c)
		}
	}
	return res, nil
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
	if err := tkn.arguments.Validate(); err != nil {
		return nil, fmt.Errorf("invalid arguments: %w", err)
	}

	tkn.proof = m.Prf
	tkn.meta = m.Meta

	tkn.expiration, err = parse.OptionalTimestamp(m.Exp)
	if err != nil {
		return nil, fmt.Errorf("parse expiration: %w", err)
	}

	tkn.invokedAt, err = parse.OptionalTimestamp(m.Iat)
	if err != nil {
		return nil, fmt.Errorf("parse invokedAt: %w", err)
	}

	tkn.cause = m.Cause

	if err := tkn.validate(); err != nil {
		return nil, err
	}

	return &tkn, nil
}
