package token

import (
	"io"

	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime/codec"
	"github.com/libp2p/go-libp2p/core/crypto"

	"github.com/ucan-wg/go-ucan/did"
	"github.com/ucan-wg/go-ucan/pkg/meta"
)

type Token interface {
	Marshaller

	// Issuer returns the did.DID representing the Token's issuer.
	Issuer() did.DID
	// Meta returns the Token's metadata.
	Meta() meta.ReadOnly
}

type Marshaller interface {
	// ToSealed wraps the token in an envelope, generates the signature, encodes
	// the result to DAG-CBOR and calculates the CID of the resulting binary data.
	ToSealed(privKey crypto.PrivKey) ([]byte, cid.Cid, error)
	// ToSealedWriter is the same as ToSealed but accepts an io.Writer.
	ToSealedWriter(w io.Writer, privKey crypto.PrivKey) (cid.Cid, error)
	// Encode marshals a Token to the format specified by the provided codec.Encoder.
	Encode(privKey crypto.PrivKey, encFn codec.Encoder) ([]byte, error)
	// EncodeWriter is the same as Encode, but accepts an io.Writer.
	EncodeWriter(w io.Writer, privKey crypto.PrivKey, encFn codec.Encoder) error
	// ToDagCbor marshals the Token to the DAG-CBOR format.
	ToDagCbor(privKey crypto.PrivKey) ([]byte, error)
	// ToDagCborWriter is the same as ToDagCbor, but it accepts an io.Writer.
	ToDagCborWriter(w io.Writer, privKey crypto.PrivKey) error
	// ToDagJson marshals the Token to the DAG-JSON format.
	ToDagJson(privKey crypto.PrivKey) ([]byte, error)
	// ToDagJsonWriter is the same as ToDagJson, but it accepts an io.Writer.
	ToDagJsonWriter(w io.Writer, privKey crypto.PrivKey) error
}