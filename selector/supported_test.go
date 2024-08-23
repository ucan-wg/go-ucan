package selector_test

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"github.com/ipld/go-ipld-prime/datamodel"
	basicnode "github.com/ipld/go-ipld-prime/node/basic"
	"github.com/storacha-network/go-ucanto/core/policy/selector"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wI2L/jsondiff"
)

//go:embed supported.json
var supported []byte

type Testcase struct {
	Name     string `json:"name"`
	Selector string `json:"selector"`
	Input    string `json:"input"`
}

func (tc Testcase) Select(t *testing.T) (datamodel.Node, []datamodel.Node, error) {
	t.Helper()

	sel, err := selector.Parse(tc.Selector)
	require.NoError(t, err)

	return selector.Select(sel, node(t, tc.Input))
}

type SuccessTestcase struct {
	Testcase
	Output *string `json:"output"`
}

func (tc SuccessTestcase) SelectAndCompare(t *testing.T) {
	t.Helper()

	exp := node(t, *tc.Output)

	node, nodes, err := tc.Select(t)
	require.NoError(t, err)
	require.NotEqual(t, node != nil, len(nodes) > 0) // XOR (only one of node or nodes should be set)

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

	equalIPLD(t, exp, node)
}

type Testcases struct {
	SuccessTestcases []SuccessTestcase `json:"pass"`
	NullTestcases    []Testcase        `json:"null"`
	ErrorTestcases   []Testcase        `json:"fail"`
}

// TestSupported Forms runs tests against the Selector according to the
// proposed "Supported Forms" presented in this GitHub issue:
// https://github.com/ucan-wg/delegation/issues/5#issue-2154766496
func TestSupportedForms(t *testing.T) {
	t.Parallel()

	var testcases Testcases

	require.NoError(t, json.Unmarshal(supported, &testcases))

	t.Run("node(s)", func(t *testing.T) {
		t.Parallel()

		for _, testcase := range testcases.SuccessTestcases {
			testcase := testcase

			t.Run(testcase.Name, func(t *testing.T) {
				t.Parallel()

				// TODO: This test case panics during Select, though Parse works - reports
				// "index out of range [-1]" so a bit of subtraction and some bounds checking
				// should fix this testcase.
				if testcase.Name == "Negative Index" {
					t.Skip()
				}

				testcase.SelectAndCompare(t)
			})
		}
	})

	t.Run("null", func(t *testing.T) {
		t.Parallel()

		for _, testcase := range testcases.NullTestcases {
			testcase := testcase

			t.Run(testcase.Name, func(t *testing.T) {
				t.Parallel()

				node, nodes, err := testcase.Select(t)
				require.NoError(t, err)
				// TODO: should Select return a single node which is sometimes a list or null?
				// require.Equal(t, datamodel.Null, node)
				assert.Nil(t, node)
				assert.Empty(t, nodes)
			})
		}
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()

		for _, testcase := range testcases.ErrorTestcases {
			testcase := testcase

			t.Run(testcase.Name, func(t *testing.T) {
				t.Parallel()

				node, nodes, err := testcase.Select(t)
				require.Error(t, err)
				assert.Nil(t, node)
				assert.Empty(t, nodes)
			})
		}
	})
}

func equalIPLD(t *testing.T, expected datamodel.Node, actual datamodel.Node, msgAndArgs ...interface{}) bool {
	t.Helper()

	if !assert.ObjectsAreEqual(expected, actual) {
		exp, act := &bytes.Buffer{}, &bytes.Buffer{}
		if err := dagjson.Encode(expected, exp); err != nil {
			return assert.Fail(t, "Failed to encode json for expected IPLD node")
		}

		if err := dagjson.Encode(actual, act); err != nil {
			return assert.Fail(t, "Failed to encode JSON for actual IPLD node")
		}

		diff, err := jsondiff.CompareJSON(act.Bytes(), exp.Bytes())
		if err != nil {
			return assert.Fail(t, "Failed to create diff of expected and actual IPLD nodes")
		}

		return assert.Fail(t, fmt.Sprintf("Not equal: \n"+
			"expected: %s\n"+
			"actual: %s\n"+
			"diff: %s", exp, act, diff), msgAndArgs)
	}

	return true
}

func node(t *testing.T, json string) ipld.Node {
	t.Helper()

	np := basicnode.Prototype.Any
	nb := np.NewBuilder()
	require.NoError(t, dagjson.Decode(nb, strings.NewReader(json)))

	node := nb.Build()
	require.NotNil(t, node)

	return node
}
