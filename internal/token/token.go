package token

import (
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/node/bindnode"
)

//go:generate go run ../cmd/token/...

func New() (*Token, error) {
	return &Token{}, nil
}

func Unwrap(node datamodel.Node) (*Token, error) {
	iface := bindnode.Unwrap(node)
	if iface == nil {
		return nil, ErrNodeNotToken
	}

	tkn, ok := iface.(*Token)
	if !ok {
		return nil, ErrNodeNotToken
	}

	return tkn, nil
}

func (t *Token) Wrap() datamodel.Node {
	return bindnode.Wrap(t, tokenType())
}
