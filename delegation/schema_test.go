package delegation_test

import (
	"bytes"
	_ "embed"
	"fmt"
	"testing"

	"github.com/ipld/go-ipld-prime"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gotest.tools/v3/golden"

	"github.com/ucan-wg/go-ucan/delegation"
	"github.com/ucan-wg/go-ucan/internal/envelope"
)

//go:embed delegation.ipldsch
var schemaBytes []byte

func TestSchemaRoundTrip(t *testing.T) {
	t.Parallel()

	delegationJson := golden.Get(t, "new.dagjson")
	privKey := privKey(t, issuerPrivKeyCfg)

	t.Run("via buffers", func(t *testing.T) {
		t.Parallel()

		// format:    dagJson   -->   PayloadModel   -->   dagCbor   -->   PayloadModel   -->   dagJson
		// function:      DecodeDagJson()           Seal()        Unseal()          EncodeDagJson()

		p1, err := delegation.FromDagJson([]byte(delegationJson))
		require.NoError(t, err)

		cborBytes, id, err := p1.Seal(privKey)
		require.NoError(t, err)
		assert.Equal(t, newCID, envelope.CIDToBase58BTC(id))
		fmt.Println("cborBytes length", len(cborBytes))
		fmt.Println("cbor", string(cborBytes))

		p2, err := delegation.Unseal(cborBytes)
		require.NoError(t, err)
		assert.Equal(t, id, p2.CID())
		fmt.Println("read Cbor", p2)

		readJson, err := p2.ToDagJson(privKey)
		require.NoError(t, err)
		fmt.Println("readJson length", len(readJson))
		fmt.Println("json: ", string(readJson))

		assert.JSONEq(t, string(delegationJson), string(readJson))
	})

	t.Run("via streaming", func(t *testing.T) {
		t.Parallel()

		buf := bytes.NewBuffer(delegationJson)

		// format:    dagJson   -->   PayloadModel   -->   dagCbor   -->   PayloadModel   -->   dagJson
		// function:      DecodeDagJson()           Seal()         Unseal()         EncodeDagJson()

		p1, err := delegation.FromDagJsonReader(buf)
		require.NoError(t, err)

		cborBytes := &bytes.Buffer{}
		id, err := p1.SealWriter(cborBytes, privKey)
		t.Log(len(id.Bytes()), id.Bytes())
		require.NoError(t, err)
		assert.Equal(t, newCID, envelope.CIDToBase58BTC(id))

		// buf = bytes.NewBuffer(cborBytes.Bytes())
		p2, err := delegation.UnsealReader(cborBytes)
		require.NoError(t, err)
		t.Log(len(p2.CID().Bytes()), p2.CID().Bytes())
		assert.Equal(t, envelope.CIDToBase58BTC(id), envelope.CIDToBase58BTC(p2.CID()))

		readJson := &bytes.Buffer{}
		require.NoError(t, p2.ToDagJsonWriter(readJson, privKey))

		assert.JSONEq(t, string(delegationJson), readJson.String())
	})
}

func BenchmarkSchemaLoad(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = ipld.LoadSchemaBytes(schemaBytes)
	}
}
