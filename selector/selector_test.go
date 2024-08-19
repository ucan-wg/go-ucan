package selector

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	t.Run("identity", func(t *testing.T) {
		sel, err := Parse(".")
		require.NoError(t, err)
		require.True(t, sel.Identity())
		require.False(t, sel.Optional())
		require.Empty(t, sel.Field())
		require.Empty(t, sel.Index())
	})

	t.Run("dotted field", func(t *testing.T) {
		sel, err := Parse(".foo")
		require.NoError(t, err)
		require.False(t, sel.Identity())
		require.False(t, sel.Optional())
		require.Equal(t, sel.Field(), "foo")
		require.Empty(t, sel.Index())
	})

	t.Run("dotted explicit field", func(t *testing.T) {
		sel, err := Parse(".[\"foo\"]")
		require.NoError(t, err)
		require.False(t, sel.Identity())
		require.False(t, sel.Optional())
		require.Equal(t, sel.Field(), "foo")
		require.Empty(t, sel.Index())
	})

	t.Run("dotted index", func(t *testing.T) {
		sel, err := Parse(".[138]")
		require.NoError(t, err)
		require.False(t, sel.Identity())
		require.False(t, sel.Optional())
		require.Empty(t, sel.Field())
		require.Equal(t, sel.Index(), 138)
	})

	t.Run("explicit field", func(t *testing.T) {
		sel, err := Parse("[\"foo\"]")
		require.NoError(t, err)
		require.False(t, sel.Identity())
		require.False(t, sel.Optional())
		require.Equal(t, sel.Field(), "foo")
		require.Empty(t, sel.Index())
	})

	t.Run("index", func(t *testing.T) {
		sel, err := Parse("[138]")
		require.NoError(t, err)
		require.False(t, sel.Identity())
		require.False(t, sel.Optional())
		require.Empty(t, sel.Field())
		require.Equal(t, sel.Index(), 138)
	})

	t.Run("negative index", func(t *testing.T) {
		sel, err := Parse("[-138]")
		require.NoError(t, err)
		require.False(t, sel.Identity())
		require.False(t, sel.Optional())
		require.Empty(t, sel.Field())
		require.Equal(t, sel.Index(), -138)
	})

	t.Run("optional dotted field", func(t *testing.T) {
		sel, err := Parse(".foo?")
		require.NoError(t, err)
		require.False(t, sel.Identity())
		require.True(t, sel.Optional())
		require.Equal(t, sel.Field(), "foo")
		require.Empty(t, sel.Index())
	})

	t.Run("optional dotted explicit field", func(t *testing.T) {
		sel, err := Parse(".[\"foo\"]?")
		require.NoError(t, err)
		require.False(t, sel.Identity())
		require.True(t, sel.Optional())
		require.Equal(t, sel.Field(), "foo")
		require.Empty(t, sel.Index())
	})

	t.Run("optional dotted index", func(t *testing.T) {
		sel, err := Parse(".[138]?")
		require.NoError(t, err)
		require.False(t, sel.Identity())
		require.True(t, sel.Optional())
		require.Empty(t, sel.Field())
		require.Equal(t, sel.Index(), 138)
	})

	t.Run("optional explicit field", func(t *testing.T) {
		sel, err := Parse("[\"foo\"]?")
		require.NoError(t, err)
		require.False(t, sel.Identity())
		require.True(t, sel.Optional())
		require.Equal(t, sel.Field(), "foo")
		require.Empty(t, sel.Index())
	})

	t.Run("optional index", func(t *testing.T) {
		sel, err := Parse("[138]?")
		require.NoError(t, err)
		require.False(t, sel.Identity())
		require.True(t, sel.Optional())
		require.Empty(t, sel.Field())
		require.Equal(t, sel.Index(), 138)
	})

	t.Run("non dotted", func(t *testing.T) {
		_, err := Parse("foo")
		if err == nil {
			t.Fatalf("expected error parsing selector")
		}
		fmt.Println(err)
	})

	t.Run("non quoted", func(t *testing.T) {
		_, err := Parse(".[foo]")
		if err == nil {
			t.Fatalf("expected error parsing selector")
		}
		fmt.Println(err)
	})
}
