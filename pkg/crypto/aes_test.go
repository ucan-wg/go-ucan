package crypto

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAESEncryption(t *testing.T) {
	t.Parallel()

	key := make([]byte, 32) // generated random 32-byte key
	_, err := rand.Read(key)
	require.NoError(t, err)

	tests := []struct {
		name    string
		data    []byte
		key     []byte
		wantErr bool
	}{
		{
			name:    "valid encryption/decryption",
			data:    []byte("hello world"),
			key:     key,
			wantErr: false,
		},
		{
			name:    "nil key returns error",
			data:    []byte("hello world"),
			key:     nil,
			wantErr: true,
		},
		{
			name:    "empty data",
			data:    []byte{},
			key:     key,
			wantErr: false,
		},
		{
			name:    "invalid key size",
			data:    []byte("hello world"),
			key:     make([]byte, 31),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			encrypted, err := EncryptWithAESKey(tt.data, tt.key)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			fmt.Println(string(encrypted))

			decrypted, err := DecryptStringWithAESKey(encrypted, tt.key)
			require.NoError(t, err)

			if tt.key == nil {
				require.Equal(t, tt.data, encrypted)
				require.Equal(t, tt.data, decrypted)
			} else {
				require.NotEqual(t, tt.data, encrypted)
				require.True(t, bytes.Equal(tt.data, decrypted))
			}
		})
	}
}

func TestDecryptionErrors(t *testing.T) {
	t.Parallel()

	key := make([]byte, 32)
	_, err := rand.Read(key)
	require.NoError(t, err)

	tests := []struct {
		name   string
		data   []byte
		key    []byte
		errMsg string
	}{
		{
			name:   "short ciphertext",
			data:   []byte("short"),
			key:    key,
			errMsg: "ciphertext too short",
		},
		{
			name:   "invalid ciphertext",
			data:   make([]byte, 16), // just nonce size
			key:    key,
			errMsg: "message authentication failed",
		},
		{
			name:   "missing key",
			data:   []byte("�`M���l\u001AIF�\u0012���=h�?�c� ��\u0012����\u001C�\u0018Ƽ(g"),
			key:    nil,
			errMsg: "encryption key is required",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			_, err := DecryptStringWithAESKey(tt.data, tt.key)
			require.Error(t, err)
			require.Contains(t, err.Error(), tt.errMsg)
		})
	}
}
