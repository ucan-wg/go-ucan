package meta

import (
	"iter"

	"github.com/ipld/go-ipld-prime"
)

// ReadOnly wraps a Meta into a read-only facade.
type ReadOnly struct {
	meta *Meta
}

func (r ReadOnly) GetBool(key string) (bool, error) {
	return r.meta.GetBool(key)
}

func (r ReadOnly) GetString(key string) (string, error) {
	return r.meta.GetString(key)
}

func (r ReadOnly) GetEncryptedString(key string, encryptionKey []byte) (string, error) {
	return r.meta.GetEncryptedString(key, encryptionKey)
}

func (r ReadOnly) GetInt64(key string) (int64, error) {
	return r.meta.GetInt64(key)
}

func (r ReadOnly) GetFloat64(key string) (float64, error) {
	return r.meta.GetFloat64(key)
}

func (r ReadOnly) GetBytes(key string) ([]byte, error) {
	return r.meta.GetBytes(key)
}

func (r ReadOnly) GetEncryptedBytes(key string, encryptionKey []byte) ([]byte, error) {
	return r.meta.GetEncryptedBytes(key, encryptionKey)
}

func (r ReadOnly) GetNode(key string) (ipld.Node, error) {
	return r.meta.GetNode(key)
}

func (r ReadOnly) Iter() iter.Seq2[string, ipld.Node] {
	return r.meta.Iter()
}

func (r ReadOnly) Equals(other ReadOnly) bool {
	return r.meta.Equals(other.meta)
}

func (r ReadOnly) String() string {
	return r.meta.String()
}

func (r ReadOnly) WriteableClone() *Meta {
	return r.meta.Clone()
}
