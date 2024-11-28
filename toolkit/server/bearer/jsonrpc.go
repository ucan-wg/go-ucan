package bearer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/INFURA/go-ethlibs/jsonrpc"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/fluent/qp"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	"github.com/multiformats/go-multihash"
	"github.com/ucan-wg/go-ucan/pkg/args"
	"github.com/ucan-wg/go-ucan/pkg/policy"
	"github.com/ucan-wg/go-ucan/pkg/policy/literal"
)

// JsonRpcArgsKey is the key in the args, used for:
// - if it exists in the invocation, holds a hash of the args derived from the JsonRpc request
// - in the final args to be evaluated against the policies, holds the args derived from the JsonRpc request
const JsonRpcArgsKey = "jsonrpc"

type JsonRpcBearer struct {
	pol          policy.Policy
	originalArgs args.ReadOnly
	req          *jsonrpc.Request

	once     sync.Once
	args     *args.Args
	argsIpld ipld.Node
}

func NewJsonRpcBearer(pol policy.Policy, originalArgs args.ReadOnly, req *jsonrpc.Request) *JsonRpcBearer {
	return &JsonRpcBearer{pol: pol, originalArgs: originalArgs, req: req}
}

func (jrc *JsonRpcBearer) Verify() error {
	if err := jrc.makeArgs(); err != nil {
		return err
	}

	if err := jrc.verifyHash(); err != nil {
		return err
	}

	ok, leaf := jrc.pol.PartialMatch(jrc.argsIpld)
	if !ok {
		return fmt.Errorf("the following UCAN policy is not satisfied: %v", leaf.String())
	}
	return nil
}

func (jrc *JsonRpcBearer) Args() (*args.Args, error) {
	if err := jrc.makeArgs(); err != nil {
		return nil, err
	}
	return jrc.args, nil
}

func (jrc *JsonRpcBearer) makeArgs() error {
	var outerErr error
	jrc.once.Do(func() {
		var err error
		jrc.args, err = makeJsonRpcArgs(jrc.req)
		if err != nil {
			outerErr = err
			return
		}

		jrc.argsIpld, err = jrc.args.ToIPLD()
		if err != nil {
			outerErr = err
			return
		}
	})
	return outerErr
}

func (jrc *JsonRpcBearer) verifyHash() error {
	n, err := jrc.originalArgs.GetNode(JsonRpcArgsKey)
	if err != nil {
		// no hash found, nothing to verify
		return nil
	}

	mhBytes, err := n.AsBytes()
	if err != nil {
		return fmt.Errorf("jsonrpc args hash should be a string")
	}

	data, err := ipld.Encode(jrc.argsIpld, dagcbor.Encode)
	if err != nil {
		return fmt.Errorf("can't encode derived args in dag-cbor: %w", err)
	}

	sum, err := multihash.Sum(data, multihash.SHA2_256, -1)
	if err != nil {
		return err
	}

	if !bytes.Equal(mhBytes, sum) {
		return fmt.Errorf("derived args from jsonrpc request don't match the expected hash")
	}

	return nil
}

// MakeJsonRpcHash compute the hash of the derived arguments from the JsonRPC request.
// If that hash is inserted at the JsonRpcArgsKey key in the invocation arguments,
// this increases the security as the UCAN token cannot be used with a different
// JsonRPC request.
func MakeJsonRpcHash(req *jsonrpc.Request) ([]byte, error) {
	computedArgs, err := makeJsonRpcArgs(req)
	if err != nil {
		return nil, err
	}

	n, err := computedArgs.ToIPLD()
	if err != nil {
		return nil, err
	}

	data, err := ipld.Encode(n, dagcbor.Encode)
	if err != nil {
		return nil, err
	}

	sum, err := multihash.Sum(data, multihash.SHA2_256, -1)
	if err != nil {
		return nil, err
	}

	return sum, nil
}

func makeJsonRpcArgs(req *jsonrpc.Request) (*args.Args, error) {
	deserialized := make([]any, len(req.Params))
	for i, param := range req.Params {
		err := json.Unmarshal(param, &deserialized[i])
		if err != nil {
			return nil, err
		}
	}

	params, err := literal.List(deserialized)
	if err != nil {
		return nil, err
	}

	n, err := qp.BuildMap(basicnode.Prototype.Any, 3, func(ma datamodel.MapAssembler) {
		qp.MapEntry(ma, "jsonrpc", qp.String(req.JSONRPC))
		qp.MapEntry(ma, "method", qp.String(req.Method))
		qp.MapEntry(ma, "params", qp.Node(params))
	})
	if err != nil {
		return nil, err
	}

	res := args.New()
	err = res.Add(JsonRpcArgsKey, n)
	if err != nil {
		return nil, err
	}
	return res, nil
}
