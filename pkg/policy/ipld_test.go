package policy

import (
	"fmt"
	"testing"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"github.com/stretchr/testify/require"

	"github.com/ucan-wg/go-ucan/pkg/policy/literal"
)

func TestIpldRoundTrip(t *testing.T) {
	const illustrativeExample = `
[
  ["==", ".status", "draft"],
  ["all", ".reviewer", ["like", ".email", "*@example.com"]],
  ["any", ".tags", 
    ["or", [
      ["==", ".", "news"], 
      ["==", ".", "press"]
    ]]
  ]
]`

	// must contain all the operators
	const allOps = `
[
  ["and", [
    ["==", ".foo1", ".bar1"],
    ["!=", ".foo2", ".bar2"]
  ]],
  ["or", [
    [">", ".foo5", 5.2],
    [">=", ".foo6", 6.2]
  ]],
  ["not", ["like", ".foo7", "*@example.com"]],
  ["all", ".foo8",
    ["<", ".foo3", 3]
  ],
  ["any", ".foo9",
    ["<=", ".foo4", 4]
  ]
]`

	for _, tc := range []struct {
		name, dagJsonStr string
	}{
		{"illustrativeExample", illustrativeExample},
		{"allOps", allOps},
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

func TestFoo(t *testing.T) {
	fmt.Println(MustConstruct(
		And(
			Equal(".foo1", literal.String(".bar1")),
			NotEqual(".foo2", literal.String(".bar2")),
		),
		Or(
			GreaterThan(".foo5", literal.Float(5.2)),
			GreaterThanOrEqual(".foo6", literal.Float(6.2)),
		),
		Not(Like(".foo7", "*@example.com")),
		All(".foo8", LessThan(".foo3", literal.Int(3))),
		Any(".foo9", LessThanOrEqual(".foo4", literal.Int(4))),
	))
}
