package selector

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"github.com/ipld/go-ipld-prime/must"
	basicnode "github.com/ipld/go-ipld-prime/node/basic"
	"github.com/ipld/go-ipld-prime/node/bindnode"
	"github.com/ipld/go-ipld-prime/printer"
	"github.com/stretchr/testify/require"
)

func TestSelect(t *testing.T) {
	type name struct {
		First  string
		Middle *string
		Last   string
	}
	type interest struct {
		Name       string
		Outdoor    bool
		Experience int
	}
	type user struct {
		Name          name
		Age           int
		Nationalities []string
		Interests     []interest
	}

	ts, err := ipld.LoadSchemaBytes([]byte(`
	  type User struct {
		  name Name
			age Int
			nationalities [String]
			interests [Interest]
		}
		type Name struct {
			first String
			middle optional String
			last String
		}
		type Interest struct {
			name String
			outdoor Bool
			experience Int
		}
	`))
	require.NoError(t, err)
	typ := ts.TypeByName("User")

	am := "Joan"
	alice := user{
		Name:          name{First: "Alice", Middle: &am, Last: "Wonderland"},
		Age:           24,
		Nationalities: []string{"British"},
		Interests: []interest{
			{Name: "Cycling", Outdoor: true, Experience: 4},
			{Name: "Chess", Outdoor: false, Experience: 2},
		},
	}
	bob := user{
		Name:          name{First: "Bob", Last: "Builder"},
		Age:           35,
		Nationalities: []string{"Canadian", "South African"},
		Interests: []interest{
			{Name: "Snowboarding", Outdoor: true, Experience: 8},
			{Name: "Reading", Outdoor: false, Experience: 25},
		},
	}

	anode := bindnode.Wrap(&alice, typ)
	bnode := bindnode.Wrap(&bob, typ)

	t.Run("identity", func(t *testing.T) {
		sel, err := Parse(".")
		require.NoError(t, err)

		res, err := sel.Select(anode)
		require.NoError(t, err)
		require.NotEmpty(t, res)

		fmt.Println(printer.Sprint(res))

		age := must.Int(must.Node(res.LookupByString("age")))
		require.Equal(t, int64(alice.Age), age)
	})

	t.Run("nested property", func(t *testing.T) {
		sel, err := Parse(".name.first")
		require.NoError(t, err)

		res, err := sel.Select(anode)
		require.NoError(t, err)
		require.NotEmpty(t, res)

		fmt.Println(printer.Sprint(res))

		name := must.String(res)
		require.Equal(t, alice.Name.First, name)

		res, err = sel.Select(bnode)
		require.NoError(t, err)
		require.NotEmpty(t, res)

		fmt.Println(printer.Sprint(res))

		name = must.String(res)
		require.Equal(t, bob.Name.First, name)
	})

	t.Run("optional nested property", func(t *testing.T) {
		sel, err := Parse(".name.middle?")
		require.NoError(t, err)

		res, err := sel.Select(anode)
		require.NoError(t, err)
		require.NotEmpty(t, res)

		fmt.Println(printer.Sprint(res))

		name := must.String(res)
		require.Equal(t, *alice.Name.Middle, name)

		res, err = sel.Select(bnode)
		require.NoError(t, err)
		require.Empty(t, res)
	})

	t.Run("not exists", func(t *testing.T) {
		sel, err := Parse(".name.foo")
		require.NoError(t, err)

		res, err := sel.Select(anode)
		require.Error(t, err)
		require.Empty(t, res)

		fmt.Println(err)

		require.ErrorAs(t, err, &resolutionerr{}, "error should be a resolution error")
		require.True(t, IsResolutionErr(err))
	})

	t.Run("optional not exists", func(t *testing.T) {
		sel, err := Parse(".name.foo?")
		require.NoError(t, err)

		one, err := sel.Select(anode)
		require.NoError(t, err)
		require.Empty(t, one)
	})

	t.Run("iterator", func(t *testing.T) {
		sel, err := Parse(".interests[]")
		require.NoError(t, err)

		res, err := sel.Select(anode)
		require.NoError(t, err)
		require.NotEmpty(t, res)

		fmt.Println(printer.Sprint(res))

		iname := must.String(must.Node(must.Node(res.LookupByIndex(0)).LookupByString("name")))
		require.Equal(t, alice.Interests[0].Name, iname)

		iname = must.String(must.Node(must.Node(res.LookupByIndex(1)).LookupByString("name")))
		require.Equal(t, alice.Interests[1].Name, iname)
	})

	// TODO: fully test slicing + negative numbers
}

// func TestMatch(t *testing.T) {
// 	for _, tc := range []struct {
// 		sel       string
// 		path      []string
// 		want      bool
// 		remaining []string
// 	}{
// 		{sel: ".foo.bar", path: []string{"foo", "bar"}, want: true, remaining: []string{}},
// 		{sel: ".foo.bar", path: []string{"foo"}, want: true, remaining: []string{}},
// 		{sel: ".foo.bar", path: []string{"foo", "bar", "baz"}, want: true, remaining: []string{"baz"}},
// 		{sel: ".foo.bar", path: []string{"foo", "faa"}, want: false},
// 		{sel: ".foo.[]", path: []string{"foo", "faa"}, want: false},
// 		{sel: ".foo.[]", path: []string{"foo"}, want: true, remaining: []string{}},
// 		{sel: ".foo.bar?", path: []string{"foo"}, want: true, remaining: []string{}},
// 		{sel: ".foo.bar?", path: []string{"foo", "bar"}, want: true, remaining: []string{}},
// 		{sel: ".foo.bar?", path: []string{"foo", "baz"}, want: false},
// 	} {
// 		t.Run(tc.sel, func(t *testing.T) {
// 			sel := MustParse(tc.sel)
// 			res, remain := sel.MatchPath(tc.path...)
// 			require.Equal(t, tc.want, res)
// 			require.EqualValues(t, tc.remaining, remain)
// 		})
// 	}
// }

func FuzzParse(f *testing.F) {
	selectorCorpus := []string{
		`.`, `.[]`, `.[]?`, `.[][]?`, `.x`, `.["x"]`, `.[0]`, `.[-1]`, `.[0]`,
		`.[0]`, `.[0:2]`, `.[1:]`, `.[:2]`, `.[0:2]`, `.[1:]`, `.x?`, `.x?`,
		`.x?`, `.["x"]?`, `.length?`, `.[4]?`, `.[]`, `.[][]`, `.x`, `.x`, `.x`,
		`.length`, `.[4]`,
	}
	for _, selector := range selectorCorpus {
		f.Add(selector)
	}
	f.Fuzz(func(t *testing.T, selector string) {
		// only look for panic()
		_, _ = Parse(selector)
	})
}

func FuzzParseAndSelect(f *testing.F) {
	selectorCorpus := []string{
		`.`, `.[]`, `.[]?`, `.[][]?`, `.x`, `.["x"]`, `.[0]`, `.[-1]`, `.[0]`,
		`.[0]`, `.[0:2]`, `.[1:]`, `.[:2]`, `.[0:2]`, `.[1:]`, `.x?`, `.x?`,
		`.x?`, `.["x"]?`, `.length?`, `.[4]?`, `.[]`, `.[][]`, `.x`, `.x`, `.x`,
		`.length`, `.[4]`,
	}
	subjectCorpus := []string{
		`{"x":1}`, `[1, 2]`, `null`, `[[1], 2, [3]]`, `{"x": 1 }`, `{"x": 1}`,
		`[1, 2]`, `[1, 2]`, `"Hi"`, `{"/":{"bytes":"AAE"}`, `[0, 1, 2]`,
		`[0, 1, 2]`, `[0, 1, 2]`, `"hello"`, `{"/":{"bytes":"AAEC"}}`, `{}`,
		`null`, `[]`, `{}`, `[1, 2]`, `[0, 1]`, `null`, `[[1], 2, [3]]`, `{}`,
		`null`, `[]`, `[1, 2]`, `[0, 1]`,
	}
	for i := 0; ; i++ {
		switch {
		case i < len(selectorCorpus) && i < len(subjectCorpus):
			f.Add(selectorCorpus[i], subjectCorpus[i])
			continue
		case i > len(selectorCorpus):
			f.Add("", subjectCorpus[i])
			continue
		case i > len(subjectCorpus):
			f.Add(selectorCorpus[i], "")
			continue
		}
		break
	}

	f.Fuzz(func(t *testing.T, selector, subject string) {
		sel, err := Parse(selector)
		if err != nil {
			t.Skip()
		}

		np := basicnode.Prototype.Any
		nb := np.NewBuilder()
		err = dagjson.Decode(nb, strings.NewReader(subject))
		if err != nil {
			t.Skip()
		}
		node := nb.Build()
		if node == nil {
			t.Skip()
		}

		// look for panic()
		_, err = sel.Select(node)
		if err != nil && !IsResolutionErr(err) {
			// not normal, we should only have resolution errors
			t.Fatal(err)
		}
	})
}
