package client

import (
	"iter"
	"math"

	"github.com/ipfs/go-cid"
	"github.com/ucan-wg/go-ucan/did"
	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/token/delegation"
)

// FindProof find in the pool the best (shortest, smallest in bytes) chain of delegation(s) matching the given invocation parameters.
// - cmd: the command to execute
// - issuer: the DID of the client, also the issuer of the invocation token
// - audience: the DID of the resource to operate on, also the subject (or audience if defined) of the invocation token
// Note: the returned delegation(s) don't have to match exactly the parameters, as long as they allow them.
// Note: the implemented algorithm won't perform well with a large number of delegations.
func FindProof(dlgs func() iter.Seq[*delegation.Bundle], cmd command.Command, iss did.DID, aud did.DID) []cid.Cid {
	// Find the possible leaf delegations, directly matching the invocation parameters
	var candidateLeaf []*delegation.Bundle

	for bundle := range dlgs() {
		dlg := bundle.Decoded

		// The Subject of each delegation must equal the invocation's Audience field. - 4f
		if dlg.Subject() != aud {
			continue
		}
		// The first proof must be issued to the Invoker (audience DID). - 4c
		// The Issuer of each delegation must be the Audience in the next one. - 4d
		if dlg.Audience() != iss {
			continue
		}
		// The command of each delegation must "allow" the one before it. - 4g
		if !dlg.Command().Covers(cmd) {
			continue
		}
		// Time bound - 3b, 3c
		if !dlg.IsValidNow() {
			continue
		}

		candidateLeaf = append(candidateLeaf, bundle)
	}

	type state struct {
		bundle *delegation.Bundle
		path   []cid.Cid
		size   int
	}

	var bestSize = math.MaxInt
	var bestProof []cid.Cid

	// Perform a depth-first search on the DAG of connected delegations, for each of our candidates
	for _, leaf := range candidateLeaf {
		var stack = []state{{bundle: leaf, path: []cid.Cid{leaf.Cid}, size: len(leaf.Sealed)}}

		for len(stack) > 0 {
			// dequeue a delegation
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			at := cur.bundle

			// if it's a root delegation, we found a valid proof
			if at.Decoded.Issuer() == at.Decoded.Subject() {
				if len(bestProof) == 0 || len(cur.path) < len(bestProof) || len(cur.path) == len(bestProof) && cur.size < bestSize {
					bestProof = append([]cid.Cid{}, cur.path...) // make a copy
					bestSize = cur.size
					continue
				}
			}

			// find parent delegation for our current delegation
			for candidate := range dlgs() {
				// The Subject of each delegation must equal the invocation's Audience field. - 4f
				if candidate.Decoded.Subject() != aud {
					continue
				}
				// The first proof must be issued to the Invoker (audience DID). - 4c
				// The Issuer of each delegation must be the Audience in the next one. - 4d
				if candidate.Decoded.Audience() != at.Decoded.Issuer() {
					continue
				}
				// The command of each delegation must "allow" the one before it. - 4g
				if !candidate.Decoded.Command().Covers(at.Decoded.Command()) {
					continue
				}
				// Time bound - 3b, 3c
				if !candidate.Decoded.IsValidNow() {
					continue
				}

				newPath := append([]cid.Cid{}, cur.path...) // make a copy
				newPath = append(newPath, candidate.Cid)
				stack = append(stack, state{bundle: candidate, path: newPath, size: cur.size + len(candidate.Sealed)})
			}
		}
	}

	return bestProof
}
