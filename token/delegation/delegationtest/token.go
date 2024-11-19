package delegationtest

import (
	"embed"
	"fmt"
	"path/filepath"
	"sync"

	"github.com/ipfs/go-cid"
	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/token/delegation"
	"github.com/ucan-wg/go-ucan/token/invocation"
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

var (
	once sync.Once
	ldr  invocation.DelegationLoader
	err  error
)

var _ invocation.DelegationLoader = (*delegationLoader)(nil)

type delegationLoader struct {
	tokens map[cid.Cid]*delegation.Token
}

// GetDelegationLoader returns a singleton instance of a test
// DelegationLoader containing all the tokens present in the data/
// directory.
func GetDelegationLoader() (invocation.DelegationLoader, error) {
	once.Do(func() {
		ldr, err = loadDelegations()
	})

	return ldr, err
}

// GetDelegation implements invocation.DelegationLoader.
func (l *delegationLoader) GetDelegation(id cid.Cid) (*delegation.Token, error) {
	tkn, ok := l.tokens[id]
	if !ok {
		return nil, fmt.Errorf("%w: CID %s", invocation.ErrMissingDelegation, id.String())
	}

	return tkn, nil
}

func loadDelegations() (invocation.DelegationLoader, error) {
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
	ldr, err := GetDelegationLoader()
	if err != nil {
		return nil, err
	}

	return ldr.GetDelegation(id)
}

func mustGetDelegation(id cid.Cid) *delegation.Token {
	tkn, err := GetDelegation(id)
	if err != nil {
		panic(err)
	}

	return tkn
}
