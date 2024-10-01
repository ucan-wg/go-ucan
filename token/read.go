package token

import (
	"fmt"
	"io"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"github.com/ipld/go-ipld-prime/datamodel"

	"github.com/ucan-wg/go-ucan/token/delegation"
	"github.com/ucan-wg/go-ucan/token/internal/envelope"
)

// Decode unmarshals the input data using the format specified by the
// provided codec.Decoder into an arbitrary UCAN token.
// An error is returned if the conversion fails, or if the resulting
// Token is invalid.
// Supported and returned types are:
// - delegation.Token
func Decode(b []byte, decFn codec.Decoder) (Token, error) {
	node, err := ipld.Decode(b, decFn)
	if err != nil {
		return nil, err
	}
	return fromIPLD(node)
}

// DecodeReader is the same as Decode, but accept an io.Reader.
func DecodeReader(r io.Reader, decFn codec.Decoder) (Token, error) {
	node, err := ipld.DecodeStreaming(r, decFn)
	if err != nil {
		return nil, err
	}
	return fromIPLD(node)
}

// FromDagCbor unmarshals an arbitrary DagCbor encoded UCAN token.
// An error is returned if the conversion fails, or if the resulting
// Token is invalid.
// Supported and returned types are:
// - delegation.Token
func FromDagCbor(b []byte) (Token, error) {
	return Decode(b, dagcbor.Decode)
}

// FromDagCborReader is the same as FromDagCbor, but accept an io.Reader.
func FromDagCborReader(r io.Reader) (Token, error) {
	return DecodeReader(r, dagcbor.Decode)
}

// FromDagCbor unmarshals an arbitrary DagJson encoded UCAN token.
// An error is returned if the conversion fails, or if the resulting
// Token is invalid.
// Supported and returned types are:
// - delegation.Token
func FromDagJson(b []byte) (Token, error) {
	return Decode(b, dagjson.Decode)
}

// FromDagJsonReader is the same as FromDagJson, but accept an io.Reader.
func FromDagJsonReader(r io.Reader) (Token, error) {
	return DecodeReader(r, dagjson.Decode)
}

func fromIPLD(node datamodel.Node) (Token, error) {
	tag, err := envelope.FindTag(node)
	if err != nil {
		return nil, err
	}

	switch tag {
	case delegation.Tag:
		return delegation.FromIPLD(node)
	default:
		return nil, fmt.Errorf(`unknown tag "%s"`, tag)
	}
}
