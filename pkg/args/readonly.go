package args

import (
	"iter"

	"github.com/ipld/go-ipld-prime"
)

type ReadOnly struct {
	args *Args
}

func (r ReadOnly) Iter() iter.Seq2[string, ipld.Node] {
	return r.args.Iter()
}

func (r ReadOnly) ToIPLD() (ipld.Node, error) {
	return r.args.ToIPLD()
}

func (r ReadOnly) Equals(other ReadOnly) bool {
	return r.args.Equals(other.args)
}

func (r ReadOnly) String() string {
	return r.args.String()
}

func (r ReadOnly) WriteableClone() *Args {
	return r.args.Clone()
}
