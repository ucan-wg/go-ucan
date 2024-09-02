package delegation

import (
	_ "embed"
	"fmt"
	"sync"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/schema"
)

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

func PayloadType() schema.Type {
	return mustLoadSchema().TypeByName("Payload")
}

type PayloadModel struct {
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
	Meta MetaModel

	// "Not before" UTC Unix Timestamp in seconds (valid from), 53-bits integer
	// optional: can be nil
	Nbf *int64
	// The timestamp at which the Invocation becomes invalid
	// optional: can be nil
	Exp *int64
}

type MetaModel struct {
	Keys   []string
	Values map[string]datamodel.Node
}
