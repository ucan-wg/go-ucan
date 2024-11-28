package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
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

func encryptWithAESKey(data, key []byte) ([]byte, error) {
	if err := validateKey(key); err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return aesGCM.Seal(nonce, nonce, data, nil), nil
}

func decryptWithAESKey(data, key []byte) ([]byte, error) {
	if err := validateKey(key); err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := aesGCM.NonceSize()
	if len(data) < nonceSize {
		return nil, ErrShortCipherText
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, errors.New("decryption failed")
	}

	return plaintext, nil
}

func BenchmarkEncryption(b *testing.B) {
	key := make([]byte, keySize)
	_, err := rand.Read(key)
	require.NoError(b, err)

	sizes := []int{16, 64, 256, 1024, 4096} // Test different payload sizes
	for _, size := range sizes {
		data := make([]byte, size)
		_, err := rand.Read(data)
		require.NoError(b, err)

		b.Run(fmt.Sprintf("Secretbox-%dB", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				encrypted, err := EncryptWithKey(data, key)
				require.NoError(b, err)
				b.SetBytes(int64(len(encrypted)))
			}
		})

		b.Run(fmt.Sprintf("AES-GCM-%dB", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				encrypted, err := encryptWithAESKey(data, key)
				require.NoError(b, err)
				b.SetBytes(int64(len(encrypted)))
			}
		})
	}
}

func BenchmarkDecryption(b *testing.B) {
	key := make([]byte, keySize)
	_, err := rand.Read(key)
	require.NoError(b, err)

	sizes := []int{16, 64, 256, 1024, 4096}
	for _, size := range sizes {
		data := make([]byte, size)
		_, err := rand.Read(data)
		require.NoError(b, err)

		secretboxEncrypted, err := EncryptWithKey(data, key)
		require.NoError(b, err)

		aesGCMEncrypted, err := encryptWithAESKey(data, key)
		require.NoError(b, err)

		b.Run(fmt.Sprintf("Secretbox-%dB", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				decrypted, err := DecryptStringWithKey(secretboxEncrypted, key)
				require.NoError(b, err)
				b.SetBytes(int64(len(decrypted)))
			}
		})

		b.Run(fmt.Sprintf("AES-GCM-%dB", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				decrypted, err := decryptWithAESKey(aesGCMEncrypted, key)
				require.NoError(b, err)
				b.SetBytes(int64(len(decrypted)))
			}
		})
	}
}

// TestCiphertextSizeComparison shows that Secretbox encryption entails
// a slightly larger ciphertext overhead of 40 bytes, compared to AES-GCM,
// whose overhead is just 28 bytes.
func TestCiphertextSizeComparison(t *testing.T) {
	key := make([]byte, keySize)
	_, err := rand.Read(key)
	require.NoError(t, err)

	sizes := []int{0, 16, 64, 256, 1024, 4096}
	for _, size := range sizes {
		t.Run(fmt.Sprintf("size-%d", size), func(t *testing.T) {
			data := make([]byte, size)
			_, err := rand.Read(data)
			require.NoError(t, err)

			sbCiphertext, err := EncryptWithKey(data, key)
			require.NoError(t, err)

			aesCiphertext, err := encryptWithAESKey(data, key)
			require.NoError(t, err)

			t.Logf("Input size: %d bytes", size)
			t.Logf("Secretbox size: %d bytes (overhead: %d bytes)", len(sbCiphertext), len(sbCiphertext)-size)
			t.Logf("AES-GCM size: %d bytes (overhead: %d bytes)", len(aesCiphertext), len(aesCiphertext)-size)
		})
	}
}
