package issuer

import (
	"context"
	"iter"
	"time"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/ucan-wg/go-ucan/did"
	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/token/delegation"

	"github.com/INFURA/go-ucan-toolkit/client"
)

var _ client.DelegationRequester = &Issuer{}

type Issuer struct {
	did     did.DID
	privKey crypto.PrivKey

	pool      *client.Pool
	requester client.DelegationRequester
	logic     IssuingLogic
}

func NewIssuer(privKey crypto.PrivKey, requester client.DelegationRequester, logic IssuingLogic) (*Issuer, error) {
	d, err := did.FromPrivKey(privKey)
	if err != nil {
		return nil, err
	}
	return &Issuer{
		did:       d,
		privKey:   privKey,
		pool:      client.NewPool(),
		requester: client.RequesterWithRetry(requester, time.Second, 3),
		logic:     logic,
	}, nil
}

// IssuingLogic is a function that decides what powers are given to a client.
// - issuer: the DID of our issuer
// - audience: the DID of the client, also the issuer of the invocation token
// - cmd: the command to execute
// - subject: the DID of the resource to operate on, also the subject (or audience if defined) of the invocation token
// Note: you can read it as "(audience) wants to do (cmd) on (subject)".
// Note: you can decide to match the input parameters exactly or issue a broader power, as long as it allows the
// expected action. If you don't want to give that power, return an error instead.
type IssuingLogic func(iss did.DID, aud did.DID, cmd command.Command, subject did.DID) (*delegation.Token, error)

// RequestDelegation retrieve chain of delegation for the given parameters.
// - audience: the DID of the client, also the issuer of the invocation token
// - cmd: the command to execute
// - subject: the DID of the resource to operate on, also the subject (or audience if defined) of the invocation token
// Note: you can read it as "(audience) does (cmd) on (subject)".
// Note: the returned delegation(s) don't have to match exactly the parameters, as long as they allow them.
func (i *Issuer) RequestDelegation(ctx context.Context, audience did.DID, cmd command.Command, subject did.DID) (iter.Seq2[*delegation.Bundle, error], error) {
	var proof []cid.Cid

	// is there already a valid proof chain?
	if proof = i.pool.FindProof(audience, cmd, subject); len(proof) > 0 {
		return i.pool.GetBundles(proof), nil
	}

	// do we have the power to delegate this?
	if proof = i.pool.FindProof(i.did, cmd, subject); len(proof) == 0 {
		// we need to request a new proof
		proofBundles, err := i.requester.RequestDelegation(ctx, i.did, cmd, subject)
		if err != nil {
			return nil, err
		}

		// cache the new proofs
		for bundle, err := range proofBundles {
			if err != nil {
				return nil, err
			}
			proof = append(proof, bundle.Cid)
			i.pool.AddBundle(bundle)
		}
	}

	// run the custom logic to get what we actually issue
	dlg, err := i.logic(i.did, audience, cmd, subject)
	if err != nil {
		return nil, err
	}

	// sign and cache the new token
	dlgBytes, dlgCid, err := dlg.ToSealed(i.privKey)
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
		for b, err := range i.pool.GetBundles(proof) {
			if !yield(b, err) {
				return
			}
		}
	}, nil
}
