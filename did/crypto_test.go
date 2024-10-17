package did_test

import (
	"crypto/elliptic"
	"crypto/rand"
	"testing"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/multiformats/go-multicodec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ucan-wg/go-ucan/did"
)

const (
	exampleDIDStr    = "did:key:z6MkhaXgBZDvotDkL5257faiztiGiC2QtKLGpbnnEGta2doK"
	examplePubKeyStr = "Lm/M42cB3HkUiODQsXRcweM6TByfzEHGO9ND274JcOY="
)

func TestFromPubKey(t *testing.T) {
	t.Parallel()

	_, ecdsaP256, err := crypto.GenerateECDSAKeyPairWithCurve(elliptic.P256(), rand.Reader)
	require.NoError(t, err)
	_, ecdsaP384, err := crypto.GenerateECDSAKeyPairWithCurve(elliptic.P384(), rand.Reader)
	require.NoError(t, err)
	_, ecdsaP521, err := crypto.GenerateECDSAKeyPairWithCurve(elliptic.P521(), rand.Reader)
	require.NoError(t, err)
	_, secp256k1PubKey1, err := crypto.GenerateSecp256k1Key(rand.Reader)
	require.NoError(t, err)

	test := func(pub crypto.PubKey, code multicodec.Code) func(t *testing.T) {
		t.Helper()

		return func(t *testing.T) {
			t.Parallel()

			id, err := did.FromPubKey(pub)
			require.NoError(t, err)
			p, err := id.PubKey()
			require.NoError(t, err)
			assert.Equal(t, pub, p)
		}
	}

	t.Run("ECDSA with P256 curve", test(ecdsaP256, did.P256))
	t.Run("ECDSA with P384 curve", test(ecdsaP384, did.P384))
	t.Run("ECDSA with P521 curve", test(ecdsaP521, did.P521))
	t.Run("With secp256k1 (secp256k1)", test(secp256k1PubKey1, did.Secp256k1))

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
