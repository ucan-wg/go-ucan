package delegation

import (
	_ "embed"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/MetaMask/go-did-it/crypto"
	"github.com/MetaMask/go-did-it/crypto/ed25519"
	"github.com/multiformats/go-varint"
	"github.com/stretchr/testify/require"
)

// This comes from https://github.com/ucan-wg/spec/blob/main/fixtures/1.0.0/delegation.json
//
//go:embed testdata/interop_delegation.json
var interopDelegation []byte

type interop struct {
	Version    string            `json:"version"`
	Comments   string            `json:"comments"`
	Principals map[string]string `json:"principals"`
	Valid      []validTestCase   `json:"valid"`
}

type validTestCase struct {
	Name     string       `json:"name"`
	Token    string       `json:"token"`
	CID      string       `json:"cid"`
	Envelope envelopeData `json:"envelope"`
}

type envelopeData struct {
	Payload   payloadData `json:"payload"`
	Signature string      `json:"signature"`
	Algorithm string      `json:"alg"`
	Encoding  string      `json:"enc"`
	Spec      string      `json:"spec"`
	Version   string      `json:"version"`
}

type payloadData struct {
	Issuer    string          `json:"iss"`
	Audience  string          `json:"aud"`
	Subject   string          `json:"sub"`
	Command   string          `json:"cmd"`
	Policies  json.RawMessage `json:"pol"`
	ExpiresAt int64           `json:"exp"`
	Nonce     string          `json:"nonce"`
}

func TestInterop(t *testing.T) {
	var testData interop
	err := json.Unmarshal(interopDelegation, &testData)
	require.NoError(t, err)

	require.Equal(t, "1.0.0-rc.1", testData.Version)

	// alice, err := decodeKey(testData.Principals["alice"])
	// require.NoError(t, err)
	// bob, err := decodeKey(testData.Principals["bob"])
	// require.NoError(t, err)
	// carol, err := decodeKey(testData.Principals["carol"])
	// require.NoError(t, err)

	t.Run("valid", func(t *testing.T) {
		for _, tc := range testData.Valid {
			t.Run(tc.Name, func(t *testing.T) {
				dlgBytes, err := base64.StdEncoding.DecodeString(tc.Token)
				require.NoError(t, err)

				dlg, c, err := FromSealed(dlgBytes)
				require.NoError(t, err)
				require.Equal(t, tc.CID, c.String())

				require.Equal(t, tc.Envelope.Payload.Issuer, dlg.Issuer().String())
				require.Equal(t, tc.Envelope.Payload.Audience, dlg.Audience().String())
				require.Equal(t, tc.Envelope.Payload.Subject, dlg.Subject().String())
				require.Equal(t, tc.Envelope.Payload.Command, dlg.Command().String())
				require.Equal(t, tc.Envelope.Payload.Command, dlg.Command().String())
				require.JSONEq(t, string(tc.Envelope.Payload.Policies), dlg.Policy().String())
				require.Equal(t, tc.Envelope.Payload.ExpiresAt, dlg.expiration.Unix())
				nonceBytes, err := base64.StdEncoding.DecodeString(tc.Envelope.Payload.Nonce)
				require.NoError(t, err)
				require.Equal(t, nonceBytes, dlg.Nonce())
			})
		}
	})
}

func decodeKey(key string) (crypto.PrivateKeySigningBytes, error) {
	bytes, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}

	code, read, err := varint.FromUvarint(bytes)
	if err != nil {
		return nil, err
	}
	if code != 0x1300 {
		return nil, fmt.Errorf("invalid varint code: %d", code)
	}

	return ed25519.PrivateKeyFromSeed(bytes[read:])
}
