package invocation_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ucan-wg/go-ucan/token/invocation"
)

func TestSealUnsealRoundtrip(t *testing.T) {
	t.Parallel()

	privKey, iss, sub, cmd, args, prf, meta, err := setupExampleNew()
	require.NoError(t, err)

	tkn1, err := invocation.New(iss, cmd, sub, prf,
		invocation.WithArgument("uri", args["uri"]),
		invocation.WithArgument("headers", args["headers"]),
		invocation.WithArgument("payload", args["payload"]),
		invocation.WithMeta("env", "development"),
		invocation.WithMeta("tags", meta["tags"]),
		invocation.WithExpirationIn(time.Minute),
		invocation.WithoutInvokedAt(),
	)
	require.NoError(t, err)

	data, cid1, err := tkn1.ToSealed(privKey)
	require.NoError(t, err)

	tkn2, cid2, err := invocation.FromSealed(data)
	require.NoError(t, err)

	assert.Equal(t, cid1, cid2)
	assert.Equal(t, tkn1, tkn2)
}
