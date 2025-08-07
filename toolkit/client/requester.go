package client

import (
	"context"
	"iter"
	"time"

	"github.com/MetaMask/go-did-it"
	"github.com/avast/retry-go/v4"

	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/token/delegation"
)

type DelegationRequester interface {
	// RequestDelegation retrieve a delegation or chain of delegation for the given parameters.
	// - cmd: the command to execute
	// - audience: the DID of the client, also the issuer of the invocation token
	// - subject: the DID of the resource to operate on, also the subject (or audience if defined) of the invocation token
	// The returned delegations MUST be ordered starting from the leaf (the one matching the invocation) to the root
	// (the one given by the service).
	// Note: you can read it as "(audience) wants to do (cmd) on (subject)".
	// Note: the returned delegation(s) don't have to match exactly the parameters, as long as they allow them.
	RequestDelegation(ctx context.Context, audience did.DID, cmd command.Command, subject did.DID) (iter.Seq2[*delegation.Bundle, error], error)
}

var _ DelegationRequester = &withRetry{}

type withRetry struct {
	requester    DelegationRequester
	initialDelay time.Duration
	maxAttempts  uint
}

// RequesterWithRetry wraps a DelegationRequester to perform exponential backoff,
// with an initial delay and a maximum attempt count.
func RequesterWithRetry(requester DelegationRequester, initialDelay time.Duration, maxAttempt uint) DelegationRequester {
	return &withRetry{
		requester:    requester,
		initialDelay: initialDelay,
		maxAttempts:  maxAttempt,
	}
}

func (w withRetry) RequestDelegation(ctx context.Context, audience did.DID, cmd command.Command, subject did.DID) (iter.Seq2[*delegation.Bundle, error], error) {
	return retry.DoWithData(func() (iter.Seq2[*delegation.Bundle, error], error) {
		return w.requester.RequestDelegation(ctx, audience, cmd, subject)
	},
		retry.Context(ctx),
		retry.Delay(w.initialDelay),
		retry.Attempts(w.maxAttempts),
	)
}
