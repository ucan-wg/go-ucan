package args_test

import (
	"sync"
	"testing"

	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ucan-wg/go-ucan/pkg/args"
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

const (
	argsSchema = "type Args { String : Any }"
	argsName   = "Args"
)

var (
	once sync.Once
	ts   *schema.TypeSystem
	err  error
)

func argsType() schema.Type {
	once.Do(func() {
		ts, err = ipld.LoadSchemaBytes([]byte(argsSchema))
	})
	if err != nil {
		panic(err)
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
