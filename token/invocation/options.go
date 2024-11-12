package invocation

import (
	"time"

	"github.com/ipfs/go-cid"

	"github.com/ucan-wg/go-ucan/did"
)

// Option is a type that allows optional fields to be set during the
// creation of an invocation Token.
type Option func(*Token) error

// WithArgument adds a key/value pair to the Token's Arguments field.
func WithArgument(key string, val any) Option {
	return func(t *Token) error {
		return t.arguments.Add(key, val)
	}
}

// WithAudience sets the Token's audience to the provided did.DID.
//
// If the provided did.DID is the same as the Token's subject, the
// audience is not set.
func WithAudience(aud did.DID) Option {
	return func(t *Token) error {
		if t.subject.String() != aud.String() {
			t.audience = aud
		}

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

// WithEncryptedMetaString adds a key/value pair in the "meta" field.
// The string value is encrypted with the given aesKey.
func WithEncryptedMetaString(key, val string, encryptionKey []byte) Option {
	return func(t *Token) error {
		return t.meta.AddEncrypted(key, val, encryptionKey)
	}
}

// WithEncryptedMetaBytes adds a key/value pair in the "meta" field.
// The []byte value is encrypted with the given aesKey.
func WithEncryptedMetaBytes(key string, val, encryptionKey []byte) Option {
	return func(t *Token) error {
		return t.meta.AddEncrypted(key, val, encryptionKey)
	}
}

// WithNonce sets the Token's nonce with the given value.
//
// If this option is not used, a random 12-byte nonce is generated for
// this required field.  If you truly want to create an invocation Token
// without a nonce, use the WithEmptyNonce Option which will set the
// nonce to an empty byte array.
func WithNonce(nonce []byte) Option {
	return func(t *Token) error {
		t.nonce = nonce

		return nil
	}
}

// WithEmptyNonce sets the Token's nonce to an empty byte slice as
// suggested by the UCAN spec for invocation tokens that represent
// idempotent operations.
func WithEmptyNonce() Option {
	return func(t *Token) error {
		t.nonce = []byte{}

		return nil
	}
}

// WithExpiration set's the Token's optional "expiration" field to the
// value of the provided time.Time.
func WithExpiration(exp time.Time) Option {
	return func(t *Token) error {
		exp = exp.Round(time.Second)
		t.expiration = &exp

		return nil
	}
}

// WithExpirationIn set's the Token's optional "expiration" field to
// Now() plus the given duration.
func WithExpirationIn(after time.Duration) Option {
	return WithExpiration(time.Now().Add(after))
}

// WithInvokedAt sets the Token's invokedAt field to the provided
// time.Time.
//
// If this Option is not provided, the invocation Token's iat field will
// be set to the value of time.Now().  If you want to create an invocation
// Token without this field being set, use the WithoutInvokedAt Option.
func WithInvokedAt(iat time.Time) Option {
	return func(t *Token) error {
		t.invokedAt = &iat

		return nil
	}
}

// WithInvokedAtIn sets the Token's invokedAt field to Now() plus the
// given duration.
func WithInvokedAtIn(after time.Duration) Option {
	return WithInvokedAt(time.Now().Add(after))
}

// WithoutInvokedAt clears the Token's invokedAt field.
func WithoutInvokedAt() Option {
	return func(t *Token) error {
		t.invokedAt = nil

		return nil
	}
}

// WithCause sets the Token's cause field to the provided cid.Cid.
func WithCause(cause *cid.Cid) Option {
	return func(t *Token) error {
		t.cause = cause

		return nil
	}
}
