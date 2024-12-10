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
		FindProof(dlgs, didtest.PersonaBob.DID(), delegationtest.NominalCommand, didtest.PersonaAlice.DID()))
	require.Equal(t, delegationtest.ProofAliceBobCarol,
		FindProof(dlgs, didtest.PersonaCarol.DID(), delegationtest.NominalCommand, didtest.PersonaAlice.DID()))
	require.Equal(t, delegationtest.ProofAliceBobCarolDan,
		FindProof(dlgs, didtest.PersonaDan.DID(), delegationtest.NominalCommand, didtest.PersonaAlice.DID()))
	require.Equal(t, delegationtest.ProofAliceBobCarolDanErin,
		FindProof(dlgs, didtest.PersonaErin.DID(), delegationtest.NominalCommand, didtest.PersonaAlice.DID()))
	require.Equal(t, delegationtest.ProofAliceBobCarolDanErinFrank,
		FindProof(dlgs, didtest.PersonaFrank.DID(), delegationtest.NominalCommand, didtest.PersonaAlice.DID()))

	// wrong command
	require.Empty(t, FindProof(dlgs, didtest.PersonaBob.DID(), command.New("foo"), didtest.PersonaAlice.DID()))
}
