package invocation

import (
	"time"

	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ucan-wg/go-ucan/did"
)

// Option is a type that allows optional fields to be set during the
// creation of a Token.
type Option func(*Token) error

// WithArgument adds a key/value pair to the Token's Arguments field.
func WithArgument(key string, val datamodel.Node) Option {
	return func(t *Token) error {
		t.arguments[key] = val

		return nil
	}
}

// WithArguments sets the Token's Arguments field to the provided map.
//
// Note that this will overwrite any existing Arguments whether provided
// by a previous call to this function or by one or more calls to
// WithArgument.
func WithArguments(args map[string]datamodel.Node) Option {
	return func(t *Token) error {
		t.arguments = args

		return nil
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

// WithNonce sets the Token's nonce with the given value.
//
// If this option is not used, a random 12-byte nonce is generated for
// this require.
func WithNonce(nonce []byte) Option {
	return func(t *Token) error {
		t.nonce = nonce

		return nil
	}
}

// WithEmptyNonce sets the Token's nonce to an empty byte slice as
// suggested by the UCAN spec for invocation tokens that represent
// idem
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
		t.expiration = &exp

		return nil
	}
}

// WithExpirationIn set's the Token's optional "expiration" field to
// Now() plus the given duration.
func WithExpirationIn(exp time.Duration) Option {
	return func(t *Token) error {
		expTime := time.Now().Add(exp)
		t.expiration = &expTime

		return nil
	}
}

// WithInvokedAt sets the Token's invokedAt field to the provided
// time.Time.
func WithInvokedAt(iat time.Time) Option {
	return func(t *Token) error {
		t.invokedAt = &iat

		return nil
	}
}

// WithInvokedAtIn sets the Token's invokedAt field to Now() plus the
// given duration.
func WithInvokedAtIn(after time.Duration) Option {
	return func(t *Token) error {
		iat := time.Now().Add(after)
		t.invokedAt = &iat

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
