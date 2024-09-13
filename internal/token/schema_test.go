package token_test

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"github.com/stretchr/testify/require"
	"github.com/ucan-wg/go-ucan/internal/token"
)

//go:embed token.ipldsch
var schemaBytes []byte

func TestSchemaRoundTrip(t *testing.T) {
	const delegationJson = `
{
  "aud":"did:key:def456",
  "cmd":"/foo/bar",
  "exp":123456,
  "iss":"did:key:abc123",
  "meta":{
    "bar":"baaar",
    "foo":"fooo"
  },
  "nbf":123456,
  "nonce":{
    "/":{
      "bytes":"c3VwZXItcmFuZG9t"
    }
  },
  "pol":[
    ["==", ".status", "draft"],
    ["all", ".reviewer", [
		["like", ".email", "*@example.com"]]
	],
    ["any", ".tags", [ 
		["or", [
			["==", ".", "news"], 
			["==", ".", "press"]]
      ]]
	]
  ],
  "sub":""
}
`
	// format:    dagJson --> IPLD node --> token --> dagCbor --> IPLD node   -->   dagJson
	// function:                     Unwrap()   Wrap()

	n1, err := ipld.DecodeUsingPrototype([]byte(delegationJson), dagjson.Decode, token.Prototype())
	require.NoError(t, err)

	cborBytes, err := ipld.Encode(n1, dagcbor.Encode)
	require.NoError(t, err)
	fmt.Println("cborBytes length", len(cborBytes))
	fmt.Println("cbor", string(cborBytes))

	n2, err := ipld.DecodeUsingPrototype(cborBytes, dagcbor.Decode, token.Prototype())
	require.NoError(t, err)
	fmt.Println("read Cbor", n2)

	t1, err := token.Unwrap(n2)
	require.NoError(t, err)

	n3 := t1.Wrap()

	readJson, err := ipld.Encode(n3, dagjson.Encode)
	require.NoError(t, err)
	fmt.Println("readJson length", len(readJson))
	fmt.Println("json: ", string(readJson))

	require.JSONEq(t, delegationJson, string(readJson))
}

func BenchmarkSchemaLoad(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, _ = ipld.LoadSchemaBytes(schemaBytes)
	}
}
