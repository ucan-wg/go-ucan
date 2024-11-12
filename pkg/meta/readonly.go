package meta

import (
	"github.com/ipld/go-ipld-prime"
)

// ReadOnly wraps a Meta into a read-only facade.
type ReadOnly struct {
	m *Meta
}

func (r ReadOnly) GetBool(key string) (bool, error) {
	return r.m.GetBool(key)
}

func (r ReadOnly) GetString(key string) (string, error) {
	return r.m.GetString(key)
}

func (r ReadOnly) GetInt64(key string) (int64, error) {
	return r.m.GetInt64(key)
}

func (r ReadOnly) GetFloat64(key string) (float64, error) {
	return r.m.GetFloat64(key)
}

func (r ReadOnly) GetBytes(key string) ([]byte, error) {
	return r.m.GetBytes(key)
}

func (r ReadOnly) GetNode(key string) (ipld.Node, error) {
	return r.m.GetNode(key)
}

func (r ReadOnly) Equals(other ReadOnly) bool {
	return r.m.Equals(other.m)
}

func (r ReadOnly) String() string {
	return r.m.String()
}
