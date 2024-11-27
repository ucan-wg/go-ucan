package args_test

import (
	"testing"

	"github.com/ipld/go-ipld-prime"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/ucan-wg/go-ucan/pkg/args"
)

func TestBuilder_XXX(t *testing.T) {
	t.Parallel()

	const (
		keyOne = "key1"
		valOne = "string"
		keyTwo = "key2"
		valTwo = 42
	)

	exp := args.New()
	exp.Add(keyOne, valOne)
	exp.Add(keyTwo, valTwo)

	expNode, err := exp.ToIPLD()
	require.NoError(t, err)

	disjointKeys := args.NewBuilder().
		Add(keyOne, valOne).
		Add(keyTwo, valTwo)

	duplicateKeys := args.NewBuilder().
		Add(keyOne, valOne).
		Add(keyTwo, valTwo).
		Add(keyOne, "oh no!")

	t.Run("MustBuild succeeds with disjoint keys", func(t *testing.T) {
		t.Parallel()

		var act *args.Args

		require.NotPanics(t, func() {
			act = disjointKeys.MustBuild()
		})
		assert.Equal(t, exp, act)
	})

	t.Run("MustBuild fails with duplicate keys", func(t *testing.T) {
		t.Parallel()

		var act *args.Args

		require.Panics(t, func() {
			act = duplicateKeys.MustBuild()
		})
		assert.Nil(t, act)
	})

	t.Run("MustBuildIPLD succeeds with disjoint keys", func(t *testing.T) {
		t.Parallel()

		var act ipld.Node

		require.NotPanics(t, func() {
			act = disjointKeys.MustBuildIPLD()
		})
		assert.Equal(t, expNode, act)
	})

	t.Run("MustBuildIPLD fails with duplicate keys", func(t *testing.T) {
		t.Parallel()

		var act ipld.Node

		require.Panics(t, func() {
			act = duplicateKeys.MustBuildIPLD()
		})
		assert.Nil(t, act)
	})
}
