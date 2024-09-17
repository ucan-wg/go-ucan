package main

import (
	"bytes"
	"log/slog"
	"os"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/bindnode"
)

const header = `// Code generated by internal/cmd/token - DO NOT EDIT.

package token

import "github.com/ipld/go-ipld-prime/datamodel"

`

func main() {
	slog.Info("Generating Go types for token.ipldsch")

	if err := Run(); err != nil {
		slog.Error(err.Error())
		slog.Error("Finished but failed to generate and write token_gen.go")

		os.Exit(1)
	}

	slog.Info("Finished generating and writing token_gen.go")
	os.Exit(0)
}

func Run() error {
	schema, err := os.ReadFile("token.ipldsch")
	if err != nil {
		return err
	}

	slog.Debug(string(schema))

	typeSystem, err := ipld.LoadSchemaBytes(schema)
	if err != nil {
		return err
	}

	buf := bytes.NewBufferString(header)

	if err := bindnode.ProduceGoTypes(buf, typeSystem); err != nil {
		return err
	}

	return os.WriteFile("token_gen.go", buf.Bytes(), 0o600)
}