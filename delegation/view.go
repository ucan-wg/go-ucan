package delegation

import (
	"fmt"
	"time"

	"github.com/ipld/go-ipld-prime/datamodel"

	"github.com/ucan-wg/go-ucan/capability/command"
	"github.com/ucan-wg/go-ucan/capability/policy"
	"github.com/ucan-wg/go-ucan/did"
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
	meta map[string]datamodel.Node
	// "Not before" UTC Unix Timestamp in seconds (valid from), 53-bits integer
	notBefore *time.Time
	// The timestamp at which the Invocation becomes invalid
	expiration *time.Time
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
func (t *Token) Meta() map[string]datamodel.Node {
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

// viewFromModel build a decoded view of the raw IPLD data.
// This function also serves as validation.
func viewFromModel(m tokenPayloadModel) (*Token, error) {
	var view Token
	var err error

	view.issuer, err = did.Parse(m.Iss)
	if err != nil {
		return nil, fmt.Errorf("parse iss: %w", err)
	}

	view.audience, err = did.Parse(m.Aud)
	if err != nil {
		return nil, fmt.Errorf("parse audience: %w", err)
	}

	if m.Sub != nil {
		view.subject, err = did.Parse(*m.Sub)
		if err != nil {
			return nil, fmt.Errorf("parse subject: %w", err)
		}
	} else {
		view.subject = did.Undef
	}

	view.command, err = command.Parse(m.Cmd)
	if err != nil {
		return nil, fmt.Errorf("parse command: %w", err)
	}

	view.policy, err = policy.FromIPLD(m.Pol)
	if err != nil {
		return nil, fmt.Errorf("parse policy: %w", err)
	}

	if len(m.Nonce) == 0 {
		return nil, fmt.Errorf("nonce is required")
	}
	view.nonce = m.Nonce

	// TODO: copy?
	view.meta = m.Meta.Values

	if m.Nbf != nil {
		t := time.Unix(*m.Nbf, 0)
		view.notBefore = &t
	}

	if m.Exp != nil {
		t := time.Unix(*m.Exp, 0)
		view.expiration = &t
	}

	return &view, nil
}
