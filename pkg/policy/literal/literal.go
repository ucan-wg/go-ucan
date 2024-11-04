// Package literal holds a collection of functions to create IPLD types to use in policies, selector and args.
package literal

import (
	"fmt"

	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/ipld/go-ipld-prime/node/basicnode"
)

var Bool = basicnode.NewBool
var Int = basicnode.NewInt
var Float = basicnode.NewFloat
var String = basicnode.NewString
var Bytes = basicnode.NewBytes
var Link = basicnode.NewLink

func LinkCid(cid cid.Cid) ipld.Node {
	return Link(cidlink.Link{Cid: cid})
}

func Null() ipld.Node {
	nb := basicnode.Prototype.Any.NewBuilder()
	nb.AssignNull()
	return nb.Build()
}

// Map creates an IPLD node from a map[string]any
func Map(m map[string]any) (ipld.Node, error) {
	nb := basicnode.Prototype.Map.NewBuilder()
	ma, err := nb.BeginMap(int64(len(m)))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		if err := ma.AssembleKey().AssignString(k); err != nil {
			return nil, err
		}

		switch x := v.(type) {
		case string:
			if err := ma.AssembleValue().AssignString(x); err != nil {
				return nil, err
			}
		case []any:
			lb := basicnode.Prototype.List.NewBuilder()
			la, err := lb.BeginList(int64(len(x)))
			if err != nil {
				return nil, err
			}
			if err := la.Finish(); err != nil {
				return nil, err
			}
			if err := ma.AssembleValue().AssignNode(lb.Build()); err != nil {
				return nil, err
			}
		case map[string]any:
			nestedNode, err := Map(x) // recursive call for nested maps
			if err != nil {
				return nil, err
			}
			if err := ma.AssembleValue().AssignNode(nestedNode); err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("unsupported value type: %T", v)
		}
	}

	if err := ma.Finish(); err != nil {
		return nil, err
	}

	return nb.Build(), nil
}
