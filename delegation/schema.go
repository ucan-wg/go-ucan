package delegation

import (
	_ "embed"
	"fmt"
	"sync"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/node/bindnode"
	"github.com/ipld/go-ipld-prime/schema"
	"github.com/ucan-wg/go-ucan/internal/envelope"
)

const Tag = "ucan/dlg@1.0.0-rc.1"

//go:embed delegation.ipldsch
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

func payloadType() schema.Type {
	return mustLoadSchema().TypeByName("Payload")
}

var _ envelope.Tokener = (*tokenPayloadModel)(nil)

type tokenPayloadModel struct {
	// Issuer DID (sender)
	Iss string
	// Audience DID (receiver)
	Aud string
	// Principal that the chain is about (the Subject)
	// optional: can be nil
	Sub *string

	// The Command to eventually invoke
	Cmd string

	// The delegation policy
	Pol datamodel.Node

	// A unique, random nonce
	Nonce []byte

	// Arbitrary Metadata
	// optional: can be nil
	Meta metaModel

	// "Not before" UTC Unix Timestamp in seconds (valid from), 53-bits integer
	// optional: can be nil
	Nbf *int64
	// The timestamp at which the Invocation becomes invalid
	// optional: can be nil
	Exp *int64
}

func (e *tokenPayloadModel) Prototype() schema.TypedPrototype {
	return bindnode.Prototype((*tokenPayloadModel)(nil), payloadType())
}

func (*tokenPayloadModel) Tag() string {
	return Tag
}

type metaModel struct {
	Keys   []string
	Values map[string]datamodel.Node
}
