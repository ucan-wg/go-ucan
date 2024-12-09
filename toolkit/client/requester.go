package client

import (
	"context"
	"iter"

	"github.com/ucan-wg/go-ucan/did"
	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/token/delegation"
)

type DelegationRequester interface {
	// RequestDelegation retrieve a delegation or chain of delegation for the given parameters.
	// - cmd: the command to execute
	// - issuer: the DID of the client, also the issuer of the invocation token
	// - audience: the DID of the resource to operate on, also the subject (or audience if defined) of the invocation token
	// Note: the returned delegation(s) don't have to match exactly the parameters, as long as they allow them.
	RequestDelegation(ctx context.Context, cmd command.Command, audience did.DID, subject did.DID) (iter.Seq2[*delegation.Bundle, error], error)
}
