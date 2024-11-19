package delegationtest_test

import (
	"testing"

	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/ucan-wg/go-ucan/token/delegation/delegationtest"
	"github.com/ucan-wg/go-ucan/token/invocation"
)

func TestGetDelegation(t *testing.T) {
	t.Parallel()

	t.Run("passes with valid CID", func(t *testing.T) {
		t.Parallel()

		tkn, err := delegationtest.GetDelegation(delegationtest.TokenAliceBobCID)
		require.NoError(t, err)
		assert.NotZero(t, tkn)
	})

	t.Run("fails with unknown CID", func(t *testing.T) {
		t.Parallel()

		tkn, err := delegationtest.GetDelegation(cid.Undef)
		require.ErrorIs(t, err, invocation.ErrMissingDelegation)
		require.ErrorContains(t, err, "CID b")
		assert.Nil(t, tkn)
	})
}
