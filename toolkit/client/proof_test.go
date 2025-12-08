package client

import (
	"iter"
	"testing"

	"github.com/MetaMask/go-did-it"
	"github.com/MetaMask/go-did-it/didtest"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"

	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/token/delegation"
	"github.com/ucan-wg/go-ucan/token/delegation/delegationtest"
)

func TestFindProof(t *testing.T) {
	dlgs := func() iter.Seq[*delegation.Bundle] {
		return func(yield func(*delegation.Bundle) bool) {
			for _, bundle := range delegationtest.AllBundles {
				if !yield(&bundle) {
					return
				}
			}
		}
	}

	for _, tc := range []struct {
		name     string
		issuer   did.DID
		command  command.Command
		subject  did.DID
		expected []cid.Cid
	}{
		// Passes
		{
			name:     "Alice --> Alice (self-delegation)",
			issuer:   didtest.PersonaAlice.DID(),
			command:  delegationtest.NominalCommand,
			subject:  didtest.PersonaAlice.DID(),
			expected: []cid.Cid{delegationtest.TokenAliceAliceCID},
		},
		{
			name:     "Alice --> Bob",
			issuer:   didtest.PersonaBob.DID(),
			command:  delegationtest.NominalCommand,
			subject:  didtest.PersonaAlice.DID(),
			expected: delegationtest.ProofAliceBob,
		},
		{
			name:     "Alice --> Bob --> Carol",
			issuer:   didtest.PersonaCarol.DID(),
			command:  delegationtest.NominalCommand,
			subject:  didtest.PersonaAlice.DID(),
			expected: delegationtest.ProofAliceBobCarol,
		},
		{
			name:     "Alice --> Bob --> Carol --> Dan",
			issuer:   didtest.PersonaDan.DID(),
			command:  delegationtest.NominalCommand,
			subject:  didtest.PersonaAlice.DID(),
			expected: delegationtest.ProofAliceBobCarolDan,
		},
		{
			name:     "Alice --> Bob --> Carol --> Dan --> Erin",
			issuer:   didtest.PersonaErin.DID(),
			command:  delegationtest.NominalCommand,
			subject:  didtest.PersonaAlice.DID(),
			expected: delegationtest.ProofAliceBobCarolDanErin,
		},
		{
			name:     "Alice --> Bob --> Carol --> Dan --> Erin --> Frank",
			issuer:   didtest.PersonaFrank.DID(),
			command:  delegationtest.NominalCommand,
			subject:  didtest.PersonaAlice.DID(),
			expected: delegationtest.ProofAliceBobCarolDanErinFrank,
		},

		// Fails
		{
			name:     "wrong command",
			issuer:   didtest.PersonaBob.DID(),
			command:  command.New("foo"),
			subject:  didtest.PersonaAlice.DID(),
			expected: nil,
		},
		{
			name:     "wrong subject",
			issuer:   didtest.PersonaBob.DID(),
			command:  delegationtest.NominalCommand,
			subject:  didtest.PersonaDan.DID(),
			expected: nil,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			res := FindProof(dlgs, tc.issuer, tc.command, tc.subject)
			require.Equal(t, tc.expected, res)
		})
	}
}
