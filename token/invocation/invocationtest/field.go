package invocationtest

import (
	"testing"

	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
	"github.com/ucan-wg/go-ucan/pkg/args"
	"github.com/ucan-wg/go-ucan/pkg/policy"
)

var (
	EmptyArguments = args.New()
	EmptyPolicy    = emptyPolicy()
)

func emptyPolicy() policy.Policy {
	pol, err := policy.FromDagJson("[]")
	if err != nil {
		panic(err)
	}

	return pol
}

func Proof(t *testing.T, cidStrs ...string) []cid.Cid {
	// t.Helper()

	prf := make([]cid.Cid, len(cidStrs))

	for i, cidStr := range cidStrs {
		id, err := cid.Parse(cidStr)
		require.NoError(t, err)

		prf[i] = id
	}

	return prf
}
