package meta_test

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/require"
	"gotest.tools/v3/assert"

	"github.com/ucan-wg/go-ucan/pkg/meta"
)

func TestMeta_Add(t *testing.T) {
	t.Parallel()

	type Unsupported struct{}

	t.Run("error if not primitive or Node", func(t *testing.T) {
		t.Parallel()

		err := (&meta.Meta{}).Add("invalid", &Unsupported{})
		require.ErrorIs(t, err, meta.ErrUnsupported)
		assert.ErrorContains(t, err, "*github.com/ucan-wg/go-ucan/pkg/meta_test.Unsupported")
	})

	t.Run("encrypted meta", func(t *testing.T) {
		t.Parallel()

		key := make([]byte, 32)
		_, err := rand.Read(key)
		require.NoError(t, err)

		m := meta.NewMeta()

		// string encryption
		err = m.AddEncrypted("secret", "hello world", key)
		require.NoError(t, err)

		encrypted, err := m.GetString("secret")
		require.NoError(t, err)
		require.NotEqual(t, "hello world", encrypted)

		decrypted, err := m.GetEncryptedString("secret", key)
		require.NoError(t, err)
		require.Equal(t, "hello world", decrypted)

		// bytes encryption
		originalBytes := make([]byte, 128)
		_, err = rand.Read(originalBytes)
		require.NoError(t, err)
		err = m.AddEncrypted("secret-bytes", originalBytes, key)
		require.NoError(t, err)

		encryptedBytes, err := m.GetBytes("secret-bytes")
		require.NoError(t, err)
		require.NotEqual(t, originalBytes, encryptedBytes)

		decryptedBytes, err := m.GetEncryptedBytes("secret-bytes", key)
		require.NoError(t, err)
		require.Equal(t, originalBytes, decryptedBytes)

		// error cases
		t.Run("error on unsupported type", func(t *testing.T) {
			err := m.AddEncrypted("invalid", 123, key)
			require.ErrorIs(t, err, meta.ErrNotEncryptable)
		})

		t.Run("error on invalid key size", func(t *testing.T) {
			err := m.AddEncrypted("invalid", "test", []byte("short-key"))
			require.Error(t, err)
			require.Contains(t, err.Error(), "invalid key size")
		})

		t.Run("error on nil key", func(t *testing.T) {
			err := m.AddEncrypted("invalid", "test", nil)
			require.NoError(t, err)
			// with nil key, value should be stored unencrypted
			val, err := m.GetString("invalid")
			require.NoError(t, err)
			require.Equal(t, "test", val)
		})
	})
}
