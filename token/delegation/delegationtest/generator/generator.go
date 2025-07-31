package main

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"slices"
	"time"

	"github.com/MetaMask/go-did-it"
	didkeyctl "github.com/MetaMask/go-did-it/controller/did-key"
	"github.com/MetaMask/go-did-it/crypto"
	"github.com/ipfs/go-cid"

	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/pkg/policy"
	"github.com/ucan-wg/go-ucan/pkg/policy/policytest"
	"github.com/ucan-wg/go-ucan/token/delegation"
	"github.com/ucan-wg/go-ucan/token/delegation/delegationtest"
	"github.com/ucan-wg/go-ucan/token/internal/didtest"
)

const (
	tokenNamePrefix      = "Token"
	proorChainNamePrefix = "Proof"
	tokenExt             = ".dagcbor"
)

var constantNonce = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b}

type newDelegationParams struct {
	privKey crypto.PrivateKeySigningBytes // iss
	aud     did.DID
	cmd     command.Command
	pol     policy.Policy
	sub     did.DID
	opts    []delegation.Option
}

type token struct {
	name string
	id   cid.Cid
}

type proof struct {
	name string
	prf  []cid.Cid
}

type acc struct {
	name  string
	chain []cid.Cid
}

type variant struct {
	name    string
	variant func(*newDelegationParams)
}

func noopVariant() variant {
	return variant{
		name:    "",
		variant: func(_ *newDelegationParams) {},
	}
}

type generator struct {
	dlgs   []token
	chains []proof
}

func (g *generator) chainPersonas(personas []didtest.Persona, acc acc, vari variant) error {
	acc.name += personas[0].Name()

	proofName := acc.name
	if len(vari.name) > 0 {
		proofName += "_" + vari.name
	}
	g.createProofChain(proofName, acc.chain)

	if len(personas) < 2 {
		return nil
	}

	name := personas[0].Name() + personas[1].Name()

	params := newDelegationParams{
		privKey: personas[0].PrivKey(),
		aud:     personas[1].DID(),
		cmd:     delegationtest.NominalCommand,
		pol:     policytest.EmptyPolicy,
		sub:     didtest.PersonaAlice.DID(),
		opts: []delegation.Option{
			delegation.WithNonce(constantNonce),
		},
	}

	// Create each nominal token and continue the chain
	id, err := g.createDelegation(params, name, vari)
	if err != nil {
		return err
	}
	acc.chain = append(acc.chain, id)
	err = g.chainPersonas(personas[1:], acc, vari)
	if err != nil {
		return err
	}

	// If the user is Carol, create variants for each invalid and/or optional
	// parameter and also continue the chain
	if personas[0] == didtest.PersonaCarol {
		variants := []variant{
			{name: "InvalidExpandedCommand", variant: func(p *newDelegationParams) {
				p.cmd = delegationtest.ExpandedCommand
			}},
			{name: "ValidAttenuatedCommand", variant: func(p *newDelegationParams) {
				p.cmd = delegationtest.AttenuatedCommand
			}},
			{name: "InvalidSubject", variant: func(p *newDelegationParams) {
				p.sub = didtest.PersonaBob.DID()
			}},
			{name: "InvalidExpired", variant: func(p *newDelegationParams) {
				// Note: this makes the generator not deterministic
				p.opts = append(p.opts, delegation.WithExpiration(time.Now().Add(time.Second)))
			}},
			{name: "InvalidInactive", variant: func(p *newDelegationParams) {
				nbf, err := time.Parse(time.RFC3339, "2070-01-01T00:00:00Z")
				if err != nil {
					panic(err)
				}
				p.opts = append(p.opts, delegation.WithNotBefore(nbf))
			}},
			{name: "ValidExamplePolicy", variant: func(p *newDelegationParams) {
				p.pol = policytest.SpecPolicy
			}},
		}

		// Start a branch in the recursion for each of the variants
		for _, v := range variants {
			id, err := g.createDelegation(params, name, v)
			if err != nil {
				return err
			}

			// replace the previous Carol token id with the one from the variant
			acc.chain[len(acc.chain)-1] = id
			err = g.chainPersonas(personas[1:], acc, v)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (g *generator) createDelegation(params newDelegationParams, name string, vari variant) (cid.Cid, error) {
	vari.variant(&params)

	issDID := didkeyctl.FromPrivateKey(params.privKey)

	tkn, err := delegation.New(issDID, params.aud, params.cmd, params.pol, params.sub, params.opts...)
	if err != nil {
		return cid.Undef, err
	}

	data, id, err := tkn.ToSealed(params.privKey)
	if err != nil {
		return cid.Undef, err
	}

	dlgName := tokenNamePrefix + name
	if len(vari.name) > 0 {
		dlgName += "_" + vari.name
	}

	err = os.WriteFile(filepath.Join("..", delegationtest.TokenDir, dlgName+tokenExt), data, 0o644)
	if err != nil {
		return cid.Undef, err
	}

	g.dlgs = append(g.dlgs, token{
		name: dlgName,
		id:   id,
	})

	return id, nil
}

func (g *generator) createProofChain(name string, prf []cid.Cid) {
	if len(prf) < 1 {
		return
	}

	clone := make([]cid.Cid, len(prf))
	copy(clone, prf)

	g.chains = append(g.chains, proof{
		name: proorChainNamePrefix + name,
		prf:  clone,
	})
}

func (g *generator) writeGoFile() error {
	buf := bytes.NewBuffer(nil)

	Println := func(a ...any) { _, _ = fmt.Fprintln(buf, a...) }
	Printf := func(format string, a ...any) { _, _ = fmt.Fprintf(buf, format, a...) }

	Println("// Code generated by delegationtest - DO NOT EDIT.")
	Println()
	Println("package delegationtest")
	Println()
	Println("import (")
	Println("\t\"github.com/ipfs/go-cid\"")
	Println()
	Println("\t\"github.com/ucan-wg/go-ucan/token/delegation\"")
	Println(")")

	refs := make(map[cid.Cid]string, len(g.dlgs))

	for _, d := range g.dlgs {
		refs[d.id] = d.name + "CID"

		Println()
		Println("var (")
		Printf("\t%sCID    = cid.MustParse(\"%s\")\n", d.name, d.id.String())
		Printf("\t%sSealed = mustGetBundle(%s).Sealed\n", d.name, d.name+"CID")
		Printf("\t%sBundle = mustGetBundle(%s)\n", d.name, d.name+"CID")
		Printf("\t%s       = mustGetBundle(%s).Decoded\n", d.name, d.name+"CID")
		Println(")")
	}

	Println()
	Println("var AllTokens = []*delegation.Token{")
	for _, d := range g.dlgs {
		Printf("\t%s,\n", d.name)
	}
	Println("}")

	Println()
	Println("var AllBundles = []delegation.Bundle{")
	for _, d := range g.dlgs {
		Printf("\t%sBundle,\n", d.name)
	}
	Println("}")

	Println()
	Println("var cidToName = map[cid.Cid]string{")
	for _, d := range g.dlgs {
		Printf("\t%sCID: \"%s\",\n", d.name, d.name)
	}
	Println("}")

	for _, c := range g.chains {
		Println()
		Printf("var %s = []cid.Cid{\n", c.name)

		for _, d := range slices.Backward(c.prf) {
			Printf("\t%s,\n", refs[d])
		}

		Println("}")
	}

	out, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	return os.WriteFile("../token_gen.go", out, 0666)
}
