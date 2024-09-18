package container

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"iter"

	"github.com/ipfs/go-cid"
	cbor "github.com/ipfs/go-ipld-cbor"
)

/*
	Note: below is essentially a re-implementation of the CAR file v1 read and write.
	This exists here for two reasons:
		- go-car's API forces to go through an IPLD getter or through a blockstore API
		- generally, go-car is a very complex and large dependency
*/

// EmptyCid is a "zero" Cid: zero-length "identity" multihash with "raw" codec
// It can be used to have at least one root in a CARv1 file (making it legal), yet
// denote that it can be ignored.
var EmptyCid = cid.MustParse([]byte{01, 55, 00, 00})

type carBlock struct {
	c    cid.Cid
	data []byte
}

// writeCar writes a CARv1 file with no roots, containing the blocks from the iterator.
func writeCar(w io.Writer, roots []cid.Cid, blocks iter.Seq[carBlock]) error {
	if len(roots) == 0 {
		roots = []cid.Cid{EmptyCid}
	}
	h := carHeader{
		Roots:   roots,
		Version: 1,
	}
	hb, err := cbor.DumpObject(h)
	if err != nil {
		return err
	}
	err = ldWrite(w, hb)
	if err != nil {
		return err
	}

	for block := range blocks {
		err = ldWrite(w, block.c.Bytes(), block.data)
		if err != nil {
			return err
		}
	}
	return nil
}

// readCar reads a CARv1 file from the reader, and return a block iterator.
// Roots are ignored.
func readCar(r io.Reader) (roots []cid.Cid, blocks iter.Seq2[carBlock, error], err error) {
	br := bufio.NewReader(r)

	hb, err := ldRead(br)
	if err != nil {
		return nil, nil, err
	}
	var h carHeader
	if err := cbor.DecodeInto(hb, &h); err != nil {
		return nil, nil, fmt.Errorf("invalid header: %v", err)
	}

	if h.Version != 1 {
		return nil, nil, fmt.Errorf("invalid car version: %d", h.Version)
	}

	return h.Roots, func(yield func(block carBlock, err error) bool) {
		for {
			block, err := readBlock(br)
			if err == io.EOF {
				return
			}
			if err != nil {
				if !yield(carBlock{}, err) {
					return
				}
			}
			if !yield(block, nil) {
				return
			}
		}
	}, nil
}

// readBlock reads a section from the reader and decode a (cid+data) block.
func readBlock(r *bufio.Reader) (carBlock, error) {
	raw, err := ldRead(r)
	if err != nil {
		return carBlock{}, err
	}

	n, c, err := cid.CidFromReader(bytes.NewReader(raw))
	if err != nil {
		return carBlock{}, err
	}
	data := raw[n:]

	// integrity check
	hashed, err := c.Prefix().Sum(data)
	if err != nil {
		return carBlock{}, err
	}

	if !hashed.Equals(c) {
		return carBlock{}, fmt.Errorf("mismatch in content integrity, name: %s, data: %s", c, hashed)
	}

	return carBlock{c: c, data: data}, nil
}

// maxAllowedSectionSize dictates the maximum number of bytes that a CARv1 header
// or section is allowed to occupy without causing a decode to error.
// This cannot be supplied as an option, only adjusted as a global. You should
// use v2#NewReader instead since it allows for options to be passed in.
var maxAllowedSectionSize uint = 32 << 20 // 32MiB

// ldRead performs a length-delimited read of a section from the reader.
// A section is composed of an uint length followed by the data.
func ldRead(r *bufio.Reader) ([]byte, error) {
	if _, err := r.Peek(1); err != nil { // no more blocks, likely clean io.EOF
		return nil, err
	}

	l, err := binary.ReadUvarint(r)
	if err != nil {
		if err == io.EOF {
			return nil, io.ErrUnexpectedEOF // don't silently pretend this is a clean EOF
		}
		return nil, err
	}

	if l > uint64(maxAllowedSectionSize) { // Don't OOM
		return nil, fmt.Errorf("malformed car; header is bigger than MaxAllowedSectionSize")
	}

	buf := make([]byte, l)
	if _, err := io.ReadFull(r, buf); err != nil {
		return nil, err
	}

	return buf, nil
}

// ldWrite performs a length-delimited write of a section on the writer.
// A section is composed of an uint length followed by the data.
func ldWrite(w io.Writer, d ...[]byte) error {
	var sum uint64
	for _, s := range d {
		sum += uint64(len(s))
	}

	buf := make([]byte, 8)
	n := binary.PutUvarint(buf, sum)
	_, err := w.Write(buf[:n])
	if err != nil {
		return err
	}

	for _, s := range d {
		_, err = w.Write(s)
		if err != nil {
			return err
		}
	}

	return nil
}

type carHeader struct {
	Roots   []cid.Cid
	Version uint64
}

func init() {
	cbor.RegisterCborType(carHeader{})
}
