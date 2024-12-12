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
// - issuer: the DID of the client, also the issuer of the invocation token
// - cmd: the command to execute
// - subject: the DID of the resource to operate on, also the subject (or audience if defined) of the invocation token
// Note: you can read it as "(issuer) wants to do (cmd) on (subject)".
// Note: the returned delegation(s) don't have to match exactly the parameters, as long as they allow them.
// Note: the implemented algorithm won't perform well with a large number of delegations.
func FindProof(dlgs func() iter.Seq[*delegation.Bundle], issuer did.DID, cmd command.Command, subject did.DID) []cid.Cid {
	// TODO: maybe that should be part of delegation.Token directly?
	dlgMatch := func(dlg *delegation.Token, issuer did.DID, cmd command.Command, subject did.DID) bool {
		// The Subject of each delegation must equal the invocation's Subject (or Audience if defined). - 4f
		if dlg.Subject() != subject {
			return false
		}
		// The first proof must be issued to the Invoker (audience DID). - 4c
		// The Issuer of each delegation must be the Audience in the next one. - 4d
		if dlg.Audience() != issuer {
			return false
		}
		// The command of each delegation must "allow" the one before it. - 4g
		if !dlg.Command().Covers(cmd) {
			return false
		}
		// Time bound - 3b, 3c
		if !dlg.IsValidNow() {
			return false
		}
		return true
	}

	// STEP 1: Find the possible leaf delegations, directly matching the invocation parameters
	var candidateLeaf []*delegation.Bundle

	for bundle := range dlgs() {
		if !dlgMatch(bundle.Decoded, issuer, cmd, subject) {
			continue
		}
		candidateLeaf = append(candidateLeaf, bundle)
	}

	// STEP 2: Perform a depth-first search on the DAG of connected delegations, for each of our candidates
	type state struct {
		bundle *delegation.Bundle
		path   []cid.Cid
		size   int
	}

	var bestSize = math.MaxInt
	var bestProof []cid.Cid

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
				if !dlgMatch(candidate.Decoded, at.Decoded.Issuer(), at.Decoded.Command(), subject) {
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
