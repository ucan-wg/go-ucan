package envelope

import (
	"bytes"
	"errors"

	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	"github.com/ipld/go-ipld-prime/node/bindnode"
	"github.com/ipld/go-ipld-prime/schema"
	crypto "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/crypto/pb"
	"github.com/ucan-wg/go-ucan/internal/token"
	"github.com/ucan-wg/go-ucan/internal/varsig"
)

// Tokener represents a type that can be wrapped in a UCAN Envelope.
type Tokener interface {
	// Tag returns the a string that indicates which of the sub-
	// specifications define the structure of the underlying type.
	Tag() string

	Prototype() schema.TypedPrototype
}

// [Envelope] is a signed enclosure for types implementing Tokener.
//
// [Envelope]: https://github.com/ucan-wg/spec#envelope
type Envelope[T Tokener] struct {
	signature  []byte
	sigPayload *sigPayload[T]
}

// New creates an Envelope containing a VarsigHeader and Signature for
// the data resulting from wrapping the provided Tokener in and IPLD
// datamodel.Node and encoding it using DAG-CBOR
func New[T Tokener](privKey crypto.PrivKey, token T) (*Envelope[T], error) {
	sigPayload, err := newSigPayload[T](privKey.Type(), token)
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

	return &Envelope[T]{
		signature:  signature,
		sigPayload: sigPayload,
	}, nil
}

// Wrap is syntactic sugar for creating an Envelope and wrapping it as an
// IPLD datamodel.Node in a single operation.
//
// Since the Envelope itself isn't returned, us this method only when
// the IPLD datamodel.Node is used directly.  If the Envelope is also
// required, use New followed by Envelope.Wrap to avoid the need to
// unwrap the newly created datamodel.Node.
func Wrap[T Tokener](privKey crypto.PrivKey, token T) (datamodel.Node, error) {
	env, err := New[T](privKey, token)
	if err != nil {
		return nil, err
	}

	return env.Wrap()
}

func Unwrap[T Tokener](node datamodel.Node) (*Envelope[T], error) {
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

	sigPayload, err := unwrapSigPayload[T](sigPayloadNode)
	if err != nil {
		return nil, err
	}

	return &Envelope[T]{
		signature:  signature,
		sigPayload: sigPayload,
	}, nil
}

// Signature is an accessor that returns the cryptographic signature
// that was created when the Envelope was created or unwrapped.
func (e *Envelope[T]) Signature() []byte {
	return e.signature
}

func (e *Envelope[T]) Token() T {
	return e.sigPayload.tokenPayload
}

// VarsigHeader is an accessor that returns the [VarsigHeader] from the
// underlying [SigPayload] from the [Envelope].
//
// [Envelope]: https://github.com/ucan-wg/spec#envelope
// [SigPayload]: https://github.com/ucan-wg/spec#envelope
// [VarsigHeader]: https://github.com/ucan-wg/spec#envelope
func (e *Envelope[T]) VarsigHeader() []byte {
	return e.sigPayload.varsigHeader
}

// Verify checks that the [Envelope]'s signature is correct for the
// data created by encoding the SigPayload as DAG-CBOR and the public
// key passed as the only argument.
//
// Note that for Delegation and Invocation Tokeners, the public key
// is retrieved from the DID's method specific identifier.
//
// [Envelope]: https://github.com/ucan-wg/spec#envelope
func (e *Envelope[T]) Verify(pubKey crypto.PubKey) (bool, error) {
	cbor, err := e.sigPayload.cbor()
	if err != nil {
		return false, err
	}

	return pubKey.Verify(cbor, e.signature)
}

// Wrap encodes the Envelope as an IPLD datamodel.Node.
func (e *Envelope[T]) Wrap() (datamodel.Node, error) {
	sn := bindnode.Wrap(&e.signature, nil)

	spn, err := e.sigPayload.wrap()
	if err != nil {
		return nil, err
	}

	np := basicnode.Prototype.Any

	lb := np.NewBuilder()

	la, err := lb.BeginList(2)
	if err != nil {
		return nil, err
	}

	if err = la.AssembleValue().AssignNode(sn); err != nil {
		return nil, err
	}

	if err := la.AssembleValue().AssignNode(spn); err != nil {
		return nil, err
	}

	if err := la.Finish(); err != nil {
		return nil, err
	}

	return lb.Build(), nil
}

//
// The types below are strictly to make it easier to Wrap and Unwrap the
// Envelope with an IPLD datamodel.Node.  The Envelope itself provides
// accessors to the internals of these types.
//

type sigPayload[T Tokener] struct {
	varsigHeader []byte
	tokenPayload T
}

func newSigPayload[T Tokener](keyType pb.KeyType, token T) (*sigPayload[T], error) {
	varsigHeader, err := varsig.Encode(keyType)
	if err != nil {
		return nil, err
	}

	return &sigPayload[T]{
		varsigHeader: varsigHeader,
		tokenPayload: token,
	}, nil
}

func unwrapSigPayload[T Tokener](node datamodel.Node) (*sigPayload[T], error) {
	tokenPayloadNode, err := node.LookupByString((*new(T)).Tag())
	if err != nil {
		return nil, err
	}

	tokenPayload := bindnode.Unwrap(tokenPayloadNode)
	if tokenPayload == nil {
		return nil, errors.New("unexpected type") // TODO
	}

	token, ok := tokenPayload.(T)
	if !ok {
		return nil, errors.New("unexpected type") // TODO
	}

	headerNode, err := node.LookupByString("h")
	if err != nil {
		return nil, err
	}

	header, err := headerNode.AsBytes()
	if err != nil {
		return nil, err
	}

	return &sigPayload[T]{
		varsigHeader: header,
		tokenPayload: token,
	}, nil // TODO
}

func (sp *sigPayload[T]) cbor() ([]byte, error) {
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

func (sp *sigPayload[T]) wrap() (datamodel.Node, error) {
	tpn := bindnode.Wrap(sp.tokenPayload, sp.tokenPayload.Prototype().Type(), token.BindnodeOptions()...)

	np := basicnode.Prototype.Any
	mb := np.NewBuilder()

	ma, err := mb.BeginMap(2)
	if err != nil {
		return nil, err
	}

	ha, err := ma.AssembleEntry("h")
	if err != nil {
		return nil, err
	}
	ha.AssignBytes(sp.varsigHeader)

	ta, err := ma.AssembleEntry(sp.tokenPayload.Tag())
	if err != nil {
		return nil, err
	}
	ta.AssignNode(tpn)

	if err := ma.Finish(); err != nil {
		return nil, err
	}

	return mb.Build(), nil
}
