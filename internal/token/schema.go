package token

import (
	_ "embed"
	"fmt"
	"sync"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/bindnode"
	"github.com/ipld/go-ipld-prime/schema"
)

const tokenTypeName = "Token"

//go:embed token.ipldsch
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
		panic(fmt.Errorf("%w: %w", ErrFailedSchemaLoad, err))
	}

	tknType := ts.TypeByName(tokenTypeName)
	if tknType == nil {
		panic(fmt.Errorf("%w: %s", ErrNoSchemaType, tokenTypeName))
	}

	return ts
}

func tokenType() schema.Type {
	return mustLoadSchema().TypeByName(tokenTypeName)
}

func Prototype() schema.TypedPrototype {
	return bindnode.Prototype((*Token)(nil), tokenType())
}
