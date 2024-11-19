// Package policytest provides values and functions that are useful when
// testing code that relies on Policies.
package policytest

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/ucan-wg/go-ucan/pkg/policy"
)

// EmptyPolicy provides a policy with no statements for testing purposes.
func EmptyPolicy(t *testing.T) policy.Policy {
	t.Helper()

	pol, err := policy.FromDagJson("[]")
	require.NoError(t, err)

	return pol
}
