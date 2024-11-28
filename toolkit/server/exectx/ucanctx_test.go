package exectx_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/INFURA/go-ethlibs/jsonrpc"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
	"github.com/ucan-wg/go-ucan/did/didtest"
	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/pkg/container"
	"github.com/ucan-wg/go-ucan/pkg/policy"
	"github.com/ucan-wg/go-ucan/pkg/policy/literal"
	"github.com/ucan-wg/go-ucan/token/delegation"
	"github.com/ucan-wg/go-ucan/token/invocation"

	exectx2 "github.com/INFURA/go-ucan-toolkit/server/exectx"
)

func ExampleContext() {
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
	)

	dlg, _ := delegation.Root(service.DID(), user.DID(), cmd, pol,
		delegation.WithExpirationIn(24*time.Hour),
	)
	dlgBytes, dlgCid, _ := dlg.ToSealed(service.PrivKey())

	// INVOCATION: the user leverages the delegation (power) to make a request.

	inv, _ := invocation.New(user.DID(), service.DID(), cmd, []cid.Cid{dlgCid},
		invocation.WithExpirationIn(10*time.Minute),
		invocation.WithArgument("myarg", "hello"), // we can specify invocation parameters
	)
	invBytes, invCid, _ := inv.ToSealed(user.PrivKey())

	// PACKAGING: no obligation for the transport, but the user needs to give the service the invocation
	// and all the proof delegations. We can use a container for that.
	cont := container.NewWriter()
	cont.AddSealed(dlgCid, dlgBytes)
	cont.AddSealed(invCid, invBytes)
	contBytes, _ := cont.ToCborBase64()

	// MAKING A REQUEST: we pass the container in the Bearer HTTP header

	jrpc := jsonrpc.NewRequest()
	jrpc.Method = "eth_call"
	jrpc.Params = jsonrpc.MustParams("0x599784", true)
	jrpcBytes, _ := jrpc.MarshalJSON()
	req, _ := http.NewRequest(http.MethodGet, "/foo/bar", bytes.NewReader(jrpcBytes))
	req.Header.Set("Authorization", "Bearer "+string(contBytes))

	// SERVER: Auth middleware
	// - decode the container
	// - create the context

	authMw := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Note: we obviously want something more robust, this is an example
			// Note: if an error occur, we'll want to return an HTTP 401 Unauthorized
			data := bytes.TrimPrefix([]byte(r.Header.Get("Authorization")), []byte("Bearer "))
			cont, _ := container.FromCborBase64(data)
			ucanCtx, _ := exectx2.FromContainer(cont)

			// insert into the go context
			req = req.WithContext(exectx2.AddUcanCtxToContext(req.Context(), ucanCtx))

			next.ServeHTTP(w, req)
		})
	}

	// SERVER: http checks

	httpMw := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ucanCtx, _ := exectx2.FromContext(req.Context())

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
			ucanCtx, _ := exectx2.FromContext(req.Context())

			var jrpc jsonrpc.Request
			_ = json.NewDecoder(r.Body).Decode(&jrpc)

			err := ucanCtx.VerifyJsonRpc(&jrpc)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}

	// SERVER: final handler

	handler := func(w http.ResponseWriter, r *http.Request) {
		ucanCtx, _ := exectx2.FromContext(req.Context())

		if err := ucanCtx.ExecutionAllowed(); err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		_, _ = fmt.Fprintln(w, "Success!")
	}

	// Ready to go!
	_ = http.ListenAndServe("", authMw(httpMw(jsonrpcMw(http.HandlerFunc(handler)))))
}

func TestGoCtx(t *testing.T) {
	ctx := context.Background()

	ucanCtx, ok := exectx2.FromContext(ctx)
	require.False(t, ok)
	require.Nil(t, ucanCtx)

	expected := &exectx2.UcanCtx{}

	ctx = exectx2.AddUcanCtxToContext(ctx, expected)

	got, ok := exectx2.FromContext(ctx)
	require.True(t, ok)
	require.Equal(t, expected, got)
}
