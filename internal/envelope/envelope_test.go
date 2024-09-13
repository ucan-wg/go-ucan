package envelope_test

import (
	"encoding/base64"
	"testing"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/fluent/qp"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	crypto "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/ucan-wg/go-ucan/did"
	"github.com/ucan-wg/go-ucan/internal/envelope"
	"github.com/ucan-wg/go-ucan/internal/token"
	"gotest.tools/v3/golden"
)

const (
	exampleDID             = "did:key:z6MkpuK2Amsu1RqcLGgmHHQHhvmeXCCBVsM4XFSg2cCyg4Nh"
	examplePrivKeyCfg      = "CAESQP9v2uqECTuIi45dyg3znQvsryvf2IXmOF/6aws6aCehm0FVrj0zHR5RZSDxWNjcpcJqsGym3sjCungX9Zt5oA4="
	exampleSignatureStr    = "PZV6A2aI7n+MlyADqcqmWhkuyNrgUCDz+qSLSnI9bpasOwOhKUTx95m5Nu5CO/INa1LqzHGioD9+PVf6qdtTBg"
	exampleTag             = "ucan/example@v1.0.0-rc.1"
	exampleVarsigHeaderStr = "NO0BcQ"

	invalidSignatureStr = "PZV6A2aI7n+MlyADqcqmWhkuyNrgUCDz+qSLSnI9bpasOwOhKUTx95m5Nu5CO/INa1LqzHGioD9+PVf6qdtTBK"

	exampleDAGCBORFilename = "example.dagcbor"
	exampleDAGJSONFilename = "example.dagjson"
)

func TestNew(t *testing.T) {
	t.Parallel()

	envel := exampleEnvelope(t)
	assert.NotZero(t, envel)

	assert.Equal(t, exampleSignature(t), envel.Signature())
	assert.Equal(t, exampleTag, envel.Tag())
	assert.Equal(t, exampleVarsigHeader(t), envel.VarsigHeader())
	assert.EqualValues(t, exampleGoldenTokenPayload(t), envel.TokenPayload())
}

func TestWrap(t *testing.T) {
	t.Parallel()

	node, err := envelope.Wrap(examplePrivKey(t), exampleToken(t), exampleTag)
	require.NoError(t, err)

	cbor, err := ipld.Encode(node, dagcbor.Encode)
	require.NoError(t, err)

	golden.AssertBytes(t, cbor, exampleDAGCBORFilename)

	json, err := ipld.Encode(node, dagjson.Encode)
	require.NoError(t, err)

	golden.Assert(t, string(json), exampleDAGJSONFilename)
}

func TestEnvelope_Verify(t *testing.T) {
	t.Parallel()

	t.Run("valid signature by issuer", func(t *testing.T) {
		t.Parallel()

		envel := exampleEnvelope(t)
		ok, err := envel.Verify()
		require.NoError(t, err)
		assert.True(t, ok)
	})

	t.Run("invalid signature by wrong issuer", func(t *testing.T) {
		t.Parallel()

		envel, err := envelope.Unwrap(invalidNodeFromGolden(t))
		require.NoError(t, err)

		ok, _ := envel.Verify()
		assert.False(t, ok)
	})
}

func TestEnvelope_Wrap(t *testing.T) {
	t.Parallel()

	envel := exampleEnvelope(t)

	node, err := envel.Wrap()
	require.NoError(t, err)

	cbor, err := ipld.Encode(node, dagcbor.Encode)
	require.NoError(t, err)

	assert.Equal(t, golden.Get(t, exampleDAGCBORFilename), cbor)
}

func exampleGoldenEnvelope(t *testing.T) *envelope.Envelope {
	t.Helper()

	envel, err := envelope.Unwrap(exampleGoldenNode(t))
	require.NoError(t, err)

	return envel
}

func exampleGoldenNode(t *testing.T) datamodel.Node {
	t.Helper()

	cbor := golden.Get(t, exampleDAGCBORFilename)

	node, err := ipld.Decode(cbor, dagcbor.Decode)
	require.NoError(t, err)

	return node
}

func exampleGoldenTokenPayload(t *testing.T) *token.Token {
	t.Helper()

	return exampleGoldenEnvelope(t).TokenPayload()
}

func examplePrivKey(t *testing.T) crypto.PrivKey {
	t.Helper()

	privKeyEnc, err := crypto.ConfigDecodeKey(examplePrivKeyCfg)
	require.NoError(t, err)

	privKey, err := crypto.UnmarshalPrivateKey(privKeyEnc)
	require.NoError(t, err)

	return privKey
}

func exampleEnvelope(t *testing.T) *envelope.Envelope {
	t.Helper()

	envel, err := envelope.New(examplePrivKey(t), exampleToken(t), exampleTag)
	require.NoError(t, err)

	return envel
}

func examplePubKey(t *testing.T) crypto.PubKey {
	t.Helper()

	return examplePrivKey(t).GetPublic()
}

func exampleSignature(t *testing.T) []byte {
	t.Helper()

	sig, err := base64.RawStdEncoding.DecodeString(exampleSignatureStr)
	require.NoError(t, err)

	return sig
}

func exampleToken(t *testing.T) *token.Token {
	t.Helper()

	id, err := did.FromPubKey(examplePubKey(t))
	require.NoError(t, err)

	return &token.Token{
		Issuer: id.String(),
	}
}

func exampleVarsigHeader(t *testing.T) []byte {
	t.Helper()

	hdr, err := base64.RawStdEncoding.DecodeString(exampleVarsigHeaderStr)
	require.NoError(t, err)

	return hdr
}

func invalidNodeFromGolden(t *testing.T) datamodel.Node {
	t.Helper()

	invalidSig, err := base64.RawStdEncoding.DecodeString(invalidSignatureStr)
	require.NoError(t, err)

	envelNode := exampleGoldenNode(t)
	sigPayloadNode, err := envelNode.LookupByIndex(1)
	require.NoError(t, err)

	node, err := qp.BuildList(basicnode.Prototype.Any, 2, func(la datamodel.ListAssembler) {
		qp.ListEntry(la, qp.Bytes(invalidSig))
		qp.ListEntry(la, qp.Node(sigPayloadNode))
	})
	require.NoError(t, err)

	return node
}
