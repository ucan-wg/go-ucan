package invocation

import (
	"fmt"
	"time"

	"github.com/ipfs/go-cid"

	"github.com/ucan-wg/go-ucan/pkg/args"
	"github.com/ucan-wg/go-ucan/pkg/policy"
	"github.com/ucan-wg/go-ucan/token/delegation"
)

// # Invocation token validation
//
// Per the specification, invocation Tokens must be validated before the command is executed.
// This validation effectively happens in multiple places in the codebase.
// Steps 1 and 2 are the same for all token types.
//
//  1. When a token is read/unsealed from its containing envelope (`envelope` package):
//     a. The envelope can be decoded.
//     b. The envelope contains a Signature, VarsigHeader and Payload.
//     c. The Payload contains an iss field that contains a valid `did:key`.
//     d. The public key can be extracted from the `did:key`.
//     e. The public key type is supported by go-ucan.
//     f. The Signature can be decoded per the VarsigHeader.
//     g. The SigPayload can be verified using the Signature and public key.
//     h. The field key of the TokenPayload matches the expected tag.
//
//  2. When the token is created or passes step one (token constructor or decoder):
//     a. All required fields are present
//     b. All populated fields respect their own rules (example: a policy is legal)
//
//  3. When an unsealed invocation passes steps one and two for execution (verifyTimeBound below):
//     a. The invocation cannot be expired (expiration in the future or absent).
//     b. All the delegation must not be expired (expiration in the future or absent).
//     c. All the delegation must be active (nbf in the past or absent).
//
//  4. When the proof chain is being validated (verifyProofs below):
//     a. There must be at least one delegation in the proof chain.
//     b. All referenced delegations must be available.
//     c. The first proof must be issued to the Invoker (audience DID).
//     d. The Issuer of each delegation must be the Audience in the next one.
//     e. The last token must be a root delegation.
//     f. The Subject of each delegation must equal the invocation's Audience field.
//     g. The command of each delegation must "allow" the one before it.
//
//  5. If steps 1-4 pass:
//     a. The policy must "match" the arguments. (verifyArgs below)
//     b. The nonce (if present) is not reused. (out of scope for go-ucan)

type DelegationLoader interface {
	GetDelegation(cid cid.Cid) (*delegation.Token, error)
}

// verifyProofs controls that the proof chain allows the invocation:
// - principal alignment
// - command alignment
func (t *Token) verifyProofs(delegations []*delegation.Token) error {
	// There must be at least one delegation referenced - 4a
	if len(delegations) < 1 {
		return ErrNoProof
	}

	cmd := t.command
	iss := t.issuer
	aud := t.audience
	if !aud.Defined() {
		aud = t.subject
	}

	// control from the invocation to the root
	for i, dlgCid := range t.proof {
		dlg := delegations[i]

		// The Subject of each delegation must equal the invocation's Audience field. - 4f
		if dlg.Subject() != aud {
			return fmt.Errorf("%w: delegation %s, expected %s, got %s", ErrWrongSub, dlgCid, aud, dlg.Subject())
		}

		// The first proof must be issued to the Invoker (audience DID). - 4c
		// The Issuer of each delegation must be the Audience in the next one. - 4d
		if dlg.Audience() != iss {
			return fmt.Errorf("%w: delegation %s, expected %s, got %s", ErrBrokenChain, dlgCid, iss, dlg.Audience())
		}
		iss = dlg.Issuer()

		// The command of each delegation must "allow" the one before it. - 4g
		if !dlg.Command().Covers(cmd) {
			return fmt.Errorf("%w: delegation %s, %s doesn't cover %s", ErrCommandNotCovered, dlgCid, dlg.Command(), cmd)
		}
		cmd = dlg.Command()
	}

	// The last prf value must be a root delegation (have the issuer field match the Subject field) - 4e
	if last := delegations[len(delegations)-1]; last.Issuer() != last.Subject() {
		return fmt.Errorf("%w: expected %s, got %s", ErrLastNotRoot, last.Subject(), last.Issuer())
	}

	return nil
}

func (t *Token) verifyTimeBound(dlgs []*delegation.Token) error {
	return t.verifyTimeBoundAt(time.Now(), dlgs)
}

func (t *Token) verifyTimeBoundAt(at time.Time, delegations []*delegation.Token) error {
	// The invocation cannot be expired (expiration in the future or absent). - 3a
	if !t.IsValidAt(at) {
		return fmt.Errorf("%w: invocation", ErrTokenInvalidNow)
	}

	for i, dlgCid := range t.proof {
		dlg := delegations[i]

		// All the delegation must not be expired (expiration in the future or absent). - 3b
		// All the delegation must be active (nbf in the past or absent). - 3c
		if !dlg.IsValidAt(at) {
			return fmt.Errorf("%w: delegation %s", ErrTokenInvalidNow, dlgCid)
		}
	}
	return nil
}

func (t *Token) verifyArgs(delegations []*delegation.Token, arguments *args.Args) error {
	var count int
	for i := range t.proof {
		count += len(delegations[i].Policy())
	}

	policies := make(policy.Policy, 0, count)
	for i := range t.proof {
		policies = append(policies, delegations[i].Policy()...)
	}

	argsIpld, err := arguments.ToIPLD()
	if err != nil {
		return err
	}

	ok, statement := policies.Match(argsIpld)
	if !ok {
		return fmt.Errorf("the following UCAN policy is not satisfied: %v", statement.String())
	}

	return nil
}
