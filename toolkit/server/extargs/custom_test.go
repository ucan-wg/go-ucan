package extargs

import (
	"fmt"
	"testing"

	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/fluent/qp"
	"github.com/stretchr/testify/require"

	"github.com/ucan-wg/go-ucan/pkg/policy"
	"github.com/ucan-wg/go-ucan/pkg/policy/literal"
)

func ExampleCustomExtArgs() {
	// Note: this is an example for how to build arguments, but you likely want to use CustomExtArgs
	// through UcanCtx.

	pol := policy.Policy{} // policies from the delegations

	// We will construct the following args:
	// "key": {
	//  "ntwk":"eth-mainnet",
	//  "quota":{
	//    "ur":1234,
	//    "udc":1234,
	//    "arch":1234,
	//    "down":1234,
	//    "store":1234,
	//    "up":1234
	//  }
	// }
	customArgs := NewCustomExtArgs("key", pol, func(ma datamodel.MapAssembler) {
		qp.MapEntry(ma, "ntwk", qp.String("eth-mainnet"))
		qp.MapEntry(ma, "quota", qp.Map(6, func(ma datamodel.MapAssembler) {
			qp.MapEntry(ma, "ur", qp.Int(1234))
			qp.MapEntry(ma, "udc", qp.Int(1234))
			qp.MapEntry(ma, "arch", qp.Int(1234))
			qp.MapEntry(ma, "down", qp.Int(1234))
			qp.MapEntry(ma, "store", qp.Int(1234))
			qp.MapEntry(ma, "up", qp.Int(1234))
		}))
	})

	err := customArgs.Verify()
	fmt.Println(err)

	// Output:
	// <nil>
}

func TestCustom(t *testing.T) {
	assembler := func(ma datamodel.MapAssembler) {
		qp.MapEntry(ma, "ntwk", qp.String("eth-mainnet"))
		qp.MapEntry(ma, "quota", qp.Map(6, func(ma datamodel.MapAssembler) {
			qp.MapEntry(ma, "ur", qp.Int(1234))
			qp.MapEntry(ma, "udc", qp.Int(1234))
			qp.MapEntry(ma, "arch", qp.Int(1234))
			qp.MapEntry(ma, "down", qp.Int(1234))
			qp.MapEntry(ma, "store", qp.Int(1234))
			qp.MapEntry(ma, "up", qp.Int(1234))
		}))
	}

	tests := []struct {
		name     string
		pol      policy.Policy
		expected bool
	}{
		{
			name:     "no policies",
			pol:      policy.Policy{},
			expected: true,
		},
		{
			name: "matching args",
			pol: policy.MustConstruct(
				policy.Equal(".key.ntwk", literal.String("eth-mainnet")),
				policy.LessThanOrEqual(".key.quota.ur", literal.Int(1234)),
			),
			expected: true,
		},
		{
			name: "wrong network",
			pol: policy.MustConstruct(
				policy.Equal(".key.ntwk", literal.String("avalanche-fuji")),
				policy.LessThanOrEqual(".key.quota.ur", literal.Int(1234)),
			),
			expected: false,
		},
		{
			name: "unrespected quota",
			pol: policy.MustConstruct(
				policy.Equal(".key.ntwk", literal.String("eth-mainnet")),
				policy.LessThanOrEqual(".key.quota.ur", literal.Int(100)),
			),
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			extArgs := NewCustomExtArgs("key", tc.pol, assembler)

			_, err := extArgs.Args()
			require.NoError(t, err)

			if tc.expected {
				require.NoError(t, extArgs.Verify())
			} else {
				require.Error(t, extArgs.Verify())
			}
		})
	}
}
