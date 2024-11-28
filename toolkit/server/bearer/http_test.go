package bearer

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/multiformats/go-multihash"
	"github.com/stretchr/testify/require"
	"github.com/ucan-wg/go-ucan/pkg/args"
	"github.com/ucan-wg/go-ucan/pkg/policy"
	"github.com/ucan-wg/go-ucan/pkg/policy/literal"
)

func TestHttp(t *testing.T) {
	pol := policy.MustConstruct(
		policy.Equal(".http.scheme", literal.String("http")),
		policy.Equal(".http.method", literal.String("GET")),
		policy.Equal(".http.host", literal.String("example.com")),
		policy.Equal(".http.path", literal.String("/foo")),
		policy.Like(".http.headers.User-Agent", "*Mozilla*"),
		policy.Equal(".http.headers.Origin", literal.String("dapps.com")),
	)

	tests := []struct {
		name     string
		method   string
		target   string
		headers  map[string]string
		expected bool
	}{
		{
			name:   "happy path",
			method: http.MethodGet,
			target: "http://example.com/foo",
			headers: map[string]string{
				"User-Agent": "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0",
				"Origin":     "dapps.com",
			},
			expected: true,
		},
		{
			name:   "wrong scheme",
			method: http.MethodGet,
			target: "https://example.com/foo",
			headers: map[string]string{
				"User-Agent": "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0",
				"Origin":     "dapps.com",
			},
			expected: false,
		},
		{
			name:   "wrong method",
			method: http.MethodPost,
			target: "http://example.com/foo",
			headers: map[string]string{
				"User-Agent": "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0",
				"Origin":     "dapps.com",
			},
			expected: false,
		},
		{
			name:   "wrong host",
			method: http.MethodGet,
			target: "http://foo.com/foo",
			headers: map[string]string{
				"User-Agent": "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0",
				"Origin":     "dapps.com",
			},
			expected: false,
		},
		{
			name:   "wrong path",
			method: http.MethodGet,
			target: "http://example.com/bar",
			headers: map[string]string{
				"User-Agent": "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0",
				"Origin":     "dapps.com",
			},
			expected: false,
		},
		{
			name:   "wrong origin",
			method: http.MethodGet,
			target: "http://example.com/foo",
			headers: map[string]string{
				"User-Agent": "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0",
				"Origin":     "foo.com",
			},
			expected: false,
		},
		{
			name:   "wrong user-agent",
			method: http.MethodGet,
			target: "http://example.com/foo",
			headers: map[string]string{
				"User-Agent": "Chrome/51.0.2704.103 Safari/537.36",
				"Origin":     "dapps.com",
			},
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			handler := func(w http.ResponseWriter, r *http.Request) {
				// we don't test the args hash here
				emptyArgs := args.New().ReadOnly()

				ctx := NewHttpBearer(pol, emptyArgs, r)

				_, err := ctx.Args()
				require.NoError(t, err)

				if tc.expected {
					require.NoError(t, ctx.Verify())
				} else {
					require.Error(t, ctx.Verify())
				}
			}

			req := httptest.NewRequest(tc.method, tc.target, nil)
			for k, v := range tc.headers {
				req.Header.Set(k, v)
			}

			w := httptest.NewRecorder()
			handler(w, req)
		})
	}
}

func TestHttpHash(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "http://example.com/foo", nil)
	require.NoError(t, err)
	req.Header.Add("User-Agent", "Chrome/51.0.2704.103 Safari/537.36")
	req.Header.Add("Origin", "dapps.com")

	pol := policy.MustConstruct(
		policy.Equal(".http.scheme", literal.String("http")),
	)

	tests := []struct {
		name     string
		hash     []byte
		expected bool
	}{
		{
			name:     "correct hash",
			hash:     must(MakeHttpHash(req)),
			expected: true,
		},
		{
			name:     "non-matching hash",
			hash:     must(multihash.Sum([]byte{1, 2, 3, 4}, multihash.SHA2_256, -1)),
			expected: false,
		},
		{
			name:     "wrong type of hash",
			hash:     must(multihash.Sum([]byte{1, 2, 3, 4}, multihash.BLAKE3, -1)),
			expected: false,
		},
		{
			name:     "no hash",
			hash:     nil,
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			invArgs := args.New()
			err := invArgs.Add(HttpArgsKey, tc.hash)
			require.NoError(t, err)

			ctx := NewHttpBearer(pol, invArgs.ReadOnly(), req)

			if tc.expected {
				require.NoError(t, ctx.Verify())
			} else {
				require.Error(t, ctx.Verify())
			}
		})
	}
}
