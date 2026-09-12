package example

import (
	"encoding/base64"

	"github.com/ucan-wg/go-did-it"
	didkeyctl "github.com/ucan-wg/go-did-it/controller/did-key"
	"github.com/ucan-wg/go-did-it/crypto"
	"github.com/ucan-wg/go-did-it/crypto/ed25519"
)

// Endpoints

var ServiceUrl = ":8080"
var ServiceIssuerUrl = ":8081"

var AliceIssuerUrl = ":8082"

// Service

var ServicePrivKey crypto.PrivateKeySigningBytesVarsig
var ServiceDid did.DID

// Alice

var AlicePrivKey crypto.PrivateKeySigningBytesVarsig
var AliceDid did.DID

func init() {
	servPrivRaw, _ := base64.StdEncoding.DecodeString("HVcbgoj30c+7zoQzUgpl7Jc7bkXoyvo9bMX5OHaAohpv036EMxuWXGqmEWhFKHPEuRAaIGSURK8pyUYOAseiiQ==")
	ServicePrivKey, _ = ed25519.PrivateKeyFromBytes(servPrivRaw)
	ServiceDid = didkeyctl.FromPrivateKey(ServicePrivKey)

	alicePrivRaw, _ := base64.StdEncoding.DecodeString("jIIk/4ZBgIzx7fU41AWYRUDjgQmgFTIXxN4WeZAPCjwE04oLfiHgNjwIIZi97a6WwSIL5tFGdkrqDkSmDx95tw==")
	AlicePrivKey, _ = ed25519.PrivateKeyFromBytes(alicePrivRaw)
	AliceDid = didkeyctl.FromPrivateKey(AlicePrivKey)
}
