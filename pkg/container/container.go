package container

import (
	"encoding/base64"
	"io"

	"github.com/ipfs/go-cid"
	cbor "github.com/ipfs/go-ipld-cbor"
	"github.com/multiformats/go-multihash"
)

// TODO: should the invocation being set as root in the car file?

type Container map[cid.Cid][]byte

func New() Container {
	return make(Container)
}

func FromCar(r io.Reader) (Container, error) {
	_, it, err := readCar(r)
	if err != nil {
		return nil, err
	}

	c := New()

	for block, err := range it {
		if err != nil {
			return nil, err
		}
		c[block.c] = block.data
	}

	return c, nil
}

func FromCarBase64(r io.Reader) (Container, error) {
	return FromCar(base64.NewDecoder(base64.StdEncoding, r))
}

func FromCbor(r io.Reader) (Container, error) {
	var raw [][]byte
	err := cbor.DecodeReader(r, &raw)
	if err != nil {
		return nil, err
	}

	// TODO: the CID computation will likely be handled in the envelope
	// TODO: the envelope should likely be able to deserialize arbitrary types based on the tag value
	// TODO: the container should likely expose the decoded token, and have search methods (simple, but also DAG reconstruction, graph path search)
	var cidBuilder = cid.V1Builder{Codec: cid.DagCBOR, MhType: multihash.SHA2_256}

	ctn := make(Container, len(raw))
	for _, bytes := range raw {
		c, err := cidBuilder.Sum(bytes)
		if err != nil {
			return nil, err
		}
		ctn[c] = bytes
	}

	return ctn, nil
}

func FromCborBase64(r io.Reader) (Container, error) {
	return FromCbor(base64.NewDecoder(base64.StdEncoding, r))
}

func (ctn Container) AddBytes(cid cid.Cid, data []byte) {
	ctn[cid] = data
}

func (ctn Container) GetBytes(cid cid.Cid) ([]byte, bool) {
	b, ok := ctn[cid]
	return b, ok
}

func (ctn Container) ToCar(w io.Writer) error {
	return writeCar(w, nil, func(yield func(carBlock) bool) {
		for c, bytes := range ctn {
			if !yield(carBlock{c: c, data: bytes}) {
				return
			}
		}
	})
}

func (ctn Container) ToCarBase64(w io.Writer) error {
	w2 := base64.NewEncoder(base64.StdEncoding, w)
	defer w2.Close()
	return ctn.ToCar(w2)
}

func (ctn Container) ToCbor(w io.Writer) error {
	raw := make([][]byte, 0, len(ctn))
	for _, bytes := range ctn {
		raw = append(raw, bytes)
	}
	return cbor.EncodeWriter(raw, w)
}

func (ctn Container) ToCborBase64(w io.Writer) error {
	w2 := base64.NewEncoder(base64.StdEncoding, w)
	defer w2.Close()
	return ctn.ToCbor(w2)
}
