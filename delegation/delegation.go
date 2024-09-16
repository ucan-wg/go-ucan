package delegation

import (
	"errors"
	"fmt"
	"time"

	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/ucan-wg/go-ucan/capability/command"
	"github.com/ucan-wg/go-ucan/capability/policy"
	"github.com/ucan-wg/go-ucan/did"
	"github.com/ucan-wg/go-ucan/internal/envelope"
	"github.com/ucan-wg/go-ucan/internal/token"
)

const (
	Tag = "ucan/dlg@1.0.0-rc.1"
)

type Delegation struct {
	envel *envelope.Envelope
}

//go:generate -command options go run github.com/selesy/go-options
//go:generate options -type=config -prefix=With -output=delegatiom_options.go -cmp=false -stringer=false -imports=time,github.com/ucan-wg/go-ucan/did,github.com/ipld/go-ipld-prime/datamodel

type config struct {
	Expiration   *time.Time
	Meta         map[string]datamodel.Node
	NoExpiration bool
	NotBefore    *time.Time
	// is a did.DID representing the Subject.
	Subject   *did.DID
	Powerline bool
}

// Required fields for delegation

// Requirements for root

func New(privKey crypto.PrivKey, iss did.DID, aud did.DID, cmd *command.Command, pol *policy.Policy, exp *time.Time, nonce []byte, opts ...Option) (*Delegation, error) {
	cfg, err := newConfig(opts...)
	if err != nil {
		return nil, err
	}

	if !iss.Defined() {
		return nil, fmt.Errorf("%w: %s", token.ErrMissingRequiredDID, "iss")
	}

	if !aud.Defined() {
		return nil, fmt.Errorf("%w: %s", token.ErrMissingRequiredDID, "aud")
	}
	audience := aud.String()

	var subject *string
	if cfg.Subject != nil && cfg.Subject.Defined() {
		s := cfg.Subject.String()
		subject = &s
	}

	policy, err := pol.ToIPLD()
	if err != nil {
		return nil, err
	}

	nonce = []uint8(nonce)

	var notBefore *int
	if cfg.NotBefore != nil {
		n := int(cfg.NotBefore.Unix())
		notBefore = &n
	}

	var meta *token.Map__String__Any
	if len(cfg.Meta) > 0 {
		m := token.ToIPLDMapStringAny(cfg.Meta)
		meta = &m
	}

	var expiration *int
	if exp != nil && !cfg.NoExpiration {
		e := int(cfg.NotBefore.Unix())
		expiration = &e
	}

	tkn := &token.Token{
		Issuer:     iss.String(),
		Audience:   &audience,
		Subject:    subject,
		Command:    cmd.String(),
		Policy:     &policy,
		Nonce:      &nonce,
		Meta:       meta,
		NotBefore:  notBefore,
		Expiration: expiration,
	}

	envel, err := envelope.New(privKey, tkn, Tag)
	if err != nil {
		return nil, err
	}

	dlg := &Delegation{envel: envel}

	if err := dlg.Validate(); err != nil {
		return nil, err
	}

	return dlg, nil
}

type validateFunc func() error

func (d *Delegation) Validate() error {
	return errors.Join(
		d.validateDID("iss", &d.envel.TokenPayload().Issuer, false),
		d.validateDID("aud", d.envel.TokenPayload().Audience, false),
		d.validateDID("sub", d.envel.TokenPayload().Subject, true),
	)
}

func (d *Delegation) validateDID(fieldName string, identity *string, nullableOrOptional bool) error {
	if identity == nil && !nullableOrOptional {
		return fmt.Errorf("a required DID is missing: %s", fieldName)
	}

	id, err := did.Parse(*identity)
	if err != nil {

	}

	if !id.Defined() && !id.Key() {
		return fmt.Errorf("a required DID is missing: %s", fieldName)
	}

	return nil
}
