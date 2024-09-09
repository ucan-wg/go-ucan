package delegation

import (
	"time"

	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/node/bindnode"
	"github.com/ipld/go-ipld-prime/schema"
	"github.com/ucan-wg/go-ucan/capability/command"
	"github.com/ucan-wg/go-ucan/capability/policy"
	"github.com/ucan-wg/go-ucan/did"
	"github.com/ucan-wg/go-ucan/internal/envelope"
	"github.com/ucan-wg/go-ucan/internal/token"
)

const (
	Tag = "ucan/dlg@"
)

//go:generate -command options go run github.com/launchdarkly/go-options
//go:generate options -type=config -prefix=With -output=delegatiom_options.go -cmp=false -imports=time,github.com/ucan-wg/go-ucan/did

type config struct {
	Expiration   *time.Time
	Meta         map[string]any
	NoExpiration bool
	NotBefore    *time.Time
	// is a did.DID representing the Subject.
	Subject   *did.DID
	Powerline bool
}

type Meta struct {
	Keys   []string
	Values map[string]datamodel.Node
}

func NewMeta(meta map[string]any) Meta {
	keys := make([]string, len(meta))
	values := make(map[string]datamodel.Node, len(meta))
	i := 0

	for k, v := range meta {
		keys[i] = k
		values[k] = bindnode.Wrap(&v, nil)
	}

	return Meta{
		Keys:   keys,
		Values: values,
	}
}

var _ envelope.Tokener = (*Token)(nil)

type Token struct {
	// Issuer DID (sender)
	Issuer did.DID
	// Audience DID (receiver)
	Audience did.DID
	// Principal that the chain is about (the Subject)
	Subject *did.DID
	// The Command to eventually invoke
	Command *command.Command
	// The delegation policy
	Policy *policy.Policy
	// A unique, random nonce
	Nonce []byte
	// Arbitrary Metadata
	Meta Meta
	// "Not before" UTC Unix Timestamp in seconds (valid from), 53-bits integer
	NotBefore *time.Time
	// The timestamp at which the Invocation becomes invalid
	Expiration *time.Time
}

func New(iss did.DID, aud did.DID, prf []Token, cmd *command.Command, pol *policy.Policy, exp *time.Time, nonce []byte, opts ...Option) (*Token, error) {
	cfg, err := newConfig(opts...)
	if err != nil {
		return nil, err
	}

	tkn := &Token{
		Issuer:    iss,
		Audience:  aud,
		Subject:   cfg.Subject,
		Command:   cmd,
		Policy:    pol,
		Nonce:     nonce,
		NotBefore: cfg.NotBefore,
	}

	if len(cfg.Meta) > 0 {
		tkn.Meta = NewMeta(cfg.Meta)
	}

	if exp != nil && !cfg.NoExpiration {
		tkn.Expiration = exp
	}

	return tkn, nil
}

func (d *Token) Tag() string {
	return Tag
}

func (d *Token) Prototype() schema.TypedPrototype {
	return bindnode.Prototype((*Token)(nil), mustLoadSchema().TypeByName("Delegation"), token.BindnodeOptions()...)
}
