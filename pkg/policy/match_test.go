package policy

import (
	"fmt"
	"testing"

	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ucan-wg/go-ucan/pkg/policy/literal"
)

func TestMatch(t *testing.T) {
	t.Run("equality", func(t *testing.T) {
		t.Run("string", func(t *testing.T) {
			np := basicnode.Prototype.String
			nb := np.NewBuilder()
			nb.AssignString("test")
			nd := nb.Build()

			pol := MustConstruct(Equal(".", literal.String("test")))
			ok := pol.Match(nd)
			require.True(t, ok)

			pol = MustConstruct(Equal(".", literal.String("test2")))
			ok = pol.Match(nd)
			require.False(t, ok)

			pol = MustConstruct(Equal(".", literal.Int(138)))
			ok = pol.Match(nd)
			require.False(t, ok)
		})

		t.Run("int", func(t *testing.T) {
			np := basicnode.Prototype.Int
			nb := np.NewBuilder()
			nb.AssignInt(138)
			nd := nb.Build()

			pol := MustConstruct(Equal(".", literal.Int(138)))
			ok := pol.Match(nd)
			require.True(t, ok)

			pol = MustConstruct(Equal(".", literal.Int(1138)))
			ok = pol.Match(nd)
			require.False(t, ok)

			pol = MustConstruct(Equal(".", literal.String("138")))
			ok = pol.Match(nd)
			require.False(t, ok)
		})

		t.Run("float", func(t *testing.T) {
			np := basicnode.Prototype.Float
			nb := np.NewBuilder()
			nb.AssignFloat(1.138)
			nd := nb.Build()

			pol := MustConstruct(Equal(".", literal.Float(1.138)))
			ok := pol.Match(nd)
			require.True(t, ok)

			pol = MustConstruct(Equal(".", literal.Float(11.38)))
			ok = pol.Match(nd)
			require.False(t, ok)

			pol = MustConstruct(Equal(".", literal.String("138")))
			ok = pol.Match(nd)
			require.False(t, ok)
		})

		t.Run("IPLD Link", func(t *testing.T) {
			l0 := cidlink.Link{Cid: cid.MustParse("bafybeif4owy5gno5lwnixqm52rwqfodklf76hsetxdhffuxnplvijskzqq")}
			l1 := cidlink.Link{Cid: cid.MustParse("bafkreifau35r7vi37tvbvfy3hdwvgb4tlflqf7zcdzeujqcjk3rsphiwte")}

			np := basicnode.Prototype.Link
			nb := np.NewBuilder()
			nb.AssignLink(l0)
			nd := nb.Build()

			pol := MustConstruct(Equal(".", literal.Link(l0)))
			ok := pol.Match(nd)
			require.True(t, ok)

			pol = MustConstruct(Equal(".", literal.Link(l1)))
			ok = pol.Match(nd)
			require.False(t, ok)

			pol = MustConstruct(Equal(".", literal.String("bafybeif4owy5gno5lwnixqm52rwqfodklf76hsetxdhffuxnplvijskzqq")))
			ok = pol.Match(nd)
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

			pol := MustConstruct(Equal(".foo", literal.String("bar")))
			ok := pol.Match(nd)
			require.True(t, ok)

			pol = MustConstruct(Equal(".[\"foo\"]", literal.String("bar")))
			ok = pol.Match(nd)
			require.True(t, ok)

			pol = MustConstruct(Equal(".foo", literal.String("baz")))
			ok = pol.Match(nd)
			require.False(t, ok)

			pol = MustConstruct(Equal(".foobar", literal.String("bar")))
			ok = pol.Match(nd)
			require.False(t, ok)
		})

		t.Run("string in list", func(t *testing.T) {
			np := basicnode.Prototype.List
			nb := np.NewBuilder()
			la, _ := nb.BeginList(1)
			la.AssembleValue().AssignString("foo")
			la.Finish()
			nd := nb.Build()

			pol := MustConstruct(Equal(".[0]", literal.String("foo")))
			ok := pol.Match(nd)
			require.True(t, ok)

			pol = MustConstruct(Equal(".[1]", literal.String("foo")))
			ok = pol.Match(nd)
			require.False(t, ok)
		})
	})

	t.Run("inequality", func(t *testing.T) {
		t.Run("gt int", func(t *testing.T) {
			np := basicnode.Prototype.Int
			nb := np.NewBuilder()
			nb.AssignInt(138)
			nd := nb.Build()

			pol := MustConstruct(GreaterThan(".", literal.Int(1)))
			ok := pol.Match(nd)
			require.True(t, ok)
		})

		t.Run("gte int", func(t *testing.T) {
			np := basicnode.Prototype.Int
			nb := np.NewBuilder()
			nb.AssignInt(138)
			nd := nb.Build()

			pol := MustConstruct(GreaterThanOrEqual(".", literal.Int(1)))
			ok := pol.Match(nd)
			require.True(t, ok)

			pol = MustConstruct(GreaterThanOrEqual(".", literal.Int(138)))
			ok = pol.Match(nd)
			require.True(t, ok)
		})

		t.Run("gt float", func(t *testing.T) {
			np := basicnode.Prototype.Float
			nb := np.NewBuilder()
			nb.AssignFloat(1.38)
			nd := nb.Build()

			pol := MustConstruct(GreaterThan(".", literal.Float(1)))
			ok := pol.Match(nd)
			require.True(t, ok)
		})

		t.Run("gte float", func(t *testing.T) {
			np := basicnode.Prototype.Float
			nb := np.NewBuilder()
			nb.AssignFloat(1.38)
			nd := nb.Build()

			pol := MustConstruct(GreaterThanOrEqual(".", literal.Float(1)))
			ok := pol.Match(nd)
			require.True(t, ok)

			pol = MustConstruct(GreaterThanOrEqual(".", literal.Float(1.38)))
			ok = pol.Match(nd)
			require.True(t, ok)
		})

		t.Run("lt int", func(t *testing.T) {
			np := basicnode.Prototype.Int
			nb := np.NewBuilder()
			nb.AssignInt(138)
			nd := nb.Build()

			pol := MustConstruct(LessThan(".", literal.Int(1138)))
			ok := pol.Match(nd)
			require.True(t, ok)
		})

		t.Run("lte int", func(t *testing.T) {
			np := basicnode.Prototype.Int
			nb := np.NewBuilder()
			nb.AssignInt(138)
			nd := nb.Build()

			pol := MustConstruct(LessThanOrEqual(".", literal.Int(1138)))
			ok := pol.Match(nd)
			require.True(t, ok)

			pol = MustConstruct(LessThanOrEqual(".", literal.Int(138)))
			ok = pol.Match(nd)
			require.True(t, ok)
		})
	})

	t.Run("negation", func(t *testing.T) {
		np := basicnode.Prototype.Bool
		nb := np.NewBuilder()
		nb.AssignBool(false)
		nd := nb.Build()

		pol := MustConstruct(Not(Equal(".", literal.Bool(true))))
		ok := pol.Match(nd)
		require.True(t, ok)

		pol = MustConstruct(Not(Equal(".", literal.Bool(false))))
		ok = pol.Match(nd)
		require.False(t, ok)
	})

	t.Run("conjunction", func(t *testing.T) {
		np := basicnode.Prototype.Int
		nb := np.NewBuilder()
		nb.AssignInt(138)
		nd := nb.Build()

		pol := MustConstruct(
			And(
				GreaterThan(".", literal.Int(1)),
				LessThan(".", literal.Int(1138)),
			),
		)
		ok := pol.Match(nd)
		require.True(t, ok)

		pol = MustConstruct(
			And(
				GreaterThan(".", literal.Int(1)),
				Equal(".", literal.Int(1138)),
			),
		)
		ok = pol.Match(nd)
		require.False(t, ok)

		pol = MustConstruct(And())
		ok = pol.Match(nd)
		require.True(t, ok)
	})

	t.Run("disjunction", func(t *testing.T) {
		np := basicnode.Prototype.Int
		nb := np.NewBuilder()
		nb.AssignInt(138)
		nd := nb.Build()

		pol := MustConstruct(
			Or(
				GreaterThan(".", literal.Int(138)),
				LessThan(".", literal.Int(1138)),
			),
		)
		ok := pol.Match(nd)
		require.True(t, ok)

		pol = MustConstruct(
			Or(
				GreaterThan(".", literal.Int(138)),
				Equal(".", literal.Int(1138)),
			),
		)
		ok = pol.Match(nd)
		require.False(t, ok)

		pol = MustConstruct(Or())
		ok = pol.Match(nd)
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

					pol := MustConstruct(Like(".", pattern))
					ok := pol.Match(nd)
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

					pol := MustConstruct(Like(".", pattern))
					ok := pol.Match(nd)
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

			pol := MustConstruct(All(".[]", GreaterThan(".value", literal.Int(2))))
			ok := pol.Match(nd)
			require.True(t, ok)

			pol = MustConstruct(All(".[]", GreaterThan(".value", literal.Int(20))))
			ok = pol.Match(nd)
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

			pol := MustConstruct(Any(".[]", GreaterThan(".value", literal.Int(60))))
			ok := pol.Match(nd)
			require.True(t, ok)

			pol = MustConstruct(Any(".[]", GreaterThan(".value", literal.Int(100))))
			ok = pol.Match(nd)
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
		return pol.Match(data)
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

	t.Run("Map", func(t *testing.T) {
		data := makeNode(`{"a": [{"b": 1}, {"b": 2}, {"z": [7, 8, 9]}]}`)

		require.True(t, evaluate(`["any", ".a", ["==", ".b", 2]]`, data))
	})
}

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

		policy.Match(dataNode)
	})
}

func TestOptionalSelectors(t *testing.T) {
	tests := []struct {
		name     string
		policy   Policy
		data     interface{}
		expected bool
	}{
		{
			name:     "missing optional field returns true",
			policy:   MustConstruct(Equal(".field?", literal.String("value"))),
			data:     map[string]interface{}{},
			expected: true,
		},
		{
			name:     "present optional field with matching value returns true",
			policy:   MustConstruct(Equal(".field?", literal.String("value"))),
			data:     map[string]interface{}{"field": "value"},
			expected: true,
		},
		{
			name:     "present optional field with non-matching value returns false",
			policy:   MustConstruct(Equal(".field?", literal.String("value"))),
			data:     map[string]interface{}{"field": "other"},
			expected: false,
		},
		{
			name:     "missing non-optional field returns false",
			policy:   MustConstruct(Equal(".field", literal.String("value"))),
			data:     map[string]interface{}{},
			expected: false,
		},
		{
			name:     "nested missing non-optional field returns false",
			policy:   MustConstruct(Equal(".outer?.inner", literal.String("value"))),
			data:     map[string]interface{}{"outer": map[string]interface{}{}},
			expected: false,
		},
		{
			name:     "completely missing nested optional path returns true",
			policy:   MustConstruct(Equal(".outer?.inner?", literal.String("value"))),
			data:     map[string]interface{}{},
			expected: true,
		},
		{
			name:     "partially present nested optional path with missing end returns true",
			policy:   MustConstruct(Equal(".outer?.inner?", literal.String("value"))),
			data:     map[string]interface{}{"outer": map[string]interface{}{}},
			expected: true,
		},
		{
			name:     "optional array index returns true when array is empty",
			policy:   MustConstruct(Equal(".array[0]?", literal.String("value"))),
			data:     map[string]interface{}{"array": []interface{}{}},
			expected: true,
		},
		{
			name:     "non-optional array index returns false when array is empty",
			policy:   MustConstruct(Equal(".array[0]", literal.String("value"))),
			data:     map[string]interface{}{"array": []interface{}{}},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nb := basicnode.Prototype.Map.NewBuilder()
			n, err := literal.Map(tt.data)
			assert.NoError(t, err)
			err = nb.AssignNode(n)
			assert.NoError(t, err)

			result := tt.policy.Match(nb.Build())
			assert.Equal(t, tt.expected, result)
		})
	}
}
