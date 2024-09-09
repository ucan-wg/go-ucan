// Package varsig implements the portion of the [varsig specification]
// that's needed for the UCAN [Envelope].
//
// While the [Envelope] specification has a field that's labelled
// "VarsigHeader", this field is actually the prefix, header and segments
// of the body excluding the signature itself (which is a different field
// in the [Envelope]).
//
// Given that [go-ucan] supports a limited number of public key types,
// and that the signature isn't part of the resulting field, the values
// that are used are constants.  Note that for key types that are fully
// specified in the [did:key], the [VarsigHeader] field isn't technically
// needed and could theoretically conflict with the DID.
//
// Treating these values as constants has no impact when issuing or
// delegating tokens.  When decoding tokens, simply matching the strings
// will allow us to detect errors but won't provide as much detail (e.g.
// we can't indicate that the signature was incorrectly generated from
// a DAG-JSON encoding.)
//
// [varsig specification]: https://github.com/ChainAgnostic/varsig
// [Envelope]:https://github.com/ucan-wg/spec#envelope
// [go-ucan]: https://github.com/ucan-wg/go-ucan
package varsig

import (
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/libp2p/go-libp2p/core/crypto/pb"
	"github.com/multiformats/go-multicodec"
)

const (
	Prefix = 0x34
)

// ErrUnknownHeader is returned when it's not possible to decode the
// provided string into a libp2p public key type.
var ErrUnknownHeader = errors.New("could not decode unknown header")

// ErrUnknownKeyType is returned when value provided is not a valid
// libp2p public key type.
var ErrUnknownKeyType = errors.New("could not encode unsupported key type")

var (
	decMap = headerToKeyType()
	encMap = keyTypeToHeader()
)

// Decode returns either the pb.KeyType associated with the provided Header
// or an error.
//
// Currently, only the four key types supported by the [go-libp2p/core/crypto]
// library are supported.
//
// [go-libp2p/core/crypto]: github.com/libp2p/go-libp2p/core/crypto
func Decode(header []byte) (pb.KeyType, error) {
	keyType, ok := decMap[base64.RawStdEncoding.EncodeToString(header)]
	if !ok {
		return -1, fmt.Errorf("%w: %s", ErrUnknownHeader, header)
	}

	return keyType, nil
}

// Encode returns either the header associated with the provided pb.KeyType
// or an error indicating the header was unknown.
//
// Currently, only the four key types supported by the [go-libp2p/core/crypto]
// library are supported.
//
// [go-libp2p/core/crypto]: github.com/libp2p/go-libp2p/core/crypto
func Encode(keyType pb.KeyType) ([]byte, error) {
	header, ok := encMap[keyType]
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrUnknownKeyType, keyType.String())
	}

	return []byte(header), nil
}

func keyTypeToHeader() map[pb.KeyType][]byte {
	const rsaSigLen = 0x100

	return map[pb.KeyType][]byte{
		pb.KeyType_RSA: header(
			Prefix,
			multicodec.RsaPub,
			multicodec.Sha2_256,
			rsaSigLen,
			multicodec.DagCbor,
		),
		pb.KeyType_Ed25519: header(
			Prefix,
			multicodec.Ed25519Pub,
			multicodec.DagCbor,
		),
		pb.KeyType_Secp256k1: header(
			Prefix,
			multicodec.Secp256k1Pub,
			multicodec.Sha2_256,
			multicodec.DagCbor,
		),
		pb.KeyType_ECDSA: header(
			Prefix,
			multicodec.Es256,
			multicodec.Sha2_256,
			multicodec.DagCbor,
		),
	}
}

func headerToKeyType() map[string]pb.KeyType {
	out := make(map[string]pb.KeyType, len(encMap))

	for keyType, header := range encMap {
		out[base64.RawStdEncoding.EncodeToString(header)] = keyType
	}

	return out
}

func header(vals ...multicodec.Code) []byte {
	var buf []byte

	for _, val := range vals {
		buf = binary.AppendUvarint(buf, uint64(val))
	}

	return buf
}
