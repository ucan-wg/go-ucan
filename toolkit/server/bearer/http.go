package bearer

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
)

// HttpArgsKey is the key in the args, used for:
// - if it exists in the invocation, holds a hash of the args derived from the HTTP request
// - in the final args to be evaluated against the policies, holds the args derived from the HTTP request
const HttpArgsKey = "http"

type HttpBearer struct {
	pol          policy.Policy
	originalArgs args.ReadOnly
	req          *http.Request

	once     sync.Once
	args     *args.Args
	argsIpld ipld.Node
}

func NewHttpBearer(pol policy.Policy, originalArgs args.ReadOnly, req *http.Request) *HttpBearer {
	return &HttpBearer{pol: pol, originalArgs: originalArgs, req: req}
}

func (hc *HttpBearer) Verify() error {
	if err := hc.makeArgs(); err != nil {
		return err
	}

	if err := hc.verifyHash(); err != nil {
		return err
	}

	ok, leaf := hc.pol.PartialMatch(hc.argsIpld)
	if !ok {
		return fmt.Errorf("the following UCAN policy is not satisfied: %v", leaf.String())
	}
	return nil
}

func (hc *HttpBearer) Args() (*args.Args, error) {
	if err := hc.makeArgs(); err != nil {
		return nil, err
	}
	return hc.args, nil
}

func (hc *HttpBearer) makeArgs() error {
	var outerErr error
	hc.once.Do(func() {
		var err error
		hc.args, err = makeHttpArgs(hc.req)
		if err != nil {
			outerErr = err
			return
		}

		hc.argsIpld, err = hc.args.ToIPLD()
		if err != nil {
			outerErr = err
			return
		}
	})
	return outerErr
}

func (hc *HttpBearer) verifyHash() error {
	n, err := hc.originalArgs.GetNode(HttpArgsKey)
	if err != nil {
		// no hash found, nothing to verify
		return nil
	}

	mhBytes, err := n.AsBytes()
	if err != nil {
		return fmt.Errorf("http args hash should be a string")
	}

	data, err := ipld.Encode(hc.argsIpld, dagcbor.Encode)
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
func MakeHttpHash(req *http.Request) ([]byte, error) {
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

	return sum, nil
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
