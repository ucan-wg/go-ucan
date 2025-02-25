package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/ucan-wg/go-ucan/did"
	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/pkg/container"
	"github.com/ucan-wg/go-ucan/pkg/policy"
	"github.com/ucan-wg/go-ucan/pkg/policy/literal"
	"github.com/ucan-wg/go-ucan/token/delegation"

	example "github.com/INFURA/go-ucan-toolkit/_example"
	protocol "github.com/INFURA/go-ucan-toolkit/_example/_protocol-issuer"
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

	err := run(ctx,
		example.AliceIssuerUrl, example.AlicePrivKey, example.AliceDid,
		example.ServiceIssuerUrl, example.ServiceUrl, example.ServiceDid)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run(ctx context.Context, ownIssuerUrl string, priv crypto.PrivKey, d did.DID,
	serviceIssuerUrl string, serviceUrl string, serviceDid did.DID) error {
	log.Printf("Alice DID is %s", d.String())

	issuingLogic := func(iss did.DID, aud did.DID, cmd command.Command, subject did.DID) (*delegation.Token, error) {
		log.Printf("issuing delegation to %v for %v to operate on %v", aud, cmd, subject)

		// As another example, we'll force Bob to use a specific HTTP sub-path
		policies, err := policy.Construct(
			policy.Equal(".http.path", literal.String(fmt.Sprintf("/%s/%s", iss.String(), aud.String()))),
		)
		if err != nil {
			return nil, err
		}

		return delegation.New(iss, aud, cmd, policies, subject)
	}

	cli, err := client.NewWithIssuer(priv, protocol.NewRequester("http://"+serviceIssuerUrl), issuingLogic)
	if err != nil {
		return err
	}

	go startIssuerHttp(ctx, ownIssuerUrl, cli)

	for {
		proofs, err := cli.PrepareInvoke(ctx, command.MustParse("/foo/bar"), serviceDid)
		if err != nil {
			return err
		}

		err = makeRequest(ctx, d, serviceUrl, proofs)
		if err != nil {
			log.Println(err)
		}

		select {
		case <-ctx.Done():
			return nil
		case <-time.After(20 * time.Second):
		}
	}
}

func startIssuerHttp(ctx context.Context, issuerUrl string, cli *client.WithIssuer) {
	handler := issuer.HttpWrapper(cli, protocol.RequestResolver)

	srv := &http.Server{
		Addr:    issuerUrl,
		Handler: handler,
	}

	go func() {
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Printf("issuer listening on %s\n", srv.Addr)

	<-ctx.Done()

	if err := srv.Shutdown(ctx); err != nil && !errors.Is(err, context.Canceled) {
		log.Printf("issuer error: %v\n", err)
	}
}

func makeRequest(ctx context.Context, clientDid did.DID, serviceUrl string, proofs container.Writer) error {
	// we construct a URL that include the client DID as path, as requested by the UCAN policy we get issued
	u, err := url.JoinPath("http://"+serviceUrl, clientDid.String())
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
