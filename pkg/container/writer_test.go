package container

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWriterDedup(t *testing.T) {
	ctn := NewWriter()

	_, _, sealed := randToken()
	ctn.AddSealed(sealed)
	require.Len(t, ctn, 1)

	ctn.AddSealed(sealed)
	require.Len(t, ctn, 1)
}
