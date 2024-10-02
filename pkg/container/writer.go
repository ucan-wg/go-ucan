package container

import (
	"compress/flate"
	"compress/gzip"
	"encoding/base64"
	"io"

	"github.com/ipfs/go-cid"
	cbor "github.com/ipfs/go-ipld-cbor"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/fluent/qp"
	"github.com/ipld/go-ipld-prime/node/basicnode"
)

type Writer map[cid.Cid][]byte

func NewWriter() Writer {
	return make(Writer)
}

func (ctn Writer) AddSealed(cid cid.Cid, data []byte) {
	ctn[cid] = data
}

func (ctn Writer) ToCar(w io.Writer) error {
	return writeCar(w, nil, func(yield func(carBlock) bool) {
		for c, bytes := range ctn {
			if !yield(carBlock{c: c, data: bytes}) {
				return
			}
		}
	})
}

func (ctn Writer) ToCarBase64(w io.Writer) error {
	w2 := base64.NewEncoder(base64.StdEncoding, w)
	defer w2.Close()
	return ctn.ToCar(w2)
}

func (ctn Writer) ToCarGzip(w io.Writer) error {
	w2 := gzip.NewWriter(w)
	defer w2.Close()
	return ctn.ToCar(w2)
}

func (ctn Writer) ToCarGzipBase64(w io.Writer) error {
	w2 := base64.NewEncoder(base64.StdEncoding, w)
	defer w2.Close()
	return ctn.ToCarGzip(w2)
}

func (ctn Writer) ToCbor2(w io.Writer) error {
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

func (ctn Writer) ToCbor(w io.Writer) error {
	raw := make([][]byte, 0, len(ctn))
	for _, bytes := range ctn {
		raw = append(raw, bytes)
	}
	return cbor.EncodeWriter(raw, w)
}

func (ctn Writer) ToCborBase64(w io.Writer) error {
	w2 := base64.NewEncoder(base64.StdEncoding, w)
	defer w2.Close()
	return ctn.ToCbor(w2)
}

func (ctn Writer) ToCborGzip(w io.Writer) error {
	w2 := gzip.NewWriter(w)
	defer w2.Close()
	return ctn.ToCbor(w2)
}

func (ctn Writer) ToCborGzipBase64(w io.Writer) error {
	w2 := base64.NewEncoder(base64.StdEncoding, w)
	defer w2.Close()
	return ctn.ToCborGzip(w2)
}

func (ctn Writer) ToCborFlate(w io.Writer) error {
	w2, err := flate.NewWriter(w, flate.DefaultCompression)
	if err != nil {
		return err
	}
	defer w2.Close()
	return ctn.ToCbor(w2)
}

func (ctn Writer) ToCborFlateBase64(w io.Writer) error {
	w2 := base64.NewEncoder(base64.StdEncoding, w)
	defer w2.Close()
	return ctn.ToCborFlate(w2)
}
