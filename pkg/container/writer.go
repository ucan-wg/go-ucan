package container

import (
	"bytes"
	"encoding/base64"
	"io"

	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/fluent/qp"
	"github.com/ipld/go-ipld-prime/node/basicnode"
)

// Writer is a token container writer. It provides a convenient way to aggregate and serialize tokens together.
type Writer map[cid.Cid][]byte

func NewWriter() Writer {
	return make(Writer)
}

// AddSealed includes a "sealed" token (serialized with a ToSealed* function) in the container.
func (ctn Writer) AddSealed(cid cid.Cid, data []byte) {
	ctn[cid] = data
}

const currentContainerVersion = "ctn-v1"

// ToCbor encode the container into a DAG-CBOR binary format.
func (ctn Writer) ToCbor() ([]byte, error) {
	var buf bytes.Buffer
	err := ctn.ToCborWriter(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// ToCborWriter is the same as ToCbor, but with an io.Writer.
func (ctn Writer) ToCborWriter(w io.Writer) error {
	node, err := qp.BuildMap(basicnode.Prototype.Any, 1, func(ma datamodel.MapAssembler) {
		qp.MapEntry(ma, currentContainerVersion, qp.List(int64(len(ctn)), func(la datamodel.ListAssembler) {
			for _, data := range ctn {
				qp.ListEntry(la, qp.Bytes(data))
			}
		}))
	})
	if err != nil {
		return err
	}
	return ipld.EncodeStreaming(w, node, dagcbor.Encode)
}

// ToCborBase64 encode the container into a base64 encoded DAG-CBOR binary format.
func (ctn Writer) ToCborBase64() ([]byte, error) {
	var buf bytes.Buffer
	err := ctn.ToCborBase64Writer(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// ToCborBase64Writer is the same as ToCborBase64, but with an io.Writer.
func (ctn Writer) ToCborBase64Writer(w io.Writer) error {
	w2 := base64.NewEncoder(base64.StdEncoding, w)
	defer w2.Close()
	return ctn.ToCborWriter(w2)
}

// ToCar encode the container into a CAR file.
func (ctn Writer) ToCar() ([]byte, error) {
	var buf bytes.Buffer
	err := ctn.ToCarWriter(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// ToCarWriter is the same as ToCar, but with an io.Writer.
func (ctn Writer) ToCarWriter(w io.Writer) error {
	return writeCar(w, nil, func(yield func(carBlock, error) bool) {
		for c, data := range ctn {
			if !yield(carBlock{c: c, data: data}, nil) {
				return
			}
		}
	})
}

// ToCarBase64 encode the container into a base64 encoded CAR file.
func (ctn Writer) ToCarBase64() ([]byte, error) {
	var buf bytes.Buffer
	err := ctn.ToCarBase64Writer(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// ToCarBase64Writer is the same as ToCarBase64, but with an io.Writer.
func (ctn Writer) ToCarBase64Writer(w io.Writer) error {
	w2 := base64.NewEncoder(base64.StdEncoding, w)
	defer w2.Close()
	return ctn.ToCarWriter(w2)
}
