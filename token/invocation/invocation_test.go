package invocation_test

import (
	_ "embed"
	"testing"
	"time"

	"github.com/MetaMask/go-did-it/didtest"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"

	"github.com/ucan-wg/go-ucan/pkg/args"
	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/pkg/policy/policytest"
	"github.com/ucan-wg/go-ucan/token/delegation/delegationtest"
	"github.com/ucan-wg/go-ucan/token/invocation"
)

//go:embed testdata/new.dagjson
var newDagJson []byte

//go:embed testdata/selfsigned.dagjson
var selfsignedDagJson []byte

const missingTknCIDStr = "bafyreigwypmw6eul6vadi6g6lnfbsfo2zck7gfzsbjoroqs3djhnzzc7mm"

var emptyArguments = args.New()

func TestToken_ExecutionAllowed(t *testing.T) {
	for _, tc := range []struct {
		name   string
		issuer didtest.Persona
		cmd    command.Command
		args   *args.Args
		proofs []cid.Cid
		opts   []invocation.Option
		err    error
	}{
		// Passes
		{
			name:   "passes - only root",
			issuer: didtest.PersonaBob,
			cmd:    delegationtest.NominalCommand,
			args:   emptyArguments,
			proofs: delegationtest.ProofAliceBob,
			err:    nil,
		},
		{
			name:   "passes - valid chain",
			issuer: didtest.PersonaFrank,
			cmd:    delegationtest.NominalCommand,
			args:   emptyArguments,
			proofs: delegationtest.ProofAliceBobCarolDanErinFrank,
			err:    nil,
		},
		{
			name:   "passes - proof chain attenuates command",
			issuer: didtest.PersonaFrank,
			cmd:    delegationtest.AttenuatedCommand,
			args:   emptyArguments,
			proofs: delegationtest.ProofAliceBobCarolDanErinFrank_ValidAttenuatedCommand,
			err:    nil,
		},
		{
			name:   "passes - invocation attenuates command",
			issuer: didtest.PersonaFrank,
			cmd:    delegationtest.AttenuatedCommand,
			args:   emptyArguments,
			proofs: delegationtest.ProofAliceBobCarolDanErinFrank,
			err:    nil,
		},
		{
			name:   "passes - arguments satisfy empty policy",
			issuer: didtest.PersonaFrank,
			cmd:    delegationtest.NominalCommand,
			args:   policytest.SpecValidArguments,
			proofs: delegationtest.ProofAliceBobCarolDanErinFrank,
			err:    nil,
		},
		{
			name:   "passes - arguments satisfy example policy",
			issuer: didtest.PersonaFrank,
			cmd:    delegationtest.NominalCommand,
			args:   policytest.SpecValidArguments,
			proofs: delegationtest.ProofAliceBobCarolDanErinFrank_ValidExamplePolicy,
			err:    nil,
		},
		{
			name:   "passes - self-signed invocation doesn't require proof",
			issuer: didtest.PersonaAlice,
			cmd:    delegationtest.NominalCommand,
			args:   emptyArguments,
			proofs: nil,
			err:    nil,
		},
		{
			name:   "passes - self-signed invocation accepts a delegation to itself",
			issuer: didtest.PersonaAlice,
			cmd:    delegationtest.NominalCommand,
			args:   emptyArguments,
			proofs: []cid.Cid{delegationtest.TokenAliceAliceCID},
			err:    nil,
		},

		// Fails
		{
			name:   "fails - no proof",
			issuer: didtest.PersonaCarol,
			cmd:    delegationtest.NominalCommand,
			args:   emptyArguments,
			proofs: delegationtest.ProofEmpty,
			err:    invocation.ErrNoProof,
		},
		{
			name:   "fails - missing referenced delegation",
			issuer: didtest.PersonaCarol,
			cmd:    delegationtest.NominalCommand,
			args:   emptyArguments,
			proofs: []cid.Cid{cid.MustParse(missingTknCIDStr), delegationtest.TokenAliceBobCID},
			err:    invocation.ErrMissingDelegation,
		},
		{
			name:   "fails - referenced delegation expired",
			issuer: didtest.PersonaFrank,
			cmd:    delegationtest.NominalCommand,
			args:   emptyArguments,
			proofs: delegationtest.ProofAliceBobCarolDanErinFrank_InvalidExpired,
			err:    invocation.ErrTokenInvalidNow,
		},
		{
			name:   "fails - referenced delegation inactive",
			issuer: didtest.PersonaFrank,
			cmd:    delegationtest.NominalCommand,
			args:   emptyArguments,
			proofs: delegationtest.ProofAliceBobCarolDanErinFrank_InvalidInactive,
			err:    invocation.ErrTokenInvalidNow,
		},
		{
			name:   "fails - last (or only) delegation not root",
			issuer: didtest.PersonaFrank,
			cmd:    delegationtest.NominalCommand,
			args:   emptyArguments,
			proofs: []cid.Cid{delegationtest.TokenErinFrankCID, delegationtest.TokenDanErinCID, delegationtest.TokenCarolDanCID},
			err:    invocation.ErrLastNotRoot,
		},
		{
			name:   "fails - broken chain",
			issuer: didtest.PersonaFrank,
			cmd:    delegationtest.NominalCommand,
			args:   emptyArguments,
			proofs: []cid.Cid{delegationtest.TokenCarolDanCID, delegationtest.TokenAliceBobCID},
			err:    invocation.ErrBrokenChain,
		},
		{
			name:   "fails - first not issued to invoker",
			issuer: didtest.PersonaFrank,
			cmd:    delegationtest.NominalCommand,
			args:   emptyArguments,
			proofs: []cid.Cid{delegationtest.TokenBobCarolCID, delegationtest.TokenAliceBobCID},
			err:    invocation.ErrBrokenChain,
		},
		{
			name:   "fails - proof chain expands command",
			issuer: didtest.PersonaFrank,
			cmd:    delegationtest.NominalCommand,
			args:   emptyArguments,
			proofs: delegationtest.ProofAliceBobCarolDanErinFrank_InvalidExpandedCommand,
			err:    invocation.ErrCommandNotCovered,
		},
		{
			name:   "fails - invocation expands command",
			issuer: didtest.PersonaFrank,
			cmd:    delegationtest.ExpandedCommand,
			args:   emptyArguments,
			proofs: delegationtest.ProofAliceBobCarolDanErinFrank,
			err:    invocation.ErrCommandNotCovered,
		},
		{
			name:   "fails - inconsistent subject",
			issuer: didtest.PersonaFrank,
			cmd:    delegationtest.ExpandedCommand,
			args:   emptyArguments,
			proofs: delegationtest.ProofAliceBobCarolDanErinFrank_InvalidSubject,
			err:    invocation.ErrWrongSub,
		},
		{
			name:   "fails - arguments don't satisfy example policy",
			issuer: didtest.PersonaFrank,
			cmd:    delegationtest.NominalCommand,
			args:   policytest.SpecInvalidArguments,
			proofs: delegationtest.ProofAliceBobCarolDanErinFrank_ValidExamplePolicy,
			err:    invocation.ErrPolicyNotSatisfied,
		},
		{
			name:   "fails - self-signed invocation refuses a delegation to itself for a different DID",
			issuer: didtest.PersonaAlice,
			cmd:    delegationtest.NominalCommand,
			args:   emptyArguments,
			proofs: []cid.Cid{delegationtest.TokenBobBobCID},
			err:    invocation.ErrBrokenChain,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			tc.opts = append(tc.opts, invocation.WithArguments(tc.args))

			tkn, err := invocation.New(tc.issuer.DID(), tc.cmd, didtest.PersonaAlice.DID(), tc.proofs, tc.opts...)
			require.NoError(t, err)

			err = tkn.ExecutionAllowed(delegationtest.GetDelegationLoader())

			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

const (
	nonce      = "6roDhGi0kiNriQAz7J3d+bOeoI/tj8ENikmQNbtjnD0"
	subjectCmd = "/foo/bar"
)

func TestConstructors(t *testing.T) {
	cmd, err := command.Parse(subjectCmd)
	require.NoError(t, err)

	iat, err := time.Parse(time.RFC3339, "2100-01-01T00:00:00Z")
	require.NoError(t, err)

	exp, err := time.Parse(time.RFC3339, "2200-01-01T00:00:00Z")
	require.NoError(t, err)

	t.Run("New", func(t *testing.T) {
		tkn, err := invocation.New(
			didtest.PersonaAlice.DID(), cmd, didtest.PersonaBob.DID(),
			delegationtest.ProofAliceBob,
			invocation.WithNonce([]byte(nonce)),
			invocation.WithIssuedAt(iat),
			invocation.WithExpiration(exp),
			invocation.WithArgument("foo", "bar"),
			invocation.WithMeta("baz", 123),
		)
		require.NoError(t, err)

		require.False(t, tkn.IsSelfSigned())

		data, err := tkn.ToDagJson(didtest.PersonaAlice.PrivKey())
		require.NoError(t, err)

		require.JSONEq(t, string(newDagJson), string(data))
	})

	t.Run("Self-Signed", func(t *testing.T) {
		tkn, err := invocation.NewSelfSigned(
			didtest.PersonaAlice.DID(), cmd,
			invocation.WithNonce([]byte(nonce)),
			invocation.WithIssuedAt(iat),
			invocation.WithExpiration(exp),
			invocation.WithArgument("foo", "bar"),
			invocation.WithMeta("baz", 123),
		)
		require.NoError(t, err)

		require.True(t, tkn.IsSelfSigned())

		data, err := tkn.ToDagJson(didtest.PersonaAlice.PrivKey())
		require.NoError(t, err)

		require.JSONEq(t, string(selfsignedDagJson), string(data))
	})
}
