package invocation

import "errors"

var (
	// ErrDelegationExpired is returned if one of the delegations in the
	// proof chain has expired.
	ErrDelegationExpired = errors.New("delegation in proof chain has expired")

	// ErrDelegationInactive is returned if one of the delegations in the
	// proof chain is not yet active.
	ErrDelegationInactive = errors.New("delegation in proof chain not yet active")

	// ErrLastNotRoot is returned if the last delegation token in the proof
	// chain is not a root delegation token.
	ErrLastNotRoot = errors.New("the last delegation token in proof chain must be a root token")

	// ErrMissingDelegation
	ErrMissingDelegation = errors.New("loader missing delegation for proof chain")

	// ErrNoProof is returned when no delegations were provided to prove
	// that the invocation should be executed.
	ErrNoProof = errors.New("at least one delegation must be provided to validate the invocation")

	// ErrNotIssuedToInvoder is returned if the first delegation token's
	// Audience DID is not the Invoker's Issuer DID.
	ErrNotIssuedToInvoker = errors.New("first delegation token is not issued to invoker")

	// ErrBrokenChain is returned when the Audience of each delegation is
	// not the Issuer of the previous one.
	ErrBrokenChain = errors.New("delegation proof chain is broken")
)
