package literal

import (
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/basicnode"
)

func Node(n ipld.Node) ipld.Node {
	return n
}

func Link(cid ipld.Link) ipld.Node {
	nb := basicnode.Prototype.Link.NewBuilder()
	nb.AssignLink(cid)
	return nb.Build()
}

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

func Null() ipld.Node {
	nb := basicnode.Prototype.Any.NewBuilder()
	nb.AssignNull()
	return nb.Build()
}
