package delegation

import (
	"fmt"
	"time"
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

// WithEncryptedMetaString adds a key/value pair in the "meta" field.
// The string value is encrypted with the given key.
// The ciphertext will be 40 bytes larger than the plaintext due to encryption overhead.
func WithEncryptedMetaString(key, val string, encryptionKey []byte) Option {
	return func(t *Token) error {
		return t.meta.AddEncrypted(key, val, encryptionKey)
	}
}

// WithEncryptedMetaBytes adds a key/value pair in the "meta" field.
// The []byte value is encrypted with the given key.
// The ciphertext will be 40 bytes larger than the plaintext due to encryption overhead.
func WithEncryptedMetaBytes(key string, val, encryptionKey []byte) Option {
	return func(t *Token) error {
		return t.meta.AddEncrypted(key, val, encryptionKey)
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

// WithNonce sets the Token's nonce with the given value.
// If this option is not used, a random 12-byte nonce is generated for this required field.
func WithNonce(nonce []byte) Option {
	return func(t *Token) error {
		t.nonce = nonce
		return nil
	}
}
