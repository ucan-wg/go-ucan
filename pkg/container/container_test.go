package container

import (
	"bytes"
	"crypto/rand"
	"io"
	"testing"

	"github.com/ipfs/go-cid"
	mh "github.com/multiformats/go-multihash"
	"github.com/stretchr/testify/require"
)

func TestContainerRoundTrip(t *testing.T) {
	for _, tc := range []struct {
		name   string
		writer func(ctn Container, w io.Writer) error
		reader func(io.Reader) (Container, error)
	}{
		{"carBytes", Container.ToCar, FromCar},
		{"carBase64", Container.ToCarBase64, FromCarBase64},
		{"cbor", Container.ToCbor, FromCbor},
		{"cborBase64", Container.ToCborBase64, FromCborBase64},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ctn := New()

			builder := cid.V1Builder{Codec: cid.DagCBOR, MhType: mh.SHA2_256}

			var dataSize int

			for i := 0; i < 10; i++ {
				data := randBytes(32)
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

func randBytes(n int) []byte {
	b := make([]byte, n)
	_, _ = rand.Read(b)
	return b
}
