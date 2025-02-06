package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ucan-wg/go-ucan/did"

	example "github.com/INFURA/go-ucan-toolkit/_example"
	"github.com/INFURA/go-ucan-toolkit/server/exectx"
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

	err := run(ctx, example.ServiceUrl, example.ServiceDid)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run(ctx context.Context, serviceUrl string, serviceDID did.DID) error {
	log.Printf("service DID is %s\n", serviceDID.String())

	// we'll make a simple handling pipeline:
	// - exectx.ExtractMW to extract and decode the UCAN context, verify the service DID
	// - exectx.HttpExtArgsVerify to verify the HTTP policies
	// - exectx.EnforceMW to perform the final UCAN checks
	// - our handler to execute the commands

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ucanCtx, ok := exectx.FromContext(r.Context())
		if !ok {
			http.Error(w, "no ucan-ctx found", http.StatusInternalServerError)
			return
		}

		switch ucanCtx.Command().String() {
		case "/foo/bar":
			log.Printf("handled command %v at %v for %v", ucanCtx.Command(), r.URL.Path, ucanCtx.Invocation().Issuer())
			log.Printf("proof is %v", ucanCtx.Invocation().Proof())
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("OK"))
		default:
			http.Error(w, "unknown UCAN commmand", http.StatusBadRequest)
			return
		}
	})

	handler = exectx.EnforceMW(handler)
	handler = exectx.HttpExtArgsVerify(handler)
	handler = exectx.ExtractMW(handler, serviceDID)

	srv := &http.Server{
		Addr:    serviceUrl,
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
