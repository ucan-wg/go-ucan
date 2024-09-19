package delegation_test

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/ipld/go-ipld-prime"
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

	// format:    dagJson   -->   PayloadModel   -->   dagCbor   -->   PayloadModel   -->   dagJson
	// function:      DecodeDagJson()      EncodeDagCbor()   DecodeDagCbor()     EncodeDagJson()

	p1, id1, err := delegation.FromDagJson([]byte(delegationJson))
	require.NoError(t, err)
	require.Equal(t, newCID, envelope.CIDToBase58BTC(id1))

	cborBytes, err := p1.ToDagCbor(privKey)
	require.NoError(t, err)
	fmt.Println("cborBytes length", len(cborBytes))
	fmt.Println("cbor", string(cborBytes))

	p2, id2, err := delegation.FromDagCbor(cborBytes)
	require.NoError(t, err)
	fmt.Println("read Cbor", p2)

	readJson, err := p2.ToDagJson(privKey)
	require.NoError(t, err)
	require.Equal(t, newCID, envelope.CIDToBase58BTC(id2))
	fmt.Println("readJson length", len(readJson))
	fmt.Println("json: ", string(readJson))

	require.JSONEq(t, string(delegationJson), string(readJson))
}

func BenchmarkSchemaLoad(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = ipld.LoadSchemaBytes(schemaBytes)
	}
}
