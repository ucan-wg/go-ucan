// Package attestation implements the UCAN [attestation] specification with
// an immutable Token type as well as methods to convert the Token to and
// from the [envelope]-enclosed, signed and DAG-CBOR-encoded form that
// should most commonly be used for transport and storage.
//
// [envelope]: https://github.com/ucan-wg/spec#envelope
// [attestation]: TBD
package attestation

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/MetaMask/go-did-it"

	"github.com/ucan-wg/go-ucan/pkg/claims"
	"github.com/ucan-wg/go-ucan/pkg/meta"
	"github.com/ucan-wg/go-ucan/token/internal/nonce"
	"github.com/ucan-wg/go-ucan/token/internal/parse"
)

// Token is an immutable type that holds the fields of a UCAN attestation.
type Token struct {
	// The DID of the Invoker
	issuer did.DID
	// TODO: should this exist?
	// audience did.DID

	// Arbitrary Claims
	claims *claims.Claims

	// Arbitrary Metadata
	meta *meta.Meta

	// A unique, random nonce
	nonce []byte
	// The timestamp at which the Invocation becomes invalid
	expiration *time.Time
	// The timestamp at which the Invocation was created
	issuedAt *time.Time
}

// New creates an attestation Token with the provided options.
//
// If no nonce is provided, a random 12-byte nonce is generated. Use the
// WithNonce or WithEmptyNonce options to specify provide your own nonce
// or to leave the nonce empty respectively.
//
// If no IssuedAt is provided, the current time is used. Use the
// IssuedAt or WithIssuedAtIn Options to specify a different time
// or the WithoutIssuedAt Option to clear the Token's IssuedAt field.
//
// With the exception of the WithMeta option, all others will overwrite
// the previous contents of their target field.
//
// You can read it as "(Issuer - I) attest (arbitrary claim)".
func New(iss did.DID, opts ...Option) (*Token, error) {
	iat := time.Now()

	tkn := Token{
		issuer:   iss,
		claims:   claims.NewClaims(),
		meta:     meta.NewMeta(),
		nonce:    nil,
		issuedAt: &iat,
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

// Issuer returns the did.DID representing the Token's issuer.
func (t *Token) Issuer() did.DID {
	return t.issuer
}

// Claims returns the Token's claims.
func (t *Token) Claims() claims.ReadOnly {
	return t.claims.ReadOnly()
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

// IssuedAt returns the time.Time at which the invocation token was
// created.
func (t *Token) IssuedAt() *time.Time {
	return t.issuedAt
}

// IsValidNow verifies that the token can be used at the current time, based on expiration or "not before" fields.
// This does NOT do any other kind of verifications.
func (t *Token) IsValidNow() bool {
	return t.IsValidAt(time.Now())
}

// IsValidAt verifies that the token can be used at the given time, based on expiration or "not before" fields.
// This does NOT do any other kind of verifications.
func (t *Token) IsValidAt(ti time.Time) bool {
	if t.expiration != nil && ti.After(*t.expiration) {
		return false
	}
	return true
}

func (t *Token) String() string {
	var res strings.Builder

	res.WriteString(fmt.Sprintf("Issuer: %s\n", t.Issuer()))
	res.WriteString(fmt.Sprintf("Nonce: %s\n", base64.StdEncoding.EncodeToString(t.Nonce())))
	res.WriteString(fmt.Sprintf("Meta: %s\n", t.Meta()))
	res.WriteString(fmt.Sprintf("Expiration: %v\n", t.Expiration()))
	res.WriteString(fmt.Sprintf("Issued At: %v\n", t.IssuedAt()))

	return res.String()
}

func (t *Token) validate() error {
	var errs error

	requiredDID := func(id did.DID, fieldname string) {
		if id == nil {
			errs = errors.Join(errs, fmt.Errorf(`a valid did is required for %s: %s`, fieldname, id.String()))
		}
	}

	requiredDID(t.issuer, "Issuer")

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

	if tkn.issuer, err = did.Parse(m.Iss); err != nil {
		return nil, fmt.Errorf("parse iss: %w", err)
	}

	tkn.claims = m.Claims
	if tkn.claims == nil {
		tkn.claims = claims.NewClaims()
	}

	tkn.meta = m.Meta
	if tkn.meta == nil {
		tkn.meta = meta.NewMeta()
	}

	if len(m.Nonce) == 0 {
		return nil, fmt.Errorf("nonce is required")
	}
	tkn.nonce = m.Nonce

	tkn.expiration, err = parse.OptionalTimestamp(m.Exp)
	if err != nil {
		return nil, fmt.Errorf("parse expiration: %w", err)
	}

	tkn.issuedAt, err = parse.OptionalTimestamp(m.Iat)
	if err != nil {
		return nil, fmt.Errorf("parse IssuedAt: %w", err)
	}

	if err := tkn.validate(); err != nil {
		return nil, err
	}

	return &tkn, nil
}
