package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MetaMask/go-did-it"
	"github.com/MetaMask/go-did-it/crypto"

	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/pkg/policy"
	"github.com/ucan-wg/go-ucan/pkg/policy/literal"
	"github.com/ucan-wg/go-ucan/token/delegation"

	example "github.com/ucan-wg/go-ucan/toolkit/_example"
	protocol "github.com/ucan-wg/go-ucan/toolkit/_example/_protocol-issuer"
	"github.com/ucan-wg/go-ucan/toolkit/issuer"
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

	err := run(ctx, example.ServiceIssuerUrl, example.ServicePrivKey)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run(ctx context.Context, issuerUrl string, servicePrivKey crypto.PrivateKeySigningBytes) error {
	issuingLogic := func(iss did.DID, aud did.DID, cmd command.Command) (*delegation.Token, error) {
		log.Printf("issuing delegation to %v for %v", aud, cmd)

		// We construct an arbitrary policy.
		// Here, we enforce that the caller uses its own DID endpoint (an arbitrary construct for this example).
		// You will notice that the server doesn't need to know about this logic to enforce it.
		policies, err := policy.Construct(
			policy.Or(
				// allow exact path
				policy.Equal(".http.path", literal.String(fmt.Sprintf("/%s", aud.String()))),
				// allow sub-path
				policy.Like(".http.path", fmt.Sprintf("/%s/*", aud.String())),
			),
		)
		if err != nil {
			return nil, err
		}

		return delegation.Root(iss, aud, cmd, policies,
			// let's add an expiration, this will force the client to renew its token.
			delegation.WithExpirationIn(10*time.Second),
		)
	}

	rootIssuer, err := issuer.NewRootIssuer(servicePrivKey, issuingLogic)
	if err != nil {
		return err
	}

	handler := issuer.HttpWrapper(rootIssuer, protocol.RequestResolver)

	srv := &http.Server{
		Addr:    issuerUrl,
		Handler: handler,
	}

	go func() {
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Printf("listening on %s\n", srv.Addr)

	<-ctx.Done()

	if err := srv.Shutdown(ctx); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}
