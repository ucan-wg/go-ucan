// Package didkey implements the did:key method. A DID Method for Static
// Cryptographic Keys. From the w3c draft spec
// https://w3c-ccg.github.io/did-method-key/
package didkey

import (
	"fmt"
	"strings"

	"github.com/libp2p/go-libp2p-core/crypto"
	mb "github.com/multiformats/go-multibase"
	varint "github.com/multiformats/go-varint"
)

const (
	// KeyPrefix indicates a decentralized identifier that uses the key method
	KeyPrefix = "did:key"
	// MulticodecKindRSAPubKey IS NOT A REAL MULTICODEC PREFIX.
	// pulled from: https://github.com/w3c-ccg/lds-ed25519-2018/issues/3 because
	// it's only slighly better than picking a random available byte prefix
	MulticodecKindRSAPubKey = 0x5d
	// MulticodecKindEd25519PubKey ed25519-pub
	MulticodecKindEd25519PubKey = 0xed
)

// ID is a DID:key identifier
type ID struct {
	crypto.PubKey
}

// MulticodecType indicates the type for this multicodec
func (id ID) MulticodecType() uint64 {
	switch id.Type() {
	case crypto.RSA:
		return MulticodecKindRSAPubKey
	case crypto.Ed25519:
		return MulticodecKindEd25519PubKey
	default:
		panic("unexpected crypto type")
	}
}

// String returns this did:key formatted as a string
func (id ID) String() string {
	raw, err := id.Raw()
	if err != nil {
		return ""
	}

	t := id.MulticodecType()
	size := varint.UvarintSize(t)
	data := make([]byte, size+len(raw))
	n := varint.PutUvarint(data, t)
	copy(data[n:], raw)

	b58BKeyStr, err := mb.Encode(mb.Base58BTC, data)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%s:%s", KeyPrefix, b58BKeyStr)
}

// Parse turns a string into a key method ID
func Parse(keystr string) (ID, error) {
	var id ID
	if !strings.HasPrefix(keystr, KeyPrefix) {
		return id, fmt.Errorf("decentralized identifier is not a 'key' type")
	}

	keystr = strings.TrimPrefix(keystr, KeyPrefix+":")

	_, data, err := mb.Decode(keystr)
	if err != nil {
		return id, err
	}

	keyType, n, err := varint.FromUvarint(data)
	if err != nil {
		return id, err
	}

	switch keyType {
	case MulticodecKindRSAPubKey:
		pub, err := crypto.UnmarshalRsaPublicKey(data[n:])
		if err != nil {
			return id, err
		}
		return ID{pub}, nil
	case MulticodecKindEd25519PubKey:
		pub, err := crypto.UnmarshalEd25519PublicKey(data[n:])
		if err != nil {
			return id, err
		}
		return ID{pub}, nil
	}

	return id, fmt.Errorf("unrecognized key type multicodec prefix: %x", data[0])
}
