package container

import (
	"encoding/base64"
	"io"

	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/fluent/qp"
	"github.com/ipld/go-ipld-prime/node/basicnode"
)

// TODO: should we have a multibase to wrap the cbor? but there is no reader/write in go-multibase :-(

const currentContainerVersion = "ctn-v1"

// Writer is a token container writer. It provides a convenient way to aggregate and serialize tokens together.
type Writer map[cid.Cid][]byte

func NewWriter() Writer {
	return make(Writer)
}

// AddSealed includes a "sealed" token (serialized with a ToSealed* function) in the container.
func (ctn Writer) AddSealed(cid cid.Cid, data []byte) {
	ctn[cid] = data
}

func (ctn Writer) ToCar(w io.Writer) error {
	return writeCar(w, nil, func(yield func(carBlock, error) bool) {
		for c, bytes := range ctn {
			if !yield(carBlock{c: c, data: bytes}, nil) {
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

func (ctn Writer) ToCbor(w io.Writer) error {
	node, err := qp.BuildMap(basicnode.Prototype.Any, 1, func(ma datamodel.MapAssembler) {
		qp.MapEntry(ma, currentContainerVersion, qp.List(int64(len(ctn)), func(la datamodel.ListAssembler) {
			for _, bytes := range ctn {
				qp.ListEntry(la, qp.Bytes(bytes))
			}
		}))
	})
	if err != nil {
		return err
	}
	return ipld.EncodeStreaming(w, node, dagcbor.Encode)
}

func (ctn Writer) ToCborBase64(w io.Writer) error {
	w2 := base64.NewEncoder(base64.StdEncoding, w)
	defer w2.Close()
	return ctn.ToCbor(w2)
}