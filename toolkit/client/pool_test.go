package client

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/ucan-wg/go-ucan/did/didtest"
	"github.com/ucan-wg/go-ucan/token/delegation/delegationtest"
)

func TestFindProof(t *testing.T) {
	p := NewPool()

	for _, bundle := range delegationtest.AllBundles {
		p.AddBundle(bundle)
	}

	require.Equal(t, delegationtest.ProofAliceBob,
		p.FindProof(didtest.PersonaBob.DID(), didtest.PersonaAlice.DID(), delegationtest.NominalCommand))
	require.Equal(t, delegationtest.ProofAliceBobCarol,
		p.FindProof(didtest.PersonaCarol.DID(), didtest.PersonaAlice.DID(), delegationtest.NominalCommand))
	require.Equal(t, delegationtest.ProofAliceBobCarolDan,
		p.FindProof(didtest.PersonaDan.DID(), didtest.PersonaAlice.DID(), delegationtest.NominalCommand))
	require.Equal(t, delegationtest.ProofAliceBobCarolDanErin,
		p.FindProof(didtest.PersonaErin.DID(), didtest.PersonaAlice.DID(), delegationtest.NominalCommand))
	require.Equal(t, delegationtest.ProofAliceBobCarolDanErinFrank,
		p.FindProof(didtest.PersonaFrank.DID(), didtest.PersonaAlice.DID(), delegationtest.NominalCommand))

}
