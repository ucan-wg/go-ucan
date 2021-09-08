// Package didkey implements the did:key method. A DID Method for Static
// Cryptographic Keys. From the w3c draft spec
// https://w3c-ccg.github.io/did-method-key/
package didkey

import (
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"strings"

	"github.com/libp2p/go-libp2p-core/crypto"
	mb "github.com/multiformats/go-multibase"
	varint "github.com/multiformats/go-varint"
)

const (
	// KeyPrefix indicates a decentralized identifier that uses the key method
	KeyPrefix = "did:key"
	// MulticodecKindRSAPubKey rsa-x509-pub https://github.com/multiformats/multicodec/pull/226
	MulticodecKindRSAPubKey = 0x1205
	// MulticodecKindEd25519PubKey ed25519-pub
	MulticodecKindEd25519PubKey = 0xed
)

// ID is a DID:key identifier
type ID struct {
	crypto.PubKey
}

// NewID constructs an Identifier from a public key
func NewID(pub crypto.PubKey) (ID, error) {
	switch pub.Type() {
	case crypto.Ed25519, crypto.RSA:
		return ID{PubKey: pub}, nil
	default:
		return ID{}, fmt.Errorf("unsupported key type: %s", pub.Type())
	}
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

// VerifyKey returns the backing implementation for a public key, one of:
// *rsa.PublicKey, ed25519.PublicKey
func (id ID) VerifyKey() (interface{}, error) {
	rawPubBytes, err := id.PubKey.Raw()
	if err != nil {
		return nil, err
	}
	switch id.PubKey.Type() {
	case crypto.RSA:
		verifyKeyiface, err := x509.ParsePKIXPublicKey(rawPubBytes)
		if err != nil {
			return nil, err
		}
		verifyKey, ok := verifyKeyiface.(*rsa.PublicKey)
		if !ok {
			return nil, fmt.Errorf("public key is not an RSA key. got type: %T", verifyKeyiface)
		}
		return verifyKey, nil
	case crypto.Ed25519:
		return ed25519.PublicKey(rawPubBytes), nil
	default:
		return nil, fmt.Errorf("unrecognized Public Key type: %s", id.PubKey.Type())
	}
}

// Parse turns a string into a key method ID
func Parse(keystr string) (ID, error) {
	var id ID
	if !strings.HasPrefix(keystr, KeyPrefix) {
		return id, fmt.Errorf("decentralized identifier is not a 'key' type")
	}

	keystr = strings.TrimPrefix(keystr, KeyPrefix+":")

	enc, data, err := mb.Decode(keystr)
	if err != nil {
		return id, fmt.Errorf("decoding multibase: %w", err)
	}

	if enc != mb.Base58BTC {
		return id, fmt.Errorf("unexpected multibase encoding: %s", mb.EncodingToStr[enc])
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
