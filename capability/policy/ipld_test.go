package policy

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"github.com/stretchr/testify/require"
)

func TestIpldRoundTrip(t *testing.T) {
	const illustrativeExample = `[
    ["==", ".status", "draft"],
    ["all", ".reviewer", [["like", ".email", "*@example.com"]]],
    ["any", ".tags", 
      ["or", [
        ["==", ".", "news"], 
        ["==", ".", "press"]]
      ]]
]`

	for _, tc := range []struct {
		name, dagjson string
	}{
		{"illustrativeExample", illustrativeExample},
	} {
		// strip all spaces and carriage return
		asDagJson := strings.Join(strings.Fields(tc.dagjson), "")

		nodes, err := ipld.Decode([]byte(asDagJson), dagjson.Decode)
		require.NoError(t, err)

		pol, err := PolicyFromIPLD(nodes)
		require.NoError(t, err)

		fmt.Println(pol)

		wroteIpld, err := pol.ToIPLD()
		require.NoError(t, err)

		wroteAsDagJson, err := ipld.Encode(wroteIpld, dagjson.Encode)
		require.NoError(t, err)

		require.Equal(t, asDagJson, string(wroteAsDagJson))
	}
}
