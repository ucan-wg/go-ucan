package selector_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"github.com/ipld/go-ipld-prime/datamodel"
	basicnode "github.com/ipld/go-ipld-prime/node/basic"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ucan-wg/go-ucan/capability/policy/selector"
)

// TestSupported Forms runs tests against the Selector according to the
// proposed "Supported Forms" presented in this GitHub issue:
// https://github.com/ucan-wg/delegation/issues/5#issue-2154766496
func TestSupportedForms(t *testing.T) {
	type Testcase struct {
		Name     string
		Selector string
		Input    string
		Output   string
	}

	// Pass
	for _, testcase := range []Testcase{
		{Name: "Identity", Selector: `.`, Input: `{"x":1}`, Output: `{"x":1}`},
		{Name: "Iterator", Selector: `.[]`, Input: `[1, 2]`, Output: `[1, 2]`},
		{Name: "Optional Null Iterator", Selector: `.[]?`, Input: `null`, Output: `()`},
		{Name: "Optional Iterator", Selector: `.[][]?`, Input: `[[1], 2, [3]]`, Output: `[1, 3]`},
		{Name: "Object Key", Selector: `.x`, Input: `{"x": 1 }`, Output: `1`},
		{Name: "Quoted Key", Selector: `.["x"]`, Input: `{"x": 1}`, Output: `1`},
		{Name: "Index", Selector: `.[0]`, Input: `[1, 2]`, Output: `1`},
		{Name: "Negative Index", Selector: `.[-1]`, Input: `[1, 2]`, Output: `2`},
		{Name: "String Index", Selector: `.[0]`, Input: `"Hi"`, Output: `"H"`},
		{Name: "Bytes Index", Selector: `.[0]`, Input: `{"/":{"bytes":"AAE"}}`, Output: `0`},
		{Name: "Array Slice", Selector: `.[0:2]`, Input: `[0, 1, 2]`, Output: `[0, 1]`},
		{Name: "Array Slice", Selector: `.[1:]`, Input: `[0, 1, 2]`, Output: `[1, 2]`},
		{Name: "Array Slice", Selector: `.[:2]`, Input: `[0, 1, 2]`, Output: `[0, 1]`},
		{Name: "String Slice", Selector: `.[0:2]`, Input: `"hello"`, Output: `"he"`},
		{Name: "Bytes Index", Selector: `.[1:]`, Input: `{"/":{"bytes":"AAEC"}}`, Output: `{"/":{"bytes":"AQI"}}`},
	} {
		tc := testcase
		t.Run("Pass: "+tc.Name, func(t *testing.T) {
			t.Parallel()

			sel, err := selector.Parse(tc.Selector)
			require.NoError(t, err)

			// attempt to select
			node, nodes, err := selector.Select(sel, makeNode(t, tc.Input))
			require.NoError(t, err)
			require.NotEqual(t, node != nil, len(nodes) > 0) // XOR (only one of node or nodes should be set)

			// make an IPLD List node from a []datamodel.Node
			if node == nil {
				nb := basicnode.Prototype.List.NewBuilder()
				la, err := nb.BeginList(int64(len(nodes)))
				require.NoError(t, err)

				for _, n := range nodes {
					// TODO: This code is probably not needed if the Select operation properly prunes nil values - e.g.: Optional Iterator
					if n == nil {
						n = datamodel.Null
					}

					require.NoError(t, la.AssembleValue().AssignNode(n))
				}
				require.NoError(t, la.Finish())

				node = nb.Build()
			}

			exp := makeNode(t, tc.Output)
			equalIPLD(t, exp, node)
		})
	}

	// null
	for _, testcase := range []Testcase{
		{Name: "Optional Missing Key", Selector: `.x?`, Input: `{}`},
		{Name: "Optional Null Key", Selector: `.x?`, Input: `null`},
		{Name: "Optional Array Key", Selector: `.x?`, Input: `[]`},
		{Name: "Optional Quoted Key", Selector: `.["x"]?`, Input: `{}`},
		{Name: ".length?", Selector: `.length?`, Input: `[1, 2]`},
		{Name: "Optional Index", Selector: `.[4]?`, Input: `[0, 1]`},
	} {
		tc := testcase
		t.Run("Null: "+tc.Name, func(t *testing.T) {
			t.Parallel()

			sel, err := selector.Parse(tc.Selector)
			require.NoError(t, err)

			// attempt to select
			node, nodes, err := selector.Select(sel, makeNode(t, tc.Input))
			require.NoError(t, err)
			// TODO: should Select return a single node which is sometimes a list or null?
			// require.Equal(t, datamodel.Null, node)
			assert.Nil(t, node)
			assert.Empty(t, nodes)
		})
	}

	// error
	for _, testcase := range []Testcase{
		{Name: "Null Iterator", Selector: `.[]`, Input: `null`},
		{Name: "Nested Iterator", Selector: `.[][]`, Input: `[[1], 2, [3]]`},
		{Name: "Missing Key", Selector: `.x`, Input: `{}`},
		{Name: "Null Key", Selector: `.x`, Input: `null`},
		{Name: "Array Key", Selector: `.x`, Input: `[]`},
		{Name: ".length", Selector: `.length`, Input: `[1, 2]`},
		{Name: "Out of bound Index", Selector: `.[4]`, Input: `[0, 1]`},
	} {
		tc := testcase
		t.Run("Null: "+tc.Name, func(t *testing.T) {
			t.Parallel()

			sel, err := selector.Parse(tc.Selector)
			require.NoError(t, err)

			// attempt to select
			node, nodes, err := selector.Select(sel, makeNode(t, tc.Input))
			require.Error(t, err)
			assert.Nil(t, node)
			assert.Empty(t, nodes)
		})
	}
}

func equalIPLD(t *testing.T, expected datamodel.Node, actual datamodel.Node) bool {
	t.Helper()

	exp, act := &bytes.Buffer{}, &bytes.Buffer{}
	if err := dagjson.Encode(expected, exp); err != nil {
		return assert.Fail(t, "Failed to encode json for expected IPLD node")
	}

	if err := dagjson.Encode(actual, act); err != nil {
		return assert.Fail(t, "Failed to encode JSON for actual IPLD node")
	}

	require.JSONEq(t, exp.String(), act.String())

	return true
}

func makeNode(t *testing.T, dagJsonInput string) ipld.Node {
	t.Helper()

	np := basicnode.Prototype.Any
	nb := np.NewBuilder()
	require.NoError(t, dagjson.Decode(nb, strings.NewReader(dagJsonInput)))

	node := nb.Build()
	require.NotNil(t, node)

	return node
}
