package selector

import (
	"fmt"
	"testing"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/must"
	"github.com/ipld/go-ipld-prime/node/bindnode"
	"github.com/ipld/go-ipld-prime/printer"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	t.Run("identity", func(t *testing.T) {
		sel, err := Parse(".")
		require.NoError(t, err)
		require.Equal(t, 1, len(sel))
		require.True(t, sel[0].Identity())
		require.False(t, sel[0].Optional())
		require.False(t, sel[0].Iterator())
		require.Empty(t, sel[0].Slice())
		require.Empty(t, sel[0].Field())
		require.Empty(t, sel[0].Index())
	})

	t.Run("field", func(t *testing.T) {
		sel, err := Parse(".foo")
		require.NoError(t, err)
		require.Equal(t, 1, len(sel))
		require.False(t, sel[0].Identity())
		require.False(t, sel[0].Optional())
		require.False(t, sel[0].Iterator())
		require.Empty(t, sel[0].Slice())
		require.Equal(t, sel[0].Field(), "foo")
		require.Empty(t, sel[0].Index())
	})

	t.Run("explicit field", func(t *testing.T) {
		sel, err := Parse(`.["foo"]`)
		require.NoError(t, err)
		require.Equal(t, 2, len(sel))
		require.True(t, sel[0].Identity())
		require.False(t, sel[0].Optional())
		require.False(t, sel[0].Iterator())
		require.Empty(t, sel[0].Slice())
		require.Empty(t, sel[0].Field())
		require.Empty(t, sel[0].Index())
		require.False(t, sel[1].Identity())
		require.False(t, sel[1].Optional())
		require.False(t, sel[1].Iterator())
		require.Empty(t, sel[1].Slice())
		require.Equal(t, sel[1].Field(), "foo")
		require.Empty(t, sel[1].Index())
	})

	t.Run("index", func(t *testing.T) {
		sel, err := Parse(".[138]")
		require.NoError(t, err)
		require.Equal(t, 2, len(sel))
		require.True(t, sel[0].Identity())
		require.False(t, sel[0].Optional())
		require.False(t, sel[0].Iterator())
		require.Empty(t, sel[0].Slice())
		require.Empty(t, sel[0].Field())
		require.Empty(t, sel[0].Index())
		require.False(t, sel[1].Identity())
		require.False(t, sel[1].Optional())
		require.False(t, sel[1].Iterator())
		require.Empty(t, sel[1].Slice())
		require.Empty(t, sel[1].Field())
		require.Equal(t, sel[1].Index(), 138)
	})

	t.Run("negative index", func(t *testing.T) {
		sel, err := Parse(".[-138]")
		require.NoError(t, err)
		require.Equal(t, 2, len(sel))
		require.True(t, sel[0].Identity())
		require.False(t, sel[0].Optional())
		require.False(t, sel[0].Iterator())
		require.Empty(t, sel[0].Slice())
		require.Empty(t, sel[0].Field())
		require.Empty(t, sel[0].Index())
		require.False(t, sel[1].Identity())
		require.False(t, sel[1].Optional())
		require.False(t, sel[1].Iterator())
		require.Empty(t, sel[1].Slice())
		require.Empty(t, sel[1].Field())
		require.Equal(t, sel[1].Index(), -138)
	})

	t.Run("iterator", func(t *testing.T) {
		sel, err := Parse(".[]")
		require.NoError(t, err)
		require.Equal(t, 2, len(sel))
		require.True(t, sel[0].Identity())
		require.False(t, sel[0].Optional())
		require.False(t, sel[0].Iterator())
		require.Empty(t, sel[0].Slice())
		require.Empty(t, sel[0].Field())
		require.Empty(t, sel[0].Index())
		require.False(t, sel[1].Identity())
		require.False(t, sel[1].Optional())
		require.True(t, sel[1].Iterator())
		require.Empty(t, sel[1].Slice())
		require.Empty(t, sel[1].Field())
		require.Empty(t, sel[1].Index())
	})

	t.Run("optional field", func(t *testing.T) {
		sel, err := Parse(".foo?")
		require.NoError(t, err)
		require.Equal(t, 1, len(sel))
		require.False(t, sel[0].Identity())
		require.True(t, sel[0].Optional())
		require.False(t, sel[0].Iterator())
		require.Empty(t, sel[0].Slice())
		require.Equal(t, sel[0].Field(), "foo")
		require.Empty(t, sel[0].Index())
	})

	t.Run("optional explicit field", func(t *testing.T) {
		sel, err := Parse(`.["foo"]?`)
		require.NoError(t, err)
		require.Equal(t, 2, len(sel))
		require.True(t, sel[0].Identity())
		require.False(t, sel[0].Optional())
		require.False(t, sel[0].Iterator())
		require.Empty(t, sel[0].Slice())
		require.Empty(t, sel[0].Field())
		require.Empty(t, sel[0].Index())
		require.False(t, sel[1].Identity())
		require.True(t, sel[1].Optional())
		require.False(t, sel[1].Iterator())
		require.Empty(t, sel[1].Slice())
		require.Equal(t, sel[1].Field(), "foo")
		require.Empty(t, sel[1].Index())
	})

	t.Run("optional index", func(t *testing.T) {
		sel, err := Parse(".[138]?")
		require.NoError(t, err)
		require.Equal(t, 2, len(sel))
		require.True(t, sel[0].Identity())
		require.False(t, sel[0].Optional())
		require.False(t, sel[0].Iterator())
		require.Empty(t, sel[0].Slice())
		require.Empty(t, sel[0].Field())
		require.Empty(t, sel[0].Index())
		require.False(t, sel[1].Identity())
		require.True(t, sel[1].Optional())
		require.False(t, sel[1].Iterator())
		require.Empty(t, sel[1].Slice())
		require.Empty(t, sel[1].Field())
		require.Equal(t, sel[1].Index(), 138)
	})

	t.Run("optional iterator", func(t *testing.T) {
		sel, err := Parse(".[]?")
		require.NoError(t, err)
		require.Equal(t, 2, len(sel))
		require.True(t, sel[0].Identity())
		require.False(t, sel[0].Optional())
		require.False(t, sel[0].Iterator())
		require.Empty(t, sel[0].Slice())
		require.Empty(t, sel[0].Field())
		require.Empty(t, sel[0].Index())
		require.False(t, sel[1].Identity())
		require.True(t, sel[1].Optional())
		require.True(t, sel[1].Iterator())
		require.Empty(t, sel[1].Slice())
		require.Empty(t, sel[1].Field())
		require.Empty(t, sel[1].Index())
	})

	t.Run("nesting", func(t *testing.T) {
		str := `.foo.["bar"].[138]?.baz[1:]`
		sel, err := Parse(str)
		require.NoError(t, err)
		printSegments(sel)
		require.Equal(t, str, sel.String())
		require.Equal(t, 7, len(sel))
		require.False(t, sel[0].Identity())
		require.False(t, sel[0].Optional())
		require.False(t, sel[0].Iterator())
		require.Empty(t, sel[0].Slice())
		require.Equal(t, sel[0].Field(), "foo")
		require.Empty(t, sel[0].Index())
		require.True(t, sel[1].Identity())
		require.False(t, sel[1].Optional())
		require.False(t, sel[1].Iterator())
		require.Empty(t, sel[1].Slice())
		require.Empty(t, sel[1].Field())
		require.Empty(t, sel[1].Index())
		require.False(t, sel[2].Identity())
		require.False(t, sel[2].Optional())
		require.False(t, sel[2].Iterator())
		require.Empty(t, sel[2].Slice())
		require.Equal(t, sel[2].Field(), "bar")
		require.Empty(t, sel[2].Index())
		require.True(t, sel[3].Identity())
		require.False(t, sel[3].Optional())
		require.False(t, sel[3].Iterator())
		require.Empty(t, sel[3].Slice())
		require.Empty(t, sel[3].Field())
		require.Empty(t, sel[3].Index())
		require.False(t, sel[4].Identity())
		require.True(t, sel[4].Optional())
		require.False(t, sel[4].Iterator())
		require.Empty(t, sel[4].Slice())
		require.Empty(t, sel[4].Field())
		require.Equal(t, sel[4].Index(), 138)
		require.False(t, sel[5].Identity())
		require.False(t, sel[5].Optional())
		require.False(t, sel[5].Iterator())
		require.Empty(t, sel[5].Slice())
		require.Equal(t, sel[5].Field(), "baz")
		require.Empty(t, sel[5].Index())
		require.False(t, sel[6].Identity())
		require.False(t, sel[6].Optional())
		require.False(t, sel[6].Iterator())
		require.Equal(t, sel[6].Slice(), []int{1})
		require.Empty(t, sel[6].Field())
		require.Empty(t, sel[6].Index())
	})

	t.Run("non dotted", func(t *testing.T) {
		_, err := Parse("foo")
		require.NotNil(t, err)
		fmt.Println(err)
	})

	t.Run("non quoted", func(t *testing.T) {
		_, err := Parse(".[foo]")
		require.NotNil(t, err)
		fmt.Println(err)
	})
}

func printSegments(s Selector) {
	for i, seg := range s {
		fmt.Printf("%d: %s\n", i, seg.String())
	}
}

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

		one, many, err := Select(sel, anode)
		require.NoError(t, err)
		require.NotEmpty(t, one)
		require.Empty(t, many)

		fmt.Println(printer.Sprint(one))

		age := must.Int(must.Node(one.LookupByString("age")))
		require.Equal(t, int64(alice.Age), age)
	})

	t.Run("nested property", func(t *testing.T) {
		sel, err := Parse(".name.first")
		require.NoError(t, err)

		one, many, err := Select(sel, anode)
		require.NoError(t, err)
		require.NotEmpty(t, one)
		require.Empty(t, many)

		fmt.Println(printer.Sprint(one))

		name := must.String(one)
		require.Equal(t, alice.Name.First, name)

		one, many, err = Select(sel, bnode)
		require.NoError(t, err)
		require.NotEmpty(t, one)
		require.Empty(t, many)

		fmt.Println(printer.Sprint(one))

		name = must.String(one)
		require.Equal(t, bob.Name.First, name)
	})

	t.Run("optional nested property", func(t *testing.T) {
		sel, err := Parse(".name.middle?")
		require.NoError(t, err)

		one, many, err := Select(sel, anode)
		require.NoError(t, err)
		require.NotEmpty(t, one)
		require.Empty(t, many)

		fmt.Println(printer.Sprint(one))

		name := must.String(one)
		require.Equal(t, *alice.Name.Middle, name)

		one, many, err = Select(sel, bnode)
		require.NoError(t, err)
		require.Empty(t, one)
		require.Empty(t, many)
	})

	t.Run("not exists", func(t *testing.T) {
		sel, err := Parse(".name.foo")
		require.NoError(t, err)

		one, many, err := Select(sel, anode)
		require.Error(t, err)
		require.Empty(t, one)
		require.Empty(t, many)

		fmt.Println(err)

		if _, ok := err.(ResolutionError); !ok {
			t.Fatalf("error was not a resolution error")
		}
	})

	t.Run("optional not exists", func(t *testing.T) {
		sel, err := Parse(".name.foo?")
		require.NoError(t, err)

		one, many, err := Select(sel, anode)
		require.NoError(t, err)
		require.Empty(t, one)
		require.Empty(t, many)
	})

	t.Run("iterator", func(t *testing.T) {
		sel, err := Parse(".interests[]")
		require.NoError(t, err)

		one, many, err := Select(sel, anode)
		require.NoError(t, err)
		require.Empty(t, one)
		require.NotEmpty(t, many)

		for _, n := range many {
			fmt.Println(printer.Sprint(n))
		}

		iname := must.String(must.Node(many[0].LookupByString("name")))
		require.Equal(t, alice.Interests[0].Name, iname)

		iname = must.String(must.Node(many[1].LookupByString("name")))
		require.Equal(t, alice.Interests[1].Name, iname)
	})

	t.Run("map iterator", func(t *testing.T) {
		sel, err := Parse(".interests[0][]")
		require.NoError(t, err)

		one, many, err := Select(sel, anode)
		require.NoError(t, err)
		require.Empty(t, one)
		require.NotEmpty(t, many)

		for _, n := range many {
			fmt.Println(printer.Sprint(n))
		}

		require.Equal(t, alice.Interests[0].Name, must.String(many[0]))
		require.Equal(t, alice.Interests[0].Experience, int(must.Int(many[2])))
	})
}
