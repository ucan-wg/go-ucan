package ucan_test

import (
	"crypto/rand"
	"fmt"
	"testing"
	"time"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"github.com/ipld/go-ipld-prime/schema"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/ucan-wg/go-ucan"
	"github.com/ucan-wg/go-ucan/capability/command"
	"github.com/ucan-wg/go-ucan/capability/policy"
	"github.com/ucan-wg/go-ucan/did"
)

const (
	ed25519PrivKeyCfg = "CAESQL1hvbXpiuk2pWr/XFbfHJcZNpJ7S90iTA3wSCTc/BPRneCwPnCZb6c0vlD6ytDWqaOt0HEOPYnqEpnzoBDprSM="
	ed25519DID        = "did:key:z6Mkq5YmbJcTrPExNDi26imrTCpKhepjBFBSHqrBDN2ArPkv"

	issuerPrivKeyCfg = "CAESQLSql38oDmQXIihFFaYIjb73mwbPsc7MIqn4o8PN4kRNnKfHkw5gRP1IV9b6d0estqkZayGZ2vqMAbhRixjgkDU="
	issuerDID        = "did:key:z6Mkpzn2n3ZGT2VaqMGSQC3tzmzV4TS9S71iFsDXE1WnoNH2"

	subjectPrivKeyCfg = "CAESQL9RtjZ4dQBeXtvDe53UyvslSd64kSGevjdNiA1IP+hey5i/3PfRXSuDr71UeJUo1fLzZ7mGldZCOZL3gsIQz5c="
	subjectDID        = "did:key:z6MktA1uBdCpq4uJBqE9jjMiLyxZBg9a6xgPPKJjMqss6Zc2"
	subJectCmd        = "/foo/bar"
	subjectPol        = `
[
		["==", ".status", "draft"],
		["all", ".reviewer", [
		["like", ".email", "*@example.com"]]
	],
		["any", ".tags", [ 
		["or", [
			["==", ".", "news"], 
			["==", ".", "press"]]
			]]
	]
]
`
)

func TestNewAuthority(t *testing.T) {
	t.Parallel()

	t.Run("with default configuration", func(t *testing.T) {
		t.Parallel()

		authority := authority(t, ed25519PrivKeyCfg)
		assert.Equal(t, ed25519DID, authority.DID().String())
		assert.Equal(t, ucan.DefaultNonceLength, authority.NonceLength())
		assert.Equal(t, ucan.DefaultExpiration, authority.Expiration())
	})
}

func TestAuthority_Nonce(t *testing.T) {
	t.Parallel()

	fixture := func(t *testing.T, exp int, opts ...ucan.AuthorityOption) {
		authority := authority(t, ed25519PrivKeyCfg, opts...)

		nonce, err := authority.Nonce()
		require.NoError(t, err)
		assert.Len(t, nonce, exp)
	}

	t.Run("with default nonce length", func(t *testing.T) {
		t.Parallel()
		fixture(t, ucan.DefaultNonceLength)
	})

	t.Run("with custom nonce length", func(t *testing.T) {
		t.Parallel()
		fixture(t, 64, ucan.WithNonceLength(64))
	})
}

func TestIssue(t *testing.T) {
	t.Parallel()

	privKey := privKey(t, issuerPrivKeyCfg)

	id, err := did.Parse(subjectDID)
	require.NoError(t, err)

	cmd, err := command.Parse(subJectCmd)
	require.NoError(t, err)

	pol, err := policy.FromDagJson(subjectPol)
	require.NoError(t, err)

	now := time.Now().Add(ucan.DefaultExpiration)

	// meta := map[string]any{
	// 	"foo": "fooo",
	// 	"bar": "barr",
	// }

	env, err := ucan.Issue(privKey, id, cmd, &pol, &now)
	require.NoError(t, err)

	node, err := env.Wrap()

	typed, ok := node.(schema.TypedNode)
	require.True(t, ok)

	json, err := ipld.Encode(typed.Representation(), dagjson.Encode)
	require.NoError(t, err)

	fmt.Println(string(json))

	t.Fail()
}

func authority(t *testing.T, privKeyCfg string, opts ...ucan.AuthorityOption) *ucan.Authority {
	t.Helper()

	privKey := privKey(t, ed25519PrivKeyCfg)

	authority, err := ucan.NewAuthority(privKey, opts...)
	require.NoError(t, err)

	return authority
}

func privKey(t *testing.T, privKeyCfg string) crypto.PrivKey {
	t.Helper()

	privKeyMar, err := crypto.ConfigDecodeKey(privKeyCfg)
	require.NoError(t, err)

	privKey, err := crypto.UnmarshalPrivateKey(privKeyMar)
	require.NoError(t, err)

	return privKey
}

func TestKey(t *testing.T) {
	t.Skip()

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
