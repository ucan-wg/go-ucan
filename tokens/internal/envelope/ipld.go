// Package envelope provides functions that convert between wire-format
// encoding of a [UCAN] token's [Envelope] and the Go type representing
// a verified [TokenPayload].
//
// Encoding functions in this package require a private key as a
// parameter so the VarsigHeader can be set and so that a
// cryptographic signature can be generated.
//
// Decoding functions in this package likewise perform the signature
// verification using a public key extracted from the TokenPayload as
// described by requirement two below.
//
// Types that wish to be marshaled and unmarshaled from the using
// is package have two requirements.
//
//  1. The type must implement the Tokener interface.
//
//  2. The IPLD Representation of the type must include an "iss"
//     field when the TokenPayload is extracted from the Envelope.
//     This field must contain the string representation of a
//     "did:key" so that a public key can be extracted from the
//
// [Envelope]:https://github.com/ucan-wg/spec#envelope
// [TokenPayload]: https://github.com/ucan-wg/spec#envelope
// [UCAN]: https://ucan.xyz
package envelope

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/fluent/qp"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	"github.com/ipld/go-ipld-prime/node/bindnode"
	"github.com/ipld/go-ipld-prime/schema"
	"github.com/libp2p/go-libp2p/core/crypto"

	"github.com/ucan-wg/go-ucan/did"
	"github.com/ucan-wg/go-ucan/tokens/internal/varsig"
)

const (
	VarsigHeaderKey = "h"
	UCANTagPrefix   = "ucan/"
)

// Tokener must be implemented by types that wish to be enclosed in a
// UCAN Envelope (presumbably one of the UCAN token types).
type Tokener interface {
	// Prototype provides the schema representation for an IPLD type so
	// that the incoming datamodel.Kinds can be mapped to the appropriate
	// schema.Kinds.
	Prototype() schema.TypedPrototype

	// Tag returns the expected key denoting the name of the IPLD node
	// that should be processed as the token payload while decoding
	// incoming bytes.
	Tag() string
}

// Decode unmarshals the input data using the format specified by the
// provided codec.Decoder into a Tokener.
//
// An error is returned if the conversion fails, or if the resulting
// Tokener is invalid.
func Decode[T Tokener](b []byte, decFn codec.Decoder) (T, error) {
	node, err := ipld.Decode(b, decFn)
	if err != nil {
		return *new(T), err
	}

	return FromIPLD[T](node)
}

// DecodeReader is the same as Decode, but accept an io.Reader.
func DecodeReader[T Tokener](r io.Reader, decFn codec.Decoder) (T, error) {
	node, err := ipld.DecodeStreaming(r, decFn)
	if err != nil {
		return *new(T), err
	}

	return FromIPLD[T](node)
}

// FromDagCbor unmarshals the input data into a Tokener.
//
// An error is returned if the conversion fails, or if the resulting
// Tokener is invalid.
func FromDagCbor[T Tokener](b []byte) (T, error) {
	undef := *new(T)

	node, err := ipld.Decode(b, dagcbor.Decode)
	if err != nil {
		return undef, err
	}

	tkn, err := fromIPLD[T](node)
	if err != nil {
		return undef, err
	}

	return tkn, nil
}

// FromDagCborReader is the same as FromDagCbor, but accept an io.Reader.
func FromDagCborReader[T Tokener](r io.Reader) (T, error) {
	return DecodeReader[T](r, dagcbor.Decode)
}

// FromDagJson unmarshals the input data into a Tokener.
//
// An error is returned if the conversion fails, or if the resulting
// Tokener is invalid.
func FromDagJson[T Tokener](b []byte) (T, error) {
	return Decode[T](b, dagjson.Decode)
}

// FromDagJsonReader is the same as FromDagJson, but accept an io.Reader.
func FromDagJsonReader[T Tokener](r io.Reader) (T, error) {
	return DecodeReader[T](r, dagjson.Decode)
}

// FromIPLD unwraps a Tokener from the provided IPLD datamodel.Node.
//
// An error is returned if the conversion fails, or if the resulting
// Tokener is invalid.
func FromIPLD[T Tokener](node datamodel.Node) (T, error) {
	undef := *new(T)

	tkn, err := fromIPLD[T](node)
	if err != nil {
		return undef, err
	}

	return tkn, nil
}

func fromIPLD[T Tokener](node datamodel.Node) (T, error) {
	zero := *new(T)

	info, err := inspect(node)
	if err != nil {
		return zero, err
	}

	if info.Tag != zero.Tag() {
		return zero, errors.New("data doesn't match the expected type")
	}

	tokenPayloadNode, err := info.SigPayloadNode.LookupByString(info.Tag)
	if err != nil {
		return zero, err
	}

	// This needs to be done before converting this node to its schema
	// representation (afterwards, the field might be renamed os it's safer
	// to use the wire name).
	issuerNode, err := tokenPayloadNode.LookupByString("iss")
	if err != nil {
		return zero, err
	}

	// Replaces the datamodel.Node in tokenPayloadNode with a
	// schema.TypedNode so that we can cast it to a *token.Token after
	// unwrapping it.
	nb := zero.Prototype().Representation().NewBuilder()

	err = nb.AssignNode(tokenPayloadNode)
	if err != nil {
		return zero, err
	}

	tokenPayloadNode = nb.Build()

	tokenPayload := bindnode.Unwrap(tokenPayloadNode)
	if tokenPayload == nil {
		return zero, errors.New("failed to Unwrap the TokenPayload")
	}

	tkn, ok := tokenPayload.(T)
	if !ok {
		return zero, errors.New("failed to assert the TokenPayload type as *token.Token")
	}

	// Check that the issuer's DID contains a public key with a type that
	// matches the VarsigHeader and then verify the SigPayload.
	issuer, err := issuerNode.AsString()
	if err != nil {
		return zero, err
	}

	issuerDID, err := did.Parse(issuer)
	if err != nil {
		return zero, err
	}

	issuerPubKey, err := issuerDID.PubKey()
	if err != nil {
		return zero, err
	}

	issuerVarsigHeader, err := varsig.Encode(issuerPubKey.Type())
	if err != nil {
		return zero, err
	}

	if string(info.VarsigHeader) != string(issuerVarsigHeader) {
		return zero, errors.New("the VarsigHeader key type doesn't match the issuer's key type")
	}

	data, err := ipld.Encode(info.SigPayloadNode, dagcbor.Encode)
	if err != nil {
		return zero, err
	}

	ok, err = issuerPubKey.Verify(data, info.Signature)
	if err != nil || !ok {
		return zero, errors.New("failed to verify the token's signature")
	}

	return tkn, nil
}

// Encode marshals a Tokener to the format specified by the provided
// codec.Encoder.
func Encode(privKey crypto.PrivKey, token Tokener, encFn codec.Encoder) ([]byte, error) {
	node, err := ToIPLD(privKey, token)
	if err != nil {
		return nil, err
	}

	return ipld.Encode(node, encFn)
}

// EncodeWriter is the same as Encode but outputs to an io.Writer instead
// of encoding into a []byte.
func EncodeWriter(w io.Writer, privKey crypto.PrivKey, token Tokener, encFn codec.Encoder) error {
	node, err := ToIPLD(privKey, token)
	if err != nil {
		return err
	}

	return ipld.EncodeStreaming(w, node, encFn)
}

// ToDagCbor marshals the Tokener to the DAG-CBOR format.
func ToDagCbor(privKey crypto.PrivKey, token Tokener) ([]byte, error) {
	return Encode(privKey, token, dagcbor.Encode)
}

// ToDagCborWriter is the same as ToDagCbor but outputs to an io.Writer
// instead of encoding into a []byte.
func ToDagCborWriter(w io.Writer, privKey crypto.PrivKey, token Tokener) error {
	return EncodeWriter(w, privKey, token, dagcbor.Encode)
}

// ToDagJson marshals the Tokener to the DAG-JSON format.
func ToDagJson(privKey crypto.PrivKey, token Tokener) ([]byte, error) {
	return Encode(privKey, token, dagjson.Encode)
}

// ToDagJsonWriter is the same as ToDagJson but outputs to an io.Writer
// instead of encoding into a []byte.
func ToDagJsonWriter(w io.Writer, privKey crypto.PrivKey, token Tokener) error {
	return EncodeWriter(w, privKey, token, dagjson.Encode)
}

// ToIPLD wraps the Tokener in an IPLD datamodel.Node.
func ToIPLD(privKey crypto.PrivKey, token Tokener) (datamodel.Node, error) {
	tokenPayloadNode := bindnode.Wrap(token, token.Prototype().Type()).Representation()

	varsigHeader, err := varsig.Encode(privKey.Type())
	if err != nil {
		return nil, err
	}

	sigPayloadNode, err := qp.BuildMap(basicnode.Prototype.Any, 2, func(ma datamodel.MapAssembler) {
		qp.MapEntry(ma, VarsigHeaderKey, qp.Bytes(varsigHeader))
		qp.MapEntry(ma, token.Tag(), qp.Node(tokenPayloadNode))
	})

	data, err := ipld.Encode(sigPayloadNode, dagcbor.Encode)
	if err != nil {
		return nil, err
	}

	signature, err := privKey.Sign(data)
	if err != nil {
		return nil, err
	}

	return qp.BuildList(basicnode.Prototype.Any, 2, func(la datamodel.ListAssembler) {
		qp.ListEntry(la, qp.Bytes(signature))
		qp.ListEntry(la, qp.Node(sigPayloadNode))
	})
}

// FindTag inspect the given token IPLD representation and extract the token tag.
func FindTag(node datamodel.Node) (string, error) {
	sigPayloadNode, err := node.LookupByIndex(1)
	if err != nil {
		return "", err
	}

	it := sigPayloadNode.MapIterator()
	i := 0

	for !it.Done() {
		if i >= 2 {
			return "", fmt.Errorf("expected two and only two fields in SigPayload")
		}
		i++

		k, _, err := it.Next()
		if err != nil {
			return "", err
		}

		key, err := k.AsString()
		if err != nil {
			return "", err
		}

		if strings.HasPrefix(key, UCANTagPrefix) {
			return key, nil
		}
	}
	return "", fmt.Errorf("no token tag found")
}

type info struct {
	Tag            string
	Signature      []byte
	SigPayloadNode datamodel.Node
	VarsigHeader   []byte
}

func inspect(node datamodel.Node) (info, error) {
	var res info

	signatureNode, err := node.LookupByIndex(0)
	if err != nil {
		return info{}, err
	}

	res.Signature, err = signatureNode.AsBytes()
	if err != nil {
		return info{}, err
	}

	res.SigPayloadNode, err = node.LookupByIndex(1)
	if err != nil {
		return info{}, err
	}

	it := res.SigPayloadNode.MapIterator()
	foundVarsigHeader := false
	foundTokenPayload := false
	i := 0

	for !it.Done() {
		if i >= 2 {
			return info{}, fmt.Errorf("expected two and only two fields in SigPayload")
		}
		i++

		k, v, err := it.Next()
		if err != nil {
			return info{}, err
		}

		key, err := k.AsString()
		if err != nil {
			return info{}, err
		}

		switch {
		case key == VarsigHeaderKey:
			foundVarsigHeader = true
			res.VarsigHeader, err = v.AsBytes()
			if err != nil {
				return info{}, err
			}
		case strings.HasPrefix(key, UCANTagPrefix):
			foundTokenPayload = true
			res.Tag = key
		default:
			return info{}, fmt.Errorf("unexpected key type %q", key)
		}
	}

	if i != 2 {
		return info{}, fmt.Errorf("expected two and only two fields in SigPayload: %d", i)
	}
	if !foundVarsigHeader {
		return info{}, errors.New("failed to find VarsigHeader field")
	}
	if !foundTokenPayload {
		return info{}, errors.New("failed to find TokenPayload field")
	}

	return res, nil
}
