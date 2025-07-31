package delegationtest

import (
	"embed"
	"path/filepath"
	"sync"

	_ "github.com/MetaMask/go-did-it/verifiers/did-key"
	"github.com/ipfs/go-cid"

	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/token/delegation"
)

var (
	// ExpandedCommand is the parent of the NominalCommand and represents
	// the cases where the delegation proof-chain or invocation token tries
	// to increase the privileges granted by the root delegation token.
	// Execution of this command is generally prohibited in tests.
	ExpandedCommand = command.MustParse("/expanded")

	// NominalCommand is the command used for most test tokens and proof-chains.
	// Execution of this command is generally allowed in tests.
	NominalCommand = ExpandedCommand.Join("nominal")

	// AttenuatedCommand is a sub-command of the NominalCommand.
	// Execution of this command is generally allowed in tests.
	AttenuatedCommand = NominalCommand.Join("attenuated")
)

// ProofEmpty provides an empty proof chain for testing purposes.
var ProofEmpty = []cid.Cid{}

const TokenDir = "data"

//go:embed data
var fs embed.FS

var _ delegation.Loader = (*DelegationLoader)(nil)

type DelegationLoader struct {
	bundles map[cid.Cid]delegation.Bundle
}

var (
	once sync.Once
	ldr  *DelegationLoader
)

// GetDelegationLoader returns a singleton instance of a test
// DelegationLoader containing all the tokens present in the data/
// directory.
func GetDelegationLoader() *DelegationLoader {
	once.Do(func() {
		var err error
		ldr, err = loadDelegations()
		if err != nil {
			panic(err)
		}
	})
	return ldr
}

// GetDelegation implements invocation.DelegationLoader.
func (l *DelegationLoader) GetDelegation(id cid.Cid) (*delegation.Token, error) {
	bundle, ok := l.bundles[id]
	if !ok {
		return nil, delegation.ErrDelegationNotFound
	}
	return bundle.Decoded, nil
}

func loadDelegations() (*DelegationLoader, error) {
	dirEntries, err := fs.ReadDir(TokenDir)
	if err != nil {
		return nil, err
	}

	bundles := make(map[cid.Cid]delegation.Bundle, len(dirEntries))

	for _, dirEntry := range dirEntries {
		data, err := fs.ReadFile(filepath.Join(TokenDir, dirEntry.Name()))
		if err != nil {
			return nil, err
		}

		tkn, id, err := delegation.FromSealed(data)
		if err != nil {
			return nil, err
		}

		bundles[id] = delegation.Bundle{Cid: id, Decoded: tkn, Sealed: data}
	}

	return &DelegationLoader{
		bundles: bundles,
	}, nil
}

// GetDelegation is a shortcut that gets (or creates) the DelegationLoader
// and attempts to return the token referenced by the provided CID.
func GetDelegation(id cid.Cid) (*delegation.Token, error) {
	return GetDelegationLoader().GetDelegation(id)
}

func CidToName(id cid.Cid) string {
	return cidToName[id]
}

func mustGetBundle(id cid.Cid) delegation.Bundle {
	bundle, ok := GetDelegationLoader().bundles[id]
	if !ok {
		panic(delegation.ErrDelegationNotFound)
	}
	return bundle
}
