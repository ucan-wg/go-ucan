package args

import (
	"errors"

	"github.com/ipld/go-ipld-prime"
)

// Builder allows the fluid construction of an Args.
type Builder struct {
	args *Args
	errs error
}

// NewBuilder returns a Builder which will assemble the Args.
func NewBuilder() *Builder {
	return &Builder{
		args: New(),
	}
}

// Add inserts a new key/val into the Args being assembled while collecting
// any errors caused by duplicate keys.
func (b *Builder) Add(key string, val any) *Builder {
	b.errs = errors.Join(b.errs, b.args.Add(key, val))

	return b
}

// Build returns the assembled Args or an error containing a list of
// errors encountered while trying to build the Args.
func (b *Builder) Build() (*Args, error) {
	if b.errs != nil {
		return nil, b.errs
	}

	return b.args, nil
}

// BuildIPLD is the same as Build except it takes the additional step of
// converting the Args to an ipld.Node.
func (b *Builder) BuildIPLD() (ipld.Node, error) {
	args, err := b.Build()
	if err != nil {
		return nil, err
	}

	return args.ToIPLD()
}

// MustBuild is the same as Build except it panics if an error occurs.
func (b *Builder) MustBuild() *Args {
	args, err := b.Build()

	if err != nil {
		panic(b.errs)
	}

	return args
}

// MustBuildIPLD is the same as BuildIPLD except it panics if an error
// occurs.
func (b *Builder) MustBuildIPLD() ipld.Node {
	node, err := b.BuildIPLD()
	if err != nil {
		panic(err)
	}

	return node
}
