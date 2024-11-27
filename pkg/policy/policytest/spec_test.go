package policytest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestInvocationValidation applies the example policy to the second
// example arguments as defined in the [Validation] section of the
// invocation specification.
//
// [Validation]: https://github.com/ucan-wg/delegation/tree/v1_ipld#validation
func TestInvocationValidationSpecExamples(t *testing.T) {
	t.Parallel()

	t.Run("with passing args", func(t *testing.T) {
		t.Parallel()

		exec, stmt := SpecPolicy.Match(specValidArgumentsIPLD)
		assert.True(t, exec)
		assert.Nil(t, stmt)
	})

	t.Run("fails on recipients (second statement)", func(t *testing.T) {
		t.Parallel()

		exec, stmt := SpecPolicy.Match(specInvalidArgumentsIPLD)
		assert.False(t, exec)
		assert.NotNil(t, stmt)
	})
}
