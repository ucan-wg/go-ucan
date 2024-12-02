package delegation_test

import (
	"encoding/base64"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"gotest.tools/v3/golden"

	"github.com/ucan-wg/go-ucan/did/didtest"
	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/pkg/policy"
	"github.com/ucan-wg/go-ucan/pkg/policy/limits"
	"github.com/ucan-wg/go-ucan/token/delegation"
)

const (
	nonce = "6roDhGi0kiNriQAz7J3d+bOeoI/tj8ENikmQNbtjnD0"

	subJectCmd = "/foo/bar"
	subjectPol = `
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

	newCID  = "zdpuAwa4qv3ncMDPeDoqVxjZy3JoyWsbqUzm94rdA1AvRFkkw"
	rootCID = "zdpuAkgGmUp5JrXvehGuuw9JA8DLQKDaxtK3R8brDQQVC2i5X"

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
		tkn, err := delegation.New(didtest.PersonaAlice.DID(), didtest.PersonaBob.DID(), cmd, pol,
			delegation.WithNonce([]byte(nonce)),
			delegation.WithSubject(didtest.PersonaAlice.DID()),
			delegation.WithExpiration(exp),
			delegation.WithMeta("foo", "fooo"),
			delegation.WithMeta("bar", "barr"),
		)
		require.NoError(t, err)

		data, err := tkn.ToDagJson(didtest.PersonaAlice.PrivKey())
		require.NoError(t, err)

		golden.Assert(t, string(data), "new.dagjson")
	})

	t.Run("Root", func(t *testing.T) {
		t.Parallel()

		tkn, err := delegation.Root(didtest.PersonaAlice.DID(), didtest.PersonaBob.DID(), cmd, pol,
			delegation.WithNonce([]byte(nonce)),
			delegation.WithExpiration(exp),
			delegation.WithMeta("foo", "fooo"),
			delegation.WithMeta("bar", "barr"),
		)
		require.NoError(t, err)

		data, err := tkn.ToDagJson(didtest.PersonaAlice.PrivKey())
		require.NoError(t, err)

		golden.Assert(t, string(data), "root.dagjson")
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

			tkn, err := delegation.New(didtest.PersonaAlice.DID(), didtest.PersonaBob.DID(), cmd, pol,
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
		tkn, err := delegation.New(didtest.PersonaAlice.DID(), didtest.PersonaBob.DID(), cmd, pol, opts...)
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

func TestTokenTimestampBounds(t *testing.T) {
	t.Parallel()

	cmd, err := command.Parse("/foo/bar")
	require.NoError(t, err)
	pol, err := policy.FromDagJson("[]")
	require.NoError(t, err)

	tomorrow := time.Now().Add(24 * time.Hour).Unix()

	tests := []struct {
		name    string
		nbf     int64
		exp     int64
		wantErr bool
	}{
		{
			name:    "valid timestamps",
			nbf:     tomorrow,
			exp:     tomorrow + 3600,
			wantErr: false,
		},
		{
			name:    "max safe integer",
			nbf:     tomorrow,
			exp:     limits.MaxInt53,
			wantErr: false,
		},
		{
			name:    "exceeds max safe integer",
			nbf:     tomorrow,
			exp:     limits.MaxInt53 + 1,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			_, err = delegation.New(didtest.PersonaAlice.DID(), didtest.PersonaBob.DID(),
				cmd, pol,
				delegation.WithNotBefore(time.Unix(tt.nbf, 0)),
				delegation.WithExpiration(time.Unix(tt.exp, 0)),
			)

			if tt.wantErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), "exceeds safe integer bounds")
			} else {
				require.NoError(t, err)
			}
		})
	}

	t.Run("nbf overflow", func(t *testing.T) {
		t.Parallel()

		futureExp := time.Now().Add(48 * time.Hour).Unix()
		_, err := delegation.New(didtest.PersonaAlice.DID(), didtest.PersonaBob.DID(),
			cmd, pol,
			delegation.WithNotBefore(time.Unix(limits.MaxInt53+1, 0)),
			delegation.WithExpiration(time.Unix(futureExp, 0)),
		)
		require.Error(t, err)
		require.Contains(t, err.Error(), "exceeds safe integer bounds")
	})
}
