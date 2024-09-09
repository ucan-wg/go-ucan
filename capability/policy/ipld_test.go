package policy

import (
	"testing"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"github.com/stretchr/testify/require"
)

func TestIpldRoundTrip(t *testing.T) {
	const illustrativeExample = `
[
  ["==", ".status", "draft"],
  ["all", ".reviewer", ["like", ".email", "*@example.com"]],
  ["any", ".tags", 
    ["or", [
      ["==", ".", "news"], 
      ["==", ".", "press"]]]
]`

	for _, tc := range []struct {
		name, dagJsonStr string
	}{
		{"illustrativeExample", illustrativeExample},
	} {
		nodes, err := ipld.Decode([]byte(tc.dagJsonStr), dagjson.Decode)
		require.NoError(t, err)

		pol, err := FromIPLD(nodes)
		require.NoError(t, err)

		wroteIpld, err := pol.ToIPLD()
		require.NoError(t, err)

		wroteAsDagJson, err := ipld.Encode(wroteIpld, dagjson.Encode)
		require.NoError(t, err)

		require.JSONEq(t, tc.dagJsonStr, string(wroteAsDagJson))
	}
}
