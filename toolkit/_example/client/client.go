package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"iter"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ucan-wg/go-ucan/did"
	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/pkg/container"
	"github.com/ucan-wg/go-ucan/token/delegation"

	example "github.com/INFURA/go-ucan-toolkit/_example"
	"github.com/INFURA/go-ucan-toolkit/client"
	"github.com/INFURA/go-ucan-toolkit/issuer"
	"github.com/INFURA/go-ucan-toolkit/server/bearer"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// register as handler of the interrupt signal to trigger the teardown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	go func() {
		<-quit
		cancel()
	}()

	err := run(ctx, example.IssuerUrl, example.ServerUrl, example.ServiceDid)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

var _ client.DelegationRequester = &requester{}

type requester struct {
	issuerURL string
}

func (r requester) RequestDelegation(ctx context.Context, audience did.DID, cmd command.Command, subject did.DID) (iter.Seq2[*delegation.Bundle, error], error) {
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

func run(ctx context.Context, issuerUrl string, serverUrl string, serviceDid did.DID) error {
	// Let's generate a keypair for our client:
	priv, d, err := did.GenerateEd25519()
	if err != nil {
		return err
	}

	log.Printf("client DID is %s", d.String())

	cli, err := client.NewClient(priv, requester{issuerURL: issuerUrl})
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
		case <-time.After(5 * time.Second):
			proofs, err := cli.PrepareInvoke(ctx, command.MustParse("/foo/bar"), serviceDid)
			if err != nil {
				return err
			}

			err = makeRequest(ctx, d, serverUrl, proofs)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func makeRequest(ctx context.Context, clientDid did.DID, serverUrl string, proofs container.Writer) error {
	// we construct a URL that include the client DID as path, as requested by the UCAN policy we get issued
	u, err := url.JoinPath("http://"+serverUrl, clientDid.String())
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return err
	}

	err = bearer.AddBearerContainerCompressed(req.Header, proofs)
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("unexpected status code: %d, error reading body: %w", res.StatusCode, err)
		}
		return fmt.Errorf("unexpected status code: %d, body: %v", res.StatusCode, string(body))
	}

	log.Printf("response status code: %d", res.StatusCode)

	return nil
}
