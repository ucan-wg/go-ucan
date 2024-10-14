package literal

import (
	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/ipld/go-ipld-prime/node/basicnode"
)

func Bool(val bool) ipld.Node {
	nb := basicnode.Prototype.Bool.NewBuilder()
	nb.AssignBool(val)
	return nb.Build()
}

func Int(val int64) ipld.Node {
	nb := basicnode.Prototype.Int.NewBuilder()
	nb.AssignInt(val)
	return nb.Build()
}

func Float(val float64) ipld.Node {
	nb := basicnode.Prototype.Float.NewBuilder()
	nb.AssignFloat(val)
	return nb.Build()
}

func String(val string) ipld.Node {
	nb := basicnode.Prototype.String.NewBuilder()
	nb.AssignString(val)
	return nb.Build()
}

func Bytes(val []byte) ipld.Node {
	nb := basicnode.Prototype.Bytes.NewBuilder()
	nb.AssignBytes(val)
	return nb.Build()
}

func Link(link ipld.Link) ipld.Node {
	nb := basicnode.Prototype.Link.NewBuilder()
	nb.AssignLink(link)
	return nb.Build()
}

func LinkCid(cid cid.Cid) ipld.Node {
	return Link(cidlink.Link{Cid: cid})
}

func Null() ipld.Node {
	nb := basicnode.Prototype.Any.NewBuilder()
	nb.AssignNull()
	return nb.Build()
}
