package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/node/bindnode"
	"github.com/ucan-wg/go-ucan/capability/command"
	"github.com/ucan-wg/go-ucan/capability/policy"
	"github.com/ucan-wg/go-ucan/did"
)

func BindnodeOptions() []bindnode.Option {
	return []bindnode.Option{
		CommandConverter(),
		DIDConverter(),
		MetaConverter(),
		PolicyConverter(),
		TimeConverter(),
	}
}

var ErrTypeAssertion = errors.New("failed to assert type")

func newErrTypeAssertion(where string) error {
	return fmt.Errorf("%w: %s", ErrTypeAssertion, where)
}

func CommandConverter() bindnode.Option {
	return bindnode.TypedStringConverter(
		(*command.Command)(nil),
		func(s string) (interface{}, error) {
			return command.Parse(s)
		},
		func(i interface{}) (string, error) {
			cmd, ok := i.(*command.Command)
			if !ok {
				return "", newErrTypeAssertion("CommandConverter")
			}

			return cmd.String(), nil
		},
	)
}

func DIDConverter() bindnode.Option {
	return bindnode.TypedStringConverter(
		(*did.DID)(nil),
		func(s string) (interface{}, error) {
			return did.Parse(s)
		},
		func(i interface{}) (string, error) {
			return i.(*did.DID).String(), nil
		},
	)
}

type Meta struct {
	Keys   []string
	Values map[string]any
}

func MetaConverter() bindnode.Option {
	return bindnode.TypedAnyConverter(
		(map[string]any)(nil),
		func(n datamodel.Node) (interface{}, error) {
			return Meta{}, nil // TODO
		},
		func(i interface{}) (datamodel.Node, error) {
			if i == nil {
				return datamodel.Null, nil
			}

			meta, ok := i.(Meta)
			if !ok {
				return nil, newErrTypeAssertion("MetaConverter")
			}

			_ = meta

			return datamodel.Null, nil // TODO
		},
	)
}

func PolicyConverter() bindnode.Option {
	return bindnode.TypedAnyConverter(
		(*policy.Policy)(nil),
		func(n datamodel.Node) (interface{}, error) {
			return policy.FromIPLD(n)
		},
		func(i interface{}) (datamodel.Node, error) {
			if i == nil {
				return datamodel.Null, nil
			}

			pol, ok := i.(*policy.Policy)
			if !ok {
				return nil, newErrTypeAssertion("PolicyConverter")
			}

			return pol.ToIPLD()
		},
	)
}

func TimeConverter() bindnode.Option {
	return bindnode.TypedIntConverter(
		(*time.Time)(nil),
		func(i int64) (interface{}, error) {
			return time.Unix(i, 0), nil
		},
		func(i interface{}) (int64, error) {
			if i == nil {
				return 0, nil
			}

			t, ok := i.(*time.Time)
			if !ok {
				return 0, newErrTypeAssertion("TimeConverter")
			}

			return t.Unix(), nil
		},
	)
}
