package exectx_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/INFURA/go-ethlibs/jsonrpc"
	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/fluent/qp"
	"github.com/stretchr/testify/require"
	"github.com/ucan-wg/go-ucan/did/didtest"
	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/pkg/container"
	"github.com/ucan-wg/go-ucan/pkg/policy"
	"github.com/ucan-wg/go-ucan/pkg/policy/literal"
	"github.com/ucan-wg/go-ucan/token/delegation"
	"github.com/ucan-wg/go-ucan/token/invocation"

	"github.com/INFURA/go-ucan-toolkit/server/exectx"
)

const (
	network = "eth-mainnet"
)

func TestCtx(t *testing.T) {
	// let's use some pre-made DID+privkey.
	// use go-ucan/did to generate or parse them.
	service := didtest.PersonaAlice
	user := didtest.PersonaBob

	// DELEGATION: the service gives some power to the user, in the form of a root UCAN token.
	// The command defines the shape of the arguments on which the policies operate, it is specific to that service.
	// Policies define what the user can do.

	cmd := command.New("foo")
	pol := policy.MustConstruct(
		// some basic HTTP constraints
		policy.Equal(".http.method", literal.String("GET")),
		policy.Like(".http.path", "/foo/*"),
		// some JsonRpc constraints
		policy.Or(
			policy.Like(".jsonrpc.method", "eth_*"),
			policy.Equal(".jsonrpc.method", literal.String("debug_traceCall")),
		),
		// some infura constraints
		// Network
		policy.Equal(".inf.ntwk", literal.String(network)),
		// Quota
		policy.LessThanOrEqual(".inf.quota.ur", literal.Int(1234)),
	)

	dlg, err := delegation.Root(service.DID(), user.DID(), cmd, pol,
		delegation.WithExpirationIn(24*time.Hour),
	)
	require.NoError(t, err)
	dlgBytes, dlgCid, err := dlg.ToSealed(service.PrivKey())
	require.NoError(t, err)

	// INVOCATION: the user leverages the delegation (power) to make a request.

	inv, err := invocation.New(user.DID(), cmd, service.DID(), []cid.Cid{dlgCid},
		invocation.WithExpirationIn(10*time.Minute),
		invocation.WithArgument("myarg", "hello"), // we can specify invocation parameters
	)
	require.NoError(t, err)
	invBytes, _, err := inv.ToSealed(user.PrivKey())
	require.NoError(t, err)

	// PACKAGING: no obligation for the transport, but the user needs to give the service the invocation
	// and all the proof delegations. We can use a container for that.
	cont := container.NewWriter()
	cont.AddSealed(dlgBytes)
	cont.AddSealed(invBytes)
	contBytes, err := cont.ToBase64StdPadding()
	require.NoError(t, err)

	// MAKING A REQUEST: we pass the container in the Bearer HTTP header

	jrpc := jsonrpc.NewRequest()
	jrpc.Method = "eth_call"
	jrpc.Params = jsonrpc.MustParams("0x599784", true)
	jrpcBytes, err := jrpc.MarshalJSON()
	require.NoError(t, err)
	req, err := http.NewRequest(http.MethodGet, "/foo/bar", bytes.NewReader(jrpcBytes))
	require.NoError(t, err)
	req.Header.Set("Authorization", "Bearer "+string(contBytes))

	// SERVER: Auth middleware
	// - decode the container
	// - create the context

	authMw := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Note: we obviously want something more robust, this is an example
			// Note: if an error occur, we'll want to return an HTTP 401 Unauthorized
			data := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
			cont, err := container.FromString(data)
			require.NoError(t, err)
			ucanCtx, err := exectx.FromContainer(cont)
			require.NoError(t, err)

			// insert into the go context
			r = r.WithContext(exectx.AddUcanCtxToContext(r.Context(), ucanCtx))

			next.ServeHTTP(w, r)
		})
	}

	// SERVER: http checks

	httpMw := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ucanCtx, ok := exectx.FromContext(r.Context())
			require.True(t, ok)

			err := ucanCtx.VerifyHttp(r)
			if err != nil {
				// This will print something like:
				// `the following UCAN policy is not satisfied: ["==", ".http.path", "/foo"]`
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}

	// SERVER: JsonRpc checks

	jsonrpcMw := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ucanCtx, ok := exectx.FromContext(r.Context())
			require.True(t, ok)

			var jrpc jsonrpc.Request
			err := json.NewDecoder(r.Body).Decode(&jrpc)
			require.NoError(t, err)

			err = ucanCtx.VerifyJsonRpc(&jrpc)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}

	// SERVER: custom infura checks

	infuraMw := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ucanCtx, ok := exectx.FromContext(r.Context())
			require.True(t, ok)
			err := ucanCtx.VerifyInfura(func(ma datamodel.MapAssembler) {
				qp.MapEntry(ma, "ntwk", qp.String(network))
				qp.MapEntry(ma, "quota", qp.Map(1, func(ma datamodel.MapAssembler) {
					qp.MapEntry(ma, "ur", qp.Int(1234))
				}))
			})
			require.NoError(t, err)
			next.ServeHTTP(w, r)
		})
	}

	// SERVER: final handler

	handler := func(w http.ResponseWriter, r *http.Request) {
		ucanCtx, ok := exectx.FromContext(r.Context())
		require.True(t, ok)

		if err := ucanCtx.ExecutionAllowed(); err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)
	}

	sut := authMw(httpMw(jsonrpcMw(infuraMw(http.HandlerFunc(handler)))))

	rec := httptest.NewRecorder()
	sut.ServeHTTP(rec, req)
	require.Equal(t, http.StatusOK, rec.Code)
}

func TestGoCtx(t *testing.T) {
	ctx := context.Background()

	ucanCtx, ok := exectx.FromContext(ctx)
	require.False(t, ok)
	require.Nil(t, ucanCtx)

	expected := &exectx.UcanCtx{}

	ctx = exectx.AddUcanCtxToContext(ctx, expected)

	got, ok := exectx.FromContext(ctx)
	require.True(t, ok)
	require.Equal(t, expected, got)
}
