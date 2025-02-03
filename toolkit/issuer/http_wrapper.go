package issuer

import (
	"fmt"
	"io"
	"iter"
	"net/http"

	"github.com/ucan-wg/go-ucan/did"
	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/pkg/container"
	"github.com/ucan-wg/go-ucan/token/delegation"

	"github.com/INFURA/go-ucan-toolkit/client"
)

type RequestResolver func(r *http.Request) (*ResolvedRequest, error)

type ResolvedRequest struct {
	Audience did.DID
	Cmd      command.Command
	Subject  did.DID
}

// HttpWrapper implements an HTTP transport for a UCAN issuer.
// It provides a common response shape, while allowing customisation of the request.
func HttpWrapper(requester client.DelegationRequester, resolver RequestResolver) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		resolved, err := resolver(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		dlgs, err := requester.RequestDelegation(r.Context(), resolved.Audience, resolved.Cmd, resolved.Subject)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		cont := container.NewWriter()
		for bundle, err := range dlgs {
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			cont.AddSealed(bundle.Sealed)
		}

		err = cont.ToBytesGzippedWriter(w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func DecodeResponse(res *http.Response) (iter.Seq2[*delegation.Bundle, error], error) {
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		msg, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("request failed with status %d, then failed to read response body: %w", res.StatusCode, err)
		}
		return nil, fmt.Errorf("request failed with status %d: %s", res.StatusCode, msg)
	}
	cont, err := container.FromReader(res.Body)
	if err != nil {
		return nil, err
	}
	return func(yield func(*delegation.Bundle, error) bool) {
		for bundle := range cont.GetAllDelegations() {
			if !yield(&bundle, nil) {
				return
			}
		}
	}, nil
}
