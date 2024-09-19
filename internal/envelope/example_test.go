package envelope_test

import (
	_ "embed"
	"encoding/base64"
	"fmt"
	"sync"
	"testing"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/fluent/qp"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	"github.com/ipld/go-ipld-prime/node/bindnode"
	"github.com/ipld/go-ipld-prime/schema"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/stretchr/testify/require"
	"github.com/ucan-wg/go-ucan/internal/envelope"
	"gotest.tools/v3/golden"
)

const (
	exampleCID             = "zdpuAyw6R5HvKSPzztuzXNYFx3ZGoMHMuAsXL6u3xLGQriRXQ"
	exampleDID             = "did:key:z6MkpuK2Amsu1RqcLGgmHHQHhvmeXCCBVsM4XFSg2cCyg4Nh"
	exampleGreeting        = "world"
	examplePrivKeyCfg      = "CAESQP9v2uqECTuIi45dyg3znQvsryvf2IXmOF/6aws6aCehm0FVrj0zHR5RZSDxWNjcpcJqsGym3sjCungX9Zt5oA4="
	exampleSignatureStr    = "PZV6A2aI7n+MlyADqcqmWhkuyNrgUCDz+qSLSnI9bpasOwOhKUTx95m5Nu5CO/INa1LqzHGioD9+PVf6qdtTBg"
	exampleTag             = "ucan/example@v1.0.0-rc.1"
	exampleTypeName        = "Example"
	exampleVarsigHeaderStr = "NO0BcQ"

	invalidSignatureStr = "PZV6A2aI7n+MlyADqcqmWhkuyNrgUCDz+qSLSnI9bpasOwOhKUTx95m5Nu5CO/INa1LqzHGioD9+PVf6qdtTBK"

	exampleDAGCBORFilename = "example.dagcbor"
	exampleDAGJSONFilename = "example.dagjson"
)

//go:embed testdata/example.ipldsch
var schemaBytes []byte

var (
	once sync.Once
	ts   *schema.TypeSystem
	err  error
)

func mustLoadSchema() *schema.TypeSystem {
	once.Do(func() {
		ts, err = ipld.LoadSchemaBytes(schemaBytes)
	})

	if err != nil {
		panic(fmt.Errorf("failed to load IPLD schema: %s", err))
	}

	return ts
}

func exampleType() schema.Type {
	return mustLoadSchema().TypeByName(exampleTypeName)
}

var _ envelope.Tokener = (*Example)(nil)

type Example struct {
	Hello  string
	Issuer string
}

func newExample(t *testing.T) *Example {
	t.Helper()

	return &Example{
		Hello:  exampleGreeting,
		Issuer: exampleDID,
	}
}

func (e *Example) Prototype() schema.TypedPrototype {
	return bindnode.Prototype(e, exampleType())
}

func (*Example) Tag() string {
	return exampleTag
}

func exampleGoldenNode(t *testing.T) datamodel.Node {
	t.Helper()

	cbor := golden.Get(t, exampleDAGCBORFilename)

	node, err := ipld.Decode(cbor, dagcbor.Decode)
	require.NoError(t, err)

	return node
}

func examplePrivKey(t *testing.T) crypto.PrivKey {
	t.Helper()

	privKeyEnc, err := crypto.ConfigDecodeKey(examplePrivKeyCfg)
	require.NoError(t, err)

	privKey, err := crypto.UnmarshalPrivateKey(privKeyEnc)
	require.NoError(t, err)

	return privKey
}

func exampleSignature(t *testing.T) []byte {
	t.Helper()

	sig, err := base64.RawStdEncoding.DecodeString(exampleSignatureStr)
	require.NoError(t, err)

	return sig
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
