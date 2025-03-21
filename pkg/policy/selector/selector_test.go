package selector

import (
	"errors"
	"math"
	"strings"
	"testing"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/fluent/qp"
	"github.com/ipld/go-ipld-prime/must"
	basicnode "github.com/ipld/go-ipld-prime/node/basic"
	"github.com/ipld/go-ipld-prime/node/bindnode"
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

		age := must.Int(must.Node(res.LookupByString("age")))
		require.Equal(t, int64(alice.Age), age)
	})

	t.Run("nested property", func(t *testing.T) {
		sel, err := Parse(".name.first")
		require.NoError(t, err)

		res, err := sel.Select(anode)
		require.NoError(t, err)
		require.NotEmpty(t, res)

		name := must.String(res)
		require.Equal(t, alice.Name.First, name)

		res, err = sel.Select(bnode)
		require.NoError(t, err)
		require.NotEmpty(t, res)

		name = must.String(res)
		require.Equal(t, bob.Name.First, name)
	})

	t.Run("optional nested property", func(t *testing.T) {
		sel, err := Parse(".name.middle?")
		require.NoError(t, err)

		res, err := sel.Select(anode)
		require.NoError(t, err)
		require.NotEmpty(t, res)

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

		require.ErrorAs(t, err, &resolutionErr{}, "error should be a resolution error")
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

		iname := must.String(must.Node(must.Node(res.LookupByIndex(0)).LookupByString("name")))
		require.Equal(t, alice.Interests[0].Name, iname)

		iname = must.String(must.Node(must.Node(res.LookupByIndex(1)).LookupByString("name")))
		require.Equal(t, alice.Interests[1].Name, iname)
	})

	t.Run("slice on string", func(t *testing.T) {
		sel, err := Parse(`.[1:3]`)
		require.NoError(t, err)

		node := basicnode.NewString("hello")
		res, err := sel.Select(node)
		require.NoError(t, err)
		require.NotEmpty(t, res)

		str, err := res.AsString()
		require.NoError(t, err)
		require.Equal(t, "el", str) // assert sliced substring
	})

	t.Run("out of bounds slicing", func(t *testing.T) {
		node, err := qp.BuildList(basicnode.Prototype.Any, 3, func(la datamodel.ListAssembler) {
			qp.ListEntry(la, qp.Int(1))
			qp.ListEntry(la, qp.Int(2))
			qp.ListEntry(la, qp.Int(3))
		})
		require.NoError(t, err)

		sel, err := Parse(`.[10:20]`)
		require.NoError(t, err)

		res, err := sel.Select(node)
		require.NoError(t, err)
		require.NotEmpty(t, res)
		require.Equal(t, int64(0), res.Length())

		_, err = res.LookupByIndex(0)
		require.ErrorIs(t, err, datamodel.ErrNotExists{}) // assert empty result for out of bounds slice
	})

	t.Run("backward slicing", func(t *testing.T) {
		node, err := qp.BuildList(basicnode.Prototype.Any, 3, func(la datamodel.ListAssembler) {
			qp.ListEntry(la, qp.Int(1))
			qp.ListEntry(la, qp.Int(2))
			qp.ListEntry(la, qp.Int(3))
		})
		require.NoError(t, err)

		sel, err := Parse(`.[5:2]`)
		require.NoError(t, err)

		res, err := sel.Select(node)
		require.NoError(t, err)
		require.NotEmpty(t, res)
		require.Equal(t, int64(0), res.Length())

		_, err = res.LookupByIndex(0)
		require.ErrorIs(t, err, datamodel.ErrNotExists{}) // assert empty result for backward slice
	})

	t.Run("slice with negative index", func(t *testing.T) {
		node, err := qp.BuildList(basicnode.Prototype.Any, 3, func(la datamodel.ListAssembler) {
			qp.ListEntry(la, qp.Int(1))
			qp.ListEntry(la, qp.Int(2))
			qp.ListEntry(la, qp.Int(3))
		})
		require.NoError(t, err)

		sel, err := Parse(`.[0:-1]`)
		require.NoError(t, err)

		res, err := sel.Select(node)
		require.NoError(t, err)
		require.NotEmpty(t, res)

		val, err := res.LookupByIndex(1)
		require.NoError(t, err)
		require.Equal(t, 2, int(must.Int(val))) // Assert sliced value at index 1
	})

	t.Run("slice on bytes", func(t *testing.T) {
		sel, err := Parse(`.[1:3]`)
		require.NoError(t, err)

		node := basicnode.NewBytes([]byte{0x01, 0x02, 0x03, 0x04, 0x05})
		res, err := sel.Select(node)
		require.NoError(t, err)
		require.NotEmpty(t, res)

		bytes, err := res.AsBytes()
		require.NoError(t, err)
		require.Equal(t, []byte{0x02, 0x03}, bytes) // assert sliced bytes
	})

	t.Run("index on bytes", func(t *testing.T) {
		sel, err := Parse(`.[2]`)
		require.NoError(t, err)

		node := basicnode.NewBytes([]byte{0x01, 0x02, 0x03, 0x04, 0x05})
		res, err := sel.Select(node)
		require.NoError(t, err)
		require.NotEmpty(t, res)

		val, err := res.AsInt()
		require.NoError(t, err)
		require.Equal(t, int64(0x03), val) // assert indexed byte value
	})

	t.Run("out of bounds slicing on bytes", func(t *testing.T) {
		sel, err := Parse(`.[10:20]`)
		require.NoError(t, err)

		node := basicnode.NewBytes([]byte{0x01, 0x02, 0x03})
		res, err := sel.Select(node)
		require.NoError(t, err)
		require.NotNil(t, res)

		bytes, err := res.AsBytes()
		require.NoError(t, err)
		require.Empty(t, bytes) // assert empty result for out of bounds slice
	})

	t.Run("out of bounds indexing on bytes", func(t *testing.T) {
		sel, err := Parse(`.[10]`)
		require.NoError(t, err)

		node := basicnode.NewBytes([]byte{0x01, 0x02, 0x03})
		_, err = sel.Select(node)
		require.Error(t, err)
		require.Contains(t, err.Error(), "can not resolve path: .10") // assert error for out of bounds index
	})
}

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
		if err != nil && !errors.As(err, &resolutionErr{}) {
			// not normal, we should only have resolution errors
			t.Fatal(err)
		}
	})
}

func TestResolveSliceIndices(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int64
		length    int64
		wantStart int64
		wantEnd   int64
	}{
		{
			name:      "normal case",
			slice:     []int64{1, 3},
			length:    5,
			wantStart: 1,
			wantEnd:   3,
		},
		{
			name:      "negative indices",
			slice:     []int64{-2, -1},
			length:    5,
			wantStart: 3,
			wantEnd:   4,
		},
		{
			name:      "overflow protection negative start",
			slice:     []int64{math.MinInt64, 3},
			length:    5,
			wantStart: 0,
			wantEnd:   3,
		},
		{
			name:      "overflow protection negative end",
			slice:     []int64{0, math.MinInt64},
			length:    5,
			wantStart: 0,
			wantEnd:   0,
		},
		{
			name:      "max bounds",
			slice:     []int64{0, math.MaxInt64},
			length:    5,
			wantStart: 0,
			wantEnd:   5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start, end := resolveSliceIndices(tt.slice, tt.length)
			require.Equal(t, tt.wantStart, start)
			require.Equal(t, tt.wantEnd, end)
		})
	}
}
