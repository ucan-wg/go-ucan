package exectx

import (
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime/datamodel"

	"github.com/ucan-wg/go-ucan/pkg/args"
	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/pkg/container"
	"github.com/ucan-wg/go-ucan/pkg/meta"
	"github.com/ucan-wg/go-ucan/pkg/policy"
	"github.com/ucan-wg/go-ucan/token/delegation"
	"github.com/ucan-wg/go-ucan/token/invocation"
	"github.com/ucan-wg/go-ucan/toolkit/server/extargs"
)

var _ delegation.Loader = &UcanCtx{}

// UcanCtx is a UCAN execution context meant to be inserted in the go context while handling a request.
// It allows to handle the control of the execution in multiple steps across different middlewares,
// as well as doing "bearer" types of controls, when arguments are derived from the request itself (HTTP, JsonRpc).
type UcanCtx struct {
	inv  *invocation.Token
	dlgs map[cid.Cid]*delegation.Token

	policies policy.Policy // all policies combined
	meta     *meta.Meta    // all meta combined, with no overwriting

	// argument sources
	http   *extargs.HttpExtArgs
	custom map[string]*extargs.CustomExtArgs
}

// FromContainer prepare a UcanCtx from a UCAN container, for further evaluation in a server pipeline.
// It is expected that the container holds a single invocation and the matching delegations. If not,
// an error is returned.
func FromContainer(cont container.Reader) (*UcanCtx, error) {
	inv, err := cont.GetInvocation()
	if err != nil {
		return nil, fmt.Errorf("no invocation: %w", err)
	}

	ctx := &UcanCtx{
		inv:  inv,
		dlgs: make(map[cid.Cid]*delegation.Token, len(cont)-1),
		meta: meta.NewMeta(),
	}

	// iterate in reverse, from the root delegation to the leaf
	for _, c := range slices.Backward(inv.Proof()) {
		// make sure we have the delegation
		dlg, err := cont.GetDelegation(c)
		if errors.Is(err, delegation.ErrDelegationNotFound) {
			return nil, fmt.Errorf("delegation proof %s is missing", c)
		}
		if err != nil {
			return nil, err
		}
		ctx.dlgs[c] = dlg
		// accumulate the policies
		ctx.policies = append(ctx.policies, dlg.Policy()...)
		// accumulate the meta values, with no overwriting
		ctx.meta.Include(dlg.Meta())
	}

	// DX: As the invocation is created without the delegation, no check is done that the proof chain (CIDs only)
	// is ordered properly and not broken. We don't check that in the container either as it doesn't make any assumption
	// on what is being carried around. That UcanCtx is the first place where we enforce having a single invocation and
	// only the matching delegation.
	// For sanity, we verify that the proofs are ordered properly. This will be checked later anyway, but it's cheap to
	// verify here and catch an easy mistake.
	chainTo := inv.Issuer()
	for _, c := range inv.Proof() {
		dlg := ctx.dlgs[c]
		if !dlg.Audience().Equal(chainTo) {
			return nil, fmt.Errorf("proof chain is broken or not ordered correctly")
		}
		chainTo = dlg.Issuer()
	}

	return ctx, nil
}

// Command returns the command triggered by the invocation.
func (ctn *UcanCtx) Command() command.Command {
	return ctn.inv.Command()
}

// Invocation returns the invocation.Token.
func (ctn *UcanCtx) Invocation() *invocation.Token {
	return ctn.inv
}

// GetDelegation returns the delegation.Token matching the given CID.
// If not found, delegation.ErrDelegationNotFound is returned.
func (ctn *UcanCtx) GetDelegation(cid cid.Cid) (*delegation.Token, error) {
	if dlg, ok := ctn.dlgs[cid]; ok {
		return dlg, nil
	}
	return nil, delegation.ErrDelegationNotFound
}

// GetRootDelegation returns the delegation.Token at the root of the proof chain.
func (ctn *UcanCtx) GetRootDelegation() *delegation.Token {
	proofs := ctn.inv.Proof()
	c := proofs[len(proofs)-1]
	return ctn.dlgs[c]
}

// Policies return the full set of policy statements to satisfy.
func (ctn *UcanCtx) Policies() policy.Policy {
	return ctn.policies
}

// Meta returns all the meta values from the delegations.
// They are accumulated from the root delegation to the leaf delegation, with no overwriting.
func (ctn *UcanCtx) Meta() meta.ReadOnly {
	return ctn.meta.ReadOnly()
}

// VerifyHttp verify the delegation's policies against arguments constructed from the HTTP request.
// These arguments will be set in the `.http` argument key, at the root.
// This function can only be called once per context.
// After being used, those constructed arguments will be used in ExecutionAllowed as well.
func (ctn *UcanCtx) VerifyHttp(req *http.Request) error {
	if ctn.http != nil {
		panic("only use once per request context")
	}
	ctn.http = extargs.NewHttpExtArgs(ctn.policies, ctn.inv.Arguments(), req)
	return ctn.http.Verify()
}

// VerifyCustom verify the delegation's policies against arbitrary arguments provider through an IPLD MapAssembler.
// These arguments will be set under the given argument key, at the root.
// This function can only be called once per context and key.
// After being used, those constructed arguments will be used in ExecutionAllowed as well.
func (ctn *UcanCtx) VerifyCustom(key string, assembler func(ma datamodel.MapAssembler)) error {
	if ctn.custom == nil {
		ctn.custom = make(map[string]*extargs.CustomExtArgs)
	}
	if _, ok := ctn.custom[key]; ok {
		panic("only use once per request context and key")
	}
	ctn.custom[key] = extargs.NewCustomExtArgs(key, ctn.policies, assembler)
	return ctn.custom[key].Verify()
}

// ExecutionAllowed does the final verification of the invocation.
// If VerifyHttp or VerifyJsonRpc has been used, those arguments are part of the verification.
func (ctn *UcanCtx) ExecutionAllowed() error {
	return ctn.inv.ExecutionAllowedWithArgsHook(ctn, func(args args.ReadOnly) (*args.Args, error) {
		newArgs := args.WriteableClone()

		if ctn.http != nil {
			httpArgs, err := ctn.http.Args()
			if err != nil {
				return nil, err
			}
			newArgs.Include(httpArgs)
		}
		if ctn.custom != nil {
			for _, cea := range ctn.custom {
				customArgs, err := cea.Args()
				if err != nil {
					return nil, err
				}
				newArgs.Include(customArgs)
			}
		}

		return newArgs, nil
	})
}
