package args

import "github.com/ipld/go-ipld-prime"

type ReadOnly struct {
	args *Args
}

func (r ReadOnly) ToIPLD() (ipld.Node, error) {
	return r.args.ToIPLD()
}

func (r ReadOnly) Equals(other *Args) bool {
	return r.args.Equals(other)
}

func (r ReadOnly) String() string {
	return r.args.String()
}

func (r ReadOnly) WriteableClone() *Args {
	return r.args.Clone()
}
