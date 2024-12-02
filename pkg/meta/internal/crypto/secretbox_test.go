package crypto

import (
	"bytes"
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSecretBoxEncryption(t *testing.T) {
	t.Parallel()

	key := make([]byte, keySize) // generate random 32-byte key
	_, errKey := rand.Read(key)
	require.NoError(t, errKey)

	tests := []struct {
		name    string
		data    []byte
		key     []byte
		wantErr error
	}{
		{
			name: "valid encryption/decryption",
			data: []byte("hello world"),
			key:  key,
		},
		{
			name:    "nil key returns error",
			data:    []byte("hello world"),
			key:     nil,
			wantErr: ErrNoEncryptionKey,
		},
		{
			name: "empty data",
			data: []byte{},
			key:  key,
		},
		{
			name:    "invalid key size",
			data:    []byte("hello world"),
			key:     make([]byte, 16), // Only 32 bytes allowed now
			wantErr: ErrInvalidKeySize,
		},
		{
			name:    "zero key returns error",
			data:    []byte("hello world"),
			key:     make([]byte, keySize),
			wantErr: ErrZeroKey,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			encrypted, err := EncryptWithKey(tt.data, tt.key)
			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)

			// Verify encrypted data is different and includes nonce
			require.Greater(t, len(encrypted), 24) // At least nonce size
			if len(tt.data) > 0 {
				require.NotEqual(t, tt.data, encrypted[24:]) // Ignore nonce prefix
			}

			decrypted, err := DecryptStringWithKey(encrypted, tt.key)
			require.NoError(t, err)
			require.True(t, bytes.Equal(tt.data, decrypted))
		})
	}
}

func TestDecryptionErrors(t *testing.T) {
	t.Parallel()

	key := make([]byte, keySize)
	_, err := rand.Read(key)
	require.NoError(t, err)

	// Create valid encrypted data for tampering tests
	validData := []byte("test message")
	encrypted, err := EncryptWithKey(validData, key)
	require.NoError(t, err)

	tests := []struct {
		name   string
		data   []byte
		key    []byte
		errMsg string
	}{
		{
			name:   "short ciphertext",
			data:   make([]byte, 23), // Less than nonce size
			key:    key,
			errMsg: "ciphertext too short",
		},
		{
			name:   "invalid ciphertext",
			data:   make([]byte, 24), // Just nonce size
			key:    key,
			errMsg: "decryption failed",
		},
		{
			name:   "tampered ciphertext",
			data:   tamperWithBytes(encrypted),
			key:    key,
			errMsg: "decryption failed",
		},
		{
			name:   "missing key",
			data:   encrypted,
			key:    nil,
			errMsg: "encryption key is required",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			_, err := DecryptStringWithKey(tt.data, tt.key)
			require.Error(t, err)
			require.Contains(t, err.Error(), tt.errMsg)
		})
	}
}

// tamperWithBytes modifies a byte in the encrypted data to simulate tampering
func tamperWithBytes(data []byte) []byte {
	if len(data) < 25 { // Need at least nonce + 1 byte
		return data
	}
	tampered := make([]byte, len(data))
	copy(tampered, data)
	tampered[24] ^= 0x01 // Modify first byte after nonce
	return tampered
}
