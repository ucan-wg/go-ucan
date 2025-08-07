package extargs

import (
	"fmt"
	"sync"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/fluent/qp"
	"github.com/ipld/go-ipld-prime/node/basicnode"

	"github.com/ucan-wg/go-ucan/pkg/args"
	"github.com/ucan-wg/go-ucan/pkg/policy"
)

type CustomExtArgs struct {
	key          string
	pol          policy.Policy
	originalArgs args.ReadOnly
	assembler    func(ma datamodel.MapAssembler)

	once     sync.Once
	args     *args.Args
	argsIpld ipld.Node
}

func NewCustomExtArgs(key string, pol policy.Policy, assembler func(ma datamodel.MapAssembler)) *CustomExtArgs {
	return &CustomExtArgs{key: key, pol: pol, assembler: assembler}
}

func (cea *CustomExtArgs) Verify() error {
	if err := cea.makeArgs(); err != nil {
		return err
	}

	// Note: CustomExtArgs doesn't support verifying a hash computed client-side like the other
	// external args, as the arguments are by nature dynamic. The client can't generate a meaningful hash.

	ok, leaf := cea.pol.PartialMatch(cea.argsIpld)
	if !ok {
		return fmt.Errorf("the following UCAN policy is not satisfied: %v", leaf.String())
	}
	return nil
}

func (cea *CustomExtArgs) Args() (*args.Args, error) {
	if err := cea.makeArgs(); err != nil {
		return nil, err
	}
	return cea.args, nil
}

func (cea *CustomExtArgs) makeArgs() error {
	var outerErr error
	cea.once.Do(func() {
		var err error

		cea.args, err = makeCustomArgs(cea.key, cea.assembler)
		if err != nil {
			outerErr = err
			return
		}

		cea.argsIpld, err = cea.args.ToIPLD()
		if err != nil {
			outerErr = err
			return
		}
	})
	return outerErr
}

func makeCustomArgs(key string, assembler func(ma datamodel.MapAssembler)) (*args.Args, error) {
	n, err := qp.BuildMap(basicnode.Prototype.Any, -1, assembler)
	if err != nil {
		return nil, err
	}

	res := args.New()
	err = res.Add(key, n)
	if err != nil {
		return nil, err
	}
	return res, nil
}
