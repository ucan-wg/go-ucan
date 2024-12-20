package selector_test

import (
	"strings"
	"testing"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	basicnode "github.com/ipld/go-ipld-prime/node/basic"
	"github.com/stretchr/testify/require"

	"github.com/ucan-wg/go-ucan/pkg/policy/selector"
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

	// Pass and return a node
	for _, testcase := range []Testcase{
		{Name: "Identity", Selector: `.`, Input: `{"x":1}`, Output: `{"x":1}`},
		{Name: "Iterator", Selector: `.[]`, Input: `[1, 2]`, Output: `[1, 2]`},
		{Name: "Optional Null Iterator", Selector: `.[]?`, Input: `null`, Output: `[]`},
		{Name: "Object Key", Selector: `.x`, Input: `{"x": 1 }`, Output: `1`},
		{Name: "Quoted Key", Selector: `.["x"]`, Input: `{"x": 1}`, Output: `1`},
		{Name: "Index", Selector: `.[0]`, Input: `[1, 2]`, Output: `1`},
		{Name: "Negative Index", Selector: `.[-1]`, Input: `[1, 2]`, Output: `2`},
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
			res, err := sel.Select(makeNode(t, tc.Input))
			require.NoError(t, err)
			require.NotNil(t, res)

			exp := makeNode(t, tc.Output)
			require.True(t, ipld.DeepEqual(exp, res))
		})
	}

	// No error and return null, as optional
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
			res, err := sel.Select(makeNode(t, tc.Input))
			require.NoError(t, err)
			require.Nil(t, res)
		})
	}

	// fail to select and return an error
	for _, testcase := range []Testcase{
		{Name: "Null Iterator", Selector: `.[]`, Input: `null`},
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
			res, err := sel.Select(makeNode(t, tc.Input))
			require.Error(t, err)
			require.Nil(t, res)
		})
	}
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
