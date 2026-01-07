package attestation

import (
	"time"
)

// Option is a type that allows optional fields to be set during the
// creation of a Token.
type Option func(*Token) error

// WithClaim adds a key/value pair in the "claims" field.
//
// WithClaims can be used multiple times in the same call.
// Accepted types for the value are: bool, string, int, int32, int64, []byte,
// and ipld.Node.
func WithClaim(key string, val any) Option {
	return func(t *Token) error {
		return t.claims.Add(key, val)
	}
}

// WithClaimsMap adds all key/value pairs in the provided map to the
// Token's "claims" field.
//
// WithClaimsMap can be used multiple times in the same call.
// Accepted types for the value are: bool, string, int, int32, int64, []byte,
// and ipld.Node.
func WithClaimMap(m map[string]any) Option {
	return func(t *Token) error {
		for k, v := range m {
			if err := t.claims.Add(k, v); err != nil {
				return err
			}
		}
		return nil
	}
}

// WithEncryptedClaimsString adds a key/value pair in the "claims" field.
// The string value is encrypted with the given aesKey.
func WithEncryptedClaimsString(key, val string, encryptionKey []byte) Option {
	return func(t *Token) error {
		return t.claims.AddEncrypted(key, val, encryptionKey)
	}
}

// WithEncryptedClaimsBytes adds a key/value pair in the "claims" field.
// The []byte value is encrypted with the given aesKey.
func WithEncryptedClaimsBytes(key string, val, encryptionKey []byte) Option {
	return func(t *Token) error {
		return t.claims.AddEncrypted(key, val, encryptionKey)
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

// WithMetaMap adds all key/value pairs in the provided map to the
// Token's "meta" field.
//
// WithMetaMap can be used multiple times in the same call.
// Accepted types for the value are: bool, string, int, int32, int64, []byte,
// and ipld.Node.
func WithMetaMap(m map[string]any) Option {
	return func(t *Token) error {
		for k, v := range m {
			if err := t.meta.Add(k, v); err != nil {
				return err
			}
		}
		return nil
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

// WithIssuedAt sets the Token's IssuedAt field to the provided
// time.Time.
//
// If this Option is not provided, the invocation Token's iat field will
// be set to the value of time.Now().  If you want to create an invocation
// Token without this field being set, use the WithoutIssuedAt Option.
func WithIssuedAt(iat time.Time) Option {
	return func(t *Token) error {
		t.issuedAt = &iat

		return nil
	}
}

// WithIssuedAtIn sets the Token's IssuedAt field to Now() plus the
// given duration.
func WithIssuedAtIn(after time.Duration) Option {
	return WithIssuedAt(time.Now().Add(after))
}

// WithoutIssuedAt clears the Token's IssuedAt field.
func WithoutIssuedAt() Option {
	return func(t *Token) error {
		t.issuedAt = nil

		return nil
	}
}
