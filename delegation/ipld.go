package delegation

import (
	"io"

	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/libp2p/go-libp2p/core/crypto"

	"github.com/ucan-wg/go-ucan/did"
	"github.com/ucan-wg/go-ucan/internal/envelope"
)

// Encode marshals a View to the format specified by the provided
// codec.Encoder.
func (t *Token) Encode(privKey crypto.PrivKey, encFn codec.Encoder) ([]byte, error) {
	node, err := t.ToIPLD(privKey)
	if err != nil {
		return nil, err
	}

	return ipld.Encode(node, encFn)
}

// EncodeWriter is the same as Encode but accepts an io.Writer.
func (t *Token) EncodeWriter(w io.Writer, privKey crypto.PrivKey, encFn codec.Encoder) error {
	node, err := t.ToIPLD(privKey)
	if err != nil {
		return err
	}

	return ipld.EncodeStreaming(w, node, encFn)
}

// ToDagCbor marshals the View to the DAG-CBOR format.
func (t *Token) ToDagCbor(privKey crypto.PrivKey) ([]byte, error) {
	return t.Encode(privKey, dagcbor.Encode)
}

// ToDagCborWriter is the same as ToDagCbor but it accepts an io.Writer.
func (t *Token) ToDagCborWriter(w io.Writer, privKey crypto.PrivKey) error {
	return t.EncodeWriter(w, privKey, dagcbor.Encode)
}

// ToDagJson marshals the View to the DAG-JSON format.
func (t *Token) ToDagJson(privKey crypto.PrivKey) ([]byte, error) {
	return t.Encode(privKey, dagjson.Encode)
}

// ToDagJsonWriter is the same as ToDagJson but it accepts an io.Writer.
func (t *Token) ToDagJsonWriter(w io.Writer, privKey crypto.PrivKey) error {
	return t.EncodeWriter(w, privKey, dagjson.Encode)
}

// ToIPLD wraps the View in an IPLD datamodel.Node.
func (t *Token) ToIPLD(privKey crypto.PrivKey) (datamodel.Node, error) {
	var sub *string
	if t.subject != did.Undef {
		s := t.subject.String()
		sub = &s
	}

	pol, err := t.policy.ToIPLD()
	if err != nil {
		return nil, err
	}

	var nbf *int64
	if t.notBefore != nil {
		u := t.notBefore.Unix()
		nbf = &u
	}

	var exp *int64
	if t.expiration != nil {
		u := t.expiration.Unix()
		exp = &u
	}

	model := &tokenPayloadModel{
		Iss:   t.issuer.String(),
		Aud:   t.audience.String(),
		Sub:   sub,
		Cmd:   t.command.String(),
		Pol:   pol,
		Nonce: t.nonce,
		Meta:  *t.meta,
		Nbf:   nbf,
		Exp:   exp,
	}

	return envelope.ToIPLD(privKey, model)
}

// Decode unmarshals the input data using the format specified by the
// provided codec.Decoder into a View.
//
// An error is returned if the conversion fails, or if the resulting
// View is invalid.
func Decode(b []byte, decFn codec.Decoder) (*Token, cid.Cid, error) {
	node, err := ipld.Decode(b, decFn)
	if err != nil {
		return nil, cid.Undef, err
	}
	return FromIPLD(node)
}

// DecodeReader is the same as Decode, but accept an io.Reader.
func DecodeReader(r io.Reader, decFn codec.Decoder) (*Token, cid.Cid, error) {
	node, err := ipld.DecodeStreaming(r, decFn)
	if err != nil {
		return nil, cid.Undef, err
	}
	return FromIPLD(node)
}

// FromDagCbor unmarshals the input data into a View.
//
// An error is returned if the conversion fails, or if the resulting
// View is invalid.
func FromDagCbor(data []byte) (*Token, cid.Cid, error) {
	pay, id, err := envelope.FromDagCbor[*tokenPayloadModel](data)
	if err != nil {
		return nil, cid.Undef, err
	}

	tkn, err := tokenFromModel(*pay)
	if err != nil {
		return nil, cid.Undef, err
	}

	return tkn, id, err
}

// FromDagCborReader is the same as FromDagCbor, but accept an io.Reader.
func FromDagCborReader(r io.Reader) (*Token, cid.Cid, error) {
	return DecodeReader(r, dagcbor.Decode)
}

// FromDagJson unmarshals the input data into a View.
//
// An error is returned if the conversion fails, or if the resulting
// View is invalid.
func FromDagJson(data []byte) (*Token, cid.Cid, error) {
	return Decode(data, dagjson.Decode)
}

// FromDagJsonReader is the same as FromDagJson, but accept an io.Reader.
func FromDagJsonReader(r io.Reader) (*Token, cid.Cid, error) {
	return DecodeReader(r, dagjson.Decode)
}

// FromIPLD unwraps a View from the provided IPLD datamodel.Node
//
// An error is returned if the conversion fails, or if the resulting
// View is invalid.
func FromIPLD(node datamodel.Node) (*Token, cid.Cid, error) {
	pay, id, err := envelope.FromIPLD[*tokenPayloadModel](node) // TODO add CID to view
	if err != nil {
		return nil, cid.Undef, err
	}

	tkn, err := tokenFromModel(*pay)
	if err != nil {
		return nil, cid.Undef, err
	}

	return tkn, id, err
}
