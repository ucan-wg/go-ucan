package example

import (
	"encoding/base64"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/ucan-wg/go-ucan/did"
)

// Endpoints

var ServiceUrl = ":8080"
var ServiceIssuerUrl = ":8081"

var AliceIssuerUrl = ":8082"

// Service

var ServicePrivKey crypto.PrivKey
var ServiceDid did.DID

// Alice

var AlicePrivKey crypto.PrivKey
var AliceDid did.DID

func init() {
	servPrivRaw, _ := base64.StdEncoding.DecodeString("CAESQGs7hPBRBmxH1UmHrdcPrBkecuFUuCWHK0kMJvZYCBqIa35SGxUdXVGuigQDkMpf7xO4C2C2Acl8QTtSrYS7Cnc=")
	ServicePrivKey, _ = crypto.UnmarshalPrivateKey(servPrivRaw)
	ServiceDid, _ = did.FromPrivKey(ServicePrivKey)

	alicePrivRaw, _ := base64.StdEncoding.DecodeString("CAESQFESA31nDYUhXXwbCNSFvg7M+TOFgyxy0tVX6o+TkJAKqAwDvtGxZeGyUjibGd/op+xOLvzE6BrTIOw62K3yLp8=")
	AlicePrivKey, _ = crypto.UnmarshalPrivateKey(alicePrivRaw)
	AliceDid, _ = did.FromPrivKey(AlicePrivKey)
}
