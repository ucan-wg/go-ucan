package meta_test

import (
	"crypto/rand"
	"maps"
	"testing"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	"github.com/stretchr/testify/require"

	"github.com/ucan-wg/go-ucan/pkg/meta"
)

func TestMeta_Add(t *testing.T) {
	t.Parallel()

	type Unsupported struct{}

	t.Run("error if not primitive or Node", func(t *testing.T) {
		t.Parallel()

		err := (&meta.Meta{}).Add("invalid", &Unsupported{})
		require.Error(t, err)
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

		_, err = m.GetString("secret")
		require.Error(t, err) // the ciphertext is saved as []byte instead of string

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
			require.Error(t, err)
			require.Contains(t, err.Error(), "encryption key is required")
		})
	})
}

func TestIterCloneEquals(t *testing.T) {
	m := meta.NewMeta()

	require.NoError(t, m.Add("foo", "bar"))
	require.NoError(t, m.Add("baz", 1234))

	expected := map[string]ipld.Node{
		"foo": basicnode.NewString("bar"),
		"baz": basicnode.NewInt(1234),
	}

	// meta -> iter
	require.Equal(t, expected, maps.Collect(m.Iter()))

	// readonly -> iter
	ro := m.ReadOnly()
	require.Equal(t, expected, maps.Collect(ro.Iter()))

	// meta -> clone -> iter
	clone := m.Clone()
	require.Equal(t, expected, maps.Collect(clone.Iter()))

	// readonly -> WriteableClone -> iter
	wclone := ro.WriteableClone()
	require.Equal(t, expected, maps.Collect(wclone.Iter()))

	require.True(t, m.Equals(wclone))
	require.True(t, ro.Equals(wclone.ReadOnly()))
}

func TestInclude(t *testing.T) {
	m1 := meta.NewMeta()

	require.NoError(t, m1.Add("samekey", "bar"))
	require.NoError(t, m1.Add("baz", 1234))

	m2 := meta.NewMeta()

	require.NoError(t, m2.Add("samekey", "othervalue")) // check no overwrite
	require.NoError(t, m2.Add("otherkey", 1234))

	m1.Include(m2)

	require.Equal(t, map[string]ipld.Node{
		"samekey":  basicnode.NewString("bar"),
		"baz":      basicnode.NewInt(1234),
		"otherkey": basicnode.NewInt(1234),
	}, maps.Collect(m1.Iter()))
}
