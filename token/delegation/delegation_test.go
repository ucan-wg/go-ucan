package delegation_test

import (
	_ "embed"
	"encoding/base64"
	"testing"
	"time"

	"github.com/MetaMask/go-did-it/didtest"
	"github.com/stretchr/testify/require"

	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/pkg/policy"
	"github.com/ucan-wg/go-ucan/token/delegation"
)

//go:embed testdata/new.dagjson
var newDagJson []byte

//go:embed testdata/powerline.dagjson
var powerlineDagJson []byte

//go:embed testdata/root.dagjson
var rootDagJson []byte

const (
	nonce = "6roDhGi0kiNriQAz7J3d+bOeoI/tj8ENikmQNbtjnD0"

	subJectCmd = "/foo/bar"
	subjectPol = `
[
  ["==", ".status", "draft"],
  ["all", ".reviewer",
    ["like", ".email", "*@example.com"]
  ],
  ["any", ".tags",
    ["or", [
      ["==", ".", "news"],
      ["==", ".", "press"]
    ]]
  ]
]
`

	aesKey = "xQklMmNTnVrmaPBq/0pwV5fEwuv/iClF5HWak9MsgI8="
)

func TestConstructors(t *testing.T) {
	t.Parallel()

	cmd, err := command.Parse(subJectCmd)
	require.NoError(t, err)

	pol, err := policy.FromDagJson(subjectPol)
	require.NoError(t, err)

	exp, err := time.Parse(time.RFC3339, "2200-01-01T00:00:00Z")
	require.NoError(t, err)

	t.Run("New", func(t *testing.T) {
		tkn, err := delegation.New(didtest.PersonaAlice.DID(), didtest.PersonaBob.DID(), cmd, pol, didtest.PersonaCarol.DID(),
			delegation.WithNonce([]byte(nonce)),
			delegation.WithExpiration(exp),
			delegation.WithMeta("foo", "fooo"),
			delegation.WithMeta("bar", "barr"),
		)
		require.NoError(t, err)

		require.False(t, tkn.IsRoot())
		require.False(t, tkn.IsPowerline())

		data, err := tkn.ToDagJson(didtest.PersonaAlice.PrivKey())
		require.NoError(t, err)

		require.Equal(t, newDagJson, data)
	})

	t.Run("Root", func(t *testing.T) {
		tkn, err := delegation.Root(didtest.PersonaAlice.DID(), didtest.PersonaBob.DID(), cmd, pol,
			delegation.WithNonce([]byte(nonce)),
			delegation.WithExpiration(exp),
			delegation.WithMeta("foo", "fooo"),
			delegation.WithMeta("bar", "barr"),
		)
		require.NoError(t, err)

		require.True(t, tkn.IsRoot())
		require.False(t, tkn.IsPowerline())

		data, err := tkn.ToDagJson(didtest.PersonaAlice.PrivKey())
		require.NoError(t, err)

		require.Equal(t, rootDagJson, data)
	})

	t.Run("Powerline", func(t *testing.T) {
		tkn, err := delegation.Powerline(didtest.PersonaAlice.DID(), didtest.PersonaBob.DID(), cmd, pol,
			delegation.WithNonce([]byte(nonce)),
			delegation.WithExpiration(exp),
			delegation.WithMeta("foo", "fooo"),
			delegation.WithMeta("bar", "barr"),
		)
		require.NoError(t, err)

		require.False(t, tkn.IsRoot())
		require.True(t, tkn.IsPowerline())

		data, err := tkn.ToDagJson(didtest.PersonaAlice.PrivKey())
		require.NoError(t, err)

		require.Equal(t, powerlineDagJson, data)
	})
}

func TestEncryptedMeta(t *testing.T) {
	t.Parallel()

	cmd, err := command.Parse(subJectCmd)
	require.NoError(t, err)
	pol, err := policy.FromDagJson(subjectPol)
	require.NoError(t, err)

	encryptionKey, err := base64.StdEncoding.DecodeString(aesKey)
	require.NoError(t, err)
	require.Len(t, encryptionKey, 32)

	tests := []struct {
		name        string
		key         string
		value       string
		expectError bool
	}{
		{
			name:  "simple string",
			key:   "secret1",
			value: "hello world",
		},
		{
			name:  "empty string",
			key:   "secret2",
			value: "",
		},
		{
			name:  "special characters",
			key:   "secret3",
			value: "!@#$%^&*()_+-=[]{}|;:,.<>?",
		},
		{
			name:  "unicode characters",
			key:   "secret4",
			value: "Hello, 世界! 👋",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tkn, err := delegation.Root(didtest.PersonaAlice.DID(), didtest.PersonaBob.DID(), cmd, pol,
				delegation.WithEncryptedMetaString(tt.key, tt.value, encryptionKey),
			)
			require.NoError(t, err)

			data, err := tkn.ToDagCbor(didtest.PersonaAlice.PrivKey())
			require.NoError(t, err)

			decodedTkn, _, err := delegation.FromSealed(data)
			require.NoError(t, err)

			_, err = decodedTkn.Meta().GetString(tt.key)
			require.Error(t, err)

			decrypted, err := decodedTkn.Meta().GetEncryptedString(tt.key, encryptionKey)
			require.NoError(t, err)
			// Verify the decrypted value is equal to the original
			require.Equal(t, tt.value, decrypted)

			// Try to decrypt with wrong key
			wrongKey := make([]byte, 32)
			_, err = decodedTkn.Meta().GetEncryptedString(tt.key, wrongKey)
			require.Error(t, err)
		})
	}

	t.Run("multiple encrypted values in the same token", func(t *testing.T) {
		values := map[string]string{
			"secret1": "value1",
			"secret2": "value2",
			"secret3": "value3",
		}
		var opts []delegation.Option
		for k, v := range values {
			opts = append(opts, delegation.WithEncryptedMetaString(k, v, encryptionKey))
		}

		// Create token with multiple encrypted values
		tkn, err := delegation.Root(didtest.PersonaAlice.DID(), didtest.PersonaBob.DID(), cmd, pol, opts...)
		require.NoError(t, err)

		data, err := tkn.ToDagCbor(didtest.PersonaAlice.PrivKey())
		require.NoError(t, err)

		decodedTkn, _, err := delegation.FromSealed(data)
		require.NoError(t, err)

		for k, v := range values {
			decrypted, err := decodedTkn.Meta().GetEncryptedString(k, encryptionKey)
			require.NoError(t, err)
			require.Equal(t, v, decrypted)
		}
	})
}
