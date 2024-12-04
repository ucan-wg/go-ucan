package args_test

import (
	"maps"
	"sync"
	"testing"

	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	"github.com/ipld/go-ipld-prime/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ucan-wg/go-ucan/pkg/args"
	"github.com/ucan-wg/go-ucan/pkg/policy/limits"
	"github.com/ucan-wg/go-ucan/pkg/policy/literal"
)

func TestArgs(t *testing.T) {
	t.Parallel()

	const (
		intKey    = "intKey"
		mapKey    = "mapKey"
		nilKey    = "nilKey"
		boolKey   = "boolKey"
		linkKey   = "linkKey"
		listKey   = "listKey"
		nodeKey   = "nodeKey"
		uintKey   = "uintKey"
		bytesKey  = "bytesKey"
		floatKey  = "floatKey"
		stringKey = "stringKey"
	)

	const (
		expIntVal    = int64(-42)
		expBoolVal   = true
		expUintVal   = uint(42)
		expStringVal = "stringVal"
	)

	var (
		expMapVal = map[string]string{"keyOne": "valOne", "keyTwo": "valTwo"}
		// expNilVal = (map[string]string)(nil)
		expLinkVal  = cid.MustParse("bafzbeigai3eoy2ccc7ybwjfz5r3rdxqrinwi4rwytly24tdbh6yk7zslrm")
		expListVal  = []string{"elem1", "elem2", "elem3"}
		expNodeVal  = literal.String("nodeVal")
		expBytesVal = []byte{0xde, 0xad, 0xbe, 0xef}
		expFloatVal = 42.0
	)

	argsIn := args.New()

	for _, a := range []struct {
		key string
		val any
	}{
		{key: intKey, val: expIntVal},
		{key: mapKey, val: expMapVal},
		// {key:  nilKey, val: expNilVal},
		{key: boolKey, val: expBoolVal},
		{key: linkKey, val: expLinkVal},
		{key: listKey, val: expListVal},
		{key: uintKey, val: expUintVal},
		{key: nodeKey, val: expNodeVal},
		{key: bytesKey, val: expBytesVal},
		{key: floatKey, val: expFloatVal},
		{key: stringKey, val: expStringVal},
	} {
		require.NoError(t, argsIn.Add(a.key, a.val))
	}

	// Round-trip to DAG-CBOR
	argsOut := roundTripThroughDAGCBOR(t, argsIn)
	assert.ElementsMatch(t, argsIn.Keys, argsOut.Keys)
	assert.Equal(t, argsIn.Values, argsOut.Values)

	actMapVal := map[string]string{}
	mit := argsOut.Values[mapKey].MapIterator()

	for !mit.Done() {
		k, v, err := mit.Next()
		require.NoError(t, err)
		ks := must(k.AsString())
		vs := must(v.AsString())
		actMapVal[ks] = vs
	}

	actListVal := []string{}
	lit := argsOut.Values[listKey].ListIterator()

	for !lit.Done() {
		_, v, err := lit.Next()
		require.NoError(t, err)
		vs := must(v.AsString())

		actListVal = append(actListVal, vs)
	}

	assert.Equal(t, expIntVal, must(argsOut.Values[intKey].AsInt()))
	assert.Equal(t, expMapVal, actMapVal) // TODO: special accessor
	// TODO: the nil map comes back empty (but the right type)
	// assert.Equal(t, expNilVal, actNilVal)
	assert.Equal(t, expBoolVal, must(argsOut.Values[boolKey].AsBool()))
	assert.Equal(t, expLinkVal.String(), must(argsOut.Values[linkKey].AsLink()).(datamodel.Link).String()) // TODO: special accessor
	assert.Equal(t, expListVal, actListVal)                                                                // TODO: special accessor
	assert.Equal(t, expNodeVal, argsOut.Values[nodeKey])
	assert.Equal(t, expUintVal, uint(must(argsOut.Values[uintKey].AsInt())))
	assert.Equal(t, expBytesVal, must(argsOut.Values[bytesKey].AsBytes()))
	assert.Equal(t, expFloatVal, must(argsOut.Values[floatKey].AsFloat()))
	assert.Equal(t, expStringVal, must(argsOut.Values[stringKey].AsString()))
}

func TestArgs_Include(t *testing.T) {
	t.Parallel()

	argsIn := args.New()
	require.NoError(t, argsIn.Add("key1", "val1"))
	require.NoError(t, argsIn.Add("key2", "val2"))

	argsOther := args.New()
	require.NoError(t, argsOther.Add("key2", "valOther")) // This should not overwrite key2 above
	require.NoError(t, argsOther.Add("key3", "val3"))
	require.NoError(t, argsOther.Add("key4", "val4"))

	argsIn.Include(argsOther)

	assert.Len(t, argsIn.Values, 4)
	assert.Equal(t, "val1", must(argsIn.Values["key1"].AsString()))
	assert.Equal(t, "val2", must(argsIn.Values["key2"].AsString()))
	assert.Equal(t, "val3", must(argsIn.Values["key3"].AsString()))
	assert.Equal(t, "val4", must(argsIn.Values["key4"].AsString()))
}

func TestIterCloneEquals(t *testing.T) {
	a := args.New()

	require.NoError(t, a.Add("foo", "bar"))
	require.NoError(t, a.Add("baz", 1234))

	expected := map[string]ipld.Node{
		"foo": basicnode.NewString("bar"),
		"baz": basicnode.NewInt(1234),
	}

	// args -> iter
	require.Equal(t, expected, maps.Collect(a.Iter()))

	// readonly -> iter
	ro := a.ReadOnly()
	require.Equal(t, expected, maps.Collect(ro.Iter()))

	// args -> clone -> iter
	clone := a.Clone()
	require.Equal(t, expected, maps.Collect(clone.Iter()))

	// readonly -> WriteableClone -> iter
	wclone := ro.WriteableClone()
	require.Equal(t, expected, maps.Collect(wclone.Iter()))

	require.True(t, a.Equals(wclone))
	require.True(t, ro.Equals(wclone.ReadOnly()))
}

func TestInclude(t *testing.T) {
	a1 := args.New()

	require.NoError(t, a1.Add("samekey", "bar"))
	require.NoError(t, a1.Add("baz", 1234))

	a2 := args.New()

	require.NoError(t, a2.Add("samekey", "othervalue")) // check no overwrite
	require.NoError(t, a2.Add("otherkey", 1234))

	a1.Include(a2)

	require.Equal(t, map[string]ipld.Node{
		"samekey":  basicnode.NewString("bar"),
		"baz":      basicnode.NewInt(1234),
		"otherkey": basicnode.NewInt(1234),
	}, maps.Collect(a1.Iter()))
}

func TestArgsIntegerBounds(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		key     string
		val     int64
		wantErr string
	}{
		{
			name: "valid int",
			key:  "valid",
			val:  42,
		},
		{
			name: "max safe integer",
			key:  "max",
			val:  limits.MaxInt53,
		},
		{
			name: "min safe integer",
			key:  "min",
			val:  limits.MinInt53,
		},
		{
			name:    "exceeds max safe integer",
			key:     "tooBig",
			val:     limits.MaxInt53 + 1,
			wantErr: "exceeds safe integer bounds",
		},
		{
			name:    "below min safe integer",
			key:     "tooSmall",
			val:     limits.MinInt53 - 1,
			wantErr: "exceeds safe integer bounds",
		},
		{
			name:    "duplicate key",
			key:     "duplicate",
			val:     42,
			wantErr: "duplicate key",
		},
	}

	a := args.New()
	require.NoError(t, a.Add("duplicate", 1)) // tests duplicate key

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := a.Add(tt.key, tt.val)
			if tt.wantErr != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.wantErr)
			} else {
				require.NoError(t, err)
				val, err := a.GetNode(tt.key)
				require.NoError(t, err)
				i, err := val.AsInt()
				require.NoError(t, err)
				require.Equal(t, tt.val, i)
			}
		})
	}
}

const (
	argsSchema = "type Args { String : Any }"
	argsName   = "Args"
)

var (
	once      sync.Once
	ts        *schema.TypeSystem
	errSchema error
)

func argsType() schema.Type {
	once.Do(func() {
		ts, errSchema = ipld.LoadSchemaBytes([]byte(argsSchema))
	})
	if errSchema != nil {
		panic(errSchema)
	}

	return ts.TypeByName(argsName)
}

func roundTripThroughDAGCBOR(t *testing.T, argsIn *args.Args) *args.Args {
	t.Helper()

	node, err := argsIn.ToIPLD()
	require.NoError(t, err)

	data, err := ipld.Encode(node, dagcbor.Encode)
	require.NoError(t, err)

	var argsOut args.Args
	_, err = ipld.Unmarshal(data, dagcbor.Decode, &argsOut, argsType())
	require.NoError(t, err)

	return &argsOut
}

func must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}
