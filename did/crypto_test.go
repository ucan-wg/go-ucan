package did_test

import (
	"testing"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/stretchr/testify/require"
	"github.com/ucan-wg/go-ucan/v1/did"
)

const (
	exampleDIDStr    = "did:key:z6MkhaXgBZDvotDkL5257faiztiGiC2QtKLGpbnnEGta2doK"
	examplePubKeyStr = "Lm/M42cB3HkUiODQsXRcweM6TByfzEHGO9ND274JcOY="
)

func TestFromPubKey(t *testing.T) {
	t.Parallel()

	id, err := did.FromPubKey(examplePubKey(t))
	require.NoError(t, err)
	require.Equal(t, exampleDID(t), id)
}

func TestToPubKey(t *testing.T) {
	t.Parallel()

	pubKey, err := did.ToPubKey(exampleDIDStr)
	require.NoError(t, err)
	require.Equal(t, examplePubKey(t), pubKey)
}

func exampleDID(t *testing.T) did.DID {
	t.Helper()

	id, err := did.Parse(exampleDIDStr)
	require.NoError(t, err)

	return id
}

func examplePubKey(t *testing.T) crypto.PubKey {
	t.Helper()

	pubKeyCfg, err := crypto.ConfigDecodeKey(examplePubKeyStr)
	require.NoError(t, err)

	pubKey, err := crypto.UnmarshalEd25519PublicKey(pubKeyCfg)
	require.NoError(t, err)

	return pubKey
}
