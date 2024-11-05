package invocation_test

import (
	"testing"
	"time"

	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/ucan-wg/go-ucan/did"
	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/token/invocation"
)

func TestSealUnsealRoundtrip(t *testing.T) {
	t.Parallel()

	privKey, iss, sub, cmd, args, prf, meta, err := setupExampleNew()
	require.NoError(t, err)

	tkn1, err := invocation.New(
		iss,
		sub,
		cmd,
		prf,
		invocation.WithArguments(args),
		invocation.WithMeta("env", "development"),
		invocation.WithMeta("tags", meta["tags"]),
		invocation.WithExpirationIn(time.Minute),
		invocation.WithoutInvokedAt())
	require.NoError(t, err)

	data, cid1, err := tkn1.ToSealed(privKey)
	require.NoError(t, err)

	tkn2, cid2, err := invocation.FromSealed(data)
	require.NoError(t, err)

	assert.Equal(t, cid1, cid2)
	assert.Equal(t, tkn1, tkn2)
}

func setupNew(t *testing.T) (privKey crypto.PrivKey, iss, sub did.DID, cmd command.Command, args map[string]datamodel.Node, prf []cid.Cid, meta map[string]datamodel.Node) {
	privKey, iss, sub, cmd, args, prf, meta, err := setupExampleNew()
	require.NoError(t, err)

	return // WARNING: named return values
}
