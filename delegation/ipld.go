package delegation

import (
	"io"

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
func (d *View) Encode(privKey crypto.PrivKey, encFn codec.Encoder) ([]byte, error) {
	node, err := d.ToIPLD(privKey)
	if err != nil {
		return nil, err
	}

	return ipld.Encode(node, encFn)
}

// EncodeWriter is the same as Encode but accepts an io.Writer.
func (d *View) EncodeWriter(w io.Writer, privKey crypto.PrivKey, encFn codec.Encoder) error {
	node, err := d.ToIPLD(privKey)
	if err != nil {
		return err
	}

	return ipld.EncodeStreaming(w, node, encFn)
}

// ToDagCbor marshals the View to the DAG-CBOR format.
func (d *View) ToDagCbor(privKey crypto.PrivKey) ([]byte, error) {
	return d.Encode(privKey, dagcbor.Encode)
}

// ToDagCborWriter is the same as ToDagCbor but it accepts an io.Writer.
func (d *View) ToDagCborWriter(w io.Writer, privKey crypto.PrivKey) error {
	return d.EncodeWriter(w, privKey, dagcbor.Encode)
}

// ToDagJson marshals the View to the DAG-JSON format.
func (d *View) ToDagJson(privKey crypto.PrivKey) ([]byte, error) {
	return d.Encode(privKey, dagjson.Encode)
}

// ToDagJsonWriter is the same as ToDagJson but it accepts an io.Writer.
func (d *View) ToDagJsonWriter(w io.Writer, privKey crypto.PrivKey) error {
	return d.EncodeWriter(w, privKey, dagjson.Encode)
}

// ToIPLD wraps the View in an IPLD datamodel.Node.
func (d *View) ToIPLD(privKey crypto.PrivKey) (datamodel.Node, error) {
	var sub *string
	if d.Subject != did.Undef {
		s := d.Subject.String()
		sub = &s
	}

	pol, err := d.Policy.ToIPLD()
	if err != nil {
		return nil, err
	}

	metaKeys := make([]string, len(d.Meta))
	i := 0

	for k := range d.Meta {
		metaKeys[i] = k
		i++
	}

	var nbf *int64
	if d.NotBefore != nil {
		u := d.NotBefore.Unix()
		nbf = &u
	}

	var exp *int64
	if d.Expiration != nil {
		u := d.Expiration.Unix()
		exp = &u
	}

	model := &tokenPayloadModel{
		Iss:   d.Issuer.String(),
		Aud:   d.Audience.String(),
		Sub:   sub,
		Cmd:   d.Command.String(),
		Pol:   pol,
		Nonce: d.Nonce,
		Meta: metaModel{
			Keys:   metaKeys,
			Values: d.Meta,
		},
		Nbf: nbf,
		Exp: exp,
	}

	return envelope.ToIPLD(privKey, model)
}

// Decode unmarshals the input data using the format specified by the
// provided codec.Decoder into a View.
//
// An error is returned if the conversion fails, or if the resulting
// View is invalid.
func Decode(b []byte, decFn codec.Decoder) (*View, error) {
	node, err := ipld.Decode(b, decFn)
	if err != nil {
		return nil, err
	}
	return FromIPLD(node)
}

// DecodeReader is the same as Decode, but accept an io.Reader.
func DecodeReader(r io.Reader, decFn codec.Decoder) (*View, error) {
	node, err := ipld.DecodeStreaming(r, decFn)
	if err != nil {
		return nil, err
	}
	return FromIPLD(node)
}

// FromDagCbor unmarshals the input data into a View.
//
// An error is returned if the conversion fails, or if the resulting
// View is invalid.
func FromDagCbor(data []byte) (*View, error) {
	return Decode(data, dagcbor.Decode)
}

// FromDagCborReader is the same as FromDagCbor, but accept an io.Reader.
func FromDagCborReader(r io.Reader) (*View, error) {
	return DecodeReader(r, dagcbor.Decode)
}

// FromDagJson unmarshals the input data into a View.
//
// An error is returned if the conversion fails, or if the resulting
// View is invalid.
func FromDagJson(data []byte) (*View, error) {
	return Decode(data, dagjson.Decode)
}

// FromDagJsonReader is the same as FromDagJson, but accept an io.Reader.
func FromDagJsonReader(r io.Reader) (*View, error) {
	return DecodeReader(r, dagjson.Decode)
}

// FromIPLD unwraps a View from the provided IPLD datamodel.Node
//
// An error is returned if the conversion fails, or if the resulting
// View is invalid.
func FromIPLD(node datamodel.Node) (*View, error) {
	tkn, _, err := envelope.FromIPLD[*tokenPayloadModel](node) // TODO add CID to view
	if err != nil {
		return nil, err
	}

	return viewFromModel(*tkn)
}
