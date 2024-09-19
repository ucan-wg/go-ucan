package envelope

import (
	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/multiformats/go-multibase"
	"github.com/multiformats/go-multicodec"
	"github.com/multiformats/go-multihash"
)

var b58BTCEnc = multibase.MustNewEncoder(multibase.Base58BTC)

// CIDToBase56BTC is a utility method to convert a CIDv1 to the canonical
// string representation used by UCAN.
func CIDToBase58BTC(id cid.Cid) string {
	return id.Encode(b58BTCEnc)
}

// CID returns the UCAN content identifier a Tokener.
func CID(privKey crypto.PrivKey, token Tokener) (cid.Cid, error) {
	data, err := ToDagCbor(privKey, token)
	if err != nil {
		return cid.Undef, err
	}

	return CIDFromBytes(data)
}

// CIDFromBytes returns the UCAN content identifier for an arbitrary slice
// of bytes.
func CIDFromBytes(b []byte) (cid.Cid, error) {
	return cid.V1Builder{
		Codec:    uint64(multicodec.DagCbor),
		MhType:   multihash.SHA2_256,
		MhLength: 0,
	}.Sum(b)
}

// CIDFromIPLD returns the UCAN content identifier for an ipld.Node.
func CIDFromIPLD(node ipld.Node) (cid.Cid, error) {
	data, err := ipld.Encode(node, dagcbor.Encode)
	if err != nil {
		return cid.Undef, nil
	}

	return CIDFromBytes(data)
}
