package envelope_test

import (
	"io"
	"testing"

	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multicodec"
	"github.com/multiformats/go-multihash"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gotest.tools/v3/golden"

	"github.com/ucan-wg/go-ucan/internal/envelope"
)

func TestCidFromBytes(t *testing.T) {
	t.Parallel()

	expData := golden.Get(t, "example.dagcbor")
	expHash, err := multihash.Sum(expData, uint64(multicodec.Sha2_256), -1)
	require.NoError(t, err)

	data, err := envelope.ToDagCbor(examplePrivKey(t), newExample(t))
	require.NoError(t, err)

	id, err := envelope.CIDFromBytes(data)
	require.NoError(t, err)
	assert.Equal(t, exampleCID, envelope.CIDToBase58BTC(id))
	assert.Equal(t, expHash, id.Hash())
}

func TestStreaming(t *testing.T) {
	t.Parallel()

	expData := []byte("this is a test")

	expCID, err := cid.V1Builder{
		Codec:    uint64(multicodec.DagCbor),
		MhType:   multihash.SHA2_256,
		MhLength: 0,
	}.Sum(expData)
	require.NoError(t, err)

	t.Run("CIDReader()", func(t *testing.T) {
		t.Parallel()

		r, w := io.Pipe() //nolint:varnamelen
		cidReader := envelope.NewCIDReader(r)

		go func() {
			_, err := w.Write(expData)
			assert.NoError(t, err)
			assert.NoError(t, w.Close())
		}()

		actData, err := io.ReadAll(cidReader)
		require.NoError(t, err)
		assert.Equal(t, expData, actData)

		actCID, err := cidReader.CID()
		require.NoError(t, err)
		assert.Equal(t, expCID, actCID)
	})

	t.Run("CIDWriter", func(t *testing.T) {
		t.Parallel()

		r, w := io.Pipe() //nolint:varnamelen
		cidWriter := envelope.NewCIDWriter(w)

		go func() {
			_, err := cidWriter.Write(expData)
			assert.NoError(t, err)
			assert.NoError(t, w.Close())
		}()

		actData, err := io.ReadAll(r)
		require.NoError(t, err)
		assert.Equal(t, expData, actData)

		actCID, err := cidWriter.CID()
		require.NoError(t, err)
		assert.Equal(t, expCID, actCID)
	})
}
