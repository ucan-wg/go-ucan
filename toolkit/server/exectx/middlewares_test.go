package exectx

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MetaMask/go-did-it/didtest"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"

	"github.com/ucan-wg/go-ucan/pkg/container"
	"github.com/ucan-wg/go-ucan/token/delegation"
	"github.com/ucan-wg/go-ucan/token/invocation"
)

func TestExtractMW(t *testing.T) {
	const service = didtest.PersonaAlice
	const client = didtest.PersonaBob
	const cmd = "/foo/bar"

	for _, tc := range []struct {
		name               string
		addHeaderFn        func(func(key string, value string))
		expectedStatusCode int
		successful         bool
	}{
		{
			name:               "no auth",
			addHeaderFn:        func(f func(key string, value string)) {},
			expectedStatusCode: http.StatusUnauthorized,
			successful:         false,
		},
		{
			name: "wrong kind of auth",
			addHeaderFn: func(f func(key string, value string)) {
				f("Authorization", "Basic foobar")
			},
			expectedStatusCode: http.StatusUnauthorized,
			successful:         false,
		},
		{
			name: "invalid container",
			addHeaderFn: func(f func(key string, value string)) {
				f("Authorization", "Bearer foobar")
			},
			expectedStatusCode: http.StatusBadRequest,
			successful:         false,
		},
		{
			name: "valid containter, for incorrect service",
			addHeaderFn: func(f func(key string, value string)) {
				cont := container.NewWriter()

				const service2 = didtest.PersonaCarol

				dlg, _ := delegation.Root(service2.DID(), client.DID(), cmd, nil)
				dlgByte, dlgCid, _ := dlg.ToSealed(service2.PrivKey())
				cont.AddSealed(dlgByte)

				inv, _ := invocation.New(client.DID(), cmd, service2.DID(), []cid.Cid{dlgCid})
				invBytes, _, _ := inv.ToSealed(client.PrivKey())
				cont.AddSealed(invBytes)

				contB64, _ := cont.ToBase64StdPadding()

				f("Authorization", "Bearer "+contB64)
			},
			expectedStatusCode: http.StatusUnauthorized,
			successful:         false,
		},
		{
			name: "valid containter, missing invocation",
			addHeaderFn: func(f func(key string, value string)) {
				cont := container.NewWriter()

				dlg, _ := delegation.Root(service.DID(), client.DID(), cmd, nil)
				dlgByte, _, _ := dlg.ToSealed(service.PrivKey())
				cont.AddSealed(dlgByte)

				contB64, _ := cont.ToBase64StdPadding()

				f("Authorization", "Bearer "+contB64)
			},
			expectedStatusCode: http.StatusUnauthorized,
			successful:         false,
		},
		{
			name: "valid containter, valid tokens",
			addHeaderFn: func(f func(key string, value string)) {
				cont := container.NewWriter()

				dlg, _ := delegation.Root(service.DID(), client.DID(), cmd, nil)
				dlgByte, dlgCid, _ := dlg.ToSealed(service.PrivKey())
				cont.AddSealed(dlgByte)

				inv, _ := invocation.New(client.DID(), cmd, service.DID(), []cid.Cid{dlgCid})
				invBytes, _, _ := inv.ToSealed(client.PrivKey())
				cont.AddSealed(invBytes)

				contB64, _ := cont.ToBase64StdPadding()

				f("Authorization", "Bearer "+contB64)
			},
			expectedStatusCode: http.StatusOK,
			successful:         true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				_, has := FromContext(r.Context())
				require.Equal(t, tc.successful, has)

				_, _ = io.WriteString(w, "OK")
			})
			handler = ExtractMW(handler, service.DID())

			req := httptest.NewRequest("GET", "https://example.com/foo", nil)
			tc.addHeaderFn(req.Header.Set)

			w := httptest.NewRecorder()
			handler.ServeHTTP(w, req)

			require.Equal(t, tc.expectedStatusCode, w.Code)

		})
	}
}
