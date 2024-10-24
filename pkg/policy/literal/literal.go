// Package literal holds a collection of functions to create IPLD types to use in policies, selector and args.
package literal

import (
	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/ipld/go-ipld-prime/node/basicnode"
)

// TODO: remove entirely?

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
