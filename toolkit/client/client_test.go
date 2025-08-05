package client

import (
	"context"
	"fmt"
	"iter"
	"time"

	"github.com/MetaMask/go-did-it"
	"github.com/MetaMask/go-did-it/didtest"

	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/pkg/policy"
	"github.com/ucan-wg/go-ucan/token/delegation"
	"github.com/ucan-wg/go-ucan/token/invocation"
)

func ExampleNewClient() {
	servicePersona := didtest.PersonaAlice
	clientPersona := didtest.PersonaBob

	// requester is an adaptor for a real world issuer, we use a mock in that example
	requester := &requesterMock{persona: servicePersona}

	client, err := NewClient(clientPersona.PrivKey(), clientPersona.DID(), requester)
	handleError(err)

	cont, err := client.PrepareInvoke(
		context.Background(),
		command.New("crud", "add"),
		servicePersona.DID(),
		// extra invocation parameters:
		invocation.WithExpirationIn(10*time.Minute),
		invocation.WithArgument("foo", "bar"),
		invocation.WithMeta("baz", 1234),
	)
	handleError(err)

	// this container holds the invocation and all the delegation proofs
	b64, err := cont.ToBase64StdPadding()
	handleError(err)

	fmt.Println(b64)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

type requesterMock struct {
	persona didtest.Persona
}

func (r requesterMock) RequestDelegation(_ context.Context, audience did.DID, cmd command.Command, _ did.DID) (iter.Seq2[*delegation.Bundle, error], error) {
	// the mock issue whatever the client asks:
	dlg, err := delegation.Root(r.persona.DID(), audience, cmd, policy.Policy{})
	if err != nil {
		return nil, err
	}

	dlgBytes, dlgCid, err := dlg.ToSealed(r.persona.PrivKey())
	if err != nil {
		return nil, err
	}

	bundle := &delegation.Bundle{
		Cid:     dlgCid,
		Decoded: dlg,
		Sealed:  dlgBytes,
	}

	return func(yield func(*delegation.Bundle, error) bool) {
		yield(bundle, nil)
	}, nil
}
