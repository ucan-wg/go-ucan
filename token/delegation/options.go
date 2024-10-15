package delegation

import (
	"fmt"
	"time"

	"github.com/ucan-wg/go-ucan/did"
)

// Option is a type that allows optional fields to be set during the
// creation of a Token.
type Option func(*Token) error

// WithExpiration set's the Token's optional "expiration" field to the
// value of the provided time.Time.
func WithExpiration(exp time.Time) Option {
	return func(t *Token) error {
		if exp.Before(time.Now()) {
			return fmt.Errorf("a Token's expiration should be set to a time in the future: %s", exp.String())
		}

		t.expiration = &exp
		return nil
	}
}

// WithExpirationIn set's the Token's optional "expiration" field to Now() plus the given duration.
func WithExpirationIn(exp time.Duration) Option {
	return func(t *Token) error {
		expTime := time.Now().Add(exp)
		t.expiration = &expTime
		return nil
	}
}

// WithMeta adds a key/value pair in the "meta" field.
//
// WithMeta can be used multiple times in the same call.
// Accepted types for the value are: bool, string, int, int32, int64, []byte,
// and ipld.Node.
func WithMeta(key string, val any) Option {
	return func(t *Token) error {
		return t.meta.Add(key, val)
	}
}

// WithNotBefore set's the Token's optional "notBefore" field to the value
// of the provided time.Time.
func WithNotBefore(nbf time.Time) Option {
	return func(t *Token) error {
		if nbf.Before(time.Now()) {
			return fmt.Errorf("a Token's \"not before\" field should be set to a time in the future: %s", nbf.String())
		}

		t.notBefore = &nbf
		return nil
	}
}

// WithNotBeforeIn set's the Token's optional "notBefore" field to the value
// of the provided time.Time.
func WithNotBeforeIn(nbf time.Duration) Option {
	return func(t *Token) error {
		nbfTime := time.Now().Add(nbf)
		t.notBefore = &nbfTime
		return nil
	}
}

// WithSubject sets the Tokens's optional "subject" field to the value of
// provided did.DID.
//
// This Option should only be used with the New constructor - since
// Subject is a required parameter when creating a Token via the Root
// constructor, any value provided via this Option will be silently
// overwritten.
func WithSubject(sub did.DID) Option {
	return func(t *Token) error {
		t.subject = sub
		return nil
	}
}

// WithNonce sets the Token's nonce with the given value.
// If this option is not used, a random 12-byte nonce is generated for this required field.
func WithNonce(nonce []byte) Option {
	return func(t *Token) error {
		t.nonce = nonce
		return nil
	}
}
