package ucan

import (
	"crypto/rand"
	"time"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/ucan-wg/go-ucan/capability/command"
	"github.com/ucan-wg/go-ucan/capability/policy"
	"github.com/ucan-wg/go-ucan/delegation"
	"github.com/ucan-wg/go-ucan/did"
	"github.com/ucan-wg/go-ucan/internal/envelope"
	"github.com/ucan-wg/go-ucan/issue"
)

const (
	DefaultExpiration  = 30 * 24 * time.Hour
	DefaultNonceLength = 32
)

//go:generate -command options go run github.com/launchdarkly/go-options
//go:generate options -type=authorityConfig -option=AuthorityOption -prefix=With -output=authority_options.go -cmp=false -new=false -imports=time

type authorityConfig struct {
	expiration  time.Duration
	nonceLength int
}

type Authority struct {
	*authorityConfig
	privKey crypto.PrivKey
	did     did.DID // TODO
}

func NewAuthority(privKey crypto.PrivKey, opts ...AuthorityOption) (*Authority, error) {
	cfg := &authorityConfig{
		expiration:  DefaultExpiration,
		nonceLength: DefaultNonceLength,
	}

	if err := applyAuthorityConfigOptions(cfg, opts...); err != nil {
		return nil, err
	}

	id, err := did.FromPubKey(privKey.GetPublic())
	if err != nil {
		return nil, err
	}

	return &Authority{
		authorityConfig: cfg,
		privKey:         privKey,
		did:             id,
	}, nil
}

func (a *Authority) DID() did.DID {
	return a.did
}

func (a *Authority) Expiration() time.Duration {
	return a.expiration
}

func (a *Authority) NonceLength() int {
	return a.nonceLength
}

func (a *Authority) Delegate(aud did.DID, prf []delegation.Token, cmd *command.Command, pol *policy.Policy, exp *time.Time, opts ...delegation.Option) (*envelope.Envelope[*delegation.Token], error) {
	nonce, err := a.Nonce()
	if err != nil {
		return nil, err
	}

	tkn, err := delegation.New(a.DID(), aud, prf, cmd, pol, nil, nonce, opts...)
	if err != nil {
		return nil, err
	}

	return envelope.New(a.privKey, tkn)
}

func (a *Authority) Nonce() ([]byte, error) {
	nonce := make([]byte, a.nonceLength)

	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	return nonce, nil
}

// Issue creates a root UCAN token that can later be delegated.
//
// A subject is required when creating a root UCAN delegation as root
// UCAN delegation tokens should never be "Powerlined" per the
// specification. Therefore, the inclusion of the WithSubject or WithPowerline
// options will result in an error.
//
// Issuing a root UCAN delegation token
// should be a relatively rare occurrence, so this method is not
// available via an Authority.
func Issue(privKey crypto.PrivKey, sub did.DID, cmd *command.Command, pol *policy.Policy, exp *time.Time, opts ...issue.Option) (*envelope.Envelope[*delegation.Token], error) { // TODO: cmd as pointer?
	delOpts, err := issue.ToDelegateOptions(sub, opts...)
	if err != nil {
		return nil, err
	}

	authority, err := NewAuthority(privKey)
	if err != nil {
		return nil, err
	}

	return authority.Delegate(authority.DID(), nil, cmd, pol, nil, delOpts...)
}

func Delegate(privKey crypto.PrivKey, aud did.DID, prf []delegation.Token, cmd *command.Command, pol *policy.Policy, exp *time.Time, opts ...delegation.Option) (*envelope.Envelope[*delegation.Token], error) {
	authority, err := NewAuthority(privKey)
	if err != nil {
		return nil, err
	}

	return authority.Delegate(aud, prf, cmd, pol, exp, opts...)
}
