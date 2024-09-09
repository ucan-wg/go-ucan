package did_test

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/multiformats/go-multicodec"
	"github.com/stretchr/testify/require"
	"github.com/ucan-wg/go-ucan/did"
)

func TestFromPubKey(t *testing.T) {
	t.Parallel()

	const example = "did:key:z6MkhaXgBZDvotDkL5257faiztiGiC2QtKLGpbnnEGta2doK"

	id, err := did.Parse(example)
	require.NoError(t, err)

	buf := bytes.NewBuffer(id.Bytes())

	code, err := binary.ReadUvarint(buf)
	require.NoError(t, err)
	require.Equal(t, uint64(multicodec.Ed25519Pub), code)

	pubKey, err := crypto.UnmarshalEd25519PublicKey(buf.Bytes())
	require.NoError(t, err)

	id2, err := did.FromPubKey(pubKey)
	require.Equal(t, id, id2)
}
