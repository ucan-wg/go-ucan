package delegation

import (
	"fmt"
	"testing"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/node/bindnode"
	"github.com/stretchr/testify/require"
)

func TestSchemaRoundTrip(t *testing.T) {
	p := &PayloadModel{
		Iss:   "did:key:abc123",
		Aud:   "did:key:def456",
		Sub:   PointerTo(""),
		Cmd:   "/foo/bar",
		Pol:   PolicyModel{}, // TODO: have something here
		Nonce: []byte("super-random"),
		Meta: MetaModel{
			Keys: []string{"foo", "bar"},
			Values: map[string]datamodel.Node{
				"foo": bindnode.Wrap(PointerTo("fooo"), nil),
				"bar": bindnode.Wrap(PointerTo("baaar"), nil),
			},
		},
		Nbf: PointerTo(int64(123456)),
		Exp: PointerTo(int64(123456)),
	}

	cborBytes, err := p.EncodeDagCbor()
	require.NoError(t, err)
	fmt.Println("cborBytes length", len(cborBytes))
	fmt.Println("cbor", string(cborBytes))

	jsonBytes, err := p.EncodeDagJson()
	require.NoError(t, err)
	fmt.Println("jsonBytes length", len(jsonBytes))
	fmt.Println("json: ", string(jsonBytes))

	fmt.Println()

	readCbor, err := DecodeDagCbor(cborBytes)
	require.NoError(t, err)
	fmt.Println("readCbor", readCbor)
	require.Equal(t, p, readCbor)

	readJson, err := DecodeDagJson(jsonBytes)
	require.NoError(t, err)
	fmt.Println("readJson", readJson)
	require.Equal(t, p, readJson)
}

func BenchmarkSchemaLoad(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = ipld.LoadSchemaBytes(schemaBytes)
	}
}
