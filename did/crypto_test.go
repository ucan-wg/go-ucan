package did_test

import (
	"crypto/elliptic"
	"crypto/rand"
	"testing"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/crypto/pb"
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
	_, ecdsaSecp256k1, err := crypto.GenerateECDSAKeyPairWithCurve(secp256k1.S256(), rand.Reader)
	require.NoError(t, err)
	_, ed25519, err := crypto.GenerateEd25519Key(rand.Reader)
	require.NoError(t, err)
	_, rsa, err := crypto.GenerateRSAKeyPair(2048, rand.Reader)
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
	t.Run("Ed25519", test(ed25519, did.Ed25519))
	t.Run("RSA", test(rsa, did.RSA))
	t.Run("secp256k1", test(secp256k1PubKey1, did.Secp256k1))

	t.Run("ECDSA with secp256k1 curve (coerced)", func(t *testing.T) {
		t.Parallel()

		id, err := did.FromPubKey(ecdsaSecp256k1)
		require.NoError(t, err)
		p, err := id.PubKey()
		require.NoError(t, err)
		require.Equal(t, pb.KeyType_Secp256k1, p.Type())
	})

	t.Run("unmarshaled example key (secp256k1)", func(t *testing.T) {
		t.Parallel()

		id, err := did.FromPubKey(examplePubKey(t))
		require.NoError(t, err)
		require.Equal(t, exampleDID(t), id)
	})
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
