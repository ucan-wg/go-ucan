package did

import (
	"fmt"
	"strings"

	crypto "github.com/libp2p/go-libp2p/core/crypto"
	mbase "github.com/multiformats/go-multibase"
	"github.com/multiformats/go-multicodec"
	varint "github.com/multiformats/go-varint"
)

const Prefix = "did:"
const KeyPrefix = "did:key:"

const DIDCore = 0x0d1d
const Ed25519 = 0xed
const RSA = uint64(multicodec.RsaPub)

var MethodOffset = varint.UvarintSize(uint64(DIDCore))

//
// [did:key format]: https://w3c-ccg.github.io/did-method-key/
type DID struct {
	key  bool
	code uint64
	str  string
}

// Undef can be used to represent a nil or undefined DID, using DID{}
// directly is also acceptable.
var Undef = DID{}

func (d DID) Defined() bool {
	return d.str != ""
}

func (d DID) Bytes() []byte {
	if !d.Defined() {
		return nil
	}
	return []byte(d.str)
}

func (d DID) Code() uint64 {
	return d.code
}

func (d DID) DID() DID {
	return d
}

func (d DID) Key() bool {
	return d.key
}

func (d DID) PubKey() (crypto.PubKey, error) {
	if !d.key {
		return nil, fmt.Errorf("unsupported did type: %s", d.String())
	}

	unmarshaler, ok := map[multicodec.Code]crypto.PubKeyUnmarshaller{
		multicodec.Ed25519Pub:   crypto.UnmarshalEd25519PublicKey,
		multicodec.RsaPub:       crypto.UnmarshalRsaPublicKey,
		multicodec.Secp256k1Pub: crypto.UnmarshalSecp256k1PublicKey,
		multicodec.Es256:        crypto.UnmarshalECDSAPublicKey,
	}[multicodec.Code(d.code)]
	if !ok {
		return nil, fmt.Errorf("unsupported multicodec: %d", d.code)
	}

	return unmarshaler(d.Bytes()[varint.UvarintSize(d.code):])
}

// String formats the decentralized identity document (DID) as a string.
func (d DID) String() string {
	if d.key {
		key, _ := mbase.Encode(mbase.Base58BTC, []byte(d.str))
		return "did:key:" + key
	}
	return "did:" + d.str[MethodOffset:]
}

func Decode(bytes []byte) (DID, error) {
	code, _, err := varint.FromUvarint(bytes)
	if err != nil {
		return Undef, err
	}
	if code == Ed25519 || code == RSA {
		return DID{str: string(bytes), code: code, key: true}, nil
	} else if code == DIDCore {
		return DID{str: string(bytes)}, nil
	}
	return Undef, fmt.Errorf("unsupported DID encoding: 0x%x", code)
}

func Parse(str string) (DID, error) {
	if !strings.HasPrefix(str, Prefix) {
		return Undef, fmt.Errorf("must start with 'did:'")
	}

	if strings.HasPrefix(str, KeyPrefix) {
		code, bytes, err := mbase.Decode(str[len(KeyPrefix):])
		if err != nil {
			return Undef, err
		}
		if code != mbase.Base58BTC {
			return Undef, fmt.Errorf("not Base58BTC encoded")
		}
		return Decode(bytes)
	}

	buf := make([]byte, MethodOffset)
	varint.PutUvarint(buf, DIDCore)
	suffix, _ := strings.CutPrefix(str, Prefix)
	buf = append(buf, suffix...)
	return DID{str: string(buf), code: DIDCore}, nil
}
