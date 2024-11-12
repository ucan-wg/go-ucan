package delegation_test

import (
	"encoding/base64"
	"testing"
	"time"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/stretchr/testify/require"
	"gotest.tools/v3/golden"

	"github.com/ucan-wg/go-ucan/did"
	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/pkg/policy"
	"github.com/ucan-wg/go-ucan/token/delegation"
)

const (
	nonce = "6roDhGi0kiNriQAz7J3d+bOeoI/tj8ENikmQNbtjnD0"

	AudiencePrivKeyCfg = "CAESQL1hvbXpiuk2pWr/XFbfHJcZNpJ7S90iTA3wSCTc/BPRneCwPnCZb6c0vlD6ytDWqaOt0HEOPYnqEpnzoBDprSM="
	AudienceDID        = "did:key:z6Mkq5YmbJcTrPExNDi26imrTCpKhepjBFBSHqrBDN2ArPkv"

	issuerPrivKeyCfg = "CAESQLSql38oDmQXIihFFaYIjb73mwbPsc7MIqn4o8PN4kRNnKfHkw5gRP1IV9b6d0estqkZayGZ2vqMAbhRixjgkDU="
	issuerDID        = "did:key:z6Mkpzn2n3ZGT2VaqMGSQC3tzmzV4TS9S71iFsDXE1WnoNH2"

	subjectPrivKeyCfg = "CAESQL9RtjZ4dQBeXtvDe53UyvslSd64kSGevjdNiA1IP+hey5i/3PfRXSuDr71UeJUo1fLzZ7mGldZCOZL3gsIQz5c="
	subjectDID        = "did:key:z6MktA1uBdCpq4uJBqE9jjMiLyxZBg9a6xgPPKJjMqss6Zc2"
	subJectCmd        = "/foo/bar"
	subjectPol        = `
[
	[
		"==",
		".status",
		"draft"
	],
	[
		"all",
		".reviewer",
		[
			"like",
			".email",
			"*@example.com"
		]
	],
	[
		"any",
		".tags",
		[
			"or",
			[
				[
					"==",
					".",
					"news"
				],
				[
					"==",
					".",
					"press"
				]
			]
		]
	]
]
`

	newCID  = "zdpuAn9JgGPvnt2WCmTaKktZdbuvcVGTg9bUT5kQaufwUtZ6e"
	rootCID = "zdpuAkgGmUp5JrXvehGuuw9JA8DLQKDaxtK3R8brDQQVC2i5X"

	aesKey = "xQklMmNTnVrmaPBq/0pwV5fEwuv/iClF5HWak9MsgI8="
)

func TestConstructors(t *testing.T) {
	t.Parallel()

	privKey := privKey(t, issuerPrivKeyCfg)

	aud, err := did.Parse(AudienceDID)

	sub, err := did.Parse(subjectDID)
	require.NoError(t, err)

	cmd, err := command.Parse(subJectCmd)
	require.NoError(t, err)

	pol, err := policy.FromDagJson(subjectPol)
	require.NoError(t, err)

	exp, err := time.Parse(time.RFC3339, "2200-01-01T00:00:00Z")
	require.NoError(t, err)

	t.Run("New", func(t *testing.T) {
		tkn, err := delegation.New(privKey, aud, cmd, pol,
			delegation.WithNonce([]byte(nonce)),
			delegation.WithSubject(sub),
			delegation.WithExpiration(exp),
			delegation.WithMeta("foo", "fooo"),
			delegation.WithMeta("bar", "barr"),
		)
		require.NoError(t, err)

		data, err := tkn.ToDagJson(privKey)
		require.NoError(t, err)

		golden.Assert(t, string(data), "new.dagjson")
	})

	t.Run("Root", func(t *testing.T) {
		t.Parallel()

		tkn, err := delegation.Root(privKey, aud, cmd, pol,
			delegation.WithNonce([]byte(nonce)),
			delegation.WithExpiration(exp),
			delegation.WithMeta("foo", "fooo"),
			delegation.WithMeta("bar", "barr"),
		)
		require.NoError(t, err)

		data, err := tkn.ToDagJson(privKey)
		require.NoError(t, err)

		golden.Assert(t, string(data), "root.dagjson")
	})
}

func TestEncryptedMeta(t *testing.T) {
	t.Parallel()

	privKey := privKey(t, issuerPrivKeyCfg)
	aud, err := did.Parse(AudienceDID)
	require.NoError(t, err)
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

			tkn, err := delegation.New(privKey, aud, cmd, pol,
				delegation.WithEncryptedMetaString(tt.key, tt.value, encryptionKey),
			)
			require.NoError(t, err)

			data, err := tkn.ToDagCbor(privKey)
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
		tkn, err := delegation.New(privKey, aud, cmd, pol, opts...)
		require.NoError(t, err)

		data, err := tkn.ToDagCbor(privKey)
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

func privKey(t require.TestingT, privKeyCfg string) crypto.PrivKey {
	privKeyMar, err := crypto.ConfigDecodeKey(privKeyCfg)
	require.NoError(t, err)

	privKey, err := crypto.UnmarshalPrivateKey(privKeyMar)
	require.NoError(t, err)

	return privKey
}
