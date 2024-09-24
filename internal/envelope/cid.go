package envelope

import (
	"crypto/sha256"
	"hash"
	"io"

	"github.com/ipfs/go-cid"
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

var _ io.Reader = (*CIDReader)(nil)

// CIDReader wraps an io.Reader and includes a hash.Hash that is updated
// as data is read from the child io.Reader.
type CIDReader struct {
	hash hash.Hash
	r    io.Reader
	err  error
}

// NewCIDReader initializes a hash.Hash to calculate the CID's hash and
// and returns a wrapped io.Reader.
func NewCIDReader(r io.Reader) *CIDReader {
	hash := sha256.New()
	hash.Reset()

	return &CIDReader{
		hash: hash,
		r:    r,
	}
}

// CID returns the UCAN-formatted cid.Cid created from the hash calculated
// as bytes are read from the inner io.Reader.
func (r *CIDReader) CID() (cid.Cid, error) {
	if r.err != nil {
		return cid.Undef, r.err // TODO: Wrap to say it's an error during streaming?
	}

	return cidFromHash(r.hash)
}

// Read implements io.Reader.
func (r *CIDReader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	if err != nil && err != io.EOF {
		r.err = err

		return
	}

	_, _ = r.hash.Write(p[:n])

	return
}

var _ io.Writer = (*CIDWriter)(nil)

type CIDWriter struct {
	hash hash.Hash
	w    io.Writer
	err  error
}

func NewCIDWriter(w io.Writer) *CIDWriter {
	hash := sha256.New()
	hash.Reset()

	return &CIDWriter{
		hash: hash,
		w:    w,
	}
}

func (w *CIDWriter) CID() (cid.Cid, error) {
	return cidFromHash(w.hash)
}

func (w *CIDWriter) Write(p []byte) (n int, err error) {
	if _, err = w.hash.Write(p); err != nil {
		w.err = err

		return
	}

	return w.w.Write(p)
}

func cidFromHash(hash hash.Hash) (cid.Cid, error) {
	mh, err := multihash.Encode(hash.Sum(nil), multihash.SHA2_256)
	if err != nil {
		return cid.Undef, err
	}

	return cid.NewCidV1(uint64(multicodec.DagCbor), mh), nil
}
