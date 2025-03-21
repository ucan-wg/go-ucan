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

	"github.com/ucan-wg/go-ucan/pkg/policy/literal"
)

func TestMatch(t *testing.T) {
	t.Run("equality", func(t *testing.T) {
		t.Run("eq string", func(t *testing.T) {
			nd := literal.String("test")

			pol := MustConstruct(Equal(".", literal.String("test")))
			ok, leaf := pol.Match(nd)
			require.True(t, ok)
			require.Nil(t, leaf)

			pol = MustConstruct(Equal(".", literal.String("test2")))
			ok, leaf = pol.Match(nd)
			require.False(t, ok)
			require.Equal(t, pol[0], leaf)

			pol = MustConstruct(Equal(".", literal.Int(138)))
			ok, leaf = pol.Match(nd)
			require.False(t, ok)
			require.Equal(t, pol[0], leaf)
		})

		t.Run("eq int", func(t *testing.T) {
			nd := literal.Int(138)

			pol := MustConstruct(Equal(".", literal.Int(138)))
			ok, leaf := pol.Match(nd)
			require.True(t, ok)
			require.Nil(t, leaf)

			pol = MustConstruct(Equal(".", literal.Int(1138)))
			ok, leaf = pol.Match(nd)
			require.False(t, ok)
			require.Equal(t, pol[0], leaf)

			pol = MustConstruct(Equal(".", literal.String("138")))
			ok, leaf = pol.Match(nd)
			require.False(t, ok)
			require.Equal(t, pol[0], leaf)
		})

		t.Run("eq float", func(t *testing.T) {
			nd := literal.Float(1.138)

			pol := MustConstruct(Equal(".", literal.Float(1.138)))
			ok, leaf := pol.Match(nd)
			require.True(t, ok)
			require.Nil(t, leaf)

			pol = MustConstruct(Equal(".", literal.Float(11.38)))
			ok, leaf = pol.Match(nd)
			require.False(t, ok)
			require.Equal(t, pol[0], leaf)

			pol = MustConstruct(Equal(".", literal.String("138")))
			ok, leaf = pol.Match(nd)
			require.False(t, ok)
			require.Equal(t, pol[0], leaf)
		})

		t.Run("eq IPLD Link", func(t *testing.T) {
			l0 := cidlink.Link{Cid: cid.MustParse("bafybeif4owy5gno5lwnixqm52rwqfodklf76hsetxdhffuxnplvijskzqq")}
			l1 := cidlink.Link{Cid: cid.MustParse("bafkreifau35r7vi37tvbvfy3hdwvgb4tlflqf7zcdzeujqcjk3rsphiwte")}

			nd := literal.Link(l0)

			pol := MustConstruct(Equal(".", literal.Link(l0)))
			ok, leaf := pol.Match(nd)
			require.True(t, ok)
			require.Nil(t, leaf)

			pol = MustConstruct(Equal(".", literal.Link(l1)))
			ok, leaf = pol.Match(nd)
			require.False(t, ok)
			require.Equal(t, pol[0], leaf)

			pol = MustConstruct(Equal(".", literal.String("bafybeif4owy5gno5lwnixqm52rwqfodklf76hsetxdhffuxnplvijskzqq")))
			ok, leaf = pol.Match(nd)
			require.False(t, ok)
			require.Equal(t, pol[0], leaf)
		})

		t.Run("eq string in map", func(t *testing.T) {
			nd, _ := literal.Map(map[string]any{
				"foo": "bar",
			})

			pol := MustConstruct(Equal(".foo", literal.String("bar")))
			ok, leaf := pol.Match(nd)
			require.True(t, ok)
			require.Nil(t, leaf)

			pol = MustConstruct(Equal(".[\"foo\"]", literal.String("bar")))
			ok, leaf = pol.Match(nd)
			require.True(t, ok)
			require.Nil(t, leaf)

			pol = MustConstruct(Equal(".foo", literal.String("baz")))
			ok, leaf = pol.Match(nd)
			require.False(t, ok)
			require.Equal(t, pol[0], leaf)

			pol = MustConstruct(Equal(".foobar", literal.String("bar")))
			ok, leaf = pol.Match(nd)
			require.False(t, ok)
			require.Equal(t, pol[0], leaf)
		})

		t.Run("eq string in list", func(t *testing.T) {
			nd, _ := literal.List([]any{"foo"})

			pol := MustConstruct(Equal(".[0]", literal.String("foo")))
			ok, leaf := pol.Match(nd)
			require.True(t, ok)
			require.Nil(t, leaf)

			pol = MustConstruct(Equal(".[1]", literal.String("foo")))
			ok, leaf = pol.Match(nd)
			require.False(t, ok)
			require.Equal(t, pol[0], leaf)
		})

		t.Run("neq string", func(t *testing.T) {
			nd := literal.String("test")

			pol := MustConstruct(NotEqual(".", literal.String("test")))
			ok, leaf := pol.Match(nd)
			require.False(t, ok)
			require.Equal(t, pol[0], leaf)

			pol = MustConstruct(NotEqual(".", literal.String("test2")))
			ok, leaf = pol.Match(nd)
			require.True(t, ok)
			require.Nil(t, leaf)

			pol = MustConstruct(NotEqual(".", literal.Int(138)))
			ok, leaf = pol.Match(nd)
			require.True(t, ok)
			require.Nil(t, leaf)
		})

		t.Run("neq int", func(t *testing.T) {
			nd := literal.Int(138)

			pol := MustConstruct(NotEqual(".", literal.Int(138)))
			ok, leaf := pol.Match(nd)
			require.False(t, ok)
			require.Equal(t, pol[0], leaf)

			pol = MustConstruct(NotEqual(".", literal.Int(1138)))
			ok, leaf = pol.Match(nd)
			require.True(t, ok)
			require.Nil(t, leaf)

			pol = MustConstruct(NotEqual(".", literal.String("138")))
			ok, leaf = pol.Match(nd)
			require.True(t, ok)
			require.Nil(t, leaf)
		})

		t.Run("neq float", func(t *testing.T) {
			nd := literal.Float(1.138)

			pol := MustConstruct(NotEqual(".", literal.Float(1.138)))
			ok, leaf := pol.Match(nd)
			require.False(t, ok)
			require.Equal(t, pol[0], leaf)

			pol = MustConstruct(NotEqual(".", literal.Float(11.38)))
			ok, leaf = pol.Match(nd)
			require.True(t, ok)
			require.Nil(t, leaf)

			pol = MustConstruct(NotEqual(".", literal.String("138")))
			ok, leaf = pol.Match(nd)
			require.True(t, ok)
			require.Nil(t, leaf)
		})

		t.Run("neq IPLD Link", func(t *testing.T) {
			l0 := cidlink.Link{Cid: cid.MustParse("bafybeif4owy5gno5lwnixqm52rwqfodklf76hsetxdhffuxnplvijskzqq")}
			l1 := cidlink.Link{Cid: cid.MustParse("bafkreifau35r7vi37tvbvfy3hdwvgb4tlflqf7zcdzeujqcjk3rsphiwte")}

			nd := literal.Link(l0)

			pol := MustConstruct(NotEqual(".", literal.Link(l0)))
			ok, leaf := pol.Match(nd)
			require.False(t, ok)
			require.Equal(t, pol[0], leaf)

			pol = MustConstruct(NotEqual(".", literal.Link(l1)))
			ok, leaf = pol.Match(nd)
			require.True(t, ok)
			require.Nil(t, leaf)

			pol = MustConstruct(NotEqual(".", literal.String("bafybeif4owy5gno5lwnixqm52rwqfodklf76hsetxdhffuxnplvijskzqq")))
			ok, leaf = pol.Match(nd)
			require.True(t, ok)
			require.Nil(t, leaf)
		})

		t.Run("neq string in map", func(t *testing.T) {
			nd, _ := literal.Map(map[string]any{
				"foo": "bar",
			})

			pol := MustConstruct(NotEqual(".foo", literal.String("bar")))
			ok, leaf := pol.Match(nd)
			require.False(t, ok)
			require.Equal(t, pol[0], leaf)

			pol = MustConstruct(NotEqual(".[\"foo\"]", literal.String("bar")))
			ok, leaf = pol.Match(nd)
			require.False(t, ok)
			require.Equal(t, pol[0], leaf)

			pol = MustConstruct(NotEqual(".foo", literal.String("baz")))
			ok, leaf = pol.Match(nd)
			require.True(t, ok)
			require.Nil(t, leaf)

			// missing data will fail, as not optional
			pol = MustConstruct(NotEqual(".foobar", literal.String("bar")))
			ok, leaf = pol.Match(nd)
			require.False(t, ok)
			require.Equal(t, pol[0], leaf)
		})

		t.Run("neq string in list", func(t *testing.T) {
			nd, _ := literal.List([]any{"foo"})

			pol := MustConstruct(NotEqual(".[0]", literal.String("foo")))
			ok, leaf := pol.Match(nd)
			require.False(t, ok)
			require.Equal(t, pol[0], leaf)

			pol = MustConstruct(NotEqual(".[0]", literal.String("bar")))
			ok, leaf = pol.Match(nd)
			require.True(t, ok)
			require.Nil(t, leaf)

			// missing data will fail, as not optional
			pol = MustConstruct(NotEqual(".[1]", literal.String("foo")))
			ok, leaf = pol.Match(nd)
			require.False(t, ok)
			require.Equal(t, pol[0], leaf)
		})
	})

	t.Run("inequality", func(t *testing.T) {
		t.Run("gt int", func(t *testing.T) {
			nd := literal.Int(138)

			pol := MustConstruct(GreaterThan(".", literal.Int(1)))
			ok, leaf := pol.Match(nd)
			require.True(t, ok)
			require.Nil(t, leaf)

			pol = MustConstruct(GreaterThan(".", literal.Int(138)))
			ok, leaf = pol.Match(nd)
			require.False(t, ok)
			require.Equal(t, pol[0], leaf)

			pol = MustConstruct(GreaterThan(".", literal.Int(140)))
			ok, leaf = pol.Match(nd)
			require.False(t, ok)
			require.Equal(t, pol[0], leaf)
		})

		t.Run("gte int", func(t *testing.T) {
			nd := literal.Int(138)

			pol := MustConstruct(GreaterThanOrEqual(".", literal.Int(1)))
			ok, leaf := pol.Match(nd)
			require.True(t, ok)
			require.Nil(t, leaf)

			pol = MustConstruct(GreaterThanOrEqual(".", literal.Int(138)))
			ok, leaf = pol.Match(nd)
			require.True(t, ok)
			require.Nil(t, leaf)

			pol = MustConstruct(GreaterThanOrEqual(".", literal.Int(140)))
			ok, leaf = pol.Match(nd)
			require.False(t, ok)
			require.Equal(t, pol[0], leaf)
		})

		t.Run("gt float", func(t *testing.T) {
			nd := literal.Float(1.38)

			pol := MustConstruct(GreaterThan(".", literal.Float(1)))
			ok, leaf := pol.Match(nd)
			require.True(t, ok)
			require.Nil(t, leaf)

			pol = MustConstruct(GreaterThan(".", literal.Float(2)))
			ok, leaf = pol.Match(nd)
			require.False(t, ok)
			require.Equal(t, pol[0], leaf)
		})

		t.Run("gte float", func(t *testing.T) {
			nd := literal.Float(1.38)

			pol := MustConstruct(GreaterThanOrEqual(".", literal.Float(1)))
			ok, leaf := pol.Match(nd)
			require.True(t, ok)
			require.Nil(t, leaf)

			pol = MustConstruct(GreaterThanOrEqual(".", literal.Float(1.38)))
			ok, leaf = pol.Match(nd)
			require.True(t, ok)
			require.Nil(t, leaf)

			pol = MustConstruct(GreaterThanOrEqual(".", literal.Float(2)))
			ok, leaf = pol.Match(nd)
			require.False(t, ok)
			require.Equal(t, pol[0], leaf)
		})

		t.Run("lt int", func(t *testing.T) {
			nd := literal.Int(138)

			pol := MustConstruct(LessThan(".", literal.Int(1138)))
			ok, leaf := pol.Match(nd)
			require.True(t, ok)
			require.Nil(t, leaf)

			pol = MustConstruct(LessThan(".", literal.Int(138)))
			ok, leaf = pol.Match(nd)
			require.False(t, ok)
			require.Equal(t, pol[0], leaf)

			pol = MustConstruct(LessThan(".", literal.Int(100)))
			ok, leaf = pol.Match(nd)
			require.False(t, ok)
			require.Equal(t, pol[0], leaf)
		})

		t.Run("lte int", func(t *testing.T) {
			nd := literal.Int(138)

			pol := MustConstruct(LessThanOrEqual(".", literal.Int(1138)))
			ok, leaf := pol.Match(nd)
			require.True(t, ok)
			require.Nil(t, leaf)

			pol = MustConstruct(LessThanOrEqual(".", literal.Int(138)))
			ok, leaf = pol.Match(nd)
			require.True(t, ok)
			require.Nil(t, leaf)

			pol = MustConstruct(LessThanOrEqual(".", literal.Int(100)))
			ok, leaf = pol.Match(nd)
			require.False(t, ok)
			require.Equal(t, pol[0], leaf)
		})

		t.Run("lt float", func(t *testing.T) {
			nd := literal.Float(1.38)

			pol := MustConstruct(LessThan(".", literal.Float(1)))
			ok, leaf := pol.Match(nd)
			require.False(t, ok)
			require.Equal(t, pol[0], leaf)

			pol = MustConstruct(LessThan(".", literal.Float(2)))
			ok, leaf = pol.Match(nd)
			require.True(t, ok)
			require.Nil(t, leaf)
		})

		t.Run("lte float", func(t *testing.T) {
			nd := literal.Float(1.38)

			pol := MustConstruct(LessThanOrEqual(".", literal.Float(1)))
			ok, leaf := pol.Match(nd)
			require.False(t, ok)
			require.Equal(t, pol[0], leaf)

			pol = MustConstruct(GreaterThanOrEqual(".", literal.Float(1.38)))
			ok, leaf = pol.Match(nd)
			require.True(t, ok)
			require.Nil(t, leaf)

			pol = MustConstruct(LessThanOrEqual(".", literal.Float(2)))
			ok, leaf = pol.Match(nd)
			require.True(t, ok)
			require.Nil(t, leaf)
		})
	})

	t.Run("negation", func(t *testing.T) {
		nd, _ := literal.Map(map[string]any{
			"foo": false,
		})

		pol := MustConstruct(Not(Equal(".foo", literal.Bool(true))))
		ok, leaf := pol.Match(nd)
		require.True(t, ok)
		require.Nil(t, leaf)

		pol = MustConstruct(Not(Equal(".foo", literal.Bool(false))))
		ok, leaf = pol.Match(nd)
		require.False(t, ok)
		require.Equal(t, pol[0], leaf)

		// missing data will fail, as not optional
		pol = MustConstruct(Not(Equal(".foobar", literal.Bool(true))))
		ok, leaf = pol.Match(nd)
		require.False(t, ok)
		require.Equal(t, MustConstruct(Equal(".foobar", literal.Bool(true)))[0], leaf)
	})

	t.Run("conjunction", func(t *testing.T) {
		nd := literal.Int(138)

		pol := MustConstruct(
			And(
				GreaterThan(".", literal.Int(1)),
				LessThan(".", literal.Int(1138)),
			),
		)
		ok, leaf := pol.Match(nd)
		require.True(t, ok)
		require.Nil(t, leaf)

		pol = MustConstruct(
			And(
				GreaterThan(".", literal.Int(1)),
				Equal(".", literal.Int(1138)),
			),
		)
		ok, leaf = pol.Match(nd)
		require.False(t, ok)
		require.Equal(t, MustConstruct(Equal(".", literal.Int(1138)))[0], leaf)

		pol = MustConstruct(And())
		ok, leaf = pol.Match(nd)
		require.True(t, ok)
		require.Nil(t, leaf)
	})

	t.Run("disjunction", func(t *testing.T) {
		nd := literal.Int(138)

		pol := MustConstruct(
			Or(
				GreaterThan(".", literal.Int(138)),
				LessThan(".", literal.Int(1138)),
			),
		)
		ok, leaf := pol.Match(nd)
		require.True(t, ok)
		require.Nil(t, leaf)

		pol = MustConstruct(
			Or(
				GreaterThan(".", literal.Int(138)),
				Equal(".", literal.Int(1138)),
			),
		)
		ok, leaf = pol.Match(nd)
		require.False(t, ok)
		require.Equal(t, pol[0], leaf)

		pol = MustConstruct(Or())
		ok, leaf = pol.Match(nd)
		require.True(t, ok)
		require.Nil(t, leaf)
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
					nd := literal.String(s)

					pol := MustConstruct(Like(".", pattern))
					ok, leaf := pol.Match(nd)
					require.True(t, ok)
					require.Nil(t, leaf)
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
					nd := literal.String(s)

					pol := MustConstruct(Like(".", pattern))
					ok, leaf := pol.Match(nd)
					require.False(t, ok)
					require.Equal(t, pol[0], leaf)
				})
			}(s)
		}
	})

	t.Run("quantification", func(t *testing.T) {
		t.Run("all", func(t *testing.T) {
			nd, _ := literal.List([]any{
				map[string]int{"value": 5},
				map[string]int{"value": 10},
				map[string]int{"value": 20},
				map[string]int{"value": 50},
				map[string]int{"value": 100},
			})

			pol := MustConstruct(All(".[]", GreaterThan(".value", literal.Int(2))))
			ok, leaf := pol.Match(nd)
			require.True(t, ok)
			require.Nil(t, leaf)

			pol = MustConstruct(All(".[]", GreaterThan(".value", literal.Int(20))))
			ok, leaf = pol.Match(nd)
			require.False(t, ok)
			require.Equal(t, MustConstruct(GreaterThan(".value", literal.Int(20)))[0], leaf)
		})

		t.Run("any", func(t *testing.T) {
			nd, _ := literal.List([]any{
				map[string]int{"value": 5},
				map[string]int{"value": 10},
				map[string]int{"value": 20},
				map[string]int{"value": 50},
				map[string]int{"value": 100},
			})

			pol := MustConstruct(Any(".[]", GreaterThan(".value", literal.Int(60))))
			ok, leaf := pol.Match(nd)
			require.True(t, ok)
			require.Nil(t, leaf)

			pol = MustConstruct(Any(".[]", GreaterThan(".value", literal.Int(100))))
			ok, leaf = pol.Match(nd)
			require.False(t, ok)
			require.Equal(t, pol[0], leaf)
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
		res, _ := pol.Match(data)
		return res
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

func FuzzMatch(f *testing.F) {
	// Policy + Data examples
	f.Add([]byte(`[["==", ".status", "draft"]]`), []byte(`{"status": "draft"}`))
	f.Add([]byte(`[["all", ".reviewer", ["like", ".email", "*@example.com"]]]`), []byte(`{"reviewer": [{"email": "alice@example.com"}, {"email": "bob@example.com"}]}`))
	f.Add([]byte(`[["any", ".tags", ["or", [["==", ".", "news"], ["==", ".", "press"]]]]]`), []byte(`{"tags": ["news", "press"]}`))
	f.Add([]byte(`[["==", ".name", "Alice"]]`), []byte(`{"name": "Alice"}`))
	f.Add([]byte(`[["!=", ".name", "Alice"]]`), []byte(`{"name": "Alice"}`))
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

		_, _ = policy.Match(dataNode)
	})
}

func TestOptionalSelectors(t *testing.T) {
	tests := []struct {
		name     string
		policy   Policy
		data     map[string]any
		expected bool
	}{
		{
			name:     "missing optional field returns true",
			policy:   MustConstruct(Equal(".field?", literal.String("value"))),
			data:     map[string]any{},
			expected: true,
		},
		{
			name:     "present optional field with matching value returns true",
			policy:   MustConstruct(Equal(".field?", literal.String("value"))),
			data:     map[string]any{"field": "value"},
			expected: true,
		},
		{
			name:     "present optional field with non-matching value returns false",
			policy:   MustConstruct(Equal(".field?", literal.String("value"))),
			data:     map[string]any{"field": "other"},
			expected: false,
		},
		{
			name:     "missing non-optional field returns false",
			policy:   MustConstruct(Equal(".field", literal.String("value"))),
			data:     map[string]any{},
			expected: false,
		},
		{
			name:     "nested missing non-optional field returns false",
			policy:   MustConstruct(Equal(".outer?.inner", literal.String("value"))),
			data:     map[string]any{"outer": map[string]any{}},
			expected: false,
		},
		{
			name:     "completely missing nested optional path returns true",
			policy:   MustConstruct(Equal(".outer?.inner?", literal.String("value"))),
			data:     map[string]any{},
			expected: true,
		},
		{
			name:     "partially present nested optional path with missing end returns true",
			policy:   MustConstruct(Equal(".outer?.inner?", literal.String("value"))),
			data:     map[string]any{"outer": map[string]any{}},
			expected: true,
		},
		{
			name:     "optional array index returns true when array is empty",
			policy:   MustConstruct(Equal(".array[0]?", literal.String("value"))),
			data:     map[string]any{"array": []any{}},
			expected: true,
		},
		{
			name:     "non-optional array index returns false when array is empty",
			policy:   MustConstruct(Equal(".array[0]", literal.String("value"))),
			data:     map[string]any{"array": []any{}},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nb := basicnode.Prototype.Map.NewBuilder()
			n, err := literal.Map(tt.data)
			require.NoError(t, err)
			err = nb.AssignNode(n)
			require.NoError(t, err)

			result, _ := tt.policy.Match(nb.Build())
			require.Equal(t, tt.expected, result)
		})
	}
}

// The unique behaviour of PartialMatch is that it should return true for missing non-optional data (unlike Match).
func TestPartialMatch(t *testing.T) {
	tests := []struct {
		name          string
		policy        Policy
		data          map[string]any
		expectedMatch bool
		expectedStmt  Statement
	}{
		{
			name: "returns true for missing non-optional field",
			policy: MustConstruct(
				Equal(".field", literal.String("value")),
			),
			data:          map[string]any{},
			expectedMatch: true,
			expectedStmt:  nil,
		},
		{
			name: "returns true when present data matches",
			policy: MustConstruct(
				Equal(".foo", literal.String("correct")),
				Equal(".missing", literal.String("whatever")),
			),
			data: map[string]any{
				"foo": "correct",
			},
			expectedMatch: true,
			expectedStmt:  nil,
		},
		{
			name: "returns false with failing statement for present but non-matching value",
			policy: MustConstruct(
				Equal(".foo", literal.String("value1")),
				Equal(".bar", literal.String("value2")),
			),
			data: map[string]any{
				"foo": "wrong",
				"bar": "value2",
			},
			expectedMatch: false,
			expectedStmt: MustConstruct(
				Equal(".foo", literal.String("value1")),
			)[0],
		},
		{
			name: "continues past missing data until finding actual mismatch",
			policy: MustConstruct(
				Equal(".missing", literal.String("value")),
				Equal(".present", literal.String("wrong")),
			),
			data: map[string]any{
				"present": "actual",
			},
			expectedMatch: false,
			expectedStmt: MustConstruct(
				Equal(".present", literal.String("wrong")),
			)[0],
		},

		// Optional fields
		{
			name: "returns false when optional field present but wrong",
			policy: MustConstruct(
				Equal(".field?", literal.String("value")),
			),
			data: map[string]any{
				"field": "wrong",
			},
			expectedMatch: false,
			expectedStmt: MustConstruct(
				Equal(".field?", literal.String("value")),
			)[0],
		},

		// Like pattern matching
		{
			name: "returns true for matching like pattern",
			policy: MustConstruct(
				Like(".pattern", "test*"),
			),
			data: map[string]any{
				"pattern": "testing123",
			},
			expectedMatch: true,
			expectedStmt:  nil,
		},
		{
			name: "returns false for non-matching like pattern",
			policy: MustConstruct(
				Like(".pattern", "test*"),
			),
			data: map[string]any{
				"pattern": "wrong123",
			},
			expectedMatch: false,
			expectedStmt: MustConstruct(
				Like(".pattern", "test*"),
			)[0],
		},

		// Array quantifiers
		{
			name: "all matches when every element satisfies condition",
			policy: MustConstruct(
				All(".numbers", Equal(".", literal.Int(1))),
			),
			data: map[string]interface{}{
				"numbers": []interface{}{1, 1, 1},
			},
			expectedMatch: true,
			expectedStmt:  nil,
		},
		{
			name: "all fails when any element doesn't satisfy",
			policy: MustConstruct(
				All(".numbers", Equal(".", literal.Int(1))),
			),
			data: map[string]interface{}{
				"numbers": []interface{}{1, 2, 1},
			},
			expectedMatch: false,
			expectedStmt: MustConstruct(
				Equal(".", literal.Int(1)),
			)[0],
		},
		{
			name: "any succeeds when one element matches",
			policy: MustConstruct(
				Any(".numbers", Equal(".", literal.Int(2))),
			),
			data: map[string]interface{}{
				"numbers": []interface{}{1, 2, 3},
			},
			expectedMatch: true,
			expectedStmt:  nil,
		},
		{
			name: "any fails when no elements match",
			policy: MustConstruct(
				Any(".numbers", Equal(".", literal.Int(4))),
			),
			data: map[string]interface{}{
				"numbers": []interface{}{1, 2, 3},
			},
			expectedMatch: false,
			expectedStmt: MustConstruct(
				Any(".numbers", Equal(".", literal.Int(4))),
			)[0],
		},

		// Complex nested case
		{
			name: "complex nested policy",
			policy: MustConstruct(
				And(
					Equal(".required", literal.String("present")),
					Equal(".optional?", literal.String("value")),
					Any(".items",
						And(
							Equal(".name", literal.String("test")),
							Like(".id", "ID*"),
						),
					),
				),
			),
			data: map[string]any{
				"required": "present",
				"items": []any{
					map[string]any{
						"name": "wrong",
						"id":   "ID123",
					},
					map[string]any{
						"name": "test",
						"id":   "ID456",
					},
				},
			},
			expectedMatch: true,
			expectedStmt:  nil,
		},

		// missing optional values for all the operators
		{
			name: "returns true for missing optional equal",
			policy: MustConstruct(
				Equal(".field?", literal.String("value")),
			),
			data:          map[string]any{},
			expectedMatch: true,
			expectedStmt:  nil,
		},
		{
			name: "returns true for missing optional like pattern",
			policy: MustConstruct(
				Like(".pattern?", "test*"),
			),
			data:          map[string]any{},
			expectedMatch: true,
			expectedStmt:  nil,
		},
		{
			name: "returns true for missing optional greater than",
			policy: MustConstruct(
				GreaterThan(".number?", literal.Int(5)),
			),
			data:          map[string]any{},
			expectedMatch: true,
			expectedStmt:  nil,
		},
		{
			name: "returns true for missing optional less than",
			policy: MustConstruct(
				LessThan(".number?", literal.Int(5)),
			),
			data:          map[string]any{},
			expectedMatch: true,
			expectedStmt:  nil,
		},
		{
			name: "returns true for missing optional array with all",
			policy: MustConstruct(
				All(".numbers?", Equal(".", literal.Int(1))),
			),
			data:          map[string]any{},
			expectedMatch: true,
			expectedStmt:  nil,
		},
		{
			name: "returns true for missing optional array with any",
			policy: MustConstruct(
				Any(".numbers?", Equal(".", literal.Int(1))),
			),
			data:          map[string]any{},
			expectedMatch: true,
			expectedStmt:  nil,
		},
		{
			name: "returns true for complex nested optional paths",
			policy: MustConstruct(
				And(
					Equal(".required", literal.String("present")),
					Any(".optional_array?",
						And(
							Equal(".name?", literal.String("test")),
							Like(".id?", "ID*"),
						),
					),
				),
			),
			data: map[string]any{
				"required": "present",
			},
			expectedMatch: true,
			expectedStmt:  nil,
		},
		{
			name: "returns true for partially present nested optional paths",
			policy: MustConstruct(
				And(
					Equal(".required", literal.String("present")),
					Any(".items",
						And(
							Equal(".name", literal.String("test")),
							Like(".optional_id?", "ID*"),
						),
					),
				),
			),
			data: map[string]any{
				"required": "present",
				"items": []any{
					map[string]any{
						"name": "test",
						// optional_id is missing
					},
				},
			},
			expectedMatch: true,
			expectedStmt:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node, err := literal.Map(tt.data)
			require.NoError(t, err)

			match, stmt := tt.policy.PartialMatch(node)
			require.Equal(t, tt.expectedMatch, match)
			if tt.expectedStmt == nil {
				require.Nil(t, stmt)
			} else {
				require.Equal(t, tt.expectedStmt, stmt)
			}
		})
	}
}
