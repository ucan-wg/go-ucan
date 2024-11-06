package policy_test

import (
	"testing"

	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/fluent/qp"
	basicnode "github.com/ipld/go-ipld-prime/node/basic"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/ucan-wg/go-ucan/pkg/policy"
)

// TestInvocationValidation applies the example policy to the example
// arguments as defined in the [Validation] section of the invocation
// specification.
//
// [Validation]: https://github.com/ucan-wg/delegation/tree/v1_ipld#validation
func TestInvocationValidation(t *testing.T) {
	t.Parallel()

	polNode, err := qp.BuildList(basicnode.Prototype.Any, 2, func(la datamodel.ListAssembler) {
		qp.ListEntry(la, qp.List(3, func(la datamodel.ListAssembler) {
			qp.ListEntry(la, qp.String("=="))
			qp.ListEntry(la, qp.String(".from"))
			qp.ListEntry(la, qp.String("alice@example.com"))
		}))
		qp.ListEntry(la, qp.List(3, func(la datamodel.ListAssembler) {
			qp.ListEntry(la, qp.String("any"))
			qp.ListEntry(la, qp.String(".to"))
			qp.ListEntry(la, qp.List(3, func(la datamodel.ListAssembler) {
				qp.ListEntry(la, qp.String("like"))
				qp.ListEntry(la, qp.String("."))
				qp.ListEntry(la, qp.String("*@example.com"))
			}))
		}))
	})
	require.NoError(t, err)

	pol, err := policy.FromIPLD(polNode)
	require.NoError(t, err)

	t.Run("with passing args", func(t *testing.T) {
		t.Parallel()

		argsNode, err := qp.BuildMap(basicnode.Prototype.Any, 2, func(ma datamodel.MapAssembler) {
			qp.MapEntry(ma, "from", qp.String("alice@example.com"))
			qp.MapEntry(ma, "to", qp.List(2, func(la datamodel.ListAssembler) {
				qp.ListEntry(la, qp.String("bob@example.com"))
				qp.ListEntry(la, qp.String("carol@not.example.com"))
			}))
			qp.MapEntry(ma, "title", qp.String("Coffee"))
			qp.MapEntry(ma, "body", qp.String("Still on for coffee"))
		})
		require.NoError(t, err)

		assert.True(t, pol.Match(argsNode))
	})

	t.Run("fails on sender (first statement)", func(t *testing.T) {
		t.Parallel()

		argsNode, err := qp.BuildMap(basicnode.Prototype.Any, 2, func(ma datamodel.MapAssembler) {
			qp.MapEntry(ma, "from", qp.String("dan@example.com"))
			qp.MapEntry(ma, "to", qp.List(2, func(la datamodel.ListAssembler) {
				qp.ListEntry(la, qp.String("bob@example.com"))
				qp.ListEntry(la, qp.String("carol@not.example.com"))
			}))
			qp.MapEntry(ma, "title", qp.String("Coffee"))
			qp.MapEntry(ma, "body", qp.String("Still on for coffee"))
		})
		require.NoError(t, err)

		assert.False(t, pol.Match(argsNode))
	})

	t.Run("fails on recipients (second statement)", func(t *testing.T) {
		t.Parallel()

		argsNode, err := qp.BuildMap(basicnode.Prototype.Any, 2, func(ma datamodel.MapAssembler) {
			qp.MapEntry(ma, "from", qp.String("alice@example.com"))
			qp.MapEntry(ma, "to", qp.List(2, func(la datamodel.ListAssembler) {
				qp.ListEntry(la, qp.String("bob@null.com"))
				qp.ListEntry(la, qp.String("carol@not.null.com"))
			}))
			qp.MapEntry(ma, "title", qp.String("Coffee"))
			qp.MapEntry(ma, "body", qp.String("Still on for coffee"))
		})
		require.NoError(t, err)

		assert.False(t, pol.Match(argsNode))
	})
}
