package envelope_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/ucan-wg/go-ucan/internal/envelope"
	"gotest.tools/v3/golden"
)

func TestDecode(t *testing.T) {
	t.Parallel()

	t.Run("via FromDagCbor", func(t *testing.T) {
		t.Parallel()

		data := golden.Get(t, "example.dagcbor")

		tkn, _, err := envelope.FromDagCbor[*Example](data)
		require.NoError(t, err)
		assert.Equal(t, exampleGreeting, tkn.Hello)
		assert.Equal(t, exampleDID, tkn.Issuer)
	})

	t.Run("via FromDagJson", func(t *testing.T) {
		t.Parallel()

		data := golden.Get(t, "example.dagjson")

		tkn, _, err := envelope.FromDagJson[*Example](data)
		require.NoError(t, err)
		assert.Equal(t, exampleGreeting, tkn.Hello)
		assert.Equal(t, exampleDID, tkn.Issuer)
	})
}

func TestEncode(t *testing.T) {
	t.Parallel()

	t.Run("via ToDagCbor", func(t *testing.T) {
		t.Parallel()

		data, err := envelope.ToDagCbor(examplePrivKey(t), newExample(t))
		require.NoError(t, err)
		golden.AssertBytes(t, data, exampleDAGCBORFilename)
	})

	t.Run("via ToDagJson", func(t *testing.T) {
		t.Parallel()

		data, err := envelope.ToDagJson(examplePrivKey(t), newExample(t))
		require.NoError(t, err)
		golden.Assert(t, string(data), exampleDAGJSONFilename)
	})
}

func TestRoundtrip(t *testing.T) {
	t.Parallel()

	t.Run("via FromDagCbor/ToDagCbor", func(t *testing.T) {
		t.Parallel()

		dataIn := golden.Get(t, exampleDAGCBORFilename)

		tkn, id, err := envelope.FromDagCbor[*Example](dataIn)
		require.NoError(t, err)
		assert.Equal(t, exampleGreeting, tkn.Hello)
		assert.Equal(t, exampleDID, tkn.Issuer)
		assert.Equal(t, exampleCID, envelope.CIDToBase58BTC(id))

		dataOut, err := envelope.ToDagCbor(examplePrivKey(t), newExample(t))
		require.NoError(t, err)
		assert.Equal(t, dataIn, dataOut)
	})

	t.Run("via FromDagCborReader/ToDagCborWriter", func(t *testing.T) {
		t.Parallel()

		data := golden.Get(t, exampleDAGCBORFilename)

		tkn, id, err := envelope.FromDagCborReader[*Example](bytes.NewReader(data))
		require.NoError(t, err)
		assert.Equal(t, exampleGreeting, tkn.Hello)
		assert.Equal(t, exampleDID, tkn.Issuer)
		assert.Equal(t, exampleCID, envelope.CIDToBase58BTC(id))

		w := &bytes.Buffer{}
		require.NoError(t, envelope.ToDagCborWriter(w, examplePrivKey(t), newExample(t)))
		assert.Equal(t, data, w.Bytes())
	})

	t.Run("via FromDagJson/ToDagJson", func(t *testing.T) {
		t.Parallel()

		dataIn := golden.Get(t, exampleDAGJSONFilename)

		tkn, id, err := envelope.FromDagJson[*Example](dataIn)
		require.NoError(t, err)
		assert.Equal(t, exampleGreeting, tkn.Hello)
		assert.Equal(t, exampleDID, tkn.Issuer)
		assert.Equal(t, exampleCID, envelope.CIDToBase58BTC(id))

		dataOut, err := envelope.ToDagJson(examplePrivKey(t), newExample(t))
		require.NoError(t, err)
		assert.Equal(t, dataIn, dataOut)
	})

	t.Run("via FromDagJsonReader/ToDagJsonrWriter", func(t *testing.T) {
		t.Parallel()

		data := golden.Get(t, exampleDAGJSONFilename)

		tkn, id, err := envelope.FromDagJsonReader[*Example](bytes.NewReader(data))
		require.NoError(t, err)
		assert.Equal(t, exampleGreeting, tkn.Hello)
		assert.Equal(t, exampleDID, tkn.Issuer)
		assert.Equal(t, exampleCID, envelope.CIDToBase58BTC(id))

		w := &bytes.Buffer{}
		require.NoError(t, envelope.ToDagJsonWriter(w, examplePrivKey(t), newExample(t)))
		assert.Equal(t, data, w.Bytes())
	})
}

func TestFromIPLD_with_invalid_signature(t *testing.T) {
	t.Parallel()

	node := invalidNodeFromGolden(t)
	tkn, _, err := envelope.FromIPLD[*Example](node)
	assert.Nil(t, tkn)
	require.EqualError(t, err, "failed to verify the token's signature")
}
