package envelope_test

import (
	"testing"

	"github.com/multiformats/go-multicodec"
	"github.com/multiformats/go-multihash"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/ucan-wg/go-ucan/internal/envelope"
	"gotest.tools/v3/golden"
)

func TestCid(t *testing.T) {
	t.Parallel()

	expData := golden.Get(t, "example.dagcbor")
	expHash, err := multihash.Sum(expData, uint64(multicodec.Sha2_256), -1)
	require.NoError(t, err)

	id, err := envelope.CID(examplePrivKey(t), newExample(t))
	require.NoError(t, err)
	assert.Equal(t, exampleCID, envelope.CIDToBase58BTC(id))
	assert.Equal(t, expHash, id.Hash())
}
