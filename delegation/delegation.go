package delegation

import (
	"errors"
	"fmt"
	"time"

	"github.com/libp2p/go-libp2p/core/crypto"

	"github.com/ucan-wg/go-ucan/capability/command"
	"github.com/ucan-wg/go-ucan/capability/policy"
	"github.com/ucan-wg/go-ucan/did"
	"github.com/ucan-wg/go-ucan/pkg/meta"
)

type Token struct {
	// Issuer DID (sender)
	issuer did.DID
	// Audience DID (receiver)
	audience did.DID
	// Principal that the chain is about (the Subject)
	subject did.DID
	// The Command to eventually invoke
	command *command.Command
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
func New(privKey crypto.PrivKey, aud did.DID, cmd *command.Command, pol policy.Policy, nonce []byte, opts ...Option) (*Token, error) {
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
		nonce:    nonce,
	}

	for _, opt := range opts {
		if err := opt(tkn); err != nil {
			return nil, err
		}
	}

	if err := tkn.validate(); err != nil {
		return nil, err
	}

	return tkn, nil
}

func Root(privKey crypto.PrivKey, aud did.DID, cmd *command.Command, pol policy.Policy, nonce []byte, opts ...Option) (*Token, error) {
	sub, err := did.FromPrivKey(privKey)
	if err != nil {
		return nil, err
	}

	opts = append(opts, WithSubject(sub))

	return New(privKey, aud, cmd, pol, nonce, opts...)
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
func (t *Token) Command() *command.Command {
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
func (t *Token) Meta() *meta.Meta {
	return t.meta
}

// NotBefore returns the time at which the Token becomes "active".
func (t *Token) NotBefore() *time.Time {
	return t.notBefore
}

// Expiration returns the time at which the Token expires.
func (t *Token) Expiration() *time.Time {
	return t.expiration
}

func (t *Token) validate() error {
	var errs error

	requiredDID := func(id did.DID, fieldname string) {
		if !id.Key() {
			errs = errors.Join(errs, fmt.Errorf("a \"did:key\" is required for %s: %s", fieldname, id.String()))
		}
	}

	requiredDID(t.issuer, "Issuer")
	requiredDID(t.audience, "Audience")
	if t.subject != did.Undef {
		requiredDID(t.subject, "Subject")
	}

	if _, err := command.Parse(t.command.String()); err != nil {
		errs = errors.Join(errs, err)
	}

	if _, err := t.policy.ToIPLD(); err != nil {
		errs = errors.Join(errs, err)
	}

	return errs
}

type Option func(*Token) error

// WithExpiration set's the Token's optional "expiration" field to the
// value of the provided time.Time.
func WithExpiration(exp time.Time) Option {
	return func(t *Token) error {
		if exp.Before(time.Now()) {
			return fmt.Errorf("a Token's expiration should be set to a time in the future: %s", exp.String())
		}

		t.expiration = &exp

		return nil
	}
}

// WithMeta adds a key/value pair in the "meta" field.
// WithMeta can be used multiple times in the same call.
// Accepted types for the value are: bool, string, int, int32, int64, []byte,
// and ipld.Node.
func WithMeta(key string, val any) Option {
	return func(t *Token) error {
		return t.meta.Add(key, val)
	}
}

// WithNotBefore set's the Token's optional "notBefore" field to the value
// of the provided time.Time.
func WithNotBefore(nbf time.Time) Option {
	return func(t *Token) error {
		if nbf.Before(time.Now()) {
			return fmt.Errorf("a Token's \"not before\" field should be set to a time in the future: %s", nbf.String())
		}

		t.notBefore = &nbf

		return nil
	}
}

// WithSubject sets the Tokens's optional "subject" field to the value of
// provided did.DID.
//
// This Option should only be used with the New constructor - since
// Subject is a required parameter when creating a Token via the  Root
// constructor, any value provided via this Option will be silently
// overwritten.
func WithSubject(sub did.DID) Option {
	return func(t *Token) error {
		t.subject = sub

		return nil
	}
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

	tkn.audience, err = did.Parse(m.Aud)
	if err != nil {
		return nil, fmt.Errorf("parse audience: %w", err)
	}

	if m.Sub != nil {
		tkn.subject, err = did.Parse(*m.Sub)
		if err != nil {
			return nil, fmt.Errorf("parse subject: %w", err)
		}
	} else {
		tkn.subject = did.Undef
	}

	tkn.command, err = command.Parse(m.Cmd)
	if err != nil {
		return nil, fmt.Errorf("parse command: %w", err)
	}

	tkn.policy, err = policy.FromIPLD(m.Pol)
	if err != nil {
		return nil, fmt.Errorf("parse policy: %w", err)
	}

	if len(m.Nonce) == 0 {
		return nil, fmt.Errorf("nonce is required")
	}
	tkn.nonce = m.Nonce

	tkn.meta = &m.Meta

	if m.Nbf != nil {
		t := time.Unix(*m.Nbf, 0)
		tkn.notBefore = &t
	}

	if m.Exp != nil {
		t := time.Unix(*m.Exp, 0)
		tkn.expiration = &t
	}

	if err := tkn.validate(); err != nil {
		return nil, err
	}

	return &tkn, nil
}
