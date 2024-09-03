package varsig_test

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/libp2p/go-libp2p/core/crypto/pb"
	"github.com/stretchr/testify/assert"
	"github.com/ucan-wg/go-ucan/v1/internal/varsig"
)

func TestDecode(t *testing.T) {
	t.Parallel()

	notAHeader := base64.RawStdEncoding.EncodeToString([]byte("not a header"))
	keyType, err := varsig.Decode(notAHeader)
	assert.Equal(t, pb.KeyType(-1), keyType)
	assert.ErrorIs(t, err, varsig.ErrUnknownHeader)
}

func ExampleDecode() {
	keyType, _ := varsig.Decode("NIUkEoACcQ")
	fmt.Println(keyType.String())
	// Output:
	// RSA
}

func TestEncode(t *testing.T) {
	t.Parallel()

	header, err := varsig.Encode(pb.KeyType(99))
	assert.Empty(t, header)
	assert.ErrorIs(t, err, varsig.ErrUnknownKeyType)
}

func ExampleEncode() {
	header, _ := varsig.Encode(pb.KeyType_RSA)
	fmt.Println(header)

	// Output:
	// NIUkEoACcQ
}
