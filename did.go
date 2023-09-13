package did

import (
	"fmt"
	"strings"

	mbase "github.com/multiformats/go-multibase"
	varint "github.com/multiformats/go-varint"
)

const Prefix = "did:"
const KeyPrefix = "did:key:"

const Ed25519 = 0xed

type DID struct {
	str string
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

func (d DID) DID() DID {
	return d
}

func (d DID) String() string {
	key, _ := mbase.Encode(mbase.Base58BTC, []byte(d.str))
	return "did:key:" + key
}

func Decode(bytes []byte) (DID, error) {
	code, _, err := varint.FromUvarint(bytes)
	if err != nil {
		return Undef, err
	}
	if code == Ed25519 {
		return DID{str: string(bytes)}, nil
	}
	return Undef, fmt.Errorf("decoding DID: unsupported DID encoding: 0x%x", code)
}

func Parse(str string) (DID, error) {
	if !strings.HasPrefix(str, Prefix) {
		return Undef, fmt.Errorf("parsing DID: must start with 'did:'")
	}

	if strings.HasPrefix(str, KeyPrefix) {
		code, bytes, err := mbase.Decode(str[len(KeyPrefix):])
		if err != nil {
			return Undef, err
		}
		if code != mbase.Base58BTC {
			return Undef, fmt.Errorf("parsing DID key: not Base58BTC encoded")
		}
		return Decode(bytes)
	}

	return Undef, fmt.Errorf("parsing DID: unsupported method")
}
