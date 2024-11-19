// Package didtest provides Personas that can be used for testing.  Each
// Persona has a name, crypto.PrivKey and associated crypto.PubKey and
// did.DID.
package didtest

import (
	"sync"
	"testing"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/stretchr/testify/require"
	"github.com/ucan-wg/go-ucan/did"
)

const (
	alicePrivKeyCfg = "CAESQHdNJLBBiuc1AdwPHBkubB2KS1p0cv2JEF7m8tfwtrcm5ajaYPm+XmVCmtcHOF2lGDlmaiDA7emfwD3IrcyES0M="
	bobPrivKeyCfg   = "CAESQHBz+AIop1g+9iBDj+ufUc/zm9/ry7c6kDFO8Wl/D0+H63V9hC6s9l4npf3pYEFCjBtlR0AMNWMoFQKSlYNKo20="
	carolPrivKeyCfg = "CAESQPrCgkcHnYFXDT9AlAydhPECBEivEuuVx9dJxLjVvDTmJIVNivfzg6H4mAiPfYS+5ryVVUZTHZBzvMuvvvG/Ks0="
	danPrivKeyCfg   = "CAESQCgNhzofKhC+7hW6x+fNd7iMPtQHeEmKRhhlduf/I7/TeOEFYAEflbJ0sAhMeDJ/HQXaAvsWgHEbJ3ZLhP8q2B0="
	erinPrivKeyCfg  = "CAESQKhCJo5UBpQcthko8DKMFsbdZ+qqQ5oc01CtLCqrE90dF2GfRlrMmot3WPHiHGCmEYi5ZMEHuiSI095e/6O4Bpw="
	frankPrivKeyCfg = "CAESQDlXPKsy3jHh7OWTWQqyZF95Ueac5DKo7xD0NOBE5F2BNr1ZVxRmJ2dBELbOt8KP9sOACcO9qlCB7uMA1UQc7sk="
)

// Persona is a generic participant used for cryptographic testing.
type Persona int

// The provided  Personas were selected from the first few generic
// participants listed in this [table].
//
// [table]: https://en.wikipedia.org/wiki/Alice_and_Bob#Cryptographic_systems
const (
	PersonaAlice Persona = iota
	PersonaBob
	PersonaCarol
	PersonaDan
	PersonaErin
	PersonaFrank
)

var (
	once     sync.Once
	privKeys = make(map[Persona]crypto.PrivKey, 6)
	err      error
)

// DID returns a did.DID based on the Persona's Ed25519 public key.
func (p Persona) DID(t *testing.T) did.DID {
	t.Helper()

	did, err := did.FromPrivKey(p.PrivKey(t))
	require.NoError(t, err)

	return did
}

// Name returns the username of the Persona.
func (p Persona) Name(t *testing.T) string {
	t.Helper()

	name, ok := map[Persona]string{
		PersonaAlice: "Alice",
		PersonaBob:   "Bob",
		PersonaCarol: "Carol",
		PersonaDan:   "Dan",
		PersonaErin:  "Erin",
		PersonaFrank: "Frank",
	}[p]
	if !ok {
		t.Fatal("Unknown persona:", p)
	}

	return name
}

// PrivKey returns the Ed25519 private key for the Persona.
func (p Persona) PrivKey(t *testing.T) crypto.PrivKey {
	t.Helper()

	once.Do(func() {
		for persona, privKeyCfg := range privKeyCfgs(t) {
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
	})
	require.NoError(t, err)

	return privKeys[p]
}

// PrivKeyConfig returns the marshaled and encoded Ed25519 private key
// for the Persona.
func (p Persona) PrivKeyConfig(t *testing.T) string {
	t.Helper()

	return privKeyCfgs(t)[p]
}

// PubKey returns the Ed25519 public key for the Persona.
func (p Persona) PubKey(t *testing.T) crypto.PubKey {
	t.Helper()

	return p.PrivKey(t).GetPublic()
}

// PubKeyConfig returns the marshaled and encoded Ed25519 public key
// for the Persona.
func (p Persona) PubKeyConfig(t *testing.T) string {
	pubKeyMar, err := crypto.MarshalPublicKey(p.PrivKey(t).GetPublic())
	require.NoError(t, err)

	return crypto.ConfigEncodeKey(pubKeyMar)
}

func privKeyCfgs(t *testing.T) map[Persona]string {
	t.Helper()

	return map[Persona]string{
		PersonaAlice: alicePrivKeyCfg,
		PersonaBob:   bobPrivKeyCfg,
		PersonaCarol: carolPrivKeyCfg,
		PersonaDan:   danPrivKeyCfg,
		PersonaErin:  erinPrivKeyCfg,
		PersonaFrank: frankPrivKeyCfg,
	}
}

// Personas returns an (alphabetically) ordered list of the defined
// Persona values.
func Personas(t *testing.T) []Persona {
	t.Helper()

	return []Persona{
		PersonaAlice,
		PersonaBob,
		PersonaCarol,
		PersonaDan,
		PersonaErin,
		PersonaFrank,
	}
}
