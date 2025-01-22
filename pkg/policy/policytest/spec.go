package policytest

import (
	"github.com/ipld/go-ipld-prime"

	"github.com/ucan-wg/go-ucan/pkg/args"
	"github.com/ucan-wg/go-ucan/pkg/policy"
	"github.com/ucan-wg/go-ucan/pkg/policy/literal"
)

// EmptyPolicy provides a Policy with no statements.
var EmptyPolicy = policy.Policy{}

// SpecPolicy provides a  valid Policy containing the statements that are included
// in the second code block of the [Validation] section of the delegation specification.
//
// [Validation]: https://github.com/ucan-wg/delegation/tree/v1_ipld#validation
var SpecPolicy = policy.MustConstruct(
	policy.Equal(".from", literal.String("alice@example.com")),
	policy.Any(".to", policy.Like(".", "*@example.com")),
)

// TODO: Replace the URL for [Validation] above when the delegation
//       specification has been finished/merged.

// SpecValidArguments provides valid, instantiated Arguments containing
// the key/value pairs that are included in portion of the second code block
// of the [Validation] section of the delegation specification.
//
// [Validation]: https://github.com/ucan-wg/delegation/tree/v1_ipld#validation
var SpecValidArguments = args.NewBuilder().
	Add("from", "alice@example.com").
	Add("to", []string{
		"bob@example.com",
		"carol@not.example.com",
	}).
	Add("title", "Coffee").
	Add("body", "Still on for coffee").
	MustBuild()

var specValidArgumentsIPLD = mustIPLD(SpecValidArguments)

// SpecInvalidArguments provides invalid, instantiated Arguments containing
// the key/value pairs that are included in portion of the second code block
// of the [Validation] section of the delegation specification.
//
// [Validation]: https://github.com/ucan-wg/delegation/tree/v1_ipld#validation
var SpecInvalidArguments = args.NewBuilder().
	Add("from", "alice@example.com").
	Add("to", []string{
		"bob@null.com",
		"carol@elsewhere.example.com",
	}).
	Add("title", "Coffee").
	Add("body", "Still on for coffee").
	MustBuild()

var specInvalidArgumentsIPLD = mustIPLD(SpecInvalidArguments)

func mustIPLD(args *args.Args) ipld.Node {
	node, err := args.ToIPLD()
	if err != nil {
		panic(err)
	}

	return node
}
