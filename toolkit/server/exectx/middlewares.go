package exectx

import (
	"errors"
	"net/http"

	"github.com/ucan-wg/go-ucan/did"

	"github.com/INFURA/go-ucan-toolkit/server/bearer"
)

// ExtractMW returns an HTTP middleware tasked with:
// - extracting UCAN credentials from the `Authorization: Bearer <data>` HTTP header
// - performing basic checks, and returning HTTP errors if necessary
// - verify that the invocation targets our service
// - exposing those credentials in the go context as a UcanCtx for further usage
func ExtractMW(next http.Handler, serviceDID did.DID) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctn, err := bearer.ExtractBearerContainer(r.Header)
		if errors.Is(err, bearer.ErrNoUcan) {
			http.Error(w, "no UCAN auth", http.StatusUnauthorized)
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
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if ucanCtx.Invocation().Subject() != serviceDID {
			http.Error(w, "UCAN delegation doesn't match the service DID", http.StatusUnauthorized)
			return
		}

		// insert into the go context
		r = r.WithContext(AddUcanCtxToContext(r.Context(), ucanCtx))

		next.ServeHTTP(w, r)
	})
}

// HttpExtArgsVerify returns an HTTP middleware tasked with verifying the UCAN policies applying on the HTTP request.
func HttpExtArgsVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ucanCtx, ok := FromContext(r.Context())
		if !ok {
			http.Error(w, "no ucan-ctx found", http.StatusInternalServerError)
			return
		}

		if err := ucanCtx.VerifyHttp(r); err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// EnforceMW returns an HTTP middleware tasked with the final verification of the UCAN policies.
func EnforceMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ucanCtx, ok := FromContext(r.Context())
		if !ok {
			http.Error(w, "no ucan-ctx found", http.StatusInternalServerError)
			return
		}

		if err := ucanCtx.ExecutionAllowed(); err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
