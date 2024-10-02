// Package delegation implements the UCAN [invocation] specification with
// an immutable Token type as well as methods to convert the Token to and
// from the [envelope]-enclosed, signed and DAG-CBOR-encoded form that
// should most commonly be used for transport and storage.
//
// [envelope]: https://github.com/ucan-wg/spec#envelope
// [invocation]: https://github.com/ucan-wg/invocation
package invocation

import (
	"crypto/rand"
	"errors"
	"fmt"
	"time"

	"github.com/ucan-wg/go-ucan/did"
	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/pkg/meta"
)

// Token is an immutable type that holds the fields of a UCAN invocation.
type Token struct {
	// Issuer DID (invoker)
	issuer did.DID
	// Audience DID (receiver/executor)
	audience did.DID
	// Subject DID (subject being invoked)
	subject did.DID
	// The Command to invoke
	command command.Command
	// TODO: args
	// TODO: prf
	// A unique, random nonce
	nonce []byte
	// Arbitrary Metadata
	meta *meta.Meta
	// The timestamp at which the Invocation becomes invalid
	expiration *time.Time
	// The timestamp at which the Invocation was created
	invokedAt *time.Time
	// TODO: cause
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

// Nonce returns the random Nonce encapsulated in this Token.
func (t *Token) Nonce() []byte {
	return t.nonce
}

// Meta returns the Token's metadata.
func (t *Token) Meta() *meta.Meta {
	return t.meta
}

// Expiration returns the time at which the Token expires.
func (t *Token) Expiration() *time.Time {
	return t.expiration
}

func (t *Token) validate() error {
	var errs error

	requiredDID := func(id did.DID, fieldname string) {
		if !id.Defined() {
			errs = errors.Join(errs, fmt.Errorf(`a valid did is required for %s: %s`, fieldname, id.String()))
		}
	}

	requiredDID(t.issuer, "Issuer")

	// TODO

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
	)

	// TODO

	return &tkn, nil
}

// generateNonce creates a 12-byte random nonce.
// TODO: some crypto scheme require more, is that our case?
func generateNonce() ([]byte, error) {
	res := make([]byte, 12)
	_, err := rand.Read(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
