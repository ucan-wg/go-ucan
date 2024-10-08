package server

import (
	"encoding/base64"
	"io"
	"net/http"
	"strings"

	"github.com/ucan-wg/go-ucan/did"
	"github.com/ucan-wg/go-ucan/token/delegation"
)

type Middleware func(http.Handler) http.Handler

func ExampleMiddleware(serviceDID did.DID) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			tokenReader, ok := extractBearerToken(r)
			if !ok {
				http.Error(w, "missing or malformed UCAN HTTP Bearer token", http.StatusUnauthorized)
				return
			}

			// decode
			// TODO: ultimately, this token should be a container with one invocation and 1+ delegations.
			//  We are doing something simpler for now.
			dlg, err := delegation.FromDagCborReader(tokenReader)
			if err != nil {
				http.Error(w, "malformed UCAN delegation", http.StatusBadRequest)
				return
			}

			// optional: http-bearer

			// validate
			if dlg.Subject() != serviceDID {
				http.Error(w, "invalid UCAN delegation", http.StatusBadRequest)
				return
			}
			// TODO: policies check + inject in context

			// extract values
			userId, err := dlg.Meta().GetString("userId")
			if err != nil {
				http.Error(w, "missing or malformed userId", http.StatusBadRequest)
				return
			}
			projectId, err := dlg.Meta().GetString("projectId")
			if err != nil {
				http.Error(w, "missing or malformed projectId", http.StatusBadRequest)
				return
			}

			// inject into context
			ctx = addUserIdToContext(ctx, userId)
			ctx = addProjectIdToContext(ctx, projectId)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func extractBearerToken(r *http.Request) (io.Reader, bool) {
	header := r.Header.Get("Authorization")
	if header == "" {
		return nil, false
	}

	if !strings.HasPrefix(header, "Bearer ") {
		return nil, false
	}

	// skip prefix
	reader := strings.NewReader(header[len("Bearer "):])

	// base64 decode
	return base64.NewDecoder(base64.StdEncoding, reader), true
}
