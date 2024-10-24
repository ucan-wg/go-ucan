package container

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/stretchr/testify/require"

	"github.com/ucan-wg/go-ucan/did"
	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/pkg/policy"
	"github.com/ucan-wg/go-ucan/pkg/policy/literal"
	"github.com/ucan-wg/go-ucan/token/delegation"
)

func TestContainerRoundTrip(t *testing.T) {
	for _, tc := range []struct {
		name   string
		writer func(ctn Writer, w io.Writer) error
		reader func(io.Reader) (Reader, error)
	}{
		{"car", Writer.ToCar, FromCar},
		{"carBase64", Writer.ToCarBase64, FromCarBase64},
		{"cbor", Writer.ToCbor, FromCbor},
		{"cborBase64", Writer.ToCborBase64, FromCborBase64},
	} {
		t.Run(tc.name, func(t *testing.T) {
			tokens := make(map[cid.Cid]*delegation.Token)
			var dataSize int

			writer := NewWriter()

			for i := 0; i < 10; i++ {
				dlg, c, data := randToken()
				writer.AddSealed(c, data)
				tokens[c] = dlg
				dataSize += len(data)
			}

			buf := bytes.NewBuffer(nil)

			err := tc.writer(writer, buf)
			require.NoError(t, err)

			t.Logf("data size %d", dataSize)
			t.Logf("container overhead: %d%%, %d bytes", int(float32(buf.Len()-dataSize)/float32(dataSize)*100.0), buf.Len()-dataSize)

			reader, err := tc.reader(bytes.NewReader(buf.Bytes()))
			require.NoError(t, err)

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
		{"car", Writer.ToCar, FromCar},
		{"carBase64", Writer.ToCarBase64, FromCarBase64},
		{"cbor", Writer.ToCbor, FromCbor},
		{"cborBase64", Writer.ToCborBase64, FromCborBase64},
	} {
		writer := NewWriter()

		for i := 0; i < 10; i++ {
			_, c, data := randToken()
			writer.AddSealed(c, data)
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

func randDID() (crypto.PrivKey, did.DID) {
	privKey, _, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		panic(err)
	}
	d, err := did.FromPrivKey(privKey)
	if err != nil {
		panic(err)
	}
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
		delegation.WithSubject(iss),
	}
	for i := 0; i < 3; i++ {
		opts = append(opts, delegation.WithMeta(randomString(8), randomString(10)))
	}

	t, err := delegation.New(priv, aud, cmd, pol, opts...)
	if err != nil {
		panic(err)
	}
	b, c, err := t.ToSealed(priv)
	if err != nil {
		panic(err)
	}
	return t, c, b
}
