package token

import (
	"fmt"
	"io"

	"github.com/MetaMask/go-did-it"
	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"github.com/ipld/go-ipld-prime/datamodel"

	"github.com/ucan-wg/go-ucan/token/delegation"
	"github.com/ucan-wg/go-ucan/token/internal/envelope"
	"github.com/ucan-wg/go-ucan/token/invocation"
)

// FromSealed decodes an arbitrary token type from the binary data,
// verifies that the envelope's signature is correct based on the public
// key taken from the issuer (iss) field and calculates the CID of the
// incoming data.
// Supported and returned types are:
// - delegation.Token
// - invocation.Token
func FromSealed(data []byte, resolvOpts ...did.ResolutionOption) (Token, cid.Cid, error) {
	tkn, err := FromDagCbor(data, resolvOpts...)
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
func FromSealedReader(r io.Reader, resolvOpts ...did.ResolutionOption) (Token, cid.Cid, error) {
	cidReader := envelope.NewCIDReader(r)

	tkn, err := FromDagCborReader(cidReader, resolvOpts...)
	if err != nil {
		return nil, cid.Undef, err
	}

	id, err := cidReader.CID()
	if err != nil {
		return nil, cid.Undef, err
	}

	return tkn, id, nil
}

// Decode unmarshals the input data using the format specified by the
// provided codec.Decoder into an arbitrary UCAN token.
// An error is returned if the conversion fails, or if the resulting
// Token is invalid.
// Supported and returned types are:
// - delegation.Token
// - invocation.Token
func Decode(b []byte, decFn codec.Decoder, resolvOpts ...did.ResolutionOption) (Token, error) {
	node, err := ipld.Decode(b, decFn)
	if err != nil {
		return nil, err
	}
	return fromIPLD(node, resolvOpts...)
}

// DecodeReader is the same as Decode, but accept an io.Reader.
func DecodeReader(r io.Reader, decFn codec.Decoder, resolvOpts ...did.ResolutionOption) (Token, error) {
	node, err := ipld.DecodeStreaming(r, decFn)
	if err != nil {
		return nil, err
	}
	return fromIPLD(node, resolvOpts...)
}

// FromDagCbor unmarshals an arbitrary DagCbor encoded UCAN token.
// An error is returned if the conversion fails, or if the resulting
// Token is invalid.
// Supported and returned types are:
// - delegation.Token
// - invocation.Token
func FromDagCbor(b []byte, resolvOpts ...did.ResolutionOption) (Token, error) {
	return Decode(b, dagcbor.Decode, resolvOpts...)
}

// FromDagCborReader is the same as FromDagCbor, but accept an io.Reader.
func FromDagCborReader(r io.Reader, resolvOpts ...did.ResolutionOption) (Token, error) {
	return DecodeReader(r, dagcbor.Decode, resolvOpts...)
}

// FromDagJson unmarshals an arbitrary DagJson encoded UCAN token.
// An error is returned if the conversion fails, or if the resulting
// Token is invalid.
// Supported and returned types are:
// - delegation.Token
// - invocation.Token
func FromDagJson(b []byte, resolvOpts ...did.ResolutionOption) (Token, error) {
	return Decode(b, dagjson.Decode, resolvOpts...)
}

// FromDagJsonReader is the same as FromDagJson, but accept an io.Reader.
func FromDagJsonReader(r io.Reader, resolvOpts ...did.ResolutionOption) (Token, error) {
	return DecodeReader(r, dagjson.Decode, resolvOpts...)
}

func fromIPLD(node datamodel.Node, resolvOpts ...did.ResolutionOption) (Token, error) {
	tag, err := envelope.FindTag(node)
	if err != nil {
		return nil, err
	}

	switch tag {
	case delegation.Tag:
		return delegation.FromIPLD(node, resolvOpts...)
	case invocation.Tag:
		return invocation.FromIPLD(node, resolvOpts...)
	default:
		return nil, fmt.Errorf(`unknown tag "%s"`, tag)
	}
}
