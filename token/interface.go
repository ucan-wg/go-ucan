package token

import (
	"io"
	"time"

	"github.com/MetaMask/go-did-it/crypto"
	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime/codec"
)

type Token interface {
	Marshaller

	// IsValidNow verifies that the token can be used at the current time, based on expiration or "not before" fields.
	// This does NOT do any other kind of verifications.
	IsValidNow() bool
	// IsValidNow verifies that the token can be used at the given time, based on expiration or "not before" fields.
	// This does NOT do any other kind of verifications.
	IsValidAt(t time.Time) bool
}

type Marshaller interface {
	// ToSealed wraps the token in an envelope, generates the signature, encodes
	// the result to DAG-CBOR and calculates the CID of the resulting binary data.
	ToSealed(privKey crypto.PrivateKeySigningBytes) ([]byte, cid.Cid, error)
	// ToSealedWriter is the same as ToSealed but accepts an io.Writer.
	ToSealedWriter(w io.Writer, privKey crypto.PrivateKeySigningBytes) (cid.Cid, error)
	// Encode marshals a Token to the format specified by the provided codec.Encoder.
	Encode(privKey crypto.PrivateKeySigningBytes, encFn codec.Encoder) ([]byte, error)
	// EncodeWriter is the same as Encode, but accepts an io.Writer.
	EncodeWriter(w io.Writer, privKey crypto.PrivateKeySigningBytes, encFn codec.Encoder) error
	// ToDagCbor marshals the Token to the DAG-CBOR format.
	ToDagCbor(privKey crypto.PrivateKeySigningBytes) ([]byte, error)
	// ToDagCborWriter is the same as ToDagCbor, but it accepts an io.Writer.
	ToDagCborWriter(w io.Writer, privKey crypto.PrivateKeySigningBytes) error
	// ToDagJson marshals the Token to the DAG-JSON format.
	ToDagJson(privKey crypto.PrivateKeySigningBytes) ([]byte, error)
	// ToDagJsonWriter is the same as ToDagJson, but it accepts an io.Writer.
	ToDagJsonWriter(w io.Writer, privKey crypto.PrivateKeySigningBytes) error
}

// Bundle carries together a decoded token with its Cid and raw signed data.
type Bundle struct {
	Cid     cid.Cid
	Decoded Token
	Sealed  []byte
}
