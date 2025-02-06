package protocol

import (
	"encoding/json"
	"net/http"

	"github.com/ucan-wg/go-ucan/did"
	"github.com/ucan-wg/go-ucan/pkg/command"

	"github.com/INFURA/go-ucan-toolkit/issuer"
)

func RequestResolver(r *http.Request) (*issuer.ResolvedRequest, error) {
	// Let's make up a simple json protocol
	req := struct {
		Audience string `json:"aud"`
		Cmd      string `json:"cmd"`
		Subject  string `json:"sub"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	aud, err := did.Parse(req.Audience)
	if err != nil {
		return nil, err
	}
	cmd, err := command.Parse(req.Cmd)
	if err != nil {
		return nil, err
	}
	sub, err := did.Parse(req.Subject)
	if err != nil {
		return nil, err
	}
	return &issuer.ResolvedRequest{
		Audience: aud,
		Cmd:      cmd,
		Subject:  sub,
	}, nil
}
