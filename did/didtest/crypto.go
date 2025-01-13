// Package didtest provides Personas that can be used for testing.  Each
// Persona has a name, crypto.PrivKey and associated crypto.PubKey and
// did.DID.
package didtest

import (
	"fmt"

	"github.com/libp2p/go-libp2p/core/crypto"

	"github.com/ucan-wg/go-ucan/did"
)

const (
	alicePrivKeyB64 = "CAESQHdNJLBBiuc1AdwPHBkubB2KS1p0cv2JEF7m8tfwtrcm5ajaYPm+XmVCmtcHOF2lGDlmaiDA7emfwD3IrcyES0M="
	bobPrivKeyB64   = "CAESQHBz+AIop1g+9iBDj+ufUc/zm9/ry7c6kDFO8Wl/D0+H63V9hC6s9l4npf3pYEFCjBtlR0AMNWMoFQKSlYNKo20="
	carolPrivKeyB64 = "CAESQPrCgkcHnYFXDT9AlAydhPECBEivEuuVx9dJxLjVvDTmJIVNivfzg6H4mAiPfYS+5ryVVUZTHZBzvMuvvvG/Ks0="
	danPrivKeyB64   = "CAESQCgNhzofKhC+7hW6x+fNd7iMPtQHeEmKRhhlduf/I7/TeOEFYAEflbJ0sAhMeDJ/HQXaAvsWgHEbJ3ZLhP8q2B0="
	erinPrivKeyB64  = "CAESQKhCJo5UBpQcthko8DKMFsbdZ+qqQ5oc01CtLCqrE90dF2GfRlrMmot3WPHiHGCmEYi5ZMEHuiSI095e/6O4Bpw="
	frankPrivKeyB64 = "CAESQDlXPKsy3jHh7OWTWQqyZF95Ueac5DKo7xD0NOBE5F2BNr1ZVxRmJ2dBELbOt8KP9sOACcO9qlCB7uMA1UQc7sk="
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

var privKeys map[Persona]crypto.PrivKey

func init() {
	privKeys = make(map[Persona]crypto.PrivKey, 6)
	for persona, privKeyCfg := range privKeyB64() {
		privKeyMar, err := crypto.ConfigDecodeKey(privKeyCfg)
		if err != nil {
			return
		}

		privKey, err := crypto.UnmarshalPrivateKey(privKeyMar)
		if err != nil {
			return
		}

		privKeys[persona] = privKey
	}
}

// DID returns a did.DID based on the Persona's Ed25519 public key.
func (p Persona) DID() did.DID {
	d, err := did.FromPrivKey(p.PrivKey())
	if err != nil {
		panic(err)
	}
	return d
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
func (p Persona) PrivKey() crypto.PrivKey {
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
func (p Persona) PubKey() crypto.PubKey {
	return p.PrivKey().GetPublic()
}

// PubKeyConfig returns the marshaled and encoded Ed25519 public key
// for the Persona.
func (p Persona) PubKeyConfig() string {
	pubKeyMar, err := crypto.MarshalPublicKey(p.PrivKey().GetPublic())
	if err != nil {
		panic(err)
	}
	return crypto.ConfigEncodeKey(pubKeyMar)
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
