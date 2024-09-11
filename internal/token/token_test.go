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

	tkn, err := token.New()
	require.NoError(t, err)

	node := tkn.Wrap()

	json, err := ipld.Encode(node, dagjson.Encode)
	require.NoError(t, err)

	t.Log(string(json))

	t.Fail()
}
