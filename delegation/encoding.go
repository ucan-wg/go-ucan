package delegation

import (
	"fmt"
	"io"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"github.com/ipld/go-ipld-prime/datamodel"

	"github.com/ucan-wg/go-ucan/internal/envelope"
)

// Encode marshals a Delegation to the format specified by the provided
// codec.Encoder.
func (d *Delegation) Encode(encFn codec.Encoder) ([]byte, error) {
	node, err := d.ToIPLD()
	if err != nil {
		return nil, err
	}

	return ipld.Encode(node, encFn)
}

// ToDagCbor marshals the Delegation to the DAG-CBOR format.
func (d *Delegation) ToDagCbor() ([]byte, error) {
	return d.Encode(dagcbor.Encode)
}

// ToDagJson marshals the Delegation to the DAG-JSON format.
func (d *Delegation) ToDagJson() ([]byte, error) {
	return d.Encode(dagjson.Encode)
}

// ToIPLD wraps the Delegation in an IPLD datamodel.Node.
func (d *Delegation) ToIPLD() (datamodel.Node, error) {
	return d.envel.Wrap()
}

// Decode unmarshals the input data using the format specified by the
// provided codec.Decoder into a Delegation.
//
// An error is returned if the conversion fails, or if the resulting
// Delegation is invalid.
func Decode(b []byte, decFn codec.Decoder) (*Delegation, error) {
	node, err := ipld.Decode(b, decFn)
	if err != nil {
		return nil, err
	}
	return FromIPLD(node)
}

// DecodeReader is the same as Decode, but accept an io.Reader.
func DecodeReader(r io.Reader, decFn codec.Decoder) (*Delegation, error) {
	node, err := ipld.DecodeStreaming(r, decFn)
	if err != nil {
		return nil, err
	}
	return FromIPLD(node)
}

// FromDagCbor unmarshals the input data into a Delegation.
//
// An error is returned if the conversion fails, or if the resulting
// Delegation is invalid.
func FromDagCbor(data []byte) (*Delegation, error) {
	return Decode(data, dagcbor.Decode)
}

// FromDagCborReader is the same as FromDagCbor, but accept an io.Reader.
func FromDagCborReader(r io.Reader) (*Delegation, error) {
	return DecodeReader(r, dagcbor.Decode)
}

// FromDagJson unmarshals the input data into a Delegation.
//
// An error is returned if the conversion fails, or if the resulting
// Delegation is invalid.
func FromDagJson(data []byte) (*Delegation, error) {
	return Decode(data, dagjson.Decode)
}

// FromDagJsonReader is the same as FromDagJson, but accept an io.Reader.
func FromDagJsonReader(r io.Reader) (*Delegation, error) {
	return DecodeReader(r, dagjson.Decode)
}

// FromIPLD unwraps a Delegation from the provided IPLD datamodel.Node
//
// An error is returned if the conversion fails, or if the resulting
// Delegation is invalid.
func FromIPLD(node datamodel.Node) (*Delegation, error) {
	envel, err := envelope.Unwrap(node)
	if err != nil {
		return nil, err
	}

	if envel.Tag() != Tag {
		return nil, fmt.Errorf("wrong tag for TokenPayload: received %s but expected %s", envel.Tag(), Tag)
	}

	dlg := &Delegation{
		envel: envel,
	}

	if err := dlg.Validate(); err != nil {
		return nil, err
	}

	return dlg, nil
}
