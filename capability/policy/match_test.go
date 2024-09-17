package policy

import (
	"fmt"
	"testing"

	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	"github.com/stretchr/testify/require"

	"github.com/ucan-wg/go-ucan/capability/policy/literal"
	"github.com/ucan-wg/go-ucan/capability/policy/selector"
)

func TestMatch(t *testing.T) {
	t.Run("equality", func(t *testing.T) {
		t.Run("string", func(t *testing.T) {
			np := basicnode.Prototype.String
			nb := np.NewBuilder()
			nb.AssignString("test")
			nd := nb.Build()

			pol := Policy{Equal(selector.MustParse("."), literal.String("test"))}
			ok := Match(pol, nd)
			require.True(t, ok)

			pol = Policy{Equal(selector.MustParse("."), literal.String("test2"))}
			ok = Match(pol, nd)
			require.False(t, ok)

			pol = Policy{Equal(selector.MustParse("."), literal.Int(138))}
			ok = Match(pol, nd)
			require.False(t, ok)
		})

		t.Run("int", func(t *testing.T) {
			np := basicnode.Prototype.Int
			nb := np.NewBuilder()
			nb.AssignInt(138)
			nd := nb.Build()

			pol := Policy{Equal(selector.MustParse("."), literal.Int(138))}
			ok := Match(pol, nd)
			require.True(t, ok)

			pol = Policy{Equal(selector.MustParse("."), literal.Int(1138))}
			ok = Match(pol, nd)
			require.False(t, ok)

			pol = Policy{Equal(selector.MustParse("."), literal.String("138"))}
			ok = Match(pol, nd)
			require.False(t, ok)
		})

		t.Run("float", func(t *testing.T) {
			np := basicnode.Prototype.Float
			nb := np.NewBuilder()
			nb.AssignFloat(1.138)
			nd := nb.Build()

			pol := Policy{Equal(selector.MustParse("."), literal.Float(1.138))}
			ok := Match(pol, nd)
			require.True(t, ok)

			pol = Policy{Equal(selector.MustParse("."), literal.Float(11.38))}
			ok = Match(pol, nd)
			require.False(t, ok)

			pol = Policy{Equal(selector.MustParse("."), literal.String("138"))}
			ok = Match(pol, nd)
			require.False(t, ok)
		})

		t.Run("IPLD Link", func(t *testing.T) {
			l0 := cidlink.Link{Cid: cid.MustParse("bafybeif4owy5gno5lwnixqm52rwqfodklf76hsetxdhffuxnplvijskzqq")}
			l1 := cidlink.Link{Cid: cid.MustParse("bafkreifau35r7vi37tvbvfy3hdwvgb4tlflqf7zcdzeujqcjk3rsphiwte")}

			np := basicnode.Prototype.Link
			nb := np.NewBuilder()
			nb.AssignLink(l0)
			nd := nb.Build()

			pol := Policy{Equal(selector.MustParse("."), literal.Link(l0))}
			ok := Match(pol, nd)
			require.True(t, ok)

			pol = Policy{Equal(selector.MustParse("."), literal.Link(l1))}
			ok = Match(pol, nd)
			require.False(t, ok)

			pol = Policy{Equal(selector.MustParse("."), literal.String("bafybeif4owy5gno5lwnixqm52rwqfodklf76hsetxdhffuxnplvijskzqq"))}
			ok = Match(pol, nd)
			require.False(t, ok)
		})

		t.Run("string in map", func(t *testing.T) {
			np := basicnode.Prototype.Map
			nb := np.NewBuilder()
			ma, _ := nb.BeginMap(1)
			ma.AssembleKey().AssignString("foo")
			ma.AssembleValue().AssignString("bar")
			ma.Finish()
			nd := nb.Build()

			pol := Policy{Equal(selector.MustParse(".foo"), literal.String("bar"))}
			ok := Match(pol, nd)
			require.True(t, ok)

			pol = Policy{Equal(selector.MustParse(".[\"foo\"]"), literal.String("bar"))}
			ok = Match(pol, nd)
			require.True(t, ok)

			pol = Policy{Equal(selector.MustParse(".foo"), literal.String("baz"))}
			ok = Match(pol, nd)
			require.False(t, ok)

			pol = Policy{Equal(selector.MustParse(".foobar"), literal.String("bar"))}
			ok = Match(pol, nd)
			require.False(t, ok)
		})

		t.Run("string in list", func(t *testing.T) {
			np := basicnode.Prototype.List
			nb := np.NewBuilder()
			la, _ := nb.BeginList(1)
			la.AssembleValue().AssignString("foo")
			la.Finish()
			nd := nb.Build()

			pol := Policy{Equal(selector.MustParse(".[0]"), literal.String("foo"))}
			ok := Match(pol, nd)
			require.True(t, ok)

			pol = Policy{Equal(selector.MustParse(".[1]"), literal.String("foo"))}
			ok = Match(pol, nd)
			require.False(t, ok)
		})
	})

	t.Run("inequality", func(t *testing.T) {
		t.Run("gt int", func(t *testing.T) {
			np := basicnode.Prototype.Int
			nb := np.NewBuilder()
			nb.AssignInt(138)
			nd := nb.Build()

			pol := Policy{GreaterThan(selector.MustParse("."), literal.Int(1))}
			ok := Match(pol, nd)
			require.True(t, ok)
		})

		t.Run("gte int", func(t *testing.T) {
			np := basicnode.Prototype.Int
			nb := np.NewBuilder()
			nb.AssignInt(138)
			nd := nb.Build()

			pol := Policy{GreaterThanOrEqual(selector.MustParse("."), literal.Int(1))}
			ok := Match(pol, nd)
			require.True(t, ok)

			pol = Policy{GreaterThanOrEqual(selector.MustParse("."), literal.Int(138))}
			ok = Match(pol, nd)
			require.True(t, ok)
		})

		t.Run("gt float", func(t *testing.T) {
			np := basicnode.Prototype.Float
			nb := np.NewBuilder()
			nb.AssignFloat(1.38)
			nd := nb.Build()

			pol := Policy{GreaterThan(selector.MustParse("."), literal.Float(1))}
			ok := Match(pol, nd)
			require.True(t, ok)
		})

		t.Run("gte float", func(t *testing.T) {
			np := basicnode.Prototype.Float
			nb := np.NewBuilder()
			nb.AssignFloat(1.38)
			nd := nb.Build()

			pol := Policy{GreaterThanOrEqual(selector.MustParse("."), literal.Float(1))}
			ok := Match(pol, nd)
			require.True(t, ok)

			pol = Policy{GreaterThanOrEqual(selector.MustParse("."), literal.Float(1.38))}
			ok = Match(pol, nd)
			require.True(t, ok)
		})

		t.Run("lt int", func(t *testing.T) {
			np := basicnode.Prototype.Int
			nb := np.NewBuilder()
			nb.AssignInt(138)
			nd := nb.Build()

			pol := Policy{LessThan(selector.MustParse("."), literal.Int(1138))}
			ok := Match(pol, nd)
			require.True(t, ok)
		})

		t.Run("lte int", func(t *testing.T) {
			np := basicnode.Prototype.Int
			nb := np.NewBuilder()
			nb.AssignInt(138)
			nd := nb.Build()

			pol := Policy{LessThanOrEqual(selector.MustParse("."), literal.Int(1138))}
			ok := Match(pol, nd)
			require.True(t, ok)

			pol = Policy{LessThanOrEqual(selector.MustParse("."), literal.Int(138))}
			ok = Match(pol, nd)
			require.True(t, ok)
		})
	})

	t.Run("negation", func(t *testing.T) {
		np := basicnode.Prototype.Bool
		nb := np.NewBuilder()
		nb.AssignBool(false)
		nd := nb.Build()

		pol := Policy{Not(Equal(selector.MustParse("."), literal.Bool(true)))}
		ok := Match(pol, nd)
		require.True(t, ok)

		pol = Policy{Not(Equal(selector.MustParse("."), literal.Bool(false)))}
		ok = Match(pol, nd)
		require.False(t, ok)
	})

	t.Run("conjunction", func(t *testing.T) {
		np := basicnode.Prototype.Int
		nb := np.NewBuilder()
		nb.AssignInt(138)
		nd := nb.Build()

		pol := Policy{
			And(
				GreaterThan(selector.MustParse("."), literal.Int(1)),
				LessThan(selector.MustParse("."), literal.Int(1138)),
			),
		}
		ok := Match(pol, nd)
		require.True(t, ok)

		pol = Policy{
			And(
				GreaterThan(selector.MustParse("."), literal.Int(1)),
				Equal(selector.MustParse("."), literal.Int(1138)),
			),
		}
		ok = Match(pol, nd)
		require.False(t, ok)

		pol = Policy{And()}
		ok = Match(pol, nd)
		require.True(t, ok)
	})

	t.Run("disjunction", func(t *testing.T) {
		np := basicnode.Prototype.Int
		nb := np.NewBuilder()
		nb.AssignInt(138)
		nd := nb.Build()

		pol := Policy{
			Or(
				GreaterThan(selector.MustParse("."), literal.Int(138)),
				LessThan(selector.MustParse("."), literal.Int(1138)),
			),
		}
		ok := Match(pol, nd)
		require.True(t, ok)

		pol = Policy{
			Or(
				GreaterThan(selector.MustParse("."), literal.Int(138)),
				Equal(selector.MustParse("."), literal.Int(1138)),
			),
		}
		ok = Match(pol, nd)
		require.False(t, ok)

		pol = Policy{Or()}
		ok = Match(pol, nd)
		require.True(t, ok)
	})

	t.Run("wildcard", func(t *testing.T) {
		pattern := `Alice\*, Bob*, Carol.`

		for _, s := range []string{
			"Alice*, Bob, Carol.",
			"Alice*, Bob, Dan, Erin, Carol.",
			"Alice*, Bob  , Carol.",
			"Alice*, Bob*, Carol.",
		} {
			func(s string) {
				t.Run(fmt.Sprintf("pass %s", s), func(t *testing.T) {
					np := basicnode.Prototype.String
					nb := np.NewBuilder()
					nb.AssignString(s)
					nd := nb.Build()

					statement, err := Like(selector.MustParse("."), pattern)
					require.NoError(t, err)

					pol := Policy{statement}
					ok := Match(pol, nd)
					require.True(t, ok)
				})
			}(s)
		}

		for _, s := range []string{
			"Alice*, Bob, Carol",
			"Alice*, Bob*, Carol!",
			"Alice Cooper, Bob, Carol.",
			"Alice, Bob, Carol.",
			" Alice*, Bob, Carol. ",
		} {
			func(s string) {
				t.Run(fmt.Sprintf("fail %s", s), func(t *testing.T) {
					np := basicnode.Prototype.String
					nb := np.NewBuilder()
					nb.AssignString(s)
					nd := nb.Build()

					statement, err := Like(selector.MustParse("."), pattern)
					require.NoError(t, err)

					pol := Policy{statement}
					ok := Match(pol, nd)
					require.False(t, ok)
				})
			}(s)
		}
	})

	t.Run("quantification", func(t *testing.T) {
		buildValueNode := func(v int64) ipld.Node {
			np := basicnode.Prototype.Map
			nb := np.NewBuilder()
			ma, _ := nb.BeginMap(1)
			ma.AssembleKey().AssignString("value")
			ma.AssembleValue().AssignInt(v)
			ma.Finish()
			return nb.Build()
		}

		t.Run("all", func(t *testing.T) {
			np := basicnode.Prototype.List
			nb := np.NewBuilder()
			la, _ := nb.BeginList(5)
			la.AssembleValue().AssignNode(buildValueNode(5))
			la.AssembleValue().AssignNode(buildValueNode(10))
			la.AssembleValue().AssignNode(buildValueNode(20))
			la.AssembleValue().AssignNode(buildValueNode(50))
			la.AssembleValue().AssignNode(buildValueNode(100))
			la.Finish()
			nd := nb.Build()

			pol := Policy{
				All(
					selector.MustParse(".[]"),
					GreaterThan(selector.MustParse(".value"), literal.Int(2)),
				),
			}
			ok := Match(pol, nd)
			require.True(t, ok)

			pol = Policy{
				All(
					selector.MustParse(".[]"),
					GreaterThan(selector.MustParse(".value"), literal.Int(20)),
				),
			}
			ok = Match(pol, nd)
			require.False(t, ok)
		})

		t.Run("any", func(t *testing.T) {
			np := basicnode.Prototype.List
			nb := np.NewBuilder()
			la, _ := nb.BeginList(5)
			la.AssembleValue().AssignNode(buildValueNode(5))
			la.AssembleValue().AssignNode(buildValueNode(10))
			la.AssembleValue().AssignNode(buildValueNode(20))
			la.AssembleValue().AssignNode(buildValueNode(50))
			la.AssembleValue().AssignNode(buildValueNode(100))
			la.Finish()
			nd := nb.Build()

			pol := Policy{
				Any(
					selector.MustParse(".[]"),
					GreaterThan(selector.MustParse(".value"), literal.Int(60)),
				),
			}
			ok := Match(pol, nd)
			require.True(t, ok)

			pol = Policy{
				Any(
					selector.MustParse(".[]"),
					GreaterThan(selector.MustParse(".value"), literal.Int(100)),
				),
			}
			ok = Match(pol, nd)
			require.False(t, ok)
		})
	})
}

func TestPolicyExamples(t *testing.T) {
	makeNode := func(data string) ipld.Node {
		nd, err := ipld.Decode([]byte(data), dagjson.Decode)
		require.NoError(t, err)
		return nd
	}

	evaluate := func(statement string, data ipld.Node) bool {
		// we need to wrap statement with [] to make them a policy
		policy := fmt.Sprintf("[%s]", statement)

		pol, err := FromDagJson(policy)
		require.NoError(t, err)
		return Match(pol, data)
	}

	t.Run("And", func(t *testing.T) {
		data := makeNode(`{ "name": "Katie", "age": 35, "nationalities": ["Canadian", "South African"] }`)

		require.True(t, evaluate(`["and", []]`, data))
		require.True(t, evaluate(`
["and", [
  ["==", ".name", "Katie"], 
  [">=", ".age", 21]
]]`, data))
		require.False(t, evaluate(`
["and", [
  ["==", ".name", "Katie"], 
  [">=", ".age", 21], 
  ["==", ".nationalities", ["American"]]
]]`, data))
	})

	t.Run("Or", func(t *testing.T) {
		data := makeNode(`{ "name": "Katie", "age": 35, "nationalities": ["Canadian", "South African"] }`)

		require.True(t, evaluate(`["or", []]`, data))
		require.True(t, evaluate(`
		["or", [
		  ["==", ".name", "Katie"],
		  [">", ".age", 45]
		]]
		`, data))

	})

	t.Run("Not", func(t *testing.T) {
		data := makeNode(`{ "name": "Katie", "nationalities": ["Canadian", "South African"] }`)

		require.True(t, evaluate(`
["not", 
  ["and", [
    ["==", ".name", "Katie"], 
    ["==", ".nationalities", ["American"]]
  ]]
]
`, data))
	})

	t.Run("All", func(t *testing.T) {
		data := makeNode(`{"a": [{"b": 1}, {"b": 2}, {"z": [7, 8, 9]}]}`)

		require.False(t, evaluate(`["all", ".a", [">", ".b", 0]]`, data))
	})

	t.Run("Any", func(t *testing.T) {
		data := makeNode(`{"a": [{"b": 1}, {"b": 2}, {"z": [7, 8, 9]}]}`)

		require.True(t, evaluate(`["any", ".a", ["==", ".b", 2]]`, data))
	})
}

const maxFuzzInputLength = 128

func FuzzMatch(f *testing.F) {
	// Policy + Data examples
	f.Add([]byte(`[["==", ".status", "draft"]]`), []byte(`{"status": "draft"}`))
	f.Add([]byte(`[["all", ".reviewer", ["like", ".email", "*@example.com"]]]`), []byte(`{"reviewer": [{"email": "alice@example.com"}, {"email": "bob@example.com"}]}`))
	f.Add([]byte(`[["any", ".tags", ["or", [["==", ".", "news"], ["==", ".", "press"]]]]]`), []byte(`{"tags": ["news", "press"]}`))
	f.Add([]byte(`[["==", ".name", "Alice"]]`), []byte(`{"name": "Alice"}`))
	f.Add([]byte(`[[">", ".age", 30]]`), []byte(`{"age": 31}`))
	f.Add([]byte(`[["<=", ".height", 180]]`), []byte(`{"height": 170}`))
	f.Add([]byte(`[["not", ["==", ".status", "inactive"]]]`), []byte(`{"status": "active"}`))
	f.Add([]byte(`[["and", [["==", ".role", "admin"], [">=", ".experience", 5]]]]`), []byte(`{"role": "admin", "experience": 6}`))
	f.Add([]byte(`[["or", [["==", ".department", "HR"], ["==", ".department", "Finance"]]]]`), []byte(`{"department": "HR"}`))
	f.Add([]byte(`[["like", ".email", "*@company.com"]]`), []byte(`{"email": "user@company.com"}`))
	f.Add([]byte(`[["all", ".projects", [">", ".budget", 10000]]]`), []byte(`{"projects": [{"budget": 15000}, {"budget": 8000}]}`))
	f.Add([]byte(`[["any", ".skills", ["==", ".", "Go"]]]`), []byte(`{"skills": ["Go", "Python", "JavaScript"]}`))
	f.Add(
		[]byte(`[["and", [
			["==", ".name", "Bob"],
			["or", [[">", ".age", 25],["==", ".status", "active"]]],
			["all", ".tasks", ["==", ".completed", true]]
		]]]`),
		[]byte(`{
			"name": "Bob",
			"age": 26,
			"status": "active",
			"tasks": [{"completed": true}, {"completed": true}, {"completed": false}]
		}`),
	)

	f.Fuzz(func(t *testing.T, policyBytes []byte, dataBytes []byte) {
		if len(policyBytes) > maxFuzzInputLength || len(dataBytes) > maxFuzzInputLength {
			t.Skip()
		}

		policyNode, err := ipld.Decode(policyBytes, dagjson.Decode)
		if err != nil {
			t.Skip()
		}

		dataNode, err := ipld.Decode(dataBytes, dagjson.Decode)
		if err != nil {
			t.Skip()
		}

		// policy node -> policy object
		policy, err := FromIPLD(policyNode)
		if err != nil {
			t.Skip()
		}

		Match(policy, dataNode)
	})
}
