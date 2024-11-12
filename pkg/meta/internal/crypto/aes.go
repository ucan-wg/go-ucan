package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
)

// KeySize represents valid AES key sizes
type KeySize int

const (
	KeySize128 KeySize = 16 // AES-128
	KeySize192 KeySize = 24 // AES-192
	KeySize256 KeySize = 32 // AES-256 (recommended)
)

// IsValid returns true if the key size is valid for AES
func (ks KeySize) IsValid() bool {
	switch ks {
	case KeySize128, KeySize192, KeySize256:
		return true
	default:
		return false
	}
}

var ErrShortCipherText = errors.New("ciphertext too short")
var ErrNoEncryptionKey = errors.New("encryption key is required")
var ErrInvalidKeySize = errors.New("invalid key size: must be 16, 24, or 32 bytes")
var ErrZeroKey = errors.New("encryption key cannot be all zeros")

// GenerateKey generates a random AES key of default size KeySize256 (32 bytes).
// Returns an error if the specified size is invalid or if key generation fails.
func GenerateKey() ([]byte, error) {
	return GenerateKeyWithSize(KeySize256)
}

// GenerateKeyWithSize generates a random AES key of the specified size.
// Returns an error if the specified size is invalid or if key generation fails.
func GenerateKeyWithSize(size KeySize) ([]byte, error) {
	if !size.IsValid() {
		return nil, ErrInvalidKeySize
	}

	key := make([]byte, size)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return nil, fmt.Errorf("failed to generate AES key: %w", err)
	}

	return key, nil
}

// EncryptWithAESKey encrypts data using AES-GCM with the provided key.
// The key must be 16, 24, or 32 bytes long (for AES-128, AES-192, or AES-256).
// Returns the encrypted data with the nonce prepended, or an error if encryption fails.
func EncryptWithAESKey(data, key []byte) ([]byte, error) {
	if err := validateAESKey(key); err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, data, nil), nil
}

// DecryptStringWithAESKey decrypts data that was encrypted with EncryptWithAESKey.
// The key must match the one used for encryption.
// Expects the input to have a prepended nonce.
// Returns the decrypted data or an error if decryption fails.
func DecryptStringWithAESKey(data, key []byte) ([]byte, error) {
	if err := validateAESKey(key); err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(data) < gcm.NonceSize() {
		return nil, ErrShortCipherText
	}

	nonce, ciphertext := data[:gcm.NonceSize()], data[gcm.NonceSize():]
	decrypted, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return decrypted, nil
}

func validateAESKey(key []byte) error {
	if key == nil {
		return ErrNoEncryptionKey
	}

	if !KeySize(len(key)).IsValid() {
		return ErrInvalidKeySize
	}

	// check if key is all zeros
	for _, b := range key {
		if b != 0 {
			return nil
		}
	}

	return ErrZeroKey
}
