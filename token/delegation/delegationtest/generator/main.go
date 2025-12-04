package main

import (
	"github.com/MetaMask/go-did-it/didtest"
)

func main() {
	gen := &generator{}
	err := gen.createSelfDelegations(didtest.Personas())
	if err != nil {
		panic(err)
	}
	err = gen.chainPersonas(didtest.Personas(), acc{}, noopVariant())
	if err != nil {
		panic(err)
	}
	err = gen.writeGoFile()
	if err != nil {
		panic(err)
	}
}
