package envelope

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/fluent/qp"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	"github.com/ipld/go-ipld-prime/node/bindnode"
	crypto "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/crypto/pb"
	"github.com/ucan-wg/go-ucan/did"
	"github.com/ucan-wg/go-ucan/internal/token"
	"github.com/ucan-wg/go-ucan/internal/varsig"
)

// [Envelope] is a signed enclosure for a UCAN v1 Token.
//
// While the types and functions in this package are not exported,
// the names used for types, fields, variables, etc generally use the
// names from the specification
//
// [Envelope]: https://github.com/ucan-wg/spec#envelope
type Envelope struct {
	signature  []byte
	sigPayload *sigPayload
}

// New creates an Envelope containing a VarsigHeader and Signature for
// the data resulting from wrapping the provided Token in an IPLD
// datamodel.Node and encoding it using DAG-CBOR.
func New(privKey crypto.PrivKey, token *token.Token, tag string) (*Envelope, error) {
	sigPayload, err := newSigPayload(privKey.Type(), token, tag)
	if err != nil {
		return nil, err
	}

	cbor, err := sigPayload.cbor()
	if err != nil {
		return nil, err
	}

	signature, err := privKey.Sign(cbor)
	if err != nil {
		return nil, err
	}

	return &Envelope{
		signature:  signature,
		sigPayload: sigPayload,
	}, nil
}

// Wrap is syntactic sugar for creating an Envelope and wrapping it as an
// IPLD datamodel.Node in a single operation.
//
// Since the Envelope itself isn't returned, use this method only when
// the IPLD datamodel.Node is used directly.  If the Envelope is also
// required, use New followed by Envelope.Wrap to avoid the need to
// unwrap the newly created datamodel.Node.
func Wrap(privKey crypto.PrivKey, token *token.Token, tag string) (datamodel.Node, error) {
	env, err := New(privKey, token, tag)
	if err != nil {
		return nil, err
	}

	return env.Wrap()
}

// Unwrap attempts to crate an Envelope from a datamodel.Node
//
// There are lots of ways that this can fail and therefore there are
// an almost excessive number of check included here and while
// attempting to extract the token.Token from one of the inner IPLD
// nodes.
func Unwrap(node datamodel.Node) (*Envelope, error) {
	signatureNode, err := node.LookupByIndex(0)
	if err != nil {
		return nil, err
	}

	signature, err := signatureNode.AsBytes()
	if err != nil {
		return nil, err
	}

	sigPayloadNode, err := node.LookupByIndex(1)
	if err != nil {
		return nil, err
	}

	sigPayload, err := unwrapSigPayload(sigPayloadNode)
	if err != nil {
		return nil, err
	}

	envel := &Envelope{
		signature:  signature,
		sigPayload: sigPayload,
	}

	if ok, err := envel.Verify(); !ok || err != nil {
		return nil, fmt.Errorf("envelope was not signed by issuer")
	}

	return envel, nil
}

// Signature returns the cryptographic signature of the Envelope's
// SigPayload.
func (e *Envelope) Signature() []byte {
	return e.signature
}

// Tag returns the key that's used to reference the TokenPayload within
// this Envelope.
func (e *Envelope) Tag() string {
	return e.sigPayload.tag
}

// TokenPayload returns the *token.Token enclosed within this Envelope.
func (e *Envelope) TokenPayload() *token.Token {
	return e.sigPayload.tokenPayload
}

// VarsigHeader is an accessor that returns the [VarsigHeader] from the
// underlying [SigPayload] from the [Envelope].
//
// [Envelope]: https://github.com/ucan-wg/spec#envelope
// [SigPayload]: https://github.com/ucan-wg/spec#envelope
// [VarsigHeader]: https://github.com/ucan-wg/spec#envelope
func (e *Envelope) VarsigHeader() []byte {
	return e.sigPayload.varsigHeader
}

// Verify checks that the [Envelope]'s signature is correct for the
// data created by encoding the SigPayload as DAG-CBOR and the public
// key passed as the only argument.
//
// Note that for Delegation and Invocation tokens, the public key
// is retrieved from the DID's method specific identifier for the
// Issuer field.
//
// [Envelope]: https://github.com/ucan-wg/spec#envelope
func (e *Envelope) Verify() (bool, error) {
	pubKey, err := did.ToPubKey(e.sigPayload.tokenPayload.Issuer)
	if err != nil {
		return false, err
	}

	cbor, err := e.sigPayload.cbor()
	if err != nil {
		return false, err
	}

	return pubKey.Verify(cbor, e.signature)
}

// Wrap encodes the Envelope as an IPLD datamodel.Node.
func (e *Envelope) Wrap() (datamodel.Node, error) {
	spn, err := e.sigPayload.wrap()
	if err != nil {
		return nil, err
	}

	return qp.BuildList(basicnode.Prototype.Any, 2, func(la datamodel.ListAssembler) {
		qp.ListEntry(la, qp.Bytes(e.signature))
		qp.ListEntry(la, qp.Node(spn))
	})
}

//
// The types below are strictly to make it easier to Wrap and Unwrap the
// Envelope with an IPLD datamodel.Node.  The Envelope itself provides
// accessors to the internals of these types.
//

type sigPayload struct {
	varsigHeader []byte
	tokenPayload *token.Token
	tag          string
}

func newSigPayload(keyType pb.KeyType, token *token.Token, tag string) (*sigPayload, error) {
	varsigHeader, err := varsig.Encode(keyType)
	if err != nil {
		return nil, err
	}

	return &sigPayload{
		varsigHeader: varsigHeader,
		tokenPayload: token,
		tag:          tag,
	}, nil
}

func unwrapSigPayload(node datamodel.Node) (*sigPayload, error) {
	// Normally we could look up the VarsigHeader and TokenPayload using
	// node.LookupByString() - this works for the "h" key used for the
	// VarsigHeader but not for the TokenPayload's key (tag) as all we
	// know is that it starts with "ucan/" and as explained below, must
	// decode to a schema.TypedNode for the representation provided by the
	// token.Prototype().
	// vvv
	mi := node.MapIterator()
	if mi == nil {
		return nil, fmt.Errorf("the SigPayload node is not a map: %s", node.Kind().String())
	}

	var (
		hdrNode datamodel.Node
		tknNode datamodel.Node
		tag     string
	)

	keyCount := 0

	for !mi.Done() {
		k, v, err := mi.Next()
		if err != nil {
			return nil, err
		}

		kStr, err := k.AsString()
		if err != nil {
			return nil, fmt.Errorf("the SigPayload keys are not strings: %w", err)
		}

		keyCount++

		if kStr == "h" {
			hdrNode = v

			continue
		}

		if strings.HasPrefix(kStr, "ucan/") {
			tknNode = v
			tag = kStr
		}
	}

	if keyCount != 2 {
		return nil, fmt.Errorf("the SigPayload map should have exactly two keys: %d", keyCount)
	}
	// ^^^

	// Replaces the datamodel.Node in tokenPayloadNode with a
	// schema.TypedNode so that we can cast it to a *token.Token after
	// unwrapping it.
	// vvv
	nb := token.Prototype().Representation().NewBuilder()

	err := nb.AssignNode(tknNode)
	if err != nil {
		return nil, err
	}

	tknNode = nb.Build()
	// ^^^

	tokenPayload := bindnode.Unwrap(tknNode)
	if tokenPayload == nil {
		return nil, errors.New("failed to Unwrap the TokenPayload")
	}

	tkn, ok := tokenPayload.(*token.Token)
	if !ok {
		return nil, errors.New("failed to assert the TokenPayload type as *token.Token")
	}

	hdr, err := hdrNode.AsBytes()
	if err != nil {
		return nil, err
	}

	return &sigPayload{
		varsigHeader: hdr,
		tokenPayload: tkn,
		tag:          tag,
	}, nil
}

func (sp *sigPayload) cbor() ([]byte, error) {
	node, err := sp.wrap()
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}
	if err = dagcbor.Encode(node, buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (sp *sigPayload) wrap() (datamodel.Node, error) {
	tpn := bindnode.Wrap(sp.tokenPayload, token.Prototype().Type())

	return qp.BuildMap(basicnode.Prototype.Any, 2, func(ma datamodel.MapAssembler) {
		qp.MapEntry(ma, "h", qp.Bytes(sp.varsigHeader))
		qp.MapEntry(ma, sp.tag, qp.Node(tpn.Representation()))
	})
}
