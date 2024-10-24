// Package invocation implements the UCAN [invocation] specification with
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

	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ucan-wg/go-ucan/did"
	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/pkg/meta"
)

// Token is an immutable type that holds the fields of a UCAN invocation.
type Token struct {
	// The DID of the Invoker
	issuer did.DID
	// The DID of Subject being invoked
	subject did.DID
	// The DID of the intended Executor if different from the Subject
	audience did.DID

	// The Command
	command command.Command
	// The Command's Arguments
	arguments map[string]datamodel.Node
	// Delegations that prove the chain of authority
	proof []cid.Cid

	// Arbitrary Metadata
	meta *meta.Meta

	// A unique, random nonce
	nonce []byte
	// The timestamp at which the Invocation becomes invalid
	expiration *time.Time
	// The timestamp at which the Invocation was created
	invokedAt *time.Time

	// An optional CID of the Receipt that enqueued the Task
	cause *cid.Cid
}

// New creates an invocation Token with the provided options.
//
// If no nonce is provided, a random 12-byte nonce is generated. Use the
// WithNonce or WithEmptyNonce options to specify provide your own nonce
// or to leave the nonce empty respectively.
//
// If no invokedAt is provided, the current time is used. Use the
// WithInvokedAt or WithInvokedAtIn options to specify a different time.
//
// With the exception of the WithMeta option, all other will overwrite
// the previous contents of their target field.
func New(iss, sub did.DID, cmd command.Command, prf []cid.Cid, opts ...Option) (*Token, error) {
	nonce := make([]byte, 12)

	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	iat := time.Now()

	tkn := Token{
		issuer:    iss,
		subject:   sub,
		command:   cmd,
		proof:     prf,
		nonce:     nonce,
		invokedAt: &iat,
	}

	for _, opt := range opts {
		if err := opt(&tkn); err != nil {
			return nil, err
		}
	}

	return &tkn, nil
}

// Issuer returns the did.DID representing the Token's issuer.
func (t *Token) Issuer() did.DID {
	return t.issuer
}

// Subject returns the did.DID representing the Token's subject.
func (t *Token) Subject() did.DID {
	return t.subject
}

// Audience returns the did.DID representing the Token's audience.
func (t *Token) Audience() did.DID {
	return t.audience
}

// Command returns the capability's command.Command.
func (t *Token) Command() command.Command {
	return t.command
}

func (t *Token) Arguments() map[string]datamodel.Node {
	return t.arguments
}

func (t *Token) Proof() []cid.Cid {
	return t.proof
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

func (t *Token) InvokedAt() *time.Time {
	return t.invokedAt
}

func (t *Token) Cause() *cid.Cid {
	return t.cause
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
