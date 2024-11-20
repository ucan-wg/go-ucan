package nonce

import "crypto/rand"

// TODO: some crypto scheme require more, is that our case?
//
// The spec mention:
// The REQUIRED nonce parameter nonce MAY be any value.
// A randomly generated string is RECOMMENDED to provide a unique UCAN, though it MAY
// also be a monotonically increasing count of the number of links in the hash chain.
// This field helps prevent replay attacks and ensures a unique CID per delegation.
// The iss, aud, and exp fields together will often ensure that UCANs are unique,
// but adding the nonce ensures uniqueness.
//
// The recommended size of the nonce differs by key type. In many cases, a random
// 12-byte nonce is sufficient. If uncertain, check the nonce in your DID's crypto suite.
//
// 12 bytes is 10^28, 16 bytes is 10^38. Both sounds like a lot of random to achieve
// those goals, but maybe the crypto voodoo require more.
//
// The rust implementation use 16 bytes nonce.

// Generate creates a 12-byte random nonce.
func Generate() ([]byte, error) {
	res := make([]byte, 12)
	_, err := rand.Read(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
