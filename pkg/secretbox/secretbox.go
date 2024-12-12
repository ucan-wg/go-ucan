package secretbox

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"

	"golang.org/x/crypto/nacl/secretbox"
)

const keySize = 32 // secretbox allows only 32-byte keys

var ErrShortCipherText = errors.New("ciphertext too short")
var ErrNoEncryptionKey = errors.New("encryption key is required")
var ErrInvalidKeySize = errors.New("invalid key size: must be 32 bytes")
var ErrZeroKey = errors.New("encryption key cannot be all zeros")

// GenerateKey generates a random 32-byte key to be used by EncryptWithKey and DecryptWithKey
func GenerateKey() ([]byte, error) {
	key := make([]byte, keySize)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return nil, fmt.Errorf("failed to generate key: %w", err)
	}
	return key, nil
}

// EncryptWithKey encrypts data using NaCl's secretbox with the provided key.
// 40 bytes of overhead (24-byte nonce + 16-byte MAC) are added to the plaintext size.
func EncryptWithKey(data, key []byte) ([]byte, error) {
	if err := validateKey(key); err != nil {
		return nil, err
	}

	var secretKey [keySize]byte
	copy(secretKey[:], key)

	// Generate 24 bytes of random data as nonce
	var nonce [24]byte
	if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
		return nil, err
	}

	// Encrypt and authenticate data
	encrypted := secretbox.Seal(nonce[:], data, &nonce, &secretKey)
	return encrypted, nil
}

// DecryptStringWithKey decrypts data using secretbox with the provided key
func DecryptStringWithKey(data, key []byte) ([]byte, error) {
	if err := validateKey(key); err != nil {
		return nil, err
	}

	if len(data) < 24 {
		return nil, ErrShortCipherText
	}

	var secretKey [keySize]byte
	copy(secretKey[:], key)

	var nonce [24]byte
	copy(nonce[:], data[:24])

	decrypted, ok := secretbox.Open(nil, data[24:], &nonce, &secretKey)
	if !ok {
		return nil, errors.New("decryption failed")
	}

	return decrypted, nil
}

func validateKey(key []byte) error {
	if key == nil {
		return ErrNoEncryptionKey
	}

	if len(key) != keySize {
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
