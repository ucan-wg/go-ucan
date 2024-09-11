package envelope

import (
	"bytes"
	"errors"

	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	"github.com/ipld/go-ipld-prime/node/bindnode"
	crypto "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/crypto/pb"
	"github.com/multiformats/go-multibase"
	"github.com/multiformats/go-multicodec"
	"github.com/multiformats/go-multihash"
	"github.com/multiformats/go-varint"
	"github.com/ucan-wg/go-ucan/internal/token"
	"github.com/ucan-wg/go-ucan/internal/varsig"
)

// [Envelope] is a signed enclosure for types implementing Tokener.
//
// [Envelope]: https://github.com/ucan-wg/spec#envelope
type Envelope struct {
	signature  []byte
	sigPayload *sigPayload
}

// New creates an Envelope containing a VarsigHeader and Signature for
// the data resulting from wrapping the provided Tokener in and IPLD
// datamodel.Node and encoding it using DAG-CBOR
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
// Since the Envelope itself isn't returned, us this method only when
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

func Unwrap(node datamodel.Node, tag string) (*Envelope, error) {
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

	sigPayload, err := unwrapSigPayload(sigPayloadNode, tag)
	if err != nil {
		return nil, err
	}

	return &Envelope{
		signature:  signature,
		sigPayload: sigPayload,
	}, nil
}

func (e *Envelope) CID() (cid.Cid, error) {
	node, err := e.Wrap()
	if err != nil {
		return cid.Undef, nil
	}

	cbor, err := ipld.Encode(node, dagcbor.Encode)
	if err != nil {
		return cid.Undef, nil
	}

	data := varint.ToUvarint(uint64(multicodec.DagCbor))
	data = append(data, cbor...)

	hash, err := multihash.Sum(data, uint64(multicodec.Sha2_256), -1)
	if err != nil {
		return cid.Undef, err
	}

	return cid.NewCidV1(multibase.Base58BTC, hash), nil
}

// Signature is an accessor that returns the cryptographic signature
// that was created when the Envelope was created or unwrapped.
func (e *Envelope) Signature() []byte {
	return e.signature
}

func (e *Envelope) Token() *token.Token {
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
// Note that for Delegation and Invocation Tokeners, the public key
// is retrieved from the DID's method specific identifier.
//
// [Envelope]: https://github.com/ucan-wg/spec#envelope
func (e *Envelope) Verify(pubKey crypto.PubKey) (bool, error) {
	cbor, err := e.sigPayload.cbor()
	if err != nil {
		return false, err
	}

	return pubKey.Verify(cbor, e.signature)
}

// Wrap encodes the Envelope as an IPLD datamodel.Node.
func (e *Envelope) Wrap() (datamodel.Node, error) {
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

func unwrapSigPayload(node datamodel.Node, tag string) (*sigPayload, error) {
	tokenPayloadNode, err := node.LookupByString(tag)
	if err != nil {
		return nil, err
	}

	tokenPayload := bindnode.Unwrap(tokenPayloadNode)
	if tokenPayload == nil {
		return nil, errors.New("unexpected type") // TODO
	}

	token, ok := tokenPayload.(*token.Token)
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

	return &sigPayload{
		varsigHeader: header,
		tokenPayload: token,
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

	ta, err := ma.AssembleEntry(sp.tag)
	if err != nil {
		return nil, err
	}
	ta.AssignNode(tpn.Representation())

	if err := ma.Finish(); err != nil {
		return nil, err
	}

	return mb.Build(), nil
}
