package invocation

import "github.com/ipfs/go-cid"

// Bundle carries together a decoded token with its Cid and raw signed data.
type Bundle struct {
	Cid     cid.Cid
	Decoded *Token
	Sealed  []byte
}
