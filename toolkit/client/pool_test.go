package client

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/ucan-wg/go-ucan/did/didtest"
	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/token/delegation/delegationtest"
)

func TestFindProof(t *testing.T) {
	p := NewPool()

	for _, bundle := range delegationtest.AllBundles {
		p.AddBundle(bundle)
	}

	require.Equal(t, delegationtest.ProofAliceBob,
		p.FindProof(delegationtest.NominalCommand, didtest.PersonaBob.DID(), didtest.PersonaAlice.DID()))
	require.Equal(t, delegationtest.ProofAliceBobCarol,
		p.FindProof(delegationtest.NominalCommand, didtest.PersonaCarol.DID(), didtest.PersonaAlice.DID()))
	require.Equal(t, delegationtest.ProofAliceBobCarolDan,
		p.FindProof(delegationtest.NominalCommand, didtest.PersonaDan.DID(), didtest.PersonaAlice.DID()))
	require.Equal(t, delegationtest.ProofAliceBobCarolDanErin,
		p.FindProof(delegationtest.NominalCommand, didtest.PersonaErin.DID(), didtest.PersonaAlice.DID()))
	require.Equal(t, delegationtest.ProofAliceBobCarolDanErinFrank,
		p.FindProof(delegationtest.NominalCommand, didtest.PersonaFrank.DID(), didtest.PersonaAlice.DID()))

	// wrong command
	require.Empty(t, p.FindProof(command.New("foo"), didtest.PersonaBob.DID(), didtest.PersonaAlice.DID()))
}
