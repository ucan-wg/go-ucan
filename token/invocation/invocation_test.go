package invocation_test

import (
	"testing"

	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
	"github.com/ucan-wg/go-ucan/did/didtest"
	"github.com/ucan-wg/go-ucan/pkg/args"
	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/token/delegation/delegationtest"
	"github.com/ucan-wg/go-ucan/token/invocation"

	"github.com/stretchr/testify/assert"
)

const (
	missingPrivKeyCfg = "CAESQMjRvrEIjpPYRQKmkAGw/pV0XgE958rYa4vlnKJjl1zz/sdnGnyV1xKLJk8D39edyjhHWyqcpgFnozQK62SG16k="
	missingTknCIDStr  = "bafyreigwypmw6eul6vadi6g6lnfbsfo2zck7gfzsbjoroqs3djhnzzc7mm"
	missingDIDStr     = "did:key:z6MkwboxFsH3kEuehBZ5fLkRmxi68yv1u38swA4r9Jm2VRma"
)

var emptyArguments = args.New()

func TestToken_ExecutionAllowed(t *testing.T) {
	t.Parallel()

	t.Run("passes - only root", func(t *testing.T) {
		t.Parallel()

		testPasses(t, didtest.PersonaBob, delegationtest.NominalCommand, emptyArguments, delegationtest.ProofAliceBob)
	})

	t.Run("passes - valid chain", func(t *testing.T) {
		t.Parallel()

		testPasses(t, didtest.PersonaFrank, delegationtest.NominalCommand, emptyArguments, delegationtest.ProofAliceBobCarolDanErinFrank)
	})

	t.Run("passes - proof chain attenuates command", func(t *testing.T) {
		t.Parallel()

		testPasses(t, didtest.PersonaFrank, delegationtest.AttenuatedCommand, emptyArguments, delegationtest.ProofAliceBobCarolDanErinFrankValidAttenuatedCommand)
	})

	t.Run("passes - invocation attenuates command", func(t *testing.T) {
		t.Parallel()

		testPasses(t, didtest.PersonaFrank, delegationtest.AttenuatedCommand, emptyArguments, delegationtest.ProofAliceBobCarolDanErinFrank)
	})

	t.Run("fails - no proof", func(t *testing.T) {
		t.Parallel()

		testFails(t, invocation.ErrNoProof, didtest.PersonaCarol, delegationtest.NominalCommand, emptyArguments, delegationtest.ProofEmpty)
	})

	t.Run("fails - missing referenced delegation", func(t *testing.T) {
		t.Parallel()

		missingTknCID, err := cid.Parse(missingTknCIDStr)
		require.NoError(t, err)

		prf := []cid.Cid{missingTknCID, delegationtest.TokenAliceBobCID}
		testFails(t, invocation.ErrMissingDelegation, didtest.PersonaCarol, delegationtest.NominalCommand, emptyArguments, prf)
	})

	t.Run("fails - referenced delegation expired", func(t *testing.T) {
		t.Parallel()

		testFails(t, invocation.ErrTokenInvalidNow, didtest.PersonaFrank, delegationtest.NominalCommand, emptyArguments, delegationtest.ProofAliceBobCarolDanErinFrankInvalidExpired)

	})

	t.Run("fails - referenced delegation inactive", func(t *testing.T) {
		t.Parallel()

		testFails(t, invocation.ErrTokenInvalidNow, didtest.PersonaFrank, delegationtest.NominalCommand, emptyArguments, delegationtest.ProofAliceBobCarolDanErinFrankInvalidInactive)
	})

	t.Run("fails - last (or only) delegation not root", func(t *testing.T) {
		t.Parallel()

		prf := []cid.Cid{delegationtest.TokenErinFrankCID, delegationtest.TokenDanErinCID, delegationtest.TokenCarolDanCID}
		testFails(t, invocation.ErrLastNotRoot, didtest.PersonaFrank, delegationtest.NominalCommand, emptyArguments, prf)
	})

	t.Run("fails - broken chain", func(t *testing.T) {
		t.Parallel()

		prf := []cid.Cid{delegationtest.TokenCarolDanCID, delegationtest.TokenAliceBobCID}
		testFails(t, invocation.ErrBrokenChain, didtest.PersonaFrank, delegationtest.NominalCommand, emptyArguments, prf)
	})

	t.Run("fails - first not issued to invoker", func(t *testing.T) {
		t.Parallel()

		prf := []cid.Cid{delegationtest.TokenBobCarolCID, delegationtest.TokenAliceBobCID}
		testFails(t, invocation.ErrBrokenChain, didtest.PersonaFrank, delegationtest.NominalCommand, emptyArguments, prf)
	})

	t.Run("fails - proof chain expands command", func(t *testing.T) {
		t.Parallel()

		testFails(t, invocation.ErrCommandNotCovered, didtest.PersonaFrank, delegationtest.NominalCommand, emptyArguments, delegationtest.ProofAliceBobCarolDanErinFrankInvalidExpandedCommand)
	})

	t.Run("fails - invocation expands command", func(t *testing.T) {
		t.Parallel()

		testFails(t, invocation.ErrCommandNotCovered, didtest.PersonaFrank, delegationtest.ExpandedCommand, emptyArguments, delegationtest.ProofAliceBobCarolDanErinFrank)
	})

	t.Run("fails - inconsistent subject", func(t *testing.T) {
		t.Parallel()

		testFails(t, invocation.ErrWrongSub, didtest.PersonaFrank, delegationtest.ExpandedCommand, emptyArguments, delegationtest.ProofAliceBobCarolDanErinFrankInvalidSubject)
	})
}

func test(t *testing.T, persona didtest.Persona, cmd command.Command, args *args.Args, prf []cid.Cid, opts ...invocation.Option) (bool, error) {
	t.Helper()

	tkn, err := invocation.New(persona.DID(t), didtest.PersonaAlice.DID(t), cmd, prf, opts...)
	require.NoError(t, err)

	ldr, err := delegationtest.GetDelegationLoader()
	require.NoError(t, err)

	return tkn.ExecutionAllowed(ldr)
}

func testFails(t *testing.T, expErr error, persona didtest.Persona, cmd command.Command, args *args.Args, prf []cid.Cid, opts ...invocation.Option) {
	ok, err := test(t, persona, cmd, args, prf, opts...)
	require.ErrorIs(t, err, expErr)
	assert.False(t, ok)
}

func testPasses(t *testing.T, persona didtest.Persona, cmd command.Command, args *args.Args, prf []cid.Cid, opts ...invocation.Option) {
	ok, err := test(t, persona, cmd, args, prf, opts...)
	require.NoError(t, err)
	assert.True(t, ok)
}
