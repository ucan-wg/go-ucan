package client

import (
	"context"
	"fmt"
	"iter"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/ucan-wg/go-ucan/did"
	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/token/delegation"
)

// DlgIssuingLogic is a function that decides what powers are given to a client.
// - issuer: the DID of our issuer
// - audience: the DID of the client, also the issuer of the invocation token
// - cmd: the command to execute
// - subject: the DID of the resource to operate on, also the subject (or audience if defined) of the invocation token
// Note: you can read it as "(audience) wants to do (cmd) on (subject)".
// Note: you can decide to match the input parameters exactly or issue a broader power, as long as it allows the
// expected action. If you don't want to give that power, return an error instead.
type DlgIssuingLogic func(iss did.DID, aud did.DID, cmd command.Command, subject did.DID) (*delegation.Token, error)

var _ DelegationRequester = &WithIssuer{}

type WithIssuer struct {
	*Client
	logic DlgIssuingLogic
}

func NewWithIssuer(privKey crypto.PrivKey, requester DelegationRequester, logic DlgIssuingLogic) (*WithIssuer, error) {
	client, err := NewClient(privKey, requester)
	if err != nil {
		return nil, err
	}
	return &WithIssuer{Client: client, logic: logic}, nil
}

// RequestDelegation retrieve chain of delegation for the given parameters.
// - audience: the DID of the client, also the issuer of the invocation token
// - cmd: the command to execute
// - subject: the DID of the resource to operate on, also the subject (or audience if defined) of the invocation token
// Note: you can read it as "(audience) does (cmd) on (subject)".
// Note: the returned delegation(s) don't have to match exactly the parameters, as long as they allow them.
func (c *WithIssuer) RequestDelegation(ctx context.Context, audience did.DID, cmd command.Command, subject did.DID) (iter.Seq2[*delegation.Bundle, error], error) {
	var proof []cid.Cid

	// is there already a valid proof chain?
	if proof = c.pool.FindProof(audience, cmd, subject); len(proof) > 0 {
		return c.pool.GetBundles(proof), nil
	}

	// do we have the power to delegate this?
	if proof = c.pool.FindProof(c.did, cmd, subject); len(proof) == 0 {
		// we need to request a new proof
		proofBundles, err := c.requester.RequestDelegation(ctx, c.did, cmd, subject)
		if err != nil {
			return nil, err
		}

		// cache the new proofs
		for bundle, err := range proofBundles {
			if err != nil {
				return nil, err
			}
			proof = append(proof, bundle.Cid)
			c.pool.AddBundle(bundle)
		}
	}

	// run the custom logic to get what we actually issue
	dlg, err := c.logic(c.did, audience, cmd, subject)
	if err != nil {
		return nil, err
	}
	if dlg.IsRoot() {
		return nil, fmt.Errorf("issuing logic should return a non-root delegation")
	}

	// sign and cache the new token
	dlgBytes, dlgCid, err := dlg.ToSealed(c.privKey)
	if err != nil {
		return nil, err
	}
	bundle := &delegation.Bundle{
		Cid:     dlgCid,
		Decoded: dlg,
		Sealed:  dlgBytes,
	}

	// output the relevant delegations
	return func(yield func(*delegation.Bundle, error) bool) {
		if !yield(bundle, nil) {
			return
		}
		for b, err := range c.pool.GetBundles(proof) {
			if !yield(b, err) {
				return
			}
		}
	}, nil
}
