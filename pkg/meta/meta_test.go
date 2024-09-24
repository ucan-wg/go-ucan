package meta_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/ucan-wg/go-ucan/pkg/meta"
	"gotest.tools/v3/assert"
)

func TestMeta_Add(t *testing.T) {
	t.Parallel()

	type Unsupported struct{}

	t.Run("error if not primative or Node", func(t *testing.T) {
		t.Parallel()

		err := (&meta.Meta{}).Add("invalid", &Unsupported{})
		require.ErrorIs(t, err, meta.ErrUnsupported)
		assert.ErrorContains(t, err, "*github.com/ucan-wg/go-ucan/pkg/meta_test.Unsupported")
	})
}
