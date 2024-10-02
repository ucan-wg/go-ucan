package invocation

import (
	_ "embed"
	"fmt"
	"sync"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/bindnode"
	"github.com/ipld/go-ipld-prime/schema"

	"github.com/ucan-wg/go-ucan/pkg/meta"
	"github.com/ucan-wg/go-ucan/token/internal/envelope"
)

// [Tag] is the string used as a key within the SigPayload that identifies
// that the TokenPayload is an invocation.
//
// [Tag]: https://github.com/ucan-wg/invocation#type-tag
const Tag = "ucan/inv@1.0.0-rc.1"

//go:embed invocation.ipldsch
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

// TODO
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

	// A unique, random nonce
	Nonce []byte

	// Arbitrary Metadata
	Meta meta.Meta

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
