package invocation_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gotest.tools/v3/golden"

	"github.com/ucan-wg/go-ucan/token/internal/envelope"
	"github.com/ucan-wg/go-ucan/token/invocation"
)

const (
	issuerPrivKeyCfg = "CAESQK45xBfqIxRp7ZdRdck3tIJZKocCqvANQc925dCJhFwO7DJNA2j94zkF0TNx5mpXV0s6utfkFdHddWTaPVU6yZc="
	newCID           = "zdpuAqY6Zypg4UnpbSUgDvYGneyFaTKaZevzxgSxV4rmv3Fpp"
)

func TestSchemaRoundTrip(t *testing.T) {
	t.Parallel()

	invocationJson := golden.Get(t, "new.dagjson")
	privKey := privKey(t, issuerPrivKeyCfg)

	t.Run("via buffers", func(t *testing.T) {
		t.Parallel()

		// format:    dagJson   -->   PayloadModel   -->   dagCbor   -->   PayloadModel   -->   dagJson
		// function:      DecodeDagJson()           Seal()        Unseal()          EncodeDagJson()

		p1, err := invocation.FromDagJson(invocationJson)
		require.NoError(t, err)

		cborBytes, id, err := p1.ToSealed(privKey)
		require.NoError(t, err)
		assert.Equal(t, newCID, envelope.CIDToBase58BTC(id))
		fmt.Println("cborBytes length", len(cborBytes))
		fmt.Println("cbor", string(cborBytes))

		p2, c2, err := invocation.FromSealed(cborBytes)
		require.NoError(t, err)
		assert.Equal(t, id, c2)
		fmt.Println("read Cbor", p2)

		readJson, err := p2.ToDagJson(privKey)
		require.NoError(t, err)
		fmt.Println("readJson length", len(readJson))
		fmt.Println("json: ", string(readJson))

		assert.JSONEq(t, string(invocationJson), string(readJson))
	})

	t.Run("via streaming", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(invocationJson)

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

		assert.JSONEq(t, string(invocationJson), readJson.String())
	})
}

func privKey(t require.TestingT, privKeyCfg string) crypto.PrivKey {
	privKeyMar, err := crypto.ConfigDecodeKey(privKeyCfg)
	require.NoError(t, err)

	privKey, err := crypto.UnmarshalPrivateKey(privKeyMar)
	require.NoError(t, err)

	return privKey
}
