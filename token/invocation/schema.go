package invocation

import (
	_ "embed"
	"fmt"
	"sync"

	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/bindnode"
	"github.com/ipld/go-ipld-prime/schema"

	"github.com/ucan-wg/go-ucan/pkg/args"
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

type tokenPayloadModel struct {
	// The DID of the Invoker
	Iss string
	// The DID of Subject being invoked
	Sub string
	// The DID of the intended Executor if different from the Subject
	Aud *string

	// The Command
	Cmd string
	// The Command's Arguments
	Args *args.Args
	// Delegations that prove the chain of authority
	Prf []cid.Cid

	// Arbitrary Metadata
	Meta *meta.Meta

	// A unique, random nonce
	Nonce []byte
	// The timestamp at which the Invocation becomes invalid
	// optional: can be nil
	Exp *int64
	// The timestamp at which the Invocation was created
	Iat *int64

	// An optional CID of the Receipt that enqueued the Task
	Cause *cid.Cid
}

func (e *tokenPayloadModel) Prototype() schema.TypedPrototype {
	return bindnode.Prototype((*tokenPayloadModel)(nil), payloadType())
}

func (*tokenPayloadModel) Tag() string {
	return Tag
}
