package container

import (
	"bytes"
	"io"
	"slices"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/cbor"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/fluent/qp"
	"github.com/ipld/go-ipld-prime/node/basicnode"
)

// Writer is a token container writer. It provides a convenient way to aggregate and serialize tokens together.
type Writer map[string]struct{}

func NewWriter() Writer {
	return make(Writer)
}

// AddSealed includes a "sealed" token (serialized with a ToSealed* function) in the container.
func (ctn Writer) AddSealed(data []byte) {
	ctn[string(data)] = struct{}{}
}

// ToBytes encode the container into raw bytes.
func (ctn Writer) ToBytes() ([]byte, error) {
	return ctn.toBytes(headerRawBytes)
}

// ToBytesWriter is the same as ToBytes, but with an io.Writer.
func (ctn Writer) ToBytesWriter(w io.Writer) error {
	return ctn.toWriter(headerRawBytes, w)
}

// ToBytesGzipped encode the container into gzipped bytes.
func (ctn Writer) ToBytesGzipped() ([]byte, error) {
	return ctn.toBytes(headerRawBytesGzip)
}

// ToBytesGzippedWriter is the same as ToBytesGzipped, but with an io.Writer.
func (ctn Writer) ToBytesGzippedWriter(w io.Writer) error {
	return ctn.toWriter(headerRawBytesGzip, w)
}

// ToBase64StdPadding encode the container into a base64 string, with standard encoding and padding.
func (ctn Writer) ToBase64StdPadding() (string, error) {
	return ctn.toString(headerBase64StdPadding)
}

// ToBase64StdPaddingWriter is the same as ToBase64StdPadding, but with an io.Writer.
func (ctn Writer) ToBase64StdPaddingWriter(w io.Writer) error {
	return ctn.toWriter(headerBase64StdPadding, w)
}

// ToBase64StdPaddingGzipped encode the container into a pre-gzipped base64 string, with standard encoding and padding.
func (ctn Writer) ToBase64StdPaddingGzipped() (string, error) {
	return ctn.toString(headerBase64StdPaddingGzip)
}

// ToBase64StdPaddingGzippedWriter is the same as ToBase64StdPaddingGzipped, but with an io.Writer.
func (ctn Writer) ToBase64StdPaddingGzippedWriter(w io.Writer) error {
	return ctn.toWriter(headerBase64StdPaddingGzip, w)
}

// ToBase64URL encode the container into base64 string, with URL-safe encoding and no padding.
func (ctn Writer) ToBase64URL() (string, error) {
	return ctn.toString(headerBase64URL)
}

// ToBase64URLWriter is the same as ToBase64URL, but with an io.Writer.
func (ctn Writer) ToBase64URLWriter(w io.Writer) error {
	return ctn.toWriter(headerBase64URL, w)
}

// ToBase64URLGzipped encode the container into pre-gzipped base64 string, with URL-safe encoding and no padding.
func (ctn Writer) ToBase64URLGzipped() (string, error) {
	return ctn.toString(headerBase64URLGzip)
}

// ToBase64URLGzipWriter is the same as ToBase64URL, but with an io.Writer.
func (ctn Writer) ToBase64URLGzipWriter(w io.Writer) error {
	return ctn.toWriter(headerBase64URLGzip, w)
}

func (ctn Writer) toBytes(header header) ([]byte, error) {
	var buf bytes.Buffer
	err := ctn.toWriter(header, &buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (ctn Writer) toString(header header) (string, error) {
	var buf bytes.Buffer
	err := ctn.toWriter(header, &buf)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (ctn Writer) toWriter(header header, w io.Writer) (err error) {
	encoder := header.encoder(w)

	defer func() {
		err = encoder.Close()
	}()
	node, err := qp.BuildMap(basicnode.Prototype.Any, 1, func(ma datamodel.MapAssembler) {
		qp.MapEntry(ma, containerVersionTag, qp.List(int64(len(ctn)), func(la datamodel.ListAssembler) {
			tokens := make([][]byte, 0, len(ctn))
			for data := range ctn {
				tokens = append(tokens, []byte(data))
			}
			slices.SortFunc(tokens, bytes.Compare)
			for _, data := range tokens {
				qp.ListEntry(la, qp.Bytes(data))
			}
		}))
	})
	if err != nil {
		return err
	}

	return ipld.EncodeStreaming(encoder, node, cbor.Encode)
}

// ToReader convert a container Writer into a Reader.
// Most likely, you only want to use this in tests for convenience.
// This is not optimized and can panic.
func (ctn Writer) ToReader() Reader {
	data, err := ctn.ToBytes()
	if err != nil {
		panic(err)
	}

	reader, err := FromBytes(data)
	if err != nil {
		panic(err)
	}

	return reader
}
