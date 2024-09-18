package container

import (
	"encoding/base64"
	"io"

	"github.com/ipfs/go-cid"
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

func (ctn Container) AddBytes(cid cid.Cid, data []byte) {
	ctn[cid] = data
}

func (ctn Container) GetBytes(cid cid.Cid) ([]byte, bool) {
	b, ok := ctn[cid]
	return b, ok
}
