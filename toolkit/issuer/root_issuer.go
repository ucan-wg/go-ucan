package issuer

import (
	"context"
	"fmt"
	"iter"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/ucan-wg/go-ucan/did"
	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/token/delegation"

	"github.com/INFURA/go-ucan-toolkit/client"
)

// RootIssuingLogic is a function that decides what powers are given to a client.
// - issuer: the DID of our issuer
// - audience: the DID of the client, also the issuer of the invocation token
// - cmd: the command to execute
// Note: you can read it as "(audience) wants to do (cmd) on (subject)".
// Note: you can decide to match the input parameters exactly or issue a broader power, as long as it allows the
// expected action. If you don't want to give that power, return an error instead.
type RootIssuingLogic func(iss did.DID, aud did.DID, cmd command.Command) (*delegation.Token, error)

var _ client.DelegationRequester = &RootIssuer{}

// RootIssuer is an implementation of a root delegation issuer.
// Note: Your actual needs for an issuer can easily be different (caching...) than the choices made here.
// Feel free to replace this component with your own flavor.
type RootIssuer struct {
	did     did.DID
	privKey crypto.PrivKey

	logic RootIssuingLogic
}

func NewRootIssuer(privKey crypto.PrivKey, logic RootIssuingLogic) (*RootIssuer, error) {
	d, err := did.FromPrivKey(privKey)
	if err != nil {
		return nil, err
	}
	return &RootIssuer{
		did:     d,
		privKey: privKey,
		logic:   logic,
	}, nil
}

func (r *RootIssuer) RequestDelegation(ctx context.Context, audience did.DID, cmd command.Command, subject did.DID) (iter.Seq2[*delegation.Bundle, error], error) {
	if subject != r.did {
		return nil, fmt.Errorf("subject DID doesn't match the issuer DID")
	}

	// run the custom logic to get what we actually issue
	dlg, err := r.logic(r.did, audience, cmd)
	if err != nil {
		return nil, err
	}
	if !dlg.IsRoot() {
		return nil, fmt.Errorf("issuing logic should return a root delegation")
	}

	// sign and cache the new token
	dlgBytes, dlgCid, err := dlg.ToSealed(r.privKey)
	if err != nil {
		return nil, err
	}
	bundle := &delegation.Bundle{
		Cid:     dlgCid,
		Decoded: dlg,
		Sealed:  dlgBytes,
	}

	// output the root delegation
	return func(yield func(*delegation.Bundle, error) bool) {
		yield(bundle, nil)
	}, nil
}
