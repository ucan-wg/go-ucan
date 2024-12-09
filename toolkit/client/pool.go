package client

import (
	"fmt"
	"iter"
	"sync"
	"time"

	"github.com/ipfs/go-cid"
	"github.com/ucan-wg/go-ucan/did"
	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/token/delegation"
)

type Pool struct {
	mu   sync.RWMutex
	dlgs map[cid.Cid]*delegation.Bundle
}

func NewPool() *Pool {
	return &Pool{dlgs: make(map[cid.Cid]*delegation.Bundle)}
}

func (p *Pool) AddBundle(bundle *delegation.Bundle) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.dlgs[bundle.Cid] = bundle
}

func (p *Pool) AddBundles(bundles iter.Seq[*delegation.Bundle]) {
	for bundle := range bundles {
		p.AddBundle(bundle)
	}
}

// FindProof find in the pool the best (shortest, smallest in bytes) chain of delegation(s) matching the given invocation parameters.
// - cmd: the command to execute
// - issuer: the DID of the client, also the issuer of the invocation token
// - audience: the DID of the resource to operate on, also the subject (or audience if defined) of the invocation token
// Note: the returned delegation(s) don't have to match exactly the parameters, as long as they allow them.
// Note: the implemented algorithm won't perform well with a large number of delegations.
func (p *Pool) FindProof(cmd command.Command, iss did.DID, aud did.DID) []cid.Cid {
	// TODO: move to some kind of background trim job?
	p.trim()

	p.mu.RLock()
	defer p.mu.RUnlock()

	return FindProof(func() iter.Seq[*delegation.Bundle] {
		return func(yield func(*delegation.Bundle) bool) {
			for _, bundle := range p.dlgs {
				if !yield(bundle) {
					return
				}
			}
		}
	}, cmd, iss, aud)
}

func (p *Pool) GetBundles(cids []cid.Cid) iter.Seq2[*delegation.Bundle, error] {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return func(yield func(*delegation.Bundle, error) bool) {
		for _, c := range cids {
			if b, ok := p.dlgs[c]; ok {
				if !yield(b, nil) {
					return
				}
			} else {
				yield(nil, fmt.Errorf("bundle not found"))
				return
			}
		}
	}
}

// trim removes expired tokens
func (p *Pool) trim() {
	p.mu.Lock()
	defer p.mu.Unlock()

	now := time.Now()
	for c, bundle := range p.dlgs {
		if bundle.Decoded.Expiration() != nil && bundle.Decoded.Expiration().Before(now) {
			delete(p.dlgs, c)
		}
	}
}
