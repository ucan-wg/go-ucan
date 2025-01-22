package client

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"iter"
	"net/http"
	"net/url"

	"github.com/ucan-wg/go-ucan/did"
	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/token/delegation"
)

var _ DelegationRequester = &InfuraRequester{}

type InfuraRequester struct {
	baseURL string
}

// NewInfuraRequester create a requester client for the Infura UCAN token issuer.
// dev: http://ucan-issuer-api.commercial-dev.eks-dev.infura.org
// prod: http://ucan-issuer-api.commercial-prod.eks.infura.org
func NewInfuraRequester(baseURL string) *InfuraRequester {
	return &InfuraRequester{baseURL: baseURL}
}

func (i InfuraRequester) RequestDelegation(ctx context.Context, audience did.DID, cmd command.Command, subject did.DID) (iter.Seq2[*delegation.Bundle, error], error) {
	p, err := url.JoinPath(i.baseURL, "v1/token/generate-with-did")
	if err != nil {
		return nil, err
	}

	payload := struct {
		Cmd string `json:"cmd"`
		Aud string `json:"aud"`
	}{
		Cmd: cmd.String(),
		Aud: audience.String(),
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	r, err := http.NewRequest(http.MethodPost, p, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(r.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		msg, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("request failed with status %d, then failed to read response body: %w", res.StatusCode, err)
		}
		return nil, fmt.Errorf("request failed with status %d: %s", res.StatusCode, msg)
	}

	resp := struct {
		Cid     string `json:"cid"`
		Content string `json:"content"`
	}{}
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return nil, err
	}

	raw, err := base64.StdEncoding.DecodeString(resp.Content)
	if err != nil {
		return nil, err
	}

	tkn, c, err := delegation.FromSealed(raw)
	if err != nil {
		return nil, err
	}

	// For sanity, we verify that the delegation we got matches the expected subject,
	// meaning that we are talking to the expected issuer.
	if tkn.Subject() != subject {
		return nil, fmt.Errorf("received token has unexpected subject: expected %s, got %s", subject, tkn.Subject())
	}

	return func(yield func(*delegation.Bundle, error) bool) {
		yield(&delegation.Bundle{
			Cid:     c,
			Decoded: tkn,
			Sealed:  raw,
		}, nil)
	}, nil
}
