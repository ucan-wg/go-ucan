package invocation_test

import (
	"bytes"
	"encoding/base64"
	"testing"

	"github.com/MetaMask/go-did-it/crypto"
	"github.com/MetaMask/go-did-it/crypto/ed25519"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ucan-wg/go-ucan/token/internal/envelope"
	"github.com/ucan-wg/go-ucan/token/invocation"
)

const (
	issuerPrivKeyCfg = "BeAgktAj8irGgWjp4PGk/fV67e5CcML/KRmmHSldco3etP5lRiuYQ+VVO/39ol3XXruJC8deSuBxoEXzgdYpYw=="
	newCID           = "zdpuB1NjhETofEUp5iYzoHjSc2KKgZvSoT6FBaLMoVzzsxiR1"
)

func TestSchemaRoundTrip(t *testing.T) {
	t.Parallel()

	privKey := privKey(t, issuerPrivKeyCfg)

	t.Run("via buffers", func(t *testing.T) {
		t.Parallel()

		// format:    dagJson   -->   PayloadModel   -->   dagCbor   -->   PayloadModel   -->   dagJson
		// function:      DecodeDagJson()           Seal()        Unseal()          EncodeDagJson()

		p1, err := invocation.FromDagJson(newDagJson)
		require.NoError(t, err)

		cborBytes, id, err := p1.ToSealed(privKey)
		require.NoError(t, err)
		assert.Equal(t, newCID, envelope.CIDToBase58BTC(id))

		p2, c2, err := invocation.FromSealed(cborBytes)
		require.NoError(t, err)
		assert.Equal(t, id, c2)

		readJson, err := p2.ToDagJson(privKey)
		require.NoError(t, err)

		assert.JSONEq(t, string(newDagJson), string(readJson))
	})

	t.Run("via streaming", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(newDagJson)

		// format:    dagJson   -->   PayloadModel   -->   dagCbor   -->   PayloadModel   -->   dagJson
		// function:      DecodeDagJson()           Seal()         Unseal()         EncodeDagJson()

		p1, err := invocation.FromDagJsonReader(buf)
		require.NoError(t, err)

		cborBytes := &bytes.Buffer{}
		id, err := p1.ToSealedWriter(cborBytes, privKey)
		require.NoError(t, err)
		assert.Equal(t, newCID, envelope.CIDToBase58BTC(id))

		p2, c2, err := invocation.FromSealedReader(cborBytes)
		require.NoError(t, err)
		assert.Equal(t, envelope.CIDToBase58BTC(id), envelope.CIDToBase58BTC(c2))

		readJson := &bytes.Buffer{}
		require.NoError(t, p2.ToDagJsonWriter(readJson, privKey))

		assert.JSONEq(t, string(newDagJson), readJson.String())
	})
}

func privKey(t require.TestingT, privKeyCfg string) crypto.PrivateKeySigningBytes {
	privBytes, err := base64.StdEncoding.DecodeString(privKeyCfg)
	require.NoError(t, err)

	privKey, err := ed25519.PrivateKeyFromBytes(privBytes)
	require.NoError(t, err)

	return privKey
}
