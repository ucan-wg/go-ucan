// Package extargs adds external arguments to the invocation's arguments before the policy is evaluated.
package extargs

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/fluent/qp"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	"github.com/multiformats/go-multihash"
	"github.com/ucan-wg/go-ucan/pkg/args"
	"github.com/ucan-wg/go-ucan/pkg/policy"
	"github.com/ucan-wg/go-ucan/token/invocation"
)

// HttpArgsKey is the key in the args, used for:
// - if it exists in the invocation, holds a hash of the args derived from the HTTP request
// - in the final args to be evaluated against the policies, holds the args derived from the HTTP request
const HttpArgsKey = "http"

type HttpExtArgs struct {
	pol          policy.Policy
	originalArgs args.ReadOnly
	req          *http.Request

	once     sync.Once
	args     *args.Args
	argsIpld ipld.Node
}

func NewHttpExtArgs(pol policy.Policy, originalArgs args.ReadOnly, req *http.Request) *HttpExtArgs {
	return &HttpExtArgs{pol: pol, originalArgs: originalArgs, req: req}
}

func (hea *HttpExtArgs) Verify() error {
	if err := hea.makeArgs(); err != nil {
		return err
	}

	if err := hea.verifyHash(); err != nil {
		return err
	}

	ok, leaf := hea.pol.PartialMatch(hea.argsIpld)
	if !ok {
		return fmt.Errorf("the following UCAN policy is not satisfied: %v", leaf.String())
	}
	return nil
}

func (hea *HttpExtArgs) Args() (*args.Args, error) {
	if err := hea.makeArgs(); err != nil {
		return nil, err
	}
	return hea.args, nil
}

func (hea *HttpExtArgs) makeArgs() error {
	var outerErr error
	hea.once.Do(func() {
		var err error
		hea.args, err = makeHttpArgs(hea.req)
		if err != nil {
			outerErr = err
			return
		}

		hea.argsIpld, err = hea.args.ToIPLD()
		if err != nil {
			outerErr = err
			return
		}
	})
	return outerErr
}

func (hea *HttpExtArgs) verifyHash() error {
	n, err := hea.originalArgs.GetNode(HttpArgsKey)
	if err != nil {
		// no hash found, nothing to verify
		return nil
	}

	mhBytes, err := n.AsBytes()
	if err != nil {
		return fmt.Errorf("http args hash should be a string")
	}

	data, err := ipld.Encode(hea.argsIpld, dagcbor.Encode)
	if err != nil {
		return fmt.Errorf("can't encode derived args in dag-cbor: %w", err)
	}

	sum, err := multihash.Sum(data, multihash.SHA2_256, -1)
	if err != nil {
		return err
	}

	if !bytes.Equal(mhBytes, sum) {
		return fmt.Errorf("derived args from http request don't match the expected hash")
	}

	return nil
}

// MakeHttpHash compute the hash of the derived arguments from the HTTP request.
// If that hash is inserted at the HttpArgsKey key in the invocation arguments,
// this increases the security as the UCAN token cannot be used with a different
// HTTP request.
// For convenience, the hash is returned as a read to use invocation argument.
func MakeHttpHash(req *http.Request) (invocation.Option, error) {
	// Note: the hash is computed on the full IPLD args, including HttpArgsKey
	computedArgs, err := makeHttpArgs(req)
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

	return invocation.WithArgument(HttpArgsKey, []byte(sum)), nil
}

func makeHttpArgs(req *http.Request) (*args.Args, error) {
	n, err := qp.BuildMap(basicnode.Prototype.Any, 4, func(ma datamodel.MapAssembler) {
		qp.MapEntry(ma, "scheme", qp.String(req.URL.Scheme)) // https
		qp.MapEntry(ma, "method", qp.String(req.Method))     // GET
		qp.MapEntry(ma, "host", qp.String(req.Host))         // example.com
		qp.MapEntry(ma, "path", qp.String(req.URL.Path))     // /foo

		qp.MapEntry(ma, "headers", qp.Map(2, func(ma datamodel.MapAssembler) {
			qp.MapEntry(ma, "Origin", qp.String(req.Header.Get("Origin")))
			qp.MapEntry(ma, "User-Agent", qp.String(req.Header.Get("User-Agent")))
		}))
	})
	if err != nil {
		return nil, err
	}

	res := args.New()
	err = res.Add(HttpArgsKey, n)
	if err != nil {
		return nil, err
	}

	return res, nil
}
