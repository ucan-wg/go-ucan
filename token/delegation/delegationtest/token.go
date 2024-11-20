package delegationtest

import (
	"embed"
	"path/filepath"
	"sync"

	"github.com/ipfs/go-cid"

	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/token/delegation"
)

const (
	tokenDir = "data"
	tokenExt = ".dagcbor"
)

var (
	// ExpandedCommand is the parent of the NominalCommand and represents
	// the cases where the delegation proof-chain or invocation token tries
	// to increase the privileges granted by the root delegation token.
	// Execution of this command is generally prohibited in tests.
	ExpandedCommand = command.MustParse("/expanded")

	// NominalCommand is the command used for most test tokens and proof-
	// chains.  Execution of this command is generally allowed in tests.
	NominalCommand = ExpandedCommand.Join("nominal")

	// AttenuatedCommand is a sub-command of the NominalCommand.  Execution
	// of this command is generally allowed in tests.
	AttenuatedCommand = NominalCommand.Join("attenuated")
)

// ProofEmpty provides an empty proof chain for testing purposes.
var ProofEmpty = []cid.Cid{}

//go:embed data
var fs embed.FS

var _ delegation.Loader = (*delegationLoader)(nil)

type delegationLoader struct {
	tokens map[cid.Cid]*delegation.Token
}

var (
	once sync.Once
	ldr  delegation.Loader
)

// GetDelegationLoader returns a singleton instance of a test
// DelegationLoader containing all the tokens present in the data/
// directory.
func GetDelegationLoader() delegation.Loader {
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
func (l *delegationLoader) GetDelegation(id cid.Cid) (*delegation.Token, error) {
	tkn, ok := l.tokens[id]
	if !ok {
		return nil, delegation.ErrDelegationNotFound
	}

	return tkn, nil
}

func loadDelegations() (delegation.Loader, error) {
	dirEntries, err := fs.ReadDir("data")
	if err != nil {
		return nil, err
	}

	tkns := make(map[cid.Cid]*delegation.Token, len(dirEntries))

	for _, dirEntry := range dirEntries {
		data, err := fs.ReadFile(filepath.Join(tokenDir, dirEntry.Name()))
		if err != nil {
			return nil, err
		}

		tkn, id, err := delegation.FromSealed(data)
		if err != nil {
			return nil, err
		}

		tkns[id] = tkn
	}

	return &delegationLoader{
		tokens: tkns,
	}, nil
}

// GetDelegation is a shortcut that gets (or creates) the DelegationLoader
// and attempts to return the token referenced by the provided CID.
func GetDelegation(id cid.Cid) (*delegation.Token, error) {
	return GetDelegationLoader().GetDelegation(id)
}

func mustGetDelegation(id cid.Cid) *delegation.Token {
	tkn, err := GetDelegation(id)
	if err != nil {
		panic(err)
	}
	return tkn
}
