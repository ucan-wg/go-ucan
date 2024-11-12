package meta

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/printer"

	"github.com/ucan-wg/go-ucan/pkg/meta/internal/crypto"


	"github.com/ucan-wg/go-ucan/pkg/policy/literal"
)
var ErrUnsupported = errors.New("failure adding unsupported type to meta")

var ErrNotFound = errors.New("key-value not found in meta")

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
	v, err := m.GetString(key)
	if err != nil {
		return "", err
	}

	decrypted, err := crypto.DecryptStringWithAESKey([]byte(v), encryptionKey)
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

	decrypted, err := crypto.DecryptStringWithAESKey(v, encryptionKey)
	if err != nil {
		return nil, err
	}

	return decrypted, nil
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

// AddEncrypted adds a key/value pair in the meta set.
// The value is encrypted with the given encryptionKey.
// Accepted types for the value are: string, []byte.
func (m *Meta) AddEncrypted(key string, val any, encryptionKey []byte) error {
	var encrypted []byte
	var err error

	switch val := val.(type) {
	case string:
		encrypted, err = crypto.EncryptWithAESKey([]byte(val), encryptionKey)
		if err != nil {
			return err
		}
		return m.Add(key, string(encrypted))
	case []byte:
		encrypted, err = crypto.EncryptWithAESKey(val, encryptionKey)
		if err != nil {
			return err
		}
		return m.Add(key, encrypted)
	default:
		return ErrNotEncryptable
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
