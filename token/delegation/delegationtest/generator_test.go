package delegationtest

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gotest.tools/v3/golden"

	"github.com/ucan-wg/go-ucan/did/didtest"
)

// TestUpdate doesn't actually run a test but uses the Go testing library
// to trigger generation of the delegation tokens and associated Go file.
func TestUpdate(t *testing.T) {
	if golden.FlagUpdate() {
		update(t)
	}
}

func update(t *testing.T) {
	t.Helper()

	gen := &generator{}
	require.NoError(t, gen.chainPersonas(didtest.Personas(), acc{}, noopVariant()))
	require.NoError(t, gen.writeGoFile())
}
