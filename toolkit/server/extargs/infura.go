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

const InfuraArgsKey = "inf"

type InfuraExtArgs struct {
	pol          policy.Policy
	originalArgs args.ReadOnly
	assembler    func(ma datamodel.MapAssembler)

	once     sync.Once
	args     *args.Args
	argsIpld ipld.Node
}

func NewInfuraExtArgs(pol policy.Policy, assembler func(ma datamodel.MapAssembler)) *InfuraExtArgs {
	return &InfuraExtArgs{pol: pol, assembler: assembler}
}

func (ia *InfuraExtArgs) Verify() error {
	if err := ia.makeArgs(); err != nil {
		return err
	}

	// Note: InfuraExtArgs doesn't support verifying a hash computed client-side like the other
	// external args, as the arguments are by nature dynamic. The client can't generate a meaningful hash.

	ok, leaf := ia.pol.PartialMatch(ia.argsIpld)
	if !ok {
		return fmt.Errorf("the following UCAN policy is not satisfied: %v", leaf.String())
	}
	return nil
}

func (ia *InfuraExtArgs) Args() (*args.Args, error) {
	if err := ia.makeArgs(); err != nil {
		return nil, err
	}
	return ia.args, nil
}

func (ia *InfuraExtArgs) makeArgs() error {
	var outerErr error
	ia.once.Do(func() {
		var err error

		ia.args, err = makeInfuraArgs(ia.assembler)
		if err != nil {
			outerErr = err
			return
		}

		ia.argsIpld, err = ia.args.ToIPLD()
		if err != nil {
			outerErr = err
			return
		}
	})
	return outerErr
}

func makeInfuraArgs(assembler func(ma datamodel.MapAssembler)) (*args.Args, error) {
	n, err := qp.BuildMap(basicnode.Prototype.Any, -1, assembler)
	if err != nil {
		return nil, err
	}

	res := args.New()
	err = res.Add(InfuraArgsKey, n)
	if err != nil {
		return nil, err
	}
	return res, nil
}
