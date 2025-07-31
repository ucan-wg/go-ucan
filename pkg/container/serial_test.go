package container

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/MetaMask/go-did-it"
	"github.com/MetaMask/go-did-it/controller/did-key"
	"github.com/MetaMask/go-did-it/crypto/ed25519"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"

	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/pkg/policy"
	"github.com/ucan-wg/go-ucan/pkg/policy/literal"
	"github.com/ucan-wg/go-ucan/token/delegation"
)

func TestContainerRoundTrip(t *testing.T) {
	for _, tc := range []struct {
		name           string
		expectedHeader header
		writer         any
	}{
		{"Bytes", headerRawBytes, Writer.ToBytes},
		{"BytesWriter", headerRawBytes, Writer.ToBytesWriter},
		{"BytesGzipped", headerRawBytesGzip, Writer.ToBytesGzipped},
		{"BytesGzippedWriter", headerRawBytesGzip, Writer.ToBytesGzippedWriter},
		{"Base64StdPadding", headerBase64StdPadding, Writer.ToBase64StdPadding},
		{"Base64StdPaddingWriter", headerBase64StdPadding, Writer.ToBase64StdPaddingWriter},
		{"Base64StdPaddingGzipped", headerBase64StdPaddingGzip, Writer.ToBase64StdPaddingGzipped},
		{"Base64StdPaddingGzippedWriter", headerBase64StdPaddingGzip, Writer.ToBase64StdPaddingGzippedWriter},
		{"Base64URL", headerBase64URL, Writer.ToBase64URL},
		{"Base64URLWriter", headerBase64URL, Writer.ToBase64URLWriter},
		{"Base64URLGzipped", headerBase64URLGzip, Writer.ToBase64URLGzipped},
		{"Base64URLGzipWriter", headerBase64URLGzip, Writer.ToBase64URLGzipWriter},
	} {
		t.Run(tc.name, func(t *testing.T) {
			tokens := make(map[cid.Cid]*delegation.Token)
			var dataSize int

			writer := NewWriter()

			for i := 0; i < 10; i++ {
				dlg, c, data := randToken()
				writer.AddSealed(data)
				tokens[c] = dlg
				dataSize += len(data)
			}

			var reader Reader
			var serialLen int

			switch fn := tc.writer.(type) {
			case func(ctn Writer, w io.Writer) error:
				buf := bytes.NewBuffer(nil)
				err := fn(writer, buf)
				require.NoError(t, err)
				serialLen = buf.Len()

				h, err := buf.ReadByte()
				require.NoError(t, err)
				require.Equal(t, byte(tc.expectedHeader), h)
				err = buf.UnreadByte()
				require.NoError(t, err)

				reader, err = FromReader(bytes.NewReader(buf.Bytes()))
				require.NoError(t, err)

			case func(ctn Writer) ([]byte, error):
				b, err := fn(writer)
				require.NoError(t, err)
				serialLen = len(b)

				require.Equal(t, byte(tc.expectedHeader), b[0])

				reader, err = FromBytes(b)
				require.NoError(t, err)

			case func(ctn Writer) (string, error):
				s, err := fn(writer)
				require.NoError(t, err)
				serialLen = len(s)

				require.Equal(t, byte(tc.expectedHeader), s[0])

				reader, err = FromString(s)
				require.NoError(t, err)
			}

			t.Logf("data size %d, container size %d, overhead: %d%%, %d bytes",
				dataSize, serialLen, int(float32(serialLen-dataSize)/float32(dataSize)*100.0), serialLen-dataSize)

			for c, dlg := range tokens {
				tknRead, err := reader.GetToken(c)
				require.NoError(t, err)

				// require.Equal fails as time.Time holds a wall time that is going to be
				// different, even if it represents the same event.
				// We need to do the following instead.

				dlgRead := tknRead.(*delegation.Token)
				require.Equal(t, dlg.Issuer(), dlgRead.Issuer())
				require.Equal(t, dlg.Audience(), dlgRead.Audience())
				require.Equal(t, dlg.Subject(), dlgRead.Subject())
				require.Equal(t, dlg.Command(), dlgRead.Command())
				require.Equal(t, dlg.Policy(), dlgRead.Policy())
				require.Equal(t, dlg.Nonce(), dlgRead.Nonce())
				require.True(t, dlg.Meta().Equals(dlgRead.Meta()))
				if dlg.NotBefore() != nil {
					// within 1s as the original value gets truncated to seconds when serialized
					require.WithinDuration(t, *dlg.NotBefore(), *dlgRead.NotBefore(), time.Second)
				}
				if dlg.Expiration() != nil {
					// within 1s as the original value gets truncated to seconds when serialized
					require.WithinDuration(t, *dlg.Expiration(), *dlgRead.Expiration(), time.Second)
				}
			}
		})
	}
}

func BenchmarkContainerSerialisation(b *testing.B) {
	var duration strings.Builder
	var allocByte strings.Builder
	var allocCount strings.Builder

	for _, builder := range []strings.Builder{duration, allocByte, allocCount} {
		builder.WriteString("car\tcarBase64\tcarGzip\tcarGzipBase64\tcbor\tcborBase64\tcborGzip\tcborGzipBase64\tcborFlate\tcborFlateBase64\n")
	}

	for _, tc := range []struct {
		name   string
		writer func(ctn Writer, w io.Writer) error
		reader func(io.Reader) (Reader, error)
	}{
		{"Bytes", Writer.ToBytesWriter, FromReader},
		{"BytesGzipped", Writer.ToBytesGzippedWriter, FromReader},
		{"Base64StdPadding", Writer.ToBase64StdPaddingWriter, FromReader},
		{"Base64StdPaddingGzipped", Writer.ToBase64StdPaddingGzippedWriter, FromReader},
		{"Base64URL", Writer.ToBase64URLWriter, FromReader},
		{"Base64URLGzip", Writer.ToBase64URLGzipWriter, FromReader},
	} {
		writer := NewWriter()

		for i := 0; i < 10; i++ {
			_, _, data := randToken()
			writer.AddSealed(data)
		}

		buf := bytes.NewBuffer(nil)
		_ = tc.writer(writer, buf)

		b.Run(tc.name+"_write", func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				buf := bytes.NewBuffer(nil)
				_ = tc.writer(writer, buf)
			}
		})

		b.Run(tc.name+"_read", func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_, _ = tc.reader(bytes.NewReader(buf.Bytes()))
			}
		})
	}
}

func randDID() (ed25519.PrivateKey, did.DID) {
	_, privKey, err := ed25519.GenerateKeyPair()
	if err != nil {
		panic(err)
	}
	d := didkeyctl.FromPrivateKey(privKey)
	return privKey, d
}

func randomString(length int) string {
	b := make([]byte, length/2+1)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x", b)[0:length]
}

func randToken() (*delegation.Token, cid.Cid, []byte) {
	priv, iss := randDID()
	_, aud := randDID()
	cmd := command.New("foo", "bar")
	pol := policy.MustConstruct(
		policy.All(".[]",
			policy.GreaterThan(".value", literal.Int(2)),
		),
	)

	opts := []delegation.Option{
		delegation.WithExpiration(time.Now().Add(time.Hour)),
	}
	for i := 0; i < 3; i++ {
		opts = append(opts, delegation.WithMeta(randomString(8), randomString(10)))
	}

	t, err := delegation.Root(iss, aud, cmd, pol, opts...)
	if err != nil {
		panic(err)
	}
	b, c, err := t.ToSealed(priv)
	if err != nil {
		panic(err)
	}
	return t, c, b
}

func FuzzContainerRead(f *testing.F) {
	// Generate a corpus
	for tokenCount := 0; tokenCount < 10; tokenCount++ {
		writer := NewWriter()
		for i := 0; i < tokenCount; i++ {
			_, _, data := randToken()
			writer.AddSealed(data)
		}
		data, err := writer.ToBytes()
		require.NoError(f, err)

		f.Add(data)
	}

	f.Fuzz(func(t *testing.T, data []byte) {
		start := time.Now()

		// search for panics
		_, _ = FromBytes(data)

		if time.Since(start) > 100*time.Millisecond {
			panic("too long")
		}
	})
}
