package token

import (
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/node/bindnode"
)

//go:generate go run ../cmd/token/...

// Unwrap creates a Token from an arbitrary IPLD node or returns an
// error if at least the required model fields are not present.
//
// It is the responsibility of the Delegation and Invocation views
// to further validate the presence of the required fields and the
// content as needed.
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

// Wrap creates an IPLD node representing the Token.
func (t *Token) Wrap() datamodel.Node {
	return bindnode.Wrap(t, tokenType())
}
