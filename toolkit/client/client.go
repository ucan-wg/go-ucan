package client

import (
	"context"
	"fmt"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/ucan-wg/go-ucan/did"
	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/pkg/container"
	"github.com/ucan-wg/go-ucan/token/invocation"
)

type Client struct {
	did     did.DID
	privKey crypto.PrivKey

	pool      *Pool
	requester DelegationRequester
}

func NewClient(privKey crypto.PrivKey, requester DelegationRequester) (*Client, error) {
	d, err := did.FromPrivKey(privKey)
	if err != nil {
		return nil, err
	}
	return &Client{
		did:       d,
		privKey:   privKey,
		pool:      NewPool(),
		requester: requester,
	}, nil
}

// PrepareInvoke returns an invocation and the proof delegation, bundled in a container.Writer.
func (c *Client) PrepareInvoke(ctx context.Context, cmd command.Command, subject did.DID, opts ...invocation.Option) (container.Writer, error) {
	var proof []cid.Cid

	// do we already have a valid proof?
	if proof = c.pool.FindProof(c.did, cmd, subject); len(proof) == 0 {
		// we need to request a new proof
		proofBundles, err := c.requester.RequestDelegation(ctx, c.did, cmd, subject)
		if err != nil {
			return nil, fmt.Errorf("requesting delegation: %w", err)
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

	inv, err := invocation.New(c.did, cmd, subject, proof, opts...)
	if err != nil {
		return nil, err
	}

	invSealed, _, err := inv.ToSealed(c.privKey)
	if err != nil {
		return nil, err
	}

	cont := container.NewWriter()
	cont.AddSealed(invSealed)
	for bundle, err := range c.pool.GetBundles(proof) {
		if err != nil {
			return nil, err
		}
		cont.AddSealed(bundle.Sealed)
	}

	return cont, nil
}
