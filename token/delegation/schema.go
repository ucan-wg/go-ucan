package delegation

import (
	_ "embed"
	"fmt"
	"sync"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/node/bindnode"
	"github.com/ipld/go-ipld-prime/schema"

	"github.com/ucan-wg/go-ucan/pkg/meta"
	"github.com/ucan-wg/go-ucan/token/internal/envelope"
)

// [Tag] is the string used as a key within the SigPayload that identifies
// that the TokenPayload is a delegation.
//
// [Tag]: https://github.com/ucan-wg/delegation/tree/v1_ipld#type-tag
const Tag = "ucan/dlg@1.0.0-rc.1"

// TODO: update the above Tag URL once the delegation specification is merged.

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
	Meta meta.Meta

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
