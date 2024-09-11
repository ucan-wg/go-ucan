package did

import (
	"errors"

	crypto "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/crypto/pb"
	"github.com/multiformats/go-multicodec"
	"github.com/multiformats/go-varint"
)

func FromPrivKey(privKey crypto.PrivKey) (DID, error) {
	return FromPubKey(privKey.GetPublic())
}

func FromPubKey(pubKey crypto.PubKey) (DID, error) {
	code, ok := map[pb.KeyType]multicodec.Code{
		pb.KeyType_Ed25519:   multicodec.Ed25519Pub,
		pb.KeyType_RSA:       multicodec.RsaPub,
		pb.KeyType_Secp256k1: multicodec.Secp256k1Pub,
		pb.KeyType_ECDSA:     multicodec.Es256,
	}[pubKey.Type()]
	if !ok {
		return Undef, errors.New("Blah")
	}

	buf := varint.ToUvarint(uint64(code))

	pubBytes, err := pubKey.Raw()
	if err != nil {
		return Undef, err
	}

	return DID{
		str:  string(append(buf, pubBytes...)),
		code: uint64(code),
		key:  true,
	}, nil
}

func ToPubKey(s string) (crypto.PubKey, error) {
	id, err := Parse(s)
	if err != nil {
		return nil, err
	}

	return id.PubKey()
}
