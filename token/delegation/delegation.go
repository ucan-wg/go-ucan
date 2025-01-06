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
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"

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

// New creates a validated delegation Token from the provided parameters and options.
// This is typically used to delegate a given power to another agent.
//
// You can read it as "(issuer) allows (audience) to perform (cmd+pol) on (subject)".
func New(iss did.DID, aud did.DID, cmd command.Command, pol policy.Policy, sub did.DID, opts ...Option) (*Token, error) {
	tkn := &Token{
		issuer:   iss,
		audience: aud,
		subject:  sub,
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

	return tkn, nil
}

// Root creates a validated UCAN delegation Token from the provided parameters and options.
// This is typically used to create and give a power to an agent.
//
// You can read it as "(issuer) allows (audience) to perform (cmd+pol) on itself".
func Root(iss did.DID, aud did.DID, cmd command.Command, pol policy.Policy, opts ...Option) (*Token, error) {
	return New(iss, aud, cmd, pol, iss, opts...)
}

// Powerline creates a validated UCAN delegation Token from the provided parameters and options.
//
// Powerline is a pattern for automatically delegating all future delegations to another agent regardless of Subject.
// This is a very powerful pattern, use it only if you understand it.
// Powerline delegations MUST NOT be used as the root delegation to a resource
//
// A very common use case for Powerline is providing a stable DID across multiple agents (e.g. representing a user with
// multiple devices). This enables the automatic sharing of authority across their devices without needing to share keys
// or set up a threshold scheme. It is also flexible, since a Powerline delegation MAY be revoked.
//
// You can read it as "(issuer) allows (audience) to perform (cmd+pol) on anything".
func Powerline(iss did.DID, aud did.DID, cmd command.Command, pol policy.Policy, opts ...Option) (*Token, error) {
	return New(iss, aud, cmd, pol, did.Undef, opts...)
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

// Covers indicate if this token has the power to allow the given sub-delegation.
// This function only verifies the principals alignment
func (t *Token) Covers(subDelegation *Token) bool {
	// The Subject of each delegation must equal the invocation's Subject (or Audience if defined). - 4f
	if t.Subject() != sub {
		return fmt.Errorf("%w: delegation %s, expected %s, got %s", ErrWrongSub, dlgCid, sub, dlg.Subject())
	}

	// The Issuer of each delegation must be the Audience in the next one. - 4d
	if t.Audience() != subDelegation.Issuer() {
		return fmt.Errorf("%w: delegation %s, expected %s, got %s", ErrBrokenChain, dlgCid, iss, dlg.Audience())
	}

	// The command of each delegation must "allow" the one before it. - 4g
	if !dlg.Command().Covers(cmd) {
		return fmt.Errorf("%w: delegation %s, %s doesn't cover %s", ErrCommandNotCovered, dlgCid, dlg.Command(), cmd)
	}
}

func (t *Token) String() string {
	var res strings.Builder

	var kind string
	switch {
	case t.issuer == t.subject:
		kind = " (root delegation)"
	case t.subject == did.Undef:
		kind = " (powerline delegation)"
	default:
		kind = " (normal delegation)"
	}

	res.WriteString(fmt.Sprintf("Issuer: %s\n", t.Issuer()))
	res.WriteString(fmt.Sprintf("Audience: %s\n", t.Audience()))
	res.WriteString(fmt.Sprintf("Subject: %s%s\n", t.Subject(), kind))
	res.WriteString(fmt.Sprintf("Command: %s\n", t.Command()))
	res.WriteString(fmt.Sprintf("Policy: %s\n", t.Policy()))
	res.WriteString(fmt.Sprintf("Nonce: %s\n", base64.StdEncoding.EncodeToString(t.Nonce())))
	res.WriteString(fmt.Sprintf("Meta: %s\n", t.Meta()))
	res.WriteString(fmt.Sprintf("NotBefore: %v\n", t.NotBefore()))
	res.WriteString(fmt.Sprintf("Expiration: %v", t.Expiration()))

	return res.String()
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
	if tkn.meta == nil {
		tkn.meta = meta.NewMeta()
	}

	tkn.notBefore, err = parse.OptionalTimestamp(m.Nbf)
	if err != nil {
		return nil, fmt.Errorf("parse notBefore: %w", err)
	}

	tkn.expiration, err = parse.OptionalTimestamp(m.Exp)
	if err != nil {
		return nil, fmt.Errorf("parse expiration: %w", err)
	}

	if err := tkn.validate(); err != nil {
		return nil, err
	}

	return &tkn, nil
}
