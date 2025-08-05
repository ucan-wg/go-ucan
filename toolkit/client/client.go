package client

import (
	"context"
	"fmt"
	"iter"

	"github.com/MetaMask/go-did-it"
	"github.com/MetaMask/go-did-it/crypto"
	"github.com/ipfs/go-cid"

	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/pkg/container"
	"github.com/ucan-wg/go-ucan/pkg/policy"
	"github.com/ucan-wg/go-ucan/token/delegation"
	"github.com/ucan-wg/go-ucan/token/invocation"
)

type Client struct {
	did     did.DID
	privKey crypto.PrivateKeySigningBytes

	pool      *Pool
	requester DelegationRequester
}

func NewClient(privKey crypto.PrivateKeySigningBytes, d did.DID, requester DelegationRequester) (*Client, error) {
	return &Client{
		did:       d,
		privKey:   privKey,
		pool:      NewPool(),
		requester: requester,
	}, nil
}

// AddBundle adds a delegation.Bundle to the client's pool.
func (c *Client) AddBundle(bundle *delegation.Bundle) {
	c.pool.AddBundle(bundle)
}

// AddBundles adds a sequence of delegation.Bundles to the client's pool.
func (c *Client) AddBundles(bundles iter.Seq[*delegation.Bundle]) {
	c.pool.AddBundles(bundles)
}

// PrepareInvoke returns an invocation, bundled in a container.Writer with the necessary proofs.
func (c *Client) PrepareInvoke(ctx context.Context, cmd command.Command, subject did.DID, opts ...invocation.Option) (container.Writer, error) {
	proof, err := c.findProof(ctx, cmd, subject)
	if err != nil {
		return nil, err
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

// PrepareDelegation returns a new delegation for a third party DID, bundled in a container.Writer with the necessary proofs.
func (c *Client) PrepareDelegation(ctx context.Context, aud did.DID, cmd command.Command, subject did.DID, policies policy.Policy, opts ...delegation.Option) (container.Writer, error) {
	proof, err := c.findProof(ctx, cmd, subject)
	if err != nil {
		return nil, err
	}

	dlg, err := delegation.New(c.did, aud, cmd, policies, subject, opts...)
	if err != nil {
		return nil, err
	}

	dlgSealed, _, err := dlg.ToSealed(c.privKey)
	if err != nil {
		return nil, err
	}

	cont := container.NewWriter()
	cont.AddSealed(dlgSealed)
	for bundle, err := range c.pool.GetBundles(proof) {
		if err != nil {
			return nil, err
		}
		cont.AddSealed(bundle.Sealed)
	}

	return cont, nil
}

func (c *Client) findProof(ctx context.Context, cmd command.Command, subject did.DID) ([]cid.Cid, error) {
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

	return proof, nil
}
