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

	"github.com/ucan-wg/go-ucan/token"
	"github.com/ucan-wg/go-ucan/token/delegation"
	"github.com/ucan-wg/go-ucan/token/invocation"
)

var ErrNotFound = fmt.Errorf("not found")

type Reader map[cid.Cid]token.Token

func (ctn Reader) GetToken(cid cid.Cid) (token.Token, error) {
	tkn, ok := ctn[cid]
	if !ok {
		return nil, ErrNotFound
	}
	return tkn, nil
}

func (ctn Reader) GetDelegation(cid cid.Cid) (*delegation.Token, error) {
	tkn, err := ctn.GetToken(cid)
	if err != nil {
		return nil, err
	}
	if tkn, ok := tkn.(*delegation.Token); ok {
		return tkn, nil
	}
	return nil, fmt.Errorf("not a delegation token")
}

func (ctn Reader) GetInvocation() (*invocation.Token, error) {
	for _, t := range ctn {
		if inv, ok := t.(*invocation.Token); ok {
			return inv, nil
		}
	}
	return nil, ErrNotFound
}

func FromCar(r io.Reader) (Reader, error) {
	_, it, err := readCar(r)
	if err != nil {
		return nil, err
	}

	ctn := make(Reader)

	for block, err := range it {
		if err != nil {
			return nil, err
		}

		err = ctn.addToken(block.data)
		if err != nil {
			return nil, err
		}
	}

	return ctn, nil
}

func FromCarBase64(r io.Reader) (Reader, error) {
	return FromCar(base64.NewDecoder(base64.StdEncoding, r))
}

func FromCarGzip(r io.Reader) (Reader, error) {
	r2, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	defer r2.Close()
	return FromCar(r2)
}

func FromCarGzipBase64(r io.Reader) (Reader, error) {
	return FromCarGzip(base64.NewDecoder(base64.StdEncoding, r))
}

func FromCbor(r io.Reader) (Reader, error) {
	var raw [][]byte
	err := cbor.DecodeReader(r, &raw)
	if err != nil {
		return nil, err
	}

	ctn := make(Reader, len(raw))
	for _, data := range raw {
		err = ctn.addToken(data)
		if err != nil {
			return nil, err
		}
	}

	return ctn, nil
}

func FromCbor2(r io.Reader) (Reader, error) {
	n, err := ipld.DecodeStreaming(r, dagcbor.Decode)
	if err != nil {
		return nil, err
	}
	if n.Kind() != datamodel.Kind_List {
		return nil, fmt.Errorf("not a list")
	}

	ctn := make(Reader, n.Length())

	it := n.ListIterator()
	for !it.Done() {
		_, val, err := it.Next()
		if err != nil {
			return nil, err
		}
		data, err := val.AsBytes()
		if err != nil {
			return nil, err
		}
		err = ctn.addToken(data)
		if err != nil {
			return nil, err
		}
	}
	return ctn, nil
}

func FromCborBase64(r io.Reader) (Reader, error) {
	return FromCbor(base64.NewDecoder(base64.StdEncoding, r))
}

func FromCborGzip(r io.Reader) (Reader, error) {
	r2, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	defer r2.Close()
	return FromCbor(r2)
}

func FromCborGzipBase64(r io.Reader) (Reader, error) {
	return FromCborGzip(base64.NewDecoder(base64.StdEncoding, r))
}

func FromCborFlate(r io.Reader) (Reader, error) {
	r2 := flate.NewReader(r)
	defer r2.Close()
	return FromCbor(r2)
}

func FromCborFlateBase64(r io.Reader) (Reader, error) {
	return FromCborFlate(base64.NewDecoder(base64.StdEncoding, r))
}

func (ctn Reader) addToken(data []byte) error {
	tkn, c, err := token.FromSealed(data)
	if err != nil {
		return err
	}
	ctn[c] = tkn
	return nil
}
