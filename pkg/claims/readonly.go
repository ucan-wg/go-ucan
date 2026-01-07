package claims

import (
	"iter"

	"github.com/ipld/go-ipld-prime"
)

// ReadOnly wraps a Claims into a read-only facade.
type ReadOnly struct {
	claims *Claims
}

func (r ReadOnly) GetBool(key string) (bool, error) {
	return r.claims.GetBool(key)
}

func (r ReadOnly) GetString(key string) (string, error) {
	return r.claims.GetString(key)
}

func (r ReadOnly) GetEncryptedString(key string, encryptionKey []byte) (string, error) {
	return r.claims.GetEncryptedString(key, encryptionKey)
}

func (r ReadOnly) GetInt64(key string) (int64, error) {
	return r.claims.GetInt64(key)
}

func (r ReadOnly) GetFloat64(key string) (float64, error) {
	return r.claims.GetFloat64(key)
}

func (r ReadOnly) GetBytes(key string) ([]byte, error) {
	return r.claims.GetBytes(key)
}

func (r ReadOnly) GetEncryptedBytes(key string, encryptionKey []byte) ([]byte, error) {
	return r.claims.GetEncryptedBytes(key, encryptionKey)
}

func (r ReadOnly) GetNode(key string) (ipld.Node, error) {
	return r.claims.GetNode(key)
}

func (r ReadOnly) Len() int {
	return r.claims.Len()
}

func (r ReadOnly) Iter() iter.Seq2[string, ipld.Node] {
	return r.claims.Iter()
}

func (r ReadOnly) Equals(other ReadOnly) bool {
	return r.claims.Equals(other.claims)
}

func (r ReadOnly) String() string {
	return r.claims.String()
}

func (r ReadOnly) WriteableClone() *Claims {
	return r.claims.Clone()
}
