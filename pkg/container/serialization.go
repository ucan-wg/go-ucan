package container

import (
	"compress/flate"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"

	"github.com/ipfs/go-cid"
	cbor "github.com/ipfs/go-ipld-cbor"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/fluent/qp"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	"github.com/multiformats/go-multihash"
)

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

func (ctn Container) ToCar(w io.Writer) error {
	return writeCar(w, nil, func(yield func(carBlock) bool) {
		for c, bytes := range ctn {
			if !yield(carBlock{c: c, data: bytes}) {
				return
			}
		}
	})
}

func FromCarBase64(r io.Reader) (Container, error) {
	return FromCar(base64.NewDecoder(base64.StdEncoding, r))
}

func (ctn Container) ToCarBase64(w io.Writer) error {
	w2 := base64.NewEncoder(base64.StdEncoding, w)
	defer w2.Close()
	return ctn.ToCar(w2)
}

func FromCarGzip(r io.Reader) (Container, error) {
	r2, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	defer r2.Close()
	return FromCar(r2)
}

func (ctn Container) ToCarGzip(w io.Writer) error {
	w2 := gzip.NewWriter(w)
	defer w2.Close()
	return ctn.ToCar(w2)
}

func FromCarGzipBase64(r io.Reader) (Container, error) {
	return FromCarGzip(base64.NewDecoder(base64.StdEncoding, r))
}

func (ctn Container) ToCarGzipBase64(w io.Writer) error {
	w2 := base64.NewEncoder(base64.StdEncoding, w)
	defer w2.Close()
	return ctn.ToCarGzip(w2)
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
	cidBuilder := cid.V1Builder{Codec: cid.DagCBOR, MhType: multihash.SHA2_256}

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

func FromCbor2(r io.Reader) (Container, error) {
	n, err := ipld.DecodeStreaming(r, dagcbor.Decode)
	if err != nil {
		return nil, err
	}
	if n.Kind() != datamodel.Kind_List {
		return nil, fmt.Errorf("not a list")
	}

	ctn := make(Container, n.Length())
	cidBuilder := cid.V1Builder{Codec: cid.DagCBOR, MhType: multihash.SHA2_256}

	it := n.ListIterator()
	for !it.Done() {
		_, val, err := it.Next()
		if err != nil {
			return nil, err
		}
		bytes, err := val.AsBytes()
		if err != nil {
			return nil, err
		}
		c, err := cidBuilder.Sum(bytes)
		if err != nil {
			return nil, err
		}
		ctn.AddBytes(c, bytes)
	}
	return ctn, nil
}

func (ctn Container) ToCbor2(w io.Writer) error {
	node, err := qp.BuildList(basicnode.Prototype.Any, int64(len(ctn)), func(la datamodel.ListAssembler) {
		for _, bytes := range ctn {
			qp.ListEntry(la, qp.Bytes(bytes))
		}
	})
	if err != nil {
		return err
	}
	return ipld.EncodeStreaming(w, node, dagcbor.Encode)
}

func (ctn Container) ToCbor(w io.Writer) error {
	raw := make([][]byte, 0, len(ctn))
	for _, bytes := range ctn {
		raw = append(raw, bytes)
	}
	return cbor.EncodeWriter(raw, w)
}

func FromCborBase64(r io.Reader) (Container, error) {
	return FromCbor(base64.NewDecoder(base64.StdEncoding, r))
}

func (ctn Container) ToCborBase64(w io.Writer) error {
	w2 := base64.NewEncoder(base64.StdEncoding, w)
	defer w2.Close()
	return ctn.ToCbor(w2)
}

func FromCborGzip(r io.Reader) (Container, error) {
	r2, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	defer r2.Close()
	return FromCbor(r2)
}

func (ctn Container) ToCborGzip(w io.Writer) error {
	w2 := gzip.NewWriter(w)
	defer w2.Close()
	return ctn.ToCbor(w2)
}

func FromCborGzipBase64(r io.Reader) (Container, error) {
	return FromCborGzip(base64.NewDecoder(base64.StdEncoding, r))
}

func (ctn Container) ToCborGzipBase64(w io.Writer) error {
	w2 := base64.NewEncoder(base64.StdEncoding, w)
	defer w2.Close()
	return ctn.ToCborGzip(w2)
}

func FromCborFlate(r io.Reader) (Container, error) {
	r2 := flate.NewReader(r)
	defer r2.Close()
	return FromCbor(r2)
}

func (ctn Container) ToCborFlate(w io.Writer) error {
	w2, err := flate.NewWriter(w, flate.DefaultCompression)
	if err != nil {
		return err
	}
	defer w2.Close()
	return ctn.ToCbor(w2)
}

func FromCborFlateBase64(r io.Reader) (Container, error) {
	return FromCborFlate(base64.NewDecoder(base64.StdEncoding, r))
}

func (ctn Container) ToCborFlateBase64(w io.Writer) error {
	w2 := base64.NewEncoder(base64.StdEncoding, w)
	defer w2.Close()
	return ctn.ToCborFlate(w2)
}
