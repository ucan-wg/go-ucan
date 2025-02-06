package protocol

import (
	"bytes"
	"context"
	"encoding/json"
	"iter"
	"log"
	"net/http"

	"github.com/ucan-wg/go-ucan/did"
	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/token/delegation"

	"github.com/INFURA/go-ucan-toolkit/client"
	"github.com/INFURA/go-ucan-toolkit/issuer"
)

var _ client.DelegationRequester = &Requester{}

type Requester struct {
	issuerURL string
}

func NewRequester(issuerURL string) *Requester {
	return &Requester{issuerURL: issuerURL}
}

func (r Requester) RequestDelegation(ctx context.Context, audience did.DID, cmd command.Command, subject did.DID) (iter.Seq2[*delegation.Bundle, error], error) {
	log.Printf("requesting delegation for %s on %s", cmd, subject)

	// we match the simple json protocol of the issuer
	data := struct {
		Audience string `json:"aud"`
		Cmd      string `json:"cmd"`
		Subject  string `json:"sub"`
	}{
		Audience: audience.String(),
		Cmd:      cmd.String(),
		Subject:  subject.String(),
	}
	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, "http://"+r.issuerURL, buf)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	return issuer.DecodeResponse(res)
}
