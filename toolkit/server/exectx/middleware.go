package exectx

import (
	"errors"
	"net/http"

	"github.com/INFURA/go-ucan-toolkit/server/bearer"
)

// Middleware returns an HTTP middleware tasked with:
// - extracting UCAN credentials from the `Authorization: Bearer <data>` HTTP header
// - performing basic checks, and returning HTTP errors if necessary
// - exposing those credentials in the go context as a UcanCtx for further usage
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctn, err := bearer.ExtractBearerContainer(r.Header)
		if errors.Is(err, bearer.ErrNoUcan) {
			http.Error(w, "no UCAN auth", http.StatusBadRequest)
			return
		}
		if errors.Is(err, bearer.ErrContainerMalformed) {
			http.Error(w, "malformed token", http.StatusBadRequest)
			return
		}
		if err != nil {
			// should not happen, defensive programming
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// prepare a UcanCtx from the container, for further evaluation in the server pipeline
		ucanCtx, err := FromContainer(ctn)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// insert into the go context
		r = r.WithContext(AddUcanCtxToContext(r.Context(), ucanCtx))

		next.ServeHTTP(w, r)
	})
}
