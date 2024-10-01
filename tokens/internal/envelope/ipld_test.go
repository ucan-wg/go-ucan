package envelope_test

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"os"
	"testing"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gotest.tools/v3/golden"

	"github.com/ucan-wg/go-ucan/tokens/internal/envelope"
)

func TestDecode(t *testing.T) {
	t.Parallel()

	t.Run("via FromDagCbor", func(t *testing.T) {
		t.Parallel()

		data := golden.Get(t, "example.dagcbor")

		tkn, err := envelope.FromDagCbor[*Example](data)
		require.NoError(t, err)
		assert.Equal(t, exampleGreeting, tkn.Hello)
		assert.Equal(t, exampleDID, tkn.Issuer)
	})

	t.Run("via FromDagJson", func(t *testing.T) {
		t.Parallel()

		data := golden.Get(t, "example.dagjson")

		tkn, err := envelope.FromDagJson[*Example](data)
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

		tkn, err := envelope.FromDagCbor[*Example](dataIn)
		require.NoError(t, err)
		assert.Equal(t, exampleGreeting, tkn.Hello)
		assert.Equal(t, exampleDID, tkn.Issuer)

		dataOut, err := envelope.ToDagCbor(examplePrivKey(t), newExample(t))
		require.NoError(t, err)
		assert.Equal(t, dataIn, dataOut)
	})

	t.Run("via FromDagCborReader/ToDagCborWriter", func(t *testing.T) {
		t.Parallel()

		data := golden.Get(t, exampleDAGCBORFilename)

		tkn, err := envelope.FromDagCborReader[*Example](bytes.NewReader(data))
		require.NoError(t, err)
		assert.Equal(t, exampleGreeting, tkn.Hello)
		assert.Equal(t, exampleDID, tkn.Issuer)

		w := &bytes.Buffer{}
		require.NoError(t, envelope.ToDagCborWriter(w, examplePrivKey(t), newExample(t)))
		assert.Equal(t, data, w.Bytes())
	})

	t.Run("via FromDagJson/ToDagJson", func(t *testing.T) {
		t.Parallel()

		dataIn := golden.Get(t, exampleDAGJSONFilename)

		tkn, err := envelope.FromDagJson[*Example](dataIn)
		require.NoError(t, err)
		assert.Equal(t, exampleGreeting, tkn.Hello)
		assert.Equal(t, exampleDID, tkn.Issuer)

		dataOut, err := envelope.ToDagJson(examplePrivKey(t), newExample(t))
		require.NoError(t, err)
		assert.Equal(t, dataIn, dataOut)
	})

	t.Run("via FromDagJsonReader/ToDagJsonrWriter", func(t *testing.T) {
		t.Parallel()

		data := golden.Get(t, exampleDAGJSONFilename)

		tkn, err := envelope.FromDagJsonReader[*Example](bytes.NewReader(data))
		require.NoError(t, err)
		assert.Equal(t, exampleGreeting, tkn.Hello)
		assert.Equal(t, exampleDID, tkn.Issuer)

		w := &bytes.Buffer{}
		require.NoError(t, envelope.ToDagJsonWriter(w, examplePrivKey(t), newExample(t)))
		assert.Equal(t, data, w.Bytes())
	})
}

func TestFromIPLD_with_invalid_signature(t *testing.T) {
	t.Parallel()

	node := invalidNodeFromGolden(t)
	tkn, err := envelope.FromIPLD[*Example](node)
	assert.Nil(t, tkn)
	require.EqualError(t, err, "failed to verify the token's signature")
}

func TestHash(t *testing.T) {
	t.Parallel()

	msg := []byte("this is a test")

	hash1 := sha256.Sum256(msg)

	hasher := sha256.New()

	for _, b := range msg {
		hasher.Write([]byte{b})
	}

	hash2 := hasher.Sum(nil)
	hash3 := hasher.Sum(nil)

	require.Equal(t, hash1[:], hash2)
	require.Equal(t, hash1[:], hash3)
}

func TestInspect(t *testing.T) {
	t.Parallel()

	data := golden.Get(t, "example.dagcbor")
	node, err := ipld.Decode(data, dagcbor.Decode)
	require.NoError(t, err)

	expSig, err := base64.RawStdEncoding.DecodeString("fPqfwL3iFpbw9SvBiq0DIbUurv9o6c36R08tC/yslGrJcwV51ghzWahxdetpEf6T5LCszXX9I/K8khvnmAxjAg")
	require.NoError(t, err)

	info, err := envelope.Inspect(node)
	require.NoError(t, err)
	assert.Equal(t, expSig, info.Signature)
	assert.Equal(t, "ucan/example@v1.0.0-rc.1", info.Tag)
	assert.Equal(t, []byte{0x34, 0xed, 0x1, 0x71}, info.VarsigHeader)
}

func FuzzInspect(f *testing.F) {
	data, err := os.ReadFile("testdata/example.dagcbor")
	require.NoError(f, err)

	f.Add(data)

	f.Fuzz(func(t *testing.T, data []byte) {
		node, err := ipld.Decode(data, dagcbor.Decode)
		if err != nil {
			t.Skip()
		}
		_, err = envelope.Inspect(node)
		if err != nil {
			t.Skip()
		}
	})
}

func FuzzFindTag(f *testing.F) {
	data, err := os.ReadFile("testdata/example.dagcbor")
	require.NoError(f, err)

	f.Add(data)

	f.Fuzz(func(t *testing.T, data []byte) {
		node, err := ipld.Decode(data, dagcbor.Decode)
		if err != nil {
			t.Skip()
		}
		_, err = envelope.FindTag(node)
		if err != nil {
			t.Skip()
		}
	})
}
