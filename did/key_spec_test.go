// go:build jwx_es256k

package did_test

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/mr-tron/base58"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/ucan-wg/go-ucan/did"
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
		vs := loadTestVectors(t, f)

		t.Run(f, func(t *testing.T) {
			t.Parallel()

			for k, v := range vs {
				f := f

				t.Run(k, func(t *testing.T) {
					t.Parallel()

					exp := vectorPubKey(t, v)

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

func loadTestVectors(t *testing.T, filename string) vectors {
	t.Helper()

	data, err := os.ReadFile(filepath.Join("testdata", filename))
	require.NoError(t, err)

	var vs vectors

	require.NoError(t, json.Unmarshal(data, &vs))

	return vs
}

func vectorPubKey(t *testing.T, v vector) crypto.PubKey {
	t.Helper()

	pubKey, err := v.pubKey()
	require.NoError(t, err)
	require.NotZero(t, pubKey)

	return pubKey
}

func vectorType(t *testing.T, v vector) string {
	vectorType, err := v.pubKeyType()
	require.NoError(t, err)

	return vectorType
}

type vectors map[string]vector

// This is pretty gross but the structure allows the repeated verifier,
// PublicKeyJwk and PublicKeyBase58 account for the fact that the test
// files are very inconsistent.
type vector struct {
	VerificationKeyPair verifier
	VerificationMethod  verifier
	PublicKeyJwk        json.RawMessage
	DidDocument         json.RawMessage // TODO: if we start producing DID documents, we should test this too
}

type verifier struct {
	ID              string
	Type            string
	PublicKeyBase58 string
	PublicKeyJwk    json.RawMessage
}

func (v vector) pubKey() (crypto.PubKey, error) {
	// If the public key is in base58
	if pubB58 := v.pubKeyBase58(); len(pubB58) > 0 {
		pubBytes, err := base58.Decode(pubB58)
		if err != nil {
			return nil, err
		}

		t, err := v.pubKeyType()
		if err != nil {
			return nil, err
		}

		var unmarshaler crypto.PubKeyUnmarshaller

		switch t {
		case "Ed25519VerificationKey2018":
			unmarshaler = crypto.UnmarshalEd25519PublicKey
		case "EcdsaSecp256k1VerificationKey2019":
			unmarshaler = crypto.UnmarshalSecp256k1PublicKey
		// This is weak as it assumes the P256 curve - that's all the vectors contain (for now)
		case "P256Key2021":
			unmarshaler = compressedEcdsaPublicKeyUnmarshaler
		default:
			return nil, errors.New("failed to resolve unmarshaler")
		}

		return unmarshaler(pubBytes)
	}

	// If the public key is in a JWK
	if pubJwk := v.pubKeyJwk(); len(pubJwk) > 0 {
		key, err := jwk.ParseKey(pubJwk)
		if err != nil {
			return nil, err
		}

		var a any

		if err := key.Raw(&a); err != nil {
			return nil, err
		}

		switch a.(type) {
		case *ecdsa.PublicKey:
			epub := a.(*ecdsa.PublicKey)

			if epub.Curve == secp256k1.S256() {
				bytes := append([]byte{0x04}, append(epub.X.Bytes(), epub.Y.Bytes()...)...)

				return crypto.UnmarshalSecp256k1PublicKey(bytes)
			}

			asn1, err := x509.MarshalPKIXPublicKey(epub)
			if err != nil {
				return nil, err
			}

			return crypto.UnmarshalECDSAPublicKey(asn1)
		case ed25519.PublicKey:
			return crypto.UnmarshalEd25519PublicKey(a.(ed25519.PublicKey))
		case *rsa.PublicKey:
			asn1, err := x509.MarshalPKIXPublicKey(a.(*rsa.PublicKey))
			if err != nil {
				return nil, err
			}

			return crypto.UnmarshalRsaPublicKey(asn1)
		default:
			return nil, errors.New("unsupported key type")
		}
	}

	// If we don't find a public key at all
	return nil, errors.New("vector's public key not found")
}

func (v vector) pubKeyBase58() string {
	if len(v.VerificationKeyPair.PublicKeyBase58) > 0 {
		return v.VerificationKeyPair.PublicKeyBase58
	}

	return v.VerificationMethod.PublicKeyBase58
}

func (v vector) pubKeyJwk() json.RawMessage {
	if len(v.VerificationKeyPair.PublicKeyJwk) > 0 {
		return v.VerificationKeyPair.PublicKeyJwk
	}

	if len(v.VerificationMethod.PublicKeyJwk) > 0 {
		return v.VerificationMethod.PublicKeyJwk
	}

	return v.PublicKeyJwk
}

func (v vector) pubKeyType() (string, error) {
	if len(v.VerificationKeyPair.Type) > 0 {
		return v.VerificationKeyPair.Type, nil
	}

	if len(v.VerificationMethod.Type) > 0 {
		return v.VerificationMethod.Type, nil
	}

	return "", errors.New("vector's type not found")
}

func compressedEcdsaPublicKeyUnmarshaler(data []byte) (crypto.PubKey, error) {
	x, y := elliptic.UnmarshalCompressed(elliptic.P256(), data)

	ecdsaPublicKey := ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}

	asn1, err := x509.MarshalPKIXPublicKey(&ecdsaPublicKey)
	if err != nil {
		return nil, err
	}

	return crypto.UnmarshalECDSAPublicKey(asn1)
}
