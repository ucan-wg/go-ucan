package tokens

import (
	"github.com/ipld/go-ipld-prime/datamodel"

	"github.com/ucan-wg/go-ucan/tokens/internal/envelope"
)

type Info = envelope.Info

// Inspect inspects the given token IPLD representation and extract some envelope facts.
func Inspect(node datamodel.Node) (Info, error) {
	return envelope.Inspect(node)
}

// FindTag inspect the given token IPLD representation and extract the token tag.
func FindTag(node datamodel.Node) (string, error) {
	return envelope.FindTag(node)
}
