package token_test

import (
	"testing"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"github.com/stretchr/testify/require"
	"github.com/ucan-wg/go-ucan/internal/token"
)

func TestEncode(t *testing.T) {
	t.Parallel()

	tkn := &token.Token{}

	node := tkn.Wrap()

	json, err := ipld.Encode(node, dagjson.Encode)
	require.NoError(t, err)

	t.Log(string(json))

	t.Fail()
}

func TestPrototype(t *testing.T) {
	t.Parallel()

	tkn := &token.Token{
		Issuer: "blah",
	}
	n1 := tkn.Wrap()
	json, err := ipld.Encode(n1, dagjson.Encode)
	require.NoError(t, err)

	t.Log(string(json))

	n2, err := ipld.Decode(json, dagjson.Decode)
	require.NoError(t, err)

	nb := token.Prototype().Representation().NewBuilder()
	require.NoError(t, nb.AssignNode(n2))

	n3 := nb.Build()

	tkn2, err := token.Unwrap(n3)
	require.NoError(t, err)

	t.Log(tkn2)

	require.Equal(t, tkn, tkn2)

	t.Fail()
}
