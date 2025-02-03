package example

import (
	"encoding/base64"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/ucan-wg/go-ucan/did"
)

// Endpoints

var ServerUrl = ":8080"
var IssuerUrl = ":8081"

// Service

var ServicePrivKey crypto.PrivKey
var ServiceDid did.DID

func init() {
	privRaw, _ := base64.StdEncoding.DecodeString("CAESQGs7hPBRBmxH1UmHrdcPrBkecuFUuCWHK0kMJvZYCBqIa35SGxUdXVGuigQDkMpf7xO4C2C2Acl8QTtSrYS7Cnc=")
	ServicePrivKey, _ = crypto.UnmarshalPrivateKey(privRaw)
	ServiceDid, _ = did.FromPrivKey(ServicePrivKey)
}
