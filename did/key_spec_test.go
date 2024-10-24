//go:build jwx_es256k

package did_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ucan-wg/go-ucan/did"
	"github.com/ucan-wg/go-ucan/did/testvectors"
)

// TestDidKeyVectors executes tests read from the [test vector files] provided
// as part of the DID Key method's [specification].
//
// [test vector files]: https://github.com/w3c-ccg/did-method-key/tree/main/test-vectors
// [specification]: https://w3c-ccg.github.io/did-method-key
func TestDidKeyVectors(t *testing.T) {
	t.Parallel()

	for _, f := range []string{
		// TODO: These test vectors are not supported by go-libp2p/core/crypto
		// "bls12381.json",
		"ed25519-x25519.json",
		"nist-curves.json",
		"rsa.json",
		"secp256k1.json",
		// This test vector only contains a DID Document
		// "x25519.json",
	} {
		vectors := loadTestVectors(t, f)

		t.Run(f, func(t *testing.T) {
			t.Parallel()

			for k, vector := range vectors {
				t.Run(k, func(t *testing.T) {
					// round-trip pubkey-->did-->pubkey, verified against the test vectors.

					exp := vectorPubKey(t, vector)

					id, err := did.FromPubKey(exp)
					require.NoError(t, err, f, k)
					act, err := id.PubKey()
					require.NoError(t, err)

					assert.Equal(t, k, id.String(), f, k)
					assert.Equal(t, exp, act)
				})
			}
		})
	}
}

func loadTestVectors(t *testing.T, filename string) testvectors.Vectors {
	t.Helper()

	data, err := os.ReadFile(filepath.Join("testvectors", filename))
	require.NoError(t, err)

	var vs testvectors.Vectors

	require.NoError(t, json.Unmarshal(data, &vs))

	return vs
}

func vectorPubKey(t *testing.T, v testvectors.Vector) crypto.PubKey {
	t.Helper()

	pubKey, err := v.PubKey()
	require.NoError(t, err)
	require.NotZero(t, pubKey)

	return pubKey
}
