package container

import (
	"compress/gzip"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
)

const containerVersionTag = "ctn-v1"

type header byte

const (
	headerRawBytes             = header(0x40)
	headerBase64StdPadding     = header(0x42)
	headerBase64URL            = header(0x43)
	headerRawBytesGzip         = header(0x4D)
	headerBase64StdPaddingGzip = header(0x4F)
	headerBase64URLGzip        = header(0x50)
)

func (h header) encoder(w io.Writer) *payloadWriter {
	res := &payloadWriter{rawWriter: w, writer: w, header: h}

	switch h {
	case headerBase64StdPadding, headerBase64StdPaddingGzip:
		b64Writer := base64.NewEncoder(base64.StdEncoding, res.writer)
		res.writer = b64Writer
		res.closers = append([]io.Closer{b64Writer}, res.closers...)
	case headerBase64URL, headerBase64URLGzip:
		b64Writer := base64.NewEncoder(base64.RawURLEncoding, res.writer)
		res.writer = b64Writer
		res.closers = append([]io.Closer{b64Writer}, res.closers...)
	}

	switch h {
	case headerRawBytesGzip, headerBase64StdPaddingGzip, headerBase64URLGzip:
		gzipWriter := gzip.NewWriter(res.writer)
		res.writer = gzipWriter
		res.closers = append([]io.Closer{gzipWriter}, res.closers...)
	}

	return res
}

func payloadDecoder(r io.Reader) (io.Reader, error) {
	headerBuf := make([]byte, 1)
	_, err := r.Read(headerBuf)
	if err != nil {
		return nil, err
	}
	h := header(headerBuf[0])

	switch h {
	case headerRawBytes,
		headerBase64StdPadding,
		headerBase64URL,
		headerRawBytesGzip,
		headerBase64StdPaddingGzip,
		headerBase64URLGzip:
	default:
		return nil, fmt.Errorf("unknown container header")
	}

	switch h {
	case headerBase64StdPadding, headerBase64StdPaddingGzip:
		r = base64.NewDecoder(base64.StdEncoding, r)
	case headerBase64URL, headerBase64URLGzip:
		r = base64.NewDecoder(base64.RawURLEncoding, r)
	}

	switch h {
	case headerRawBytesGzip, headerBase64StdPaddingGzip, headerBase64URLGzip:
		gzipReader, err := gzip.NewReader(r)
		if err != nil {
			return nil, err
		}
		r = gzipReader
	}

	return r, nil
}

var _ io.WriteCloser = &payloadWriter{}

// payloadWriter is tasked with two things:
// - prepend the header byte
// - call Close() on all the underlying io.Writer
type payloadWriter struct {
	rawWriter   io.Writer
	writer      io.Writer
	header      header
	headerWrote bool
	closers     []io.Closer
}

func (w *payloadWriter) Write(p []byte) (n int, err error) {
	if !w.headerWrote {
		_, err := w.rawWriter.Write([]byte{byte(w.header)})
		if err != nil {
			return 0, err
		}
		w.headerWrote = true
	}
	return w.writer.Write(p)
}

func (w *payloadWriter) Close() error {
	var errs error
	for _, closer := range w.closers {
		if err := closer.Close(); err != nil {
			errs = errors.Join(errs, err)
		}
	}
	return errs
}
