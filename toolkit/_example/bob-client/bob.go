package main

import (
	"context"
	"fmt"
	"io"
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

	example "github.com/INFURA/go-ucan-toolkit/_example"
	protocol "github.com/INFURA/go-ucan-toolkit/_example/_protocol-issuer"
	"github.com/INFURA/go-ucan-toolkit/client"
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

	err := run(ctx, example.AliceIssuerUrl, example.AliceDid, example.ServiceUrl, example.ServiceDid)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run(ctx context.Context, aliceUrl string, aliceDid did.DID, serverUrl string, serviceDid did.DID) error {
	// Let's generate a keypair for our client:
	priv, d, err := did.GenerateEd25519()
	if err != nil {
		return err
	}

	log.Printf("Bob DID is %s", d.String())

	cli, err := client.NewClient(priv, protocol.NewRequester(aliceUrl))
	if err != nil {
		return err
	}

	for {
		proofs, err := cli.PrepareInvoke(ctx, command.MustParse("/foo/bar"), serviceDid)
		if err != nil {
			return err
		}

		err = makeRequest(ctx, d, serverUrl, aliceDid, proofs)
		if err != nil {
			log.Println(err)
		}

		select {
		case <-ctx.Done():
			return nil
		case <-time.After(5 * time.Second):
		}
	}
}

func makeRequest(ctx context.Context, clientDid did.DID, serviceUrl string, aliceDid did.DID, proofs container.Writer) error {
	// we construct a URL that include the our DID and Alice DID as path, as requested by the UCAN policy we get issued
	u, err := url.JoinPath("http://"+serviceUrl, aliceDid.String(), clientDid.String())
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
