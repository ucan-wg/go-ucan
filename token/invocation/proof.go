package invocation

import (
	"fmt"
	"time"

	"github.com/ipfs/go-cid"

	"github.com/ucan-wg/go-ucan/pkg/args"
	"github.com/ucan-wg/go-ucan/pkg/policy"
	"github.com/ucan-wg/go-ucan/token/delegation"
)

type DelegationLoader interface {
	GetDelegation(cid cid.Cid) (*delegation.Token, error)
}

// verifyProofs controls that the proof chain allows the invocation:
// - principal alignment
// - command alignment
func (t *Token) verifyProofs(delegations []*delegation.Token) error {
	cmd := t.command
	iss := t.issuer
	aud := t.audience
	if !aud.Defined() {
		aud = t.subject
	}

	// control from the invocation to the root
	for i, dlgCid := range t.proof {
		dlg := delegations[i]

		if dlg.Subject() != aud {
			return fmt.Errorf("%w: delegation %s, expected %s, got %s", ErrWrongSub, dlgCid, aud, dlg.Subject())
		}

		if dlg.Audience() != iss {
			return fmt.Errorf("%w: delegation %s, expected %s, got %s", ErrBrokenChain, dlgCid, iss, dlg.Audience())
		}
		iss = dlg.Audience()

		if !dlg.Command().Covers(cmd) {
			return fmt.Errorf("%w: delegation %s, %s doesn't cover %s", ErrCommandNotCovered, dlgCid, dlg.Command(), cmd)
		}
		cmd = dlg.Command()

		iss = dlg.Issuer()
	}

	// There must be at least one delegation referenced
	// (yes, it's an odd way to test this, but it allows for the static check to not be mad about "last"
	// being possibly nil below).
	if len(delegations) < 1 {
		return ErrNoProof
	}

	// The last prf value must be a root delegation (have the issuer field
	// match the Subject field) - 4g
	if last := delegations[len(delegations)-1]; last.Issuer() != last.Subject() {
		return fmt.Errorf("%w: expected %s, got %s", ErrLastNotRoot, last.Subject(), last.Issuer())
	}

	return nil
}

func (t *Token) verifyTimeBound(dlgs []*delegation.Token) error {
	return t.verifyTimeBoundAt(time.Now(), dlgs)
}

func (t *Token) verifyTimeBoundAt(at time.Time, delegations []*delegation.Token) error {
	for i, dlgCid := range t.proof {
		dlg := delegations[i]

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
