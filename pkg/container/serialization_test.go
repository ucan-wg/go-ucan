package container

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p/core/crypto"
	mh "github.com/multiformats/go-multihash"
	"github.com/stretchr/testify/require"

	"github.com/ucan-wg/go-ucan/did"
	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/pkg/policy"
	"github.com/ucan-wg/go-ucan/pkg/policy/literal"
	"github.com/ucan-wg/go-ucan/pkg/policy/selector"
	"github.com/ucan-wg/go-ucan/token/delegation"
)

func TestContainerRoundTrip(t *testing.T) {
	for _, tc := range []struct {
		name   string
		writer func(ctn Container, w io.Writer) error
		reader func(io.Reader) (Container, error)
	}{
		{"car", Container.ToCar, FromCar},
		{"carBase64", Container.ToCarBase64, FromCarBase64},
		{"carGzip", Container.ToCarGzip, FromCarGzip},
		{"carGzipBase64", Container.ToCarGzipBase64, FromCarGzipBase64},
		{"cbor", Container.ToCbor, FromCbor},
		{"cborBase64", Container.ToCborBase64, FromCborBase64},
		{"cborGzip", Container.ToCborGzip, FromCborGzip},
		{"cborGzipBase64", Container.ToCborGzipBase64, FromCborGzipBase64},
		{"cborFlate", Container.ToCborFlate, FromCborFlate},
		{"cborFlateBase64", Container.ToCborFlateBase64, FromCborFlateBase64},
		{"cbor2", Container.ToCbor2, FromCbor2},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ctn := New()

			builder := cid.V1Builder{Codec: cid.DagCBOR, MhType: mh.SHA2_256}

			var dataSize int

			for i := 0; i < 10; i++ {
				data := randTokenBytes()
				c, err := builder.Sum(data)
				require.NoError(t, err)
				ctn.AddBytes(c, data)
				dataSize += len(data)
			}

			buf := bytes.NewBuffer(nil)

			err := tc.writer(ctn, buf)
			require.NoError(t, err)

			t.Logf("data size %d", dataSize)
			t.Logf("container overhead: %d%%, %d bytes", int(float32(buf.Len()-dataSize)/float32(dataSize)*100.0), buf.Len()-dataSize)

			ctn2, err := tc.reader(bytes.NewReader(buf.Bytes()))
			require.NoError(t, err)

			require.Equal(t, ctn, ctn2)
		})
	}
}

func BenchmarkContainerSerialisation(b *testing.B) {
	for _, tc := range []struct {
		name   string
		writer func(ctn Container, w io.Writer) error
		reader func(io.Reader) (Container, error)
	}{
		{"car", Container.ToCar, FromCar},
		{"carBase64", Container.ToCarBase64, FromCarBase64},
		{"carGzip", Container.ToCarGzip, FromCarGzip},
		{"carGzipBase64", Container.ToCarGzipBase64, FromCarGzipBase64},
		{"cbor", Container.ToCbor, FromCbor},
		{"cborBase64", Container.ToCborBase64, FromCborBase64},
		{"cborGzip", Container.ToCborGzip, FromCborGzip},
		{"cborGzipBase64", Container.ToCborGzipBase64, FromCborGzipBase64},
		{"cborFlate", Container.ToCborFlate, FromCborFlate},
		{"cborFlateBase64", Container.ToCborFlateBase64, FromCborFlateBase64},
		{"cbor2", Container.ToCbor2, FromCbor2},
	} {
		ctn := New()

		builder := cid.V1Builder{Codec: cid.DagCBOR, MhType: mh.SHA2_256}

		var dataSize int

		for i := 0; i < 10; i++ {
			data := randTokenBytes()
			c, err := builder.Sum(data)
			require.NoError(b, err)
			ctn.AddBytes(c, data)
			dataSize += len(data)
		}

		buf := bytes.NewBuffer(nil)
		_ = tc.writer(ctn, buf)

		b.Run(tc.name+"_write", func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				buf := bytes.NewBuffer(nil)
				_ = tc.writer(ctn, buf)
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

func randBytes(n int) []byte {
	b := make([]byte, n)
	_, _ = rand.Read(b)
	return b
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
	rand.Read(b)
	return fmt.Sprintf("%x", b)[0:length]
}

func randTokenBytes() []byte {
	priv, iss := randDID()
	_, aud := randDID()
	cmd := command.New("foo", "bar")
	pol := policy.Policy{policy.All(
		selector.MustParse(".[]"),
		policy.GreaterThan(selector.MustParse(".value"), literal.Int(2)),
	)}

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
	b, _, err := t.ToSealed(priv)
	if err != nil {
		panic(err)
	}
	return b
}
