package meta

import (
	"fmt"
	"strings"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/printer"

	"github.com/ucan-wg/go-ucan/pkg/policy/literal"
)

var ErrNotFound = fmt.Errorf("key-value not found in meta")

// Meta is a container for meta key-value pairs in a UCAN token.
// This also serves as a way to construct the underlying IPLD data with minimum allocations
// and transformations, while hiding the IPLD complexity from the caller.
type Meta struct {
	// This type must be compatible with the IPLD type represented by the IPLD
	// schema { String : Any }.

	Keys   []string
	Values map[string]ipld.Node
}

// NewMeta constructs a new Meta.
func NewMeta() *Meta {
	return &Meta{Values: map[string]ipld.Node{}}
}

// GetBool retrieves a value as a bool.
// Returns ErrNotFound if the given key is missing.
// Returns datamodel.ErrWrongKind if the value has the wrong type.
func (m *Meta) GetBool(key string) (bool, error) {
	v, ok := m.Values[key]
	if !ok {
		return false, ErrNotFound
	}
	return v.AsBool()
}

// GetString retrieves a value as a string.
// Returns ErrNotFound if the given key is missing.
// Returns datamodel.ErrWrongKind if the value has the wrong type.
func (m *Meta) GetString(key string) (string, error) {
	v, ok := m.Values[key]
	if !ok {
		return "", ErrNotFound
	}
	return v.AsString()
}

// GetInt64 retrieves a value as an int64.
// Returns ErrNotFound if the given key is missing.
// Returns datamodel.ErrWrongKind if the value has the wrong type.
func (m *Meta) GetInt64(key string) (int64, error) {
	v, ok := m.Values[key]
	if !ok {
		return 0, ErrNotFound
	}
	return v.AsInt()
}

// GetFloat64 retrieves a value as a float64.
// Returns ErrNotFound if the given key is missing.
// Returns datamodel.ErrWrongKind if the value has the wrong type.
func (m *Meta) GetFloat64(key string) (float64, error) {
	v, ok := m.Values[key]
	if !ok {
		return 0, ErrNotFound
	}
	return v.AsFloat()
}

// GetBytes retrieves a value as a []byte.
// Returns ErrNotFound if the given key is missing.
// Returns datamodel.ErrWrongKind if the value has the wrong type.
func (m *Meta) GetBytes(key string) ([]byte, error) {
	v, ok := m.Values[key]
	if !ok {
		return nil, ErrNotFound
	}
	return v.AsBytes()
}

// GetNode retrieves a value as a raw IPLD node.
// Returns ErrNotFound if the given key is missing.
// Returns datamodel.ErrWrongKind if the value has the wrong type.
func (m *Meta) GetNode(key string) (ipld.Node, error) {
	v, ok := m.Values[key]
	if !ok {
		return nil, ErrNotFound
	}
	return v, nil
}

// Add adds a key/value pair in the meta set.
// Accepted types for val are any CBOR compatible type, or directly IPLD values.
func (m *Meta) Add(key string, val any) error {
	if _, ok := m.Values[key]; ok {
		return fmt.Errorf("duplicate key %q", key)
	}

	node, err := literal.Any(val)
	if err != nil {
		return err
	}

	m.Keys = append(m.Keys, key)
	m.Values[key] = node

	return nil
}

// Equals tells if two Meta hold the same key/values.
func (m *Meta) Equals(other *Meta) bool {
	if len(m.Keys) != len(other.Keys) {
		return false
	}
	if len(m.Values) != len(other.Values) {
		return false
	}
	for _, key := range m.Keys {
		if !ipld.DeepEqual(m.Values[key], other.Values[key]) {
			return false
		}
	}
	return true
}

func (m *Meta) String() string {
	buf := strings.Builder{}
	buf.WriteString("{")

	var i int
	for key, node := range m.Values {
		if i > 0 {
			buf.WriteString(", ")
		}
		i++
		buf.WriteString(key)
		buf.WriteString(":")
		buf.WriteString(printer.Sprint(node))
	}

	buf.WriteString("}")
	return buf.String()
}

// ReadOnly returns a read-only version of Meta.
func (m *Meta) ReadOnly() ReadOnly {
	return ReadOnly{m: m}
}
