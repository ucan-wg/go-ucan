package invocation

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
	"github.com/ucan-wg/go-ucan/token/internal/envelope"
)

// ToSealed wraps the invocation token in an envelope, generates the
// signature, encodes the result to DAG-CBOR and calculates the CID of
// the resulting binary data.
func (t *Token) ToSealed(privKey crypto.PrivKey) ([]byte, cid.Cid, error) {
	data, err := t.ToDagCbor(privKey)
	if err != nil {
		return nil, cid.Undef, err
	}

	id, err := envelope.CIDFromBytes(data)
	if err != nil {
		return nil, cid.Undef, err
	}

	return data, id, nil
}

// ToSealedWriter is the same as ToSealed but accepts an io.Writer.
func (t *Token) ToSealedWriter(w io.Writer, privKey crypto.PrivKey) (cid.Cid, error) {
	cidWriter := envelope.NewCIDWriter(w)

	if err := t.ToDagCborWriter(cidWriter, privKey); err != nil {
		return cid.Undef, err
	}

	return cidWriter.CID()
}

// FromSealed decodes the provided binary data from the DAG-CBOR format,
// verifies that the envelope's signature is correct based on the public
// key taken from the issuer (iss) field and calculates the CID of the
// incoming data.
func FromSealed(data []byte) (*Token, cid.Cid, error) {
	tkn, err := FromDagCbor(data)
	if err != nil {
		return nil, cid.Undef, err
	}

	id, err := envelope.CIDFromBytes(data)
	if err != nil {
		return nil, cid.Undef, err
	}

	return tkn, id, nil
}

// FromSealedReader is the same as Unseal but accepts an io.Reader.
func FromSealedReader(r io.Reader) (*Token, cid.Cid, error) {
	cidReader := envelope.NewCIDReader(r)

	tkn, err := FromDagCborReader(cidReader)
	if err != nil {
		return nil, cid.Undef, err
	}

	id, err := cidReader.CID()
	if err != nil {
		return nil, cid.Undef, err
	}

	return tkn, id, nil
}

// Encode marshals a Token to the format specified by the provided
// codec.Encoder.
func (t *Token) Encode(privKey crypto.PrivKey, encFn codec.Encoder) ([]byte, error) {
	node, err := t.toIPLD(privKey)
	if err != nil {
		return nil, err
	}

	return ipld.Encode(node, encFn)
}

// EncodeWriter is the same as Encode, but accepts an io.Writer.
func (t *Token) EncodeWriter(w io.Writer, privKey crypto.PrivKey, encFn codec.Encoder) error {
	node, err := t.toIPLD(privKey)
	if err != nil {
		return err
	}

	return ipld.EncodeStreaming(w, node, encFn)
}

// ToDagCbor marshals the Token to the DAG-CBOR format.
func (t *Token) ToDagCbor(privKey crypto.PrivKey) ([]byte, error) {
	return t.Encode(privKey, dagcbor.Encode)
}

// ToDagCborWriter is the same as ToDagCbor, but it accepts an io.Writer.
func (t *Token) ToDagCborWriter(w io.Writer, privKey crypto.PrivKey) error {
	return t.EncodeWriter(w, privKey, dagcbor.Encode)
}

// ToDagJson marshals the Token to the DAG-JSON format.
func (t *Token) ToDagJson(privKey crypto.PrivKey) ([]byte, error) {
	return t.Encode(privKey, dagjson.Encode)
}

// ToDagJsonWriter is the same as ToDagJson, but it accepts an io.Writer.
func (t *Token) ToDagJsonWriter(w io.Writer, privKey crypto.PrivKey) error {
	return t.EncodeWriter(w, privKey, dagjson.Encode)
}

// Decode unmarshals the input data using the format specified by the
// provided codec.Decoder into a Token.
//
// An error is returned if the conversion fails, or if the resulting
// Token is invalid.
func Decode(b []byte, decFn codec.Decoder) (*Token, error) {
	node, err := ipld.Decode(b, decFn)
	if err != nil {
		return nil, err
	}
	return FromIPLD(node)
}

// DecodeReader is the same as Decode, but accept an io.Reader.
func DecodeReader(r io.Reader, decFn codec.Decoder) (*Token, error) {
	node, err := ipld.DecodeStreaming(r, decFn)
	if err != nil {
		return nil, err
	}
	return FromIPLD(node)
}

// FromDagCbor unmarshals the input data into a Token.
//
// An error is returned if the conversion fails, or if the resulting
// Token is invalid.
func FromDagCbor(data []byte) (*Token, error) {
	pay, err := envelope.FromDagCbor[*tokenPayloadModel](data)
	if err != nil {
		return nil, err
	}

	tkn, err := tokenFromModel(*pay)
	if err != nil {
		return nil, err
	}

	return tkn, err
}

// FromDagCborReader is the same as FromDagCbor, but accept an io.Reader.
func FromDagCborReader(r io.Reader) (*Token, error) {
	return DecodeReader(r, dagcbor.Decode)
}

// FromDagJson unmarshals the input data into a Token.
//
// An error is returned if the conversion fails, or if the resulting
// Token is invalid.
func FromDagJson(data []byte) (*Token, error) {
	return Decode(data, dagjson.Decode)
}

// FromDagJsonReader is the same as FromDagJson, but accept an io.Reader.
func FromDagJsonReader(r io.Reader) (*Token, error) {
	return DecodeReader(r, dagjson.Decode)
}

// FromIPLD decode the given IPLD representation into a Token.
func FromIPLD(node datamodel.Node) (*Token, error) {
	pay, err := envelope.FromIPLD[*tokenPayloadModel](node)
	if err != nil {
		return nil, err
	}

	tkn, err := tokenFromModel(*pay)
	if err != nil {
		return nil, err
	}

	return tkn, err
}

func (t *Token) toIPLD(privKey crypto.PrivKey) (datamodel.Node, error) {
	var aud *string

	if t.audience != did.Undef {
		a := t.audience.String()
		aud = &a
	}

	var exp *int64
	if t.expiration != nil {
		u := t.expiration.Unix()
		exp = &u
	}

	var iat *int64
	if t.invokedAt != nil {
		i := t.invokedAt.Unix()
		iat = &i
	}

	model := &tokenPayloadModel{
		Iss:   t.issuer.String(),
		Aud:   aud,
		Sub:   t.subject.String(),
		Cmd:   t.command.String(),
		Args:  t.arguments,
		Prf:   t.proof,
		Meta:  t.meta,
		Nonce: t.nonce,
		Exp:   exp,
		Iat:   iat,
		Cause: t.cause,
	}

	// seems like it's a requirement to have a null meta if there are no values?
	if len(model.Meta.Keys) == 0 {
		model.Meta = nil
	}

	return envelope.ToIPLD(privKey, model)
}
