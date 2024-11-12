package literal

import (
	"testing"

	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime/datamodel"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/ipld/go-ipld-prime/printer"
	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	n, err := List([]int{1, 2, 3})
	require.NoError(t, err)
	require.Equal(t, datamodel.Kind_List, n.Kind())
	require.Equal(t, int64(3), n.Length())
	require.Equal(t, `list{
	0: int{1}
	1: int{2}
	2: int{3}
}`, printer.Sprint(n))

	n, err = List([][]int{{1, 2, 3}, {4, 5, 6}})
	require.NoError(t, err)
	require.Equal(t, datamodel.Kind_List, n.Kind())
	require.Equal(t, int64(2), n.Length())
	require.Equal(t, `list{
	0: list{
		0: int{1}
		1: int{2}
		2: int{3}
	}
	1: list{
		0: int{4}
		1: int{5}
		2: int{6}
	}
}`, printer.Sprint(n))
}

func TestMap(t *testing.T) {
	n, err := Map(map[string]any{
		"bool":   true,
		"string": "foobar",
		"bytes":  []byte{1, 2, 3, 4},
		"int":    1234,
		"uint":   uint(12345),
		"float":  1.45,
		"slice":  []int{1, 2, 3},
		"array":  [2]int{1, 2},
		"map": map[string]any{
			"foo": "bar",
			"foofoo": map[string]string{
				"barbar": "foo",
			},
		},
		"link": cid.MustParse("bafzbeigai3eoy2ccc7ybwjfz5r3rdxqrinwi4rwytly24tdbh6yk7zslrm"),
	})
	require.NoError(t, err)

	v, err := n.LookupByString("bool")
	require.NoError(t, err)
	require.Equal(t, datamodel.Kind_Bool, v.Kind())
	require.Equal(t, true, must(v.AsBool()))

	v, err = n.LookupByString("string")
	require.NoError(t, err)
	require.Equal(t, datamodel.Kind_String, v.Kind())
	require.Equal(t, "foobar", must(v.AsString()))

	v, err = n.LookupByString("bytes")
	require.NoError(t, err)
	require.Equal(t, datamodel.Kind_Bytes, v.Kind())
	require.Equal(t, []byte{1, 2, 3, 4}, must(v.AsBytes()))

	v, err = n.LookupByString("int")
	require.NoError(t, err)
	require.Equal(t, datamodel.Kind_Int, v.Kind())
	require.Equal(t, int64(1234), must(v.AsInt()))

	v, err = n.LookupByString("uint")
	require.NoError(t, err)
	require.Equal(t, datamodel.Kind_Int, v.Kind())
	require.Equal(t, int64(12345), must(v.AsInt()))

	v, err = n.LookupByString("float")
	require.NoError(t, err)
	require.Equal(t, datamodel.Kind_Float, v.Kind())
	require.Equal(t, 1.45, must(v.AsFloat()))

	v, err = n.LookupByString("slice")
	require.NoError(t, err)
	require.Equal(t, datamodel.Kind_List, v.Kind())
	require.Equal(t, int64(3), v.Length())
	require.Equal(t, `list{
	0: int{1}
	1: int{2}
	2: int{3}
}`, printer.Sprint(v))

	v, err = n.LookupByString("array")
	require.NoError(t, err)
	require.Equal(t, datamodel.Kind_List, v.Kind())
	require.Equal(t, int64(2), v.Length())
	require.Equal(t, `list{
	0: int{1}
	1: int{2}
}`, printer.Sprint(v))

	v, err = n.LookupByString("map")
	require.NoError(t, err)
	require.Equal(t, datamodel.Kind_Map, v.Kind())
	require.Equal(t, int64(2), v.Length())
	require.Equal(t, `map{
	string{"foo"}: string{"bar"}
	string{"foofoo"}: map{
		string{"barbar"}: string{"foo"}
	}
}`, printer.Sprint(v))

	v, err = n.LookupByString("link")
	require.NoError(t, err)
	require.Equal(t, datamodel.Kind_Link, v.Kind())
	asLink, err := v.AsLink()
	require.NoError(t, err)
	require.True(t, asLink.(cidlink.Link).Equals(cid.MustParse("bafzbeigai3eoy2ccc7ybwjfz5r3rdxqrinwi4rwytly24tdbh6yk7zslrm")))
}

func TestAny(t *testing.T) {
	data := map[string]any{
		"bool":   true,
		"string": "foobar",
		"bytes":  []byte{1, 2, 3, 4},
		"int":    1234,
		"uint":   uint(12345),
		"float":  1.45,
		"slice":  []int{1, 2, 3},
		"array":  [2]int{1, 2},
		"map": map[string]any{
			"foo": "bar",
			"foofoo": map[string]string{
				"barbar": "foo",
			},
		},
		"link": cid.MustParse("bafzbeigai3eoy2ccc7ybwjfz5r3rdxqrinwi4rwytly24tdbh6yk7zslrm"),
		"func": func() {},
	}

	v, err := Any(data["bool"])
	require.NoError(t, err)
	require.Equal(t, datamodel.Kind_Bool, v.Kind())
	require.Equal(t, true, must(v.AsBool()))

	v, err = Any(data["string"])
	require.NoError(t, err)
	require.Equal(t, datamodel.Kind_String, v.Kind())
	require.Equal(t, "foobar", must(v.AsString()))

	v, err = Any(data["bytes"])
	require.NoError(t, err)
	require.Equal(t, datamodel.Kind_Bytes, v.Kind())
	require.Equal(t, []byte{1, 2, 3, 4}, must(v.AsBytes()))

	v, err = Any(data["int"])
	require.NoError(t, err)
	require.Equal(t, datamodel.Kind_Int, v.Kind())
	require.Equal(t, int64(1234), must(v.AsInt()))

	v, err = Any(data["uint"])
	require.NoError(t, err)
	require.Equal(t, datamodel.Kind_Int, v.Kind())
	require.Equal(t, int64(12345), must(v.AsInt()))

	v, err = Any(data["float"])
	require.NoError(t, err)
	require.Equal(t, datamodel.Kind_Float, v.Kind())
	require.Equal(t, 1.45, must(v.AsFloat()))

	v, err = Any(data["slice"])
	require.NoError(t, err)
	require.Equal(t, datamodel.Kind_List, v.Kind())
	require.Equal(t, int64(3), v.Length())
	require.Equal(t, `list{
	0: int{1}
	1: int{2}
	2: int{3}
}`, printer.Sprint(v))

	v, err = Any(data["array"])
	require.NoError(t, err)
	require.Equal(t, datamodel.Kind_List, v.Kind())
	require.Equal(t, int64(2), v.Length())
	require.Equal(t, `list{
	0: int{1}
	1: int{2}
}`, printer.Sprint(v))

	v, err = Any(data["map"])
	require.NoError(t, err)
	require.Equal(t, datamodel.Kind_Map, v.Kind())
	require.Equal(t, int64(2), v.Length())
	require.Equal(t, `map{
	string{"foo"}: string{"bar"}
	string{"foofoo"}: map{
		string{"barbar"}: string{"foo"}
	}
}`, printer.Sprint(v))

	v, err = Any(data["link"])
	require.NoError(t, err)
	require.Equal(t, datamodel.Kind_Link, v.Kind())
	asLink, err := v.AsLink()
	require.NoError(t, err)
	require.True(t, asLink.(cidlink.Link).Equals(cid.MustParse("bafzbeigai3eoy2ccc7ybwjfz5r3rdxqrinwi4rwytly24tdbh6yk7zslrm")))

	v, err = Any(data["func"])
	require.Error(t, err)
}

func BenchmarkAny(b *testing.B) {
	b.Run("bool", func(b *testing.B) {
		b.ReportAllocs()

		for n := 0; n < b.N; n++ {
			_, _ = Any(true)
		}
	})

	b.Run("string", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			_, _ = Any("foobar")
		}
	})

	b.Run("bytes", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			_, _ = Any([]byte{1, 2, 3, 4})
		}
	})

	b.Run("map", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			_, _ = Any(map[string]any{
				"foo": "bar",
				"foofoo": map[string]string{
					"barbar": "foo",
				},
			})
		}
	})
}

func must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}
