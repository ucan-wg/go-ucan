package container

import (
	"fmt"

	"github.com/ipfs/go-cid"

	"github.com/ucan-wg/go-ucan/token"
	"github.com/ucan-wg/go-ucan/token/delegation"
)

// TODO: should the invocation being set as root in the car file?

var ErrNotFound = fmt.Errorf("not found")

type Container map[cid.Cid][]byte

func New() Container {
	return make(Container)
}

func (ctn Container) AddBytes(cid cid.Cid, data []byte) {
	ctn[cid] = data
}

func (ctn Container) GetBytes(cid cid.Cid) ([]byte, bool) {
	b, ok := ctn[cid]
	return b, ok
}

func (ctn Container) GetToken(cid cid.Cid) (token.Token, error) {
	b, ok := ctn[cid]
	if !ok {
		return nil, ErrNotFound
	}
	return token.FromDagCbor(b)
}

func (ctn Container) GetDelegation(cid cid.Cid) (*delegation.Token, error) {
	tkn, err := ctn.GetToken(cid)
	if err != nil {
		return nil, err
	}
	if tkn, ok := tkn.(*delegation.Token); ok {
		return tkn, nil
	}
	return nil, fmt.Errorf("not a delegation token")
}
