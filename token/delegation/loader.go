package delegation

import (
	"fmt"

	"github.com/ipfs/go-cid"
)

// ErrDelegationNotFound is returned if a delegation token is not found
var ErrDelegationNotFound = fmt.Errorf("delegation not found")

// Loader is a delegation token loader.
type Loader interface {
	// GetDelegation returns the delegation.Token matching the given CID.
	// If not found, ErrDelegationNotFound is returned.
	GetDelegation(cid cid.Cid) (*Token, error)
}
