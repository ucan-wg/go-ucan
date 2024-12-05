package client

import (
	"iter"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/ucan-wg/go-ucan/did/didtest"
	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/token/delegation"
	"github.com/ucan-wg/go-ucan/token/delegation/delegationtest"
)

func TestFindProof(t *testing.T) {
	dlgs := func() iter.Seq[*delegation.Bundle] {
		return slices.Values(delegationtest.AllBundles)
	}

	require.Equal(t, delegationtest.ProofAliceBob,
		FindProof(dlgs, delegationtest.NominalCommand, didtest.PersonaBob.DID(), didtest.PersonaAlice.DID()))
	require.Equal(t, delegationtest.ProofAliceBobCarol,
		FindProof(dlgs, delegationtest.NominalCommand, didtest.PersonaCarol.DID(), didtest.PersonaAlice.DID()))
	require.Equal(t, delegationtest.ProofAliceBobCarolDan,
		FindProof(dlgs, delegationtest.NominalCommand, didtest.PersonaDan.DID(), didtest.PersonaAlice.DID()))
	require.Equal(t, delegationtest.ProofAliceBobCarolDanErin,
		FindProof(dlgs, delegationtest.NominalCommand, didtest.PersonaErin.DID(), didtest.PersonaAlice.DID()))
	require.Equal(t, delegationtest.ProofAliceBobCarolDanErinFrank,
		FindProof(dlgs, delegationtest.NominalCommand, didtest.PersonaFrank.DID(), didtest.PersonaAlice.DID()))

	// wrong command
	require.Empty(t, FindProof(dlgs, command.New("foo"), didtest.PersonaBob.DID(), didtest.PersonaAlice.DID()))
}
