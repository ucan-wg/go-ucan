package nonce

import "crypto/rand"

// Generate creates a 12-byte random nonce.
// TODO: some crypto scheme require more, is that our case?
func Generate() ([]byte, error) {
	res := make([]byte, 12)
	_, err := rand.Read(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
