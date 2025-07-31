package envelope_test

import (
	_ "embed"
	"encoding/base64"
	"fmt"
	"sync"
	"testing"

	"github.com/MetaMask/go-did-it/crypto"
	"github.com/MetaMask/go-did-it/crypto/ed25519"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/fluent/qp"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	"github.com/ipld/go-ipld-prime/node/bindnode"
	"github.com/ipld/go-ipld-prime/schema"
	"github.com/stretchr/testify/require"

	"github.com/ucan-wg/go-ucan/token/internal/envelope"
)

const (
	exampleCID        = "zdpuAn4jksvc1gc9PLDqHw2NoFq8CBkRVTTo2xFuW2JUPS5DY"
	exampleDID        = "did:key:z6MkuqvEtTW9L1E91CY3GmL83muetLAA2h8A5fUHjJgqq2Ab"
	exampleGreeting   = "world"
	examplePrivKeyB64 = "V4hh1lcFV43Y6vyOBEVOFTwl1XS/DR0F/kYcz5i6W/DkrUTG8yx09lOwSf36NCHPKSFYv/T1R3WKjNfndgVucA=="
	exampleTag        = "ucan/example@v1.0.0-rc.1"

	invalidSignatureStr = "PZV6A2aI7n+MlyADqcqmWhkuyNrgUCDz+qSLSnI9bpasOwOhKUTx95m5Nu5CO/INa1LqzHGioD9+PVf6qdtTBK"
)

//go:embed testdata/example.ipldsch
var schemaBytes []byte

//go:embed testdata/example.dagcbor
var exampleDagCbor []byte

//go:embed testdata/example.dagjson
var exampleDagJson []byte

var (
	once      sync.Once
	ts        *schema.TypeSystem
	errSchema error
)

func mustLoadSchema() *schema.TypeSystem {
	once.Do(func() {
		ts, errSchema = ipld.LoadSchemaBytes(schemaBytes)
	})

	if errSchema != nil {
		panic(fmt.Errorf("failed to load IPLD schema: %s", errSchema))
	}

	return ts
}

func exampleType() schema.Type {
	return mustLoadSchema().TypeByName("Example")
}

var _ envelope.Tokener = (*Example)(nil)

type Example struct {
	Hello  string
	Issuer string
}

func newExample() *Example {
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

func examplePrivKey(t *testing.T) crypto.PrivateKeySigningBytes {
	t.Helper()

	privBytes, err := base64.StdEncoding.DecodeString(examplePrivKeyB64)
	require.NoError(t, err)

	privKey, err := ed25519.PrivateKeyFromBytes(privBytes)
	require.NoError(t, err)

	return privKey
}

// nodeWithInvalidSignature creates an IPLD node of a token, with an invalid signature
func nodeWithInvalidSignature(t *testing.T) datamodel.Node {
	t.Helper()

	invalidSig, err := base64.RawStdEncoding.DecodeString(invalidSignatureStr)
	require.NoError(t, err)

	cbor := exampleDagCbor

	envelNode, err := ipld.Decode(cbor, dagcbor.Decode)
	require.NoError(t, err)

	sigPayloadNode, err := envelNode.LookupByIndex(1)
	require.NoError(t, err)

	node, err := qp.BuildList(basicnode.Prototype.Any, 2, func(la datamodel.ListAssembler) {
		qp.ListEntry(la, qp.Bytes(invalidSig))
		qp.ListEntry(la, qp.Node(sigPayloadNode))
	})
	require.NoError(t, err)

	return node
}
