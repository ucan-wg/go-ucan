package delegation

import (
	"fmt"
	"time"

	"github.com/ipld/go-ipld-prime/datamodel"

	"github.com/ucan-wg/go-ucan/capability/command"
	"github.com/ucan-wg/go-ucan/capability/policy"
	"github.com/ucan-wg/go-ucan/did"
)

type View struct {
	// Issuer DID (sender)
	Issuer did.DID
	// Audience DID (receiver)
	Audience did.DID
	// Principal that the chain is about (the Subject)
	Subject did.DID
	// The Command to eventually invoke
	Command *command.Command
	// The delegation policy
	Policy policy.Policy
	// A unique, random nonce
	Nonce []byte
	// Arbitrary Metadata
	Meta map[string]datamodel.Node
	// "Not before" UTC Unix Timestamp in seconds (valid from), 53-bits integer
	NotBefore *time.Time
	// The timestamp at which the Invocation becomes invalid
	Expiration *time.Time
}

// ViewFromModel build a decoded view of the raw IPLD data.
// This function also serves as validation.
func ViewFromModel(m PayloadModel) (*View, error) {
	var view View
	var err error

	view.Issuer, err = did.Parse(m.Iss)
	if err != nil {
		return nil, fmt.Errorf("parse iss: %w", err)
	}

	view.Audience, err = did.Parse(m.Aud)
	if err != nil {
		return nil, fmt.Errorf("parse audience: %w", err)
	}

	if m.Sub != nil {
		view.Subject, err = did.Parse(*m.Sub)
		if err != nil {
			return nil, fmt.Errorf("parse subject: %w", err)
		}
	} else {
		view.Subject = did.Undef
	}

	view.Command, err = command.Parse(m.Cmd)
	if err != nil {
		return nil, fmt.Errorf("parse command: %w", err)
	}

	view.Policy, err = policy.FromIPLD(m.Pol)
	if err != nil {
		return nil, fmt.Errorf("parse policy: %w", err)
	}

	if len(m.Nonce) == 0 {
		return nil, fmt.Errorf("nonce is required")
	}
	view.Nonce = m.Nonce

	// TODO: copy?
	view.Meta = m.Meta.Values

	if m.Nbf != nil {
		t := time.Unix(*m.Nbf, 0)
		view.NotBefore = &t
	}

	if m.Exp != nil {
		t := time.Unix(*m.Exp, 0)
		view.Expiration = &t
	}

	return &view, nil
}
