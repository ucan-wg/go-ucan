package meta

import (
	"errors"
	"fmt"
	"iter"
	"sort"
	"strings"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/printer"

	"github.com/ucan-wg/go-ucan/pkg/meta/internal/crypto"
	"github.com/ucan-wg/go-ucan/pkg/policy/literal"
)

var ErrNotFound = errors.New("key not found in meta")

var ErrNotEncryptable = errors.New("value of this type cannot be encrypted")

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

// GetEncryptedString decorates GetString and decrypt its output with the given symmetric encryption key.
func (m *Meta) GetEncryptedString(key string, encryptionKey []byte) (string, error) {
	v, err := m.GetBytes(key)
	if err != nil {
		return "", err
	}

	decrypted, err := crypto.DecryptStringWithKey(v, encryptionKey)
	if err != nil {
		return "", err
	}

	return string(decrypted), nil
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

// GetEncryptedBytes decorates GetBytes and decrypt its output with the given symmetric encryption key.
func (m *Meta) GetEncryptedBytes(key string, encryptionKey []byte) ([]byte, error) {
	v, err := m.GetBytes(key)
	if err != nil {
		return nil, err
	}

	decrypted, err := crypto.DecryptStringWithKey(v, encryptionKey)
	if err != nil {
		return nil, err
	}

	return decrypted, nil
}

// GetNode retrieves a value as a raw IPLD node.
// Returns ErrNotFound if the given key is missing.
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

// AddEncrypted adds a key/value pair in the meta set.
// The value is encrypted with the given encryptionKey.
// Accepted types for the value are: string, []byte.
func (m *Meta) AddEncrypted(key string, val any, encryptionKey []byte) error {
	var encrypted []byte
	var err error

	switch val := val.(type) {
	case string:
		encrypted, err = crypto.EncryptWithKey([]byte(val), encryptionKey)
		if err != nil {
			return err
		}
	case []byte:
		encrypted, err = crypto.EncryptWithKey(val, encryptionKey)
		if err != nil {
			return err
		}
	default:
		return ErrNotEncryptable
	}

	return m.Add(key, encrypted)
}

type Iterator interface {
	Iter() iter.Seq2[string, ipld.Node]
}

// Include merges the provided meta into the existing one.
//
// If duplicate keys are encountered, the new value is silently dropped
// without causing an error.
func (m *Meta) Include(other Iterator) {
	for key, value := range other.Iter() {
		if _, ok := m.Values[key]; ok {
			// don't overwrite
			continue
		}
		m.Values[key] = value
		m.Keys = append(m.Keys, key)
	}
}

// Iter iterates over the meta key/values
func (m *Meta) Iter() iter.Seq2[string, ipld.Node] {
	return func(yield func(string, ipld.Node) bool) {
		for _, key := range m.Keys {
			if !yield(key, m.Values[key]) {
				return
			}
		}
	}
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
	sort.Strings(m.Keys)

	buf := strings.Builder{}
	buf.WriteString("{")

	for key, node := range m.Values {
		buf.WriteString("\n\t")
		buf.WriteString(key)
		buf.WriteString(": ")
		buf.WriteString(strings.ReplaceAll(printer.Sprint(node), "\n", "\n\t"))
		buf.WriteString(",")
	}

	if len(m.Values) > 0 {
		buf.WriteString("\n")
	}
	buf.WriteString("}")

	return buf.String()
}

// ReadOnly returns a read-only version of Meta.
func (m *Meta) ReadOnly() ReadOnly {
	return ReadOnly{meta: m}
}

// Clone makes a deep copy.
func (m *Meta) Clone() *Meta {
	res := &Meta{
		Keys:   make([]string, len(m.Keys)),
		Values: make(map[string]ipld.Node, len(m.Values)),
	}
	copy(res.Keys, m.Keys)
	for k, v := range m.Values {
		res.Values[k] = v
	}
	return res
}
