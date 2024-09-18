package delegation_test

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/ipld/go-ipld-prime"
	"github.com/stretchr/testify/require"
	"github.com/ucan-wg/go-ucan/delegation"
	"gotest.tools/v3/golden"
)

//go:embed delegation.ipldsch
var schemaBytes []byte

func TestSchemaRoundTrip(t *testing.T) {
	// 	const delegationJson = `
	// {
	//   "aud":"did:key:def456",
	//   "cmd":"/foo/bar",
	//   "exp":123456,
	//   "iss":"did:key:abc123",
	//   "meta":{
	//     "bar":"baaar",
	//     "foo":"fooo"
	//   },
	//   "nbf":123456,
	//   "nonce":{
	//     "/":{
	//       "bytes":"c3VwZXItcmFuZG9t"
	//     }
	//   },
	//   "pol":[
	//     ["==", ".status", "draft"],
	//     ["all", ".reviewer", [
	// 		["like", ".email", "*@example.com"]]
	// 	],
	//     ["any", ".tags", [
	// 		["or", [
	// 			["==", ".", "news"],
	// 			["==", ".", "press"]]
	//       ]]
	// 	]
	//   ],
	//   "sub":""
	// }
	// `

	delegationJson := golden.Get(t, "new.dagjson")
	privKey := privKey(t, issuerPrivKeyCfg)

	// format:    dagJson   -->   PayloadModel   -->   dagCbor   -->   PayloadModel   -->   dagJson
	// function:      DecodeDagJson()      EncodeDagCbor()   DecodeDagCbor()     EncodeDagJson()

	p1, err := delegation.FromDagJson([]byte(delegationJson))
	require.NoError(t, err)

	cborBytes, err := p1.ToDagCbor(privKey)
	require.NoError(t, err)
	fmt.Println("cborBytes length", len(cborBytes))
	fmt.Println("cbor", string(cborBytes))

	p2, err := delegation.FromDagCbor(cborBytes)
	require.NoError(t, err)
	fmt.Println("read Cbor", p2)

	readJson, err := p2.ToDagJson(privKey)
	require.NoError(t, err)
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
