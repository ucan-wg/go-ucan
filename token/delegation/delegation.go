// Package delegation implements the UCAN [delegation] specification with
// an immutable Token type as well as methods to convert the Token to and
// from the [envelope]-enclosed, signed and DAG-CBOR-encoded form that
// should most commonly be used for transport and storage.
//
// [delegation]: https://github.com/ucan-wg/delegation/tree/v1_ipld
// [envelope]: https://github.com/ucan-wg/spec#envelope
package delegation

// TODO: change the "delegation" link above when the specification is merged

import (
	"errors"
	"fmt"
	"time"

	"github.com/libp2p/go-libp2p/core/crypto"

	"github.com/ucan-wg/go-ucan/did"
	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/pkg/meta"
	"github.com/ucan-wg/go-ucan/pkg/policy"
	"github.com/ucan-wg/go-ucan/token/internal/nonce"
	"github.com/ucan-wg/go-ucan/token/internal/parse"
)

// Token is an immutable type that holds the fields of a UCAN delegation.
type Token struct {
	// Issuer DID (sender)
	issuer did.DID
	// Audience DID (receiver)
	audience did.DID
	// Principal that the chain is about (the Subject)
	subject did.DID
	// The Command to eventually invoke
	command command.Command
	// The delegation policy
	policy policy.Policy
	// A unique, random nonce
	nonce []byte
	// Arbitrary Metadata
	meta *meta.Meta
	// "Not before" UTC Unix Timestamp in seconds (valid from), 53-bits integer
	notBefore *time.Time
	// The timestamp at which the Invocation becomes invalid
	expiration *time.Time
}

// New creates a validated Token from the provided parameters and options.
//
// When creating a delegated token, the Issuer's (iss) DID is assembled
// using the public key associated with the private key sent as the first
// parameter.
func New(privKey crypto.PrivKey, aud did.DID, cmd command.Command, pol policy.Policy, opts ...Option) (*Token, error) {
	iss, err := did.FromPrivKey(privKey)
	if err != nil {
		return nil, err
	}

	tkn := &Token{
		issuer:   iss,
		audience: aud,
		subject:  did.Undef,
		command:  cmd,
		policy:   pol,
		meta:     meta.NewMeta(),
		nonce:    nil,
	}

	for _, opt := range opts {
		if err := opt(tkn); err != nil {
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

	return tkn, nil
}

// Root creates a validated UCAN delegation Token from the provided
// parameters and options.
//
// When creating a root token, both the Issuer's (iss) and Subject's
// (sub) DIDs are assembled from the public key associated with the
// private key passed as the first argument.
func Root(privKey crypto.PrivKey, aud did.DID, cmd command.Command, pol policy.Policy, opts ...Option) (*Token, error) {
	sub, err := did.FromPrivKey(privKey)
	if err != nil {
		return nil, err
	}

	opts = append(opts, WithSubject(sub))

	return New(privKey, aud, cmd, pol, opts...)
}

// Issuer returns the did.DID representing the Token's issuer.
func (t *Token) Issuer() did.DID {
	return t.issuer
}

// Audience returns the did.DID representing the Token's audience.
func (t *Token) Audience() did.DID {
	return t.audience
}

// Subject returns the did.DID representing the Token's subject.
//
// This field may be did.Undef for delegations that are [Powerlined] but
// must be equal to the value returned by the Issuer method for root
// tokens.
func (t *Token) Subject() did.DID {
	return t.subject
}

// Command returns the capability's command.Command.
func (t *Token) Command() command.Command {
	return t.command
}

// Policy returns the capability's policy.Policy.
func (t *Token) Policy() policy.Policy {
	return t.policy
}

// Nonce returns the random Nonce encapsulated in this Token.
func (t *Token) Nonce() []byte {
	return t.nonce
}

// Meta returns the Token's metadata.
func (t *Token) Meta() meta.ReadOnly {
	return t.meta.ReadOnly()
}

// NotBefore returns the time at which the Token becomes "active".
func (t *Token) NotBefore() *time.Time {
	return t.notBefore
}

// Expiration returns the time at which the Token expires.
func (t *Token) Expiration() *time.Time {
	return t.expiration
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
	if t.notBefore != nil && ti.Before(*t.notBefore) {
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
	requiredDID(t.audience, "Audience")

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

	tkn.issuer, err = did.Parse(m.Iss)
	if err != nil {
		return nil, fmt.Errorf("parse iss: %w", err)
	}

	if tkn.audience, err = did.Parse(m.Aud); err != nil {
		return nil, fmt.Errorf("parse audience: %w", err)
	}

	if tkn.subject, err = parse.OptionalDID(m.Sub); err != nil {
		return nil, fmt.Errorf("parse subject: %w", err)
	}

	if tkn.command, err = command.Parse(m.Cmd); err != nil {
		return nil, fmt.Errorf("parse command: %w", err)
	}

	if tkn.policy, err = policy.FromIPLD(m.Pol); err != nil {
		return nil, fmt.Errorf("parse policy: %w", err)
	}

	if len(m.Nonce) == 0 {
		return nil, fmt.Errorf("nonce is required")
	}
	tkn.nonce = m.Nonce

	tkn.meta = m.Meta

	tkn.notBefore = parse.OptionalTimestamp(m.Nbf)
	tkn.expiration = parse.OptionalTimestamp(m.Exp)

	if err := tkn.validate(); err != nil {
		return nil, err
	}

	return &tkn, nil
}
