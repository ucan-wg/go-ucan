package delegation

import (
	"crypto/rand"
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
//go:generate options -type=config -prefix=With -output=delegatiom_options.go -cmp=false -stringer=false -imports=time,github.com/ipld/go-ipld-prime/datamodel

type config struct {
	Meta      map[string]datamodel.Node
	NotBefore *time.Time
}

func New(privKey crypto.PrivKey, aud did.DID, sub *did.DID, cmd *command.Command, pol policy.Policy, nonce []byte, exp *time.Time, opts ...Option) (*Delegation, error) {
	cfg, err := newConfig(opts...)
	if err != nil {
		return nil, err
	}

	issuer, err := did.FromPrivKey(privKey)
	if err != nil {
		return nil, err
	}

	if !aud.Defined() {
		return nil, fmt.Errorf("%w: %s", token.ErrMissingRequiredDID, "aud")
	}
	audience := aud.String()

	var subject *string
	if sub != nil {
		s := sub.String()
		subject = &s
	}

	policy, err := pol.ToIPLD()
	if err != nil {
		return nil, err
	}

	var meta *token.Map__String__Any
	if len(cfg.Meta) > 0 {
		m := token.ToIPLDMapStringAny(cfg.Meta)
		meta = &m
	}

	var notBefore *int
	if cfg.NotBefore != nil {
		n := int(cfg.NotBefore.Unix())
		notBefore = &n
	}

	var expiration *int
	if exp != nil {
		e := int(exp.Unix())
		expiration = &e
	}

	tkn := &token.Token{
		Issuer:     issuer.String(),
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

func Root(privKey crypto.PrivKey, aud did.DID, cmd *command.Command, pol policy.Policy, nonce []byte, exp *time.Time, opts ...Option) (*Delegation, error) {
	sub, err := did.FromPrivKey(privKey)
	if err != nil {
		return nil, err
	}

	return New(privKey, aud, &sub, cmd, pol, nonce, exp, opts...)
}

func (d *Delegation) Audience() did.DID {
	id, _ := did.Parse(*d.envel.TokenPayload().Audience)

	return id
}

func (d *Delegation) Command() *command.Command {
	cmd, _ := command.Parse(d.envel.TokenPayload().Command)

	return cmd
}

func (d *Delegation) IsPowerline() bool {
	return d.envel.TokenPayload().Subject == nil
}

func (d *Delegation) IsRoot() bool {
	return &d.envel.TokenPayload().Issuer == d.envel.TokenPayload().Subject
}

func (d *Delegation) Issuer() did.DID {
	id, _ := did.Parse(d.envel.TokenPayload().Issuer)

	return id
}

func (d *Delegation) Meta() map[string]datamodel.Node {
	return d.envel.TokenPayload().Meta.Values
}

func (d *Delegation) Nonce() []byte {
	return *d.envel.TokenPayload().Nonce
}

func (d *Delegation) Policy() policy.Policy {
	pol, _ := policy.FromIPLD(*d.envel.TokenPayload().Policy)

	return pol
}

func (d *Delegation) Subject() *did.DID {
	if d.envel.TokenPayload().Subject == nil {
		return nil
	}

	id, _ := did.Parse(*d.envel.TokenPayload().Subject)

	return &id
}

func (d *Delegation) Validate() error {
	return errors.Join(
		d.validateDID("iss", &d.envel.TokenPayload().Issuer, false),
		d.validateDID("aud", d.envel.TokenPayload().Audience, false),
		d.validateDID("sub", d.envel.TokenPayload().Subject, true),
		d.validateCommand(),
		d.validatePolicy(),
		d.validateNonce(),
	)
}

func (d *Delegation) validateCommand() error {
	_, err := command.Parse(d.envel.TokenPayload().Command)

	return err
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

func (d *Delegation) validateNonce() error {
	if d.envel.TokenPayload().Nonce == nil || len(*d.envel.TokenPayload().Nonce) < 1 {
		return fmt.Errorf("nonce is required: must not be nil or empty")
	}

	return nil
}

func (d *Delegation) validatePolicy() error {
	if d.envel.TokenPayload().Policy == nil {
		return fmt.Errorf("the \"pol\" field is required")
	}

	_, err := policy.FromIPLD(*d.envel.TokenPayload().Policy)

	return err
}

func Nonce() ([]byte, error) {
	nonce := make([]byte, 32)

	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	return nonce, nil
}
