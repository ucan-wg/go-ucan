package extargs

import (
	"testing"

	"github.com/INFURA/go-ethlibs/jsonrpc"
	"github.com/multiformats/go-multihash"
	"github.com/stretchr/testify/require"
	"github.com/ucan-wg/go-ucan/did/didtest"
	"github.com/ucan-wg/go-ucan/pkg/args"
	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/pkg/policy"
	"github.com/ucan-wg/go-ucan/pkg/policy/literal"
	"github.com/ucan-wg/go-ucan/token/invocation"
)

func TestJsonRpc(t *testing.T) {
	tests := []struct {
		name     string
		req      *jsonrpc.Request
		pol      policy.Policy
		expected bool
	}{
		{
			name: "or on method, not matching",
			req: jsonrpc.MustRequest(1839673506133526, "eth_getBlockByNumber",
				"0x599784", true,
			),
			pol: policy.MustConstruct(
				policy.Or(
					policy.Equal(".jsonrpc.method", literal.String("eth_getCode")),
					policy.Equal(".jsonrpc.method", literal.String("eth_getBalance")),
					policy.Equal(".jsonrpc.method", literal.String("eth_call")),
					policy.Equal(".jsonrpc.method", literal.String("eth_blockNumber")),
				),
			),
			expected: false,
		},
		{
			name: "or on method, matching",
			req: jsonrpc.MustRequest(1839673506133526, "eth_call",
				map[string]string{"to": "0xBADBADBADBADBADBADBADBADBADBADBADBADBAD1"},
			),
			pol: policy.MustConstruct(
				policy.Or(
					policy.Equal(".jsonrpc.method", literal.String("eth_getCode")),
					policy.Equal(".jsonrpc.method", literal.String("eth_getBalance")),
					policy.Equal(".jsonrpc.method", literal.String("eth_call")),
					policy.Equal(".jsonrpc.method", literal.String("eth_blockNumber")),
				),
			),
			expected: true,
		},
		{
			name: "complex, optional parameter, matching",
			req: jsonrpc.MustRequest(1839673506133526, "debug_traceCall",
				true, false, 1234, "callTracer",
			),
			pol: policy.MustConstruct(
				policy.Equal(".jsonrpc.method", literal.String("debug_traceCall")),
				policy.Or(
					policy.Equal(".jsonrpc.params[3]?", literal.String("callTracer")),
					policy.Equal(".jsonrpc.params[3]?", literal.String("prestateTracer")),
				),
			),
			expected: true,
		},
		{
			name: "complex, optional parameter, missing parameter",
			req: jsonrpc.MustRequest(1839673506133526, "debug_traceCall",
				true, false, 1234,
			),
			pol: policy.MustConstruct(
				policy.Equal(".jsonrpc.method", literal.String("debug_traceCall")),
				policy.Or(
					policy.Equal(".jsonrpc.params[3]?", literal.String("callTracer")),
					policy.Equal(".jsonrpc.params[3]?", literal.String("prestateTracer")),
				),
			),
			expected: true,
		},
		{
			name: "complex, parameter not matching",
			req: jsonrpc.MustRequest(1839673506133526, "debug_traceCall",
				true, false, 1234, "ho_no",
			),
			pol: policy.MustConstruct(
				policy.Equal(".jsonrpc.method", literal.String("debug_traceCall")),
				policy.Or(
					policy.Equal(".jsonrpc.params[3]?", literal.String("callTracer")),
					policy.Equal(".jsonrpc.params[3]?", literal.String("prestateTracer")),
				),
			),
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// we don't test the args hash here
			emptyArgs := args.New().ReadOnly()

			ctx := NewJsonRpcExtArgs(tc.pol, emptyArgs, tc.req)

			_, err := ctx.Args()
			require.NoError(t, err)

			if tc.expected {
				require.NoError(t, ctx.Verify())
			} else {
				require.Error(t, ctx.Verify())
			}
		})
	}
}

func TestJsonRpcHash(t *testing.T) {
	servicePersona := didtest.PersonaAlice
	clientPersona := didtest.PersonaBob

	req := jsonrpc.MustRequest(1839673506133526, "debug_traceCall",
		true, false, 1234, "ho_no",
	)
	pol := policy.MustConstruct(
		policy.Equal(".jsonrpc.method", literal.String("debug_traceCall")),
	)

	makeArg := func(data []byte, code uint64) invocation.Option {
		mh, err := multihash.Sum(data, code, -1)
		require.NoError(t, err)
		return invocation.WithArgument(JsonRpcArgsKey, []byte(mh))
	}

	tests := []struct {
		name       string
		argOptions []invocation.Option
		expected   bool
	}{
		{
			name:       "correct hash",
			argOptions: []invocation.Option{must(MakeJsonRpcHash(req))},
			expected:   true,
		},
		{
			name:       "non-matching hash",
			argOptions: []invocation.Option{makeArg([]byte{1, 2, 3, 4}, multihash.SHA2_256)},
			expected:   false,
		},
		{
			name:       "wrong type of hash",
			argOptions: []invocation.Option{makeArg([]byte{1, 2, 3, 4}, multihash.BLAKE3)},
			expected:   false,
		},
		{
			name:       "no hash",
			argOptions: nil,
			expected:   true, // having a hash is not enforced
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			inv, err := invocation.New(
				clientPersona.DID(),
				command.MustParse("/foo"),
				servicePersona.DID(),
				nil,
				tc.argOptions..., // inject hash argument, if any
			)
			require.NoError(t, err)

			ctx := NewJsonRpcExtArgs(pol, inv.Arguments(), req)

			if tc.expected {
				require.NoError(t, ctx.Verify())
			} else {
				require.Error(t, ctx.Verify())
			}
		})
	}
}

func must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}
