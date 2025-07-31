package envelope_test

import (
	"bytes"
	"crypto/sha256"
	_ "embed"
	"encoding/base64"
	"os"
	"testing"

	_ "github.com/MetaMask/go-did-it/verifiers/did-key"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ucan-wg/go-ucan/token/internal/envelope"
)

func TestDecode(t *testing.T) {
	t.Parallel()

	t.Run("via FromDagCbor", func(t *testing.T) {
		t.Parallel()

		tkn, err := envelope.FromDagCbor[*Example](exampleDagCbor)
		require.NoError(t, err)
		assert.Equal(t, exampleGreeting, tkn.Hello)
		assert.Equal(t, exampleDID, tkn.Issuer)
	})

	t.Run("via FromDagJson", func(t *testing.T) {
		t.Parallel()

		tkn, err := envelope.FromDagJson[*Example](exampleDagJson)
		require.NoError(t, err)
		assert.Equal(t, exampleGreeting, tkn.Hello)
		assert.Equal(t, exampleDID, tkn.Issuer)
	})
}

func TestEncode(t *testing.T) {
	t.Parallel()

	t.Run("via ToDagCbor", func(t *testing.T) {
		t.Parallel()

		data, err := envelope.ToDagCbor(examplePrivKey(t), newExample())
		require.NoError(t, err)
		require.Equal(t, exampleDagCbor, data)
	})

	t.Run("via ToDagJson", func(t *testing.T) {
		t.Parallel()

		data, err := envelope.ToDagJson(examplePrivKey(t), newExample())
		require.NoError(t, err)
		require.Equal(t, exampleDagJson, data)
	})
}

func TestRoundtrip(t *testing.T) {
	t.Parallel()

	t.Run("via FromDagCbor/ToDagCbor", func(t *testing.T) {
		t.Parallel()

		dataIn := exampleDagCbor

		tkn, err := envelope.FromDagCbor[*Example](dataIn)
		require.NoError(t, err)
		assert.Equal(t, exampleGreeting, tkn.Hello)
		assert.Equal(t, exampleDID, tkn.Issuer)

		dataOut, err := envelope.ToDagCbor(examplePrivKey(t), newExample())
		require.NoError(t, err)
		assert.Equal(t, dataIn, dataOut)
	})

	t.Run("via FromDagCborReader/ToDagCborWriter", func(t *testing.T) {
		t.Parallel()

		data := exampleDagCbor

		tkn, err := envelope.FromDagCborReader[*Example](bytes.NewReader(data))
		require.NoError(t, err)
		assert.Equal(t, exampleGreeting, tkn.Hello)
		assert.Equal(t, exampleDID, tkn.Issuer)

		w := &bytes.Buffer{}
		require.NoError(t, envelope.ToDagCborWriter(w, examplePrivKey(t), newExample()))
		assert.Equal(t, data, w.Bytes())
	})

	t.Run("via FromDagJson/ToDagJson", func(t *testing.T) {
		t.Parallel()

		dataIn := exampleDagJson

		tkn, err := envelope.FromDagJson[*Example](dataIn)
		require.NoError(t, err)
		assert.Equal(t, exampleGreeting, tkn.Hello)
		assert.Equal(t, exampleDID, tkn.Issuer)

		dataOut, err := envelope.ToDagJson(examplePrivKey(t), newExample())
		require.NoError(t, err)
		assert.Equal(t, dataIn, dataOut)
	})

	t.Run("via FromDagJsonReader/ToDagJsonrWriter", func(t *testing.T) {
		t.Parallel()

		data := exampleDagJson

		tkn, err := envelope.FromDagJsonReader[*Example](bytes.NewReader(data))
		require.NoError(t, err)
		assert.Equal(t, exampleGreeting, tkn.Hello)
		assert.Equal(t, exampleDID, tkn.Issuer)

		w := &bytes.Buffer{}
		require.NoError(t, envelope.ToDagJsonWriter(w, examplePrivKey(t), newExample()))
		assert.Equal(t, data, w.Bytes())
	})
}

func TestFromIPLD_with_invalid_signature(t *testing.T) {
	t.Parallel()

	node := nodeWithInvalidSignature(t)
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

	node, err := ipld.Decode(exampleDagCbor, dagcbor.Decode)
	require.NoError(t, err)

	expSig, err := base64.RawStdEncoding.DecodeString("+xUwgl/5VZcTxx6iePmkrIaZAlxuelHTbeQ5lQIgIV3ZgHS+Jf5BUERB0fvmFfiIfa5A3yMPfEA/7rswYsRRCg")
	require.NoError(t, err)

	info, err := envelope.Inspect(node)
	require.NoError(t, err)
	assert.Equal(t, expSig, info.Signature)
	assert.Equal(t, "ucan/example@v1.0.0-rc.1", info.Tag)
	assert.Equal(t, []byte{0x34, 0x1, 0xed, 0x1, 0xed, 0x1, 0x13, 0x71}, info.VarsigBytes)
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
