package attestation_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ucan-wg/go-ucan/token/attestation"
)

func TestSealUnsealRoundtrip(t *testing.T) {
	t.Parallel()

	privKey, iss, claims, meta, err := setupExampleNew()
	require.NoError(t, err)

	tkn1, err := attestation.New(iss,
		attestation.WithClaimMap(claims),
		attestation.WithMetaMap(meta),
		attestation.WithExpirationIn(time.Minute),
		attestation.WithoutIssuedAt(),
	)
	require.NoError(t, err)

	data, cid1, err := tkn1.ToSealed(privKey)
	require.NoError(t, err)

	tkn2, cid2, err := attestation.FromSealed(data)
	require.NoError(t, err)

	assert.Equal(t, cid1, cid2)
	assert.Equal(t, tkn1, tkn2)
}
