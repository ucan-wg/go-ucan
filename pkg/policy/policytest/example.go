package policytest

import (
	"errors"

	"github.com/ipld/go-ipld-prime"
	"github.com/ucan-wg/go-ucan/pkg/args"
	"github.com/ucan-wg/go-ucan/pkg/policy"
	"github.com/ucan-wg/go-ucan/pkg/policy/literal"
)

// EmptyPolicy provides a Policy with no statements.
var EmptyPolicy = policy.Policy{}

// ExampleValidationPolicy provides a instantiated Policy containing the
// statements that are included in the second code block of the [Validation]
// section of the delegation specification.
//
// [Validation]: https://github.com/ucan-wg/delegation/tree/v1_ipld#validation
var ExamplePolicy = policy.MustConstruct(
	policy.Equal(".from", literal.String("alice@example.com")),
	policy.Any(".to", policy.Like(".", "*@example.com")),
)

// TODO: Replace the URL for [Validation] above when the delegation
//       specification has been finished/merged.

// ExampleValidArguments provides valid, instantiated Arguments containing
// the key/value pairs that are included in portion of the the second code
// block of the [Validation] section of the delegation specification.
//
// [Validation]: https://github.com/ucan-wg/delegation/tree/v1_ipld#validation
var ExampleValidArguments = newBuilder(nil).
	add("from", "alice@example.com").
	add("to", []string{
		"bob@example.com",
		"carol@not.example.com",
	}).
	add("title", "Coffee").
	add("body", "Still on for coffee").
	mustBuild()

var exampleValidArgumentsIPLD = mustIPLD(ExampleValidArguments)

// ExampleInvalidArguments provides invalid, instantiated Arguments containing
// the key/value pairs that are included in portion of the the second code
// block of the [Validation] section of the delegation specification.
//
// [Validation]: https://github.com/ucan-wg/delegation/tree/v1_ipld#validation
var ExampleInvalidArguments = newBuilder(nil).
	add("from", "alice@example.com").
	add("to", []string{
		"bob@null.com",
		"carol@elsewhere.example.com",
	}).
	add("title", "Coffee").
	add("body", "Still on for coffee").
	mustBuild()

var exampleInvalidArgumentsIPLD = mustIPLD(ExampleInvalidArguments)

type builder struct {
	args *args.Args
	errs error
}

func newBuilder(a *args.Args) *builder {
	if a == nil {
		a = args.New()
	}

	return &builder{
		args: a,
	}
}

func (b *builder) add(key string, val any) *builder {
	b.errs = errors.Join(b.errs, b.args.Add(key, val))

	return b
}

func (b *builder) mustBuild() *args.Args {
	if b.errs != nil {
		panic(b.errs)
	}

	return b.args
}

func mustIPLD(args *args.Args) ipld.Node {
	node, err := args.ToIPLD()
	if err != nil {
		panic(err)
	}

	return node
}
