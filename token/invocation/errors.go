package invocation

import "errors"

// Loading errors
var (
	// ErrMissingDelegation
	ErrMissingDelegation = errors.New("loader missing delegation for proof chain")
)

// Time bound errors
var (
	// ErrTokenExpired is returned if a token is invalid at execution time
	ErrTokenInvalidNow = errors.New("token has expired")
)

// Principal alignment errors
var (
	// ErrNoProof is returned when no delegations were provided to prove
	// that the invocation should be executed.
	ErrNoProof = errors.New("at least one delegation must be provided to validate the invocation")

	// ErrLastNotRoot is returned if the last delegation token in the proof
	// chain is not a root delegation token.
	ErrLastNotRoot = errors.New("the last delegation token in proof chain must be a root token")

	// ErrBrokenChain is returned when the Audience of a delegation is
	// not the Issuer of the previous one.
	ErrBrokenChain = errors.New("delegation proof chain doesn't connect the invocation to the subject")

	// ErrWrongSub is returned when the Subject of a delegation is not the invocation audience.
	ErrWrongSub = errors.New("delegation subject need to match the invocation audience")

	// ErrCommandNotCovered is returned when a delegation command doesn't cover (identical or parent of) the
	// next delegation or invocation's command.
	ErrCommandNotCovered = errors.New("allowed command doesn't cover the next delegation or invocation")
)

// ErrPolicyNotSatisfied is returned when the provided Arguments don't
// satisfy one or more of the aggregated Policy Statements
var ErrPolicyNotSatisfied = errors.New("the following UCAN policy is not satisfied")
