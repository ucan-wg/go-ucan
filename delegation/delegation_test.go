package delegation_test

import (
	"crypto/rand"
	"testing"
	"time"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gotest.tools/v3/golden"

	"github.com/ucan-wg/go-ucan/capability/command"
	"github.com/ucan-wg/go-ucan/capability/policy"
	"github.com/ucan-wg/go-ucan/delegation"
	"github.com/ucan-wg/go-ucan/did"
	"github.com/ucan-wg/go-ucan/internal/envelope"
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
		dlg, err := delegation.New(privKey, aud, cmd, pol,
			delegation.WithNonce([]byte(nonce)),
			delegation.WithSubject(sub),
			delegation.WithExpiration(exp),
			delegation.WithMeta("foo", "fooo"),
			delegation.WithMeta("bar", "barr"),
		)
		require.NoError(t, err)

		data, err := dlg.ToDagJson(privKey)
		require.NoError(t, err)

		t.Log(string(data))

		golden.Assert(t, string(data), "new.dagjson")
		assert.Equal(t, newCID, envelope.CIDToBase58BTC(dlg.CID()))
	})

	t.Run("Root", func(t *testing.T) {
		t.Parallel()

		dlg, err := delegation.Root(privKey, aud, cmd, pol,
			delegation.WithNonce([]byte(nonce)),
			delegation.WithExpiration(exp),
			delegation.WithMeta("foo", "fooo"),
			delegation.WithMeta("bar", "barr"),
		)
		require.NoError(t, err)

		data, err := dlg.ToDagJson(privKey)
		require.NoError(t, err)

		t.Log(string(data))

		golden.Assert(t, string(data), "root.dagjson")
		assert.Equal(t, rootCID, envelope.CIDToBase58BTC(dlg.CID()))
	})
}

func privKey(t require.TestingT, privKeyCfg string) crypto.PrivKey {
	privKeyMar, err := crypto.ConfigDecodeKey(privKeyCfg)
	require.NoError(t, err)

	privKey, err := crypto.UnmarshalPrivateKey(privKeyMar)
	require.NoError(t, err)

	return privKey
}

func TestKey(t *testing.T) {
	// TODO: why is this broken?
	t.Skip("TODO: why is this broken?")

	priv, _, err := crypto.GenerateEd25519Key(rand.Reader)
	require.NoError(t, err)

	privMar, err := crypto.MarshalPrivateKey(priv)
	require.NoError(t, err)

	privCfg := crypto.ConfigEncodeKey(privMar)
	t.Log(privCfg)

	id, err := did.FromPubKey(priv.GetPublic())
	require.NoError(t, err)
	t.Log(id)

	t.Fail()
}
