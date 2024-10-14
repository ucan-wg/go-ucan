package policy_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ucan-wg/go-ucan/pkg/policy"
	"github.com/ucan-wg/go-ucan/pkg/policy/literal"
)

func ExamplePolicy() {
	pol := policy.MustConstruct(
		policy.Equal(".status", literal.String("draft")),
		policy.All(".reviewer",
			policy.Like(".email", "*@example.com"),
		),
		policy.Any(".tags", policy.Or(
			policy.Equal(".", literal.String("news")),
			policy.Equal(".", literal.String("press")),
		)),
	)

	fmt.Println(pol)

	// Output:
	// [
	//   ["==", ".status", "draft"],
	//   ["all", ".reviewer",
	//     ["like", ".email", "*@example.com"]],
	//   ["any", ".tags",
	//     ["or", [
	//       ["==", ".", "news"],
	//       ["==", ".", "press"]]]
	//     ]
	// ]
}

func TestConstruct(t *testing.T) {
	pol, err := policy.Construct(
		policy.Equal(".status", literal.String("draft")),
		policy.All(".reviewer",
			policy.Like(".email", "*@example.com"),
		),
	)
	require.NoError(t, err)
	require.NotNil(t, pol)

	// check if errors cascade correctly
	pol, err = policy.Construct(
		policy.Equal(".status", literal.String("draft")),
		policy.All(".reviewer", policy.Or(
			policy.Like(".email", "*@example.com"),
			policy.Like(".", "\\"), // invalid pattern
		)),
	)
	require.Error(t, err)
	require.Nil(t, pol)
}
