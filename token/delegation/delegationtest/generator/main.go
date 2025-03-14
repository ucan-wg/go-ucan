package main

import (
	"github.com/ucan-wg/go-ucan/did/didtest"
)

func main() {
	gen := &generator{}
	err := gen.chainPersonas(didtest.Personas(), acc{}, noopVariant())
	if err != nil {
		panic(err)
	}
	err = gen.writeGoFile()
	if err != nil {
		panic(err)
	}
}
