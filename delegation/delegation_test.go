package delegation_test

import (
	"testing"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"github.com/ipld/go-ipld-prime/node/bindnode"
	"github.com/ipld/go-ipld-prime/schema"
	"github.com/stretchr/testify/require"
	"github.com/ucan-wg/go-ucan/v1/delegation"
	"github.com/ucan-wg/go-ucan/v1/internal/token"
)

func TestToken_Proto(t *testing.T) {
	t.Parallel()

	const delegationJson = `
{
  "aud":"did:key:z6MkhaXgBZDvotDkL5257faiztiGiC2QtKLGpbnnEGta2doK",
  "cmd":"/foo/bar",
  "exp":123456,
  "iss":"did:key:z6MkhaXgBZDvotDkL5257faiztiGiC2QtKLGpbnnEGta2doK",
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
  "sub":"did:key:z6MkhaXgBZDvotDkL5257faiztiGiC2QtKLGpbnnEGta2doK"
}
`

	proto := (*delegation.Token)(nil).Prototype()

	node, err := ipld.DecodeUsingPrototype([]byte(delegationJson), dagjson.Decode, proto)
	require.NoError(t, err)

	tkn, ok := bindnode.Unwrap(node).(*delegation.Token)
	require.True(t, ok)

	t.Log("Token:")
	t.Log("  Audience:", tkn.Audience)
	t.Log("  Command: ", tkn.Command)
	// t.Log("  Expiration: ", token.Expiration)
	t.Log("  Issuer:", tkn.Issuer)
	// t.Log("  Meta:", token.Meta)
	// t.Log("  NotBefore", token.NotBefore)
	// t.Log("  Nonce:", token.Nonce)
	// t.Log("  Policy:", token.Policy)
	t.Log("  Subject:", tkn.Subject)

	// token.Command = nil
	// token.Meta = nil
	// token.Policy = nil
	// token.Expiration = nil
	// token.NotBefore = nil

	_ = bindnode.Wrap(tkn, proto.Type(), token.BindnodeOptions()...)

	typed, ok := node.(schema.TypedNode)
	require.True(t, ok)

	json, err := ipld.Encode(typed.Representation(), dagjson.Encode)
	require.NoError(t, err)

	require.JSONEq(t, delegationJson, string(json))

	t.Log(string(json))
	t.Fail()
}
