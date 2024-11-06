package meta

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	"github.com/ipld/go-ipld-prime/printer"
)

var ErrUnsupported = fmt.Errorf("failure adding unsupported type to meta")

var ErrNotFound = fmt.Errorf("key-value not found in meta")

// Meta is a container for meta key-value pairs in a UCAN token.
// This also serves as a way to construct the underlying IPLD data with minimum allocations and transformations,
// while hiding the IPLD complexity from the caller.
type Meta struct {
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
// Accepted types for the value are: bool, string, int, int32, int64, []byte,
// and ipld.Node.
func (m *Meta) Add(key string, val any) error {
	if _, ok := m.Values[key]; ok {
		return fmt.Errorf("duplicate key %q", key)
	}
	switch val := val.(type) {
	case bool:
		m.Values[key] = basicnode.NewBool(val)
	case string:
		m.Values[key] = basicnode.NewString(val)
	case int:
		m.Values[key] = basicnode.NewInt(int64(val))
	case int32:
		m.Values[key] = basicnode.NewInt(int64(val))
	case int64:
		m.Values[key] = basicnode.NewInt(val)
	case float32:
		m.Values[key] = basicnode.NewFloat(float64(val))
	case float64:
		m.Values[key] = basicnode.NewFloat(val)
	case []byte:
		m.Values[key] = basicnode.NewBytes(val)
	case datamodel.Node:
		m.Values[key] = val
	default:
		return fmt.Errorf("%w: %s", ErrUnsupported, fqtn(val))
	}
	m.Keys = append(m.Keys, key)
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

func fqtn(val any) string {
	var name string

	t := reflect.TypeOf(val)
	for t.Kind() == reflect.Pointer {
		name += "*"
		t = t.Elem()
	}

	return name + t.PkgPath() + "." + t.Name()
}
