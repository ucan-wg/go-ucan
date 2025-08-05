// Package didtest provides Personas that can be used for testing.  Each
// Persona has a name, crypto.PrivKey and associated crypto.PubKey and
// did.DID.
package didtest

import (
	"encoding/base64"
	"fmt"

	"github.com/MetaMask/go-did-it"
	didkeyctl "github.com/MetaMask/go-did-it/controller/did-key"
	"github.com/MetaMask/go-did-it/crypto"
	"github.com/MetaMask/go-did-it/crypto/ed25519"
)

const (
	// all are ed25519 as base64
	alicePrivKeyB64 = "zth/9cTSUVwlLzfEWwLCcOkaEmjrRGPOI6mOJksWAYZ3Toe7ymxAzDeiseyxbmEpJ81qYM3dZ8XrXqgonnTTEw=="
	bobPrivKeyB64   = "+p1REV3MkUnLhUMbFe9RcSsmo33TT/FO85yaV+c6fiYJCBsdiwfMwodlkzSAG3sHQIuZj8qnJ678oJucYy7WEg=="
	carolPrivKeyB64 = "aSu3vTwE7z3pXaTaAhVLeizuqnZUJZQHTCSLMLxyZh5LDoZQn80uoQgMEdsbOhR+zIqrjBn5WviGurDkKYVfug=="
	danPrivKeyB64   = "s1zM1av6og3o0UMNbEs/RyezS7Nk/jbSYL2Z+xPEw9Cho/KuEAa75Sf4yJHclLwpKXNucbrZ2scE8Iy8K05KWQ=="
	erinPrivKeyB64  = "+qHpaAR3iivWMEl+pkXmq+uJeHtqFiY++XOXtZ9Tu/WPABCO+eRFrTCLJykJEzAPGFmkJF8HQ7DMwOH7Ry3Aqw=="
	frankPrivKeyB64 = "4k/1N0+Fq73DxmNbGis9PY2KgKxWmtDWhmi1E6sBLuGd7DS0TWjCn1Xa3lXkY49mFszMjhWC+V6DCBf7R68u4Q=="
)

// Persona is a generic participant used for cryptographic testing.
type Persona int

// The provided Personas were selected from the first few generic
// participants listed in this [table].
//
// [table]: https://en.wikipedia.org/wiki/Alice_and_Bob#Cryptographic_systems
const (
	PersonaAlice Persona = iota + 1
	PersonaBob
	PersonaCarol
	PersonaDan
	PersonaErin
	PersonaFrank
)

var privKeys map[Persona]crypto.PrivateKeySigningBytes

func init() {
	privKeys = make(map[Persona]crypto.PrivateKeySigningBytes, 6)
	for persona, pB64 := range privKeyB64() {
		privBytes, err := base64.StdEncoding.DecodeString(pB64)
		if err != nil {
			return
		}

		privKey, err := ed25519.PrivateKeyFromBytes(privBytes)
		if err != nil {
			return
		}

		privKeys[persona] = privKey
	}
}

// DID returns a did.DID based on the Persona's Ed25519 public key.
func (p Persona) DID() did.DID {
	return didkeyctl.FromPrivateKey(p.PrivKey())
}

// Name returns the username of the Persona.
func (p Persona) Name() string {
	name, ok := map[Persona]string{
		PersonaAlice: "Alice",
		PersonaBob:   "Bob",
		PersonaCarol: "Carol",
		PersonaDan:   "Dan",
		PersonaErin:  "Erin",
		PersonaFrank: "Frank",
	}[p]
	if !ok {
		panic(fmt.Sprintf("Unknown persona: %v", p))
	}

	return name
}

// PrivKey returns the Ed25519 private key for the Persona.
func (p Persona) PrivKey() crypto.PrivateKeySigningBytes {
	res, ok := privKeys[p]
	if !ok {
		panic(fmt.Sprintf("Unknown persona: %v", p))
	}
	return res
}

func (p Persona) PrivKeyConfig() string {
	res, ok := privKeyB64()[p]
	if !ok {
		panic(fmt.Sprintf("Unknown persona: %v", p))
	}
	return res
}

// PubKey returns the Ed25519 public key for the Persona.
func (p Persona) PubKey() crypto.PublicKey {
	return p.PrivKey().Public()
}

func privKeyB64() map[Persona]string {
	return map[Persona]string{
		PersonaAlice: alicePrivKeyB64,
		PersonaBob:   bobPrivKeyB64,
		PersonaCarol: carolPrivKeyB64,
		PersonaDan:   danPrivKeyB64,
		PersonaErin:  erinPrivKeyB64,
		PersonaFrank: frankPrivKeyB64,
	}
}

// Personas returns an (alphabetically) ordered list of the defined
// Persona values.
func Personas() []Persona {
	return []Persona{
		PersonaAlice,
		PersonaBob,
		PersonaCarol,
		PersonaDan,
		PersonaErin,
		PersonaFrank,
	}
}

// DidToName retrieve the persona's name from its DID.
func DidToName(d did.DID) string {
	return map[did.DID]string{
		PersonaAlice.DID(): "Alice",
		PersonaBob.DID():   "Bob",
		PersonaCarol.DID(): "Carol",
		PersonaDan.DID():   "Dan",
		PersonaErin.DID():  "Erin",
		PersonaFrank.DID(): "Frank",
	}[d]
}
