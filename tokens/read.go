package tokens

import (
	"fmt"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"

	"github.com/ucan-wg/go-ucan/tokens/delegation"
	"github.com/ucan-wg/go-ucan/tokens/internal/envelope"
)

func FromDagCbor(b []byte) (any, error) {
	node, err := ipld.Decode(b, dagcbor.Decode)
	if err != nil {
		return nil, err
	}

	tag, err := envelope.FindTag(node)
	if err != nil {
		return nil, err
	}

	switch tag {
	case delegation.Tag:
		return delegation.FromIPLD(node)
	default:
		return nil, fmt.Errorf(`unknown tag "%s"`, tag)
	}
}
