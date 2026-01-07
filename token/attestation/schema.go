package attestation

import (
	_ "embed"
	"fmt"
	"sync"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/bindnode"
	"github.com/ipld/go-ipld-prime/schema"

	"github.com/ucan-wg/go-ucan/pkg/claims"
	"github.com/ucan-wg/go-ucan/pkg/meta"
	"github.com/ucan-wg/go-ucan/token/internal/envelope"
)

// [Tag] is the string used as a key within the SigPayload that identifies
// that the TokenPayload is an attestation.
//
// [Tag]: TODO: TBD
const Tag = "ucan/att@tbd" // TODO: TBD

//go:embed attestation.ipldsch
var schemaBytes []byte

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

func payloadType() schema.Type {
	return mustLoadSchema().TypeByName("Payload")
}

var _ envelope.Tokener = (*tokenPayloadModel)(nil)

type tokenPayloadModel struct {
	// The DID of the Invoker
	Iss string

	// Arbitrary claims
	Claims *claims.Claims

	// Arbitrary Metadata
	Meta *meta.Meta

	// A unique, random nonce
	Nonce []byte
	// The timestamp at which the Invocation becomes invalid
	// optional: can be nil
	Exp *int64
	// The timestamp at which the Invocation was created
	Iat *int64
}

func (e *tokenPayloadModel) Prototype() schema.TypedPrototype {
	return bindnode.Prototype((*tokenPayloadModel)(nil), payloadType())
}

func (*tokenPayloadModel) Tag() string {
	return Tag
}
