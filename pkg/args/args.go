// Package args provides the type that represents the Arguments passed to
// a command within an invocation.Token as well as a convenient Add method
// to incrementally build the underlying map.
package args

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/bindnode"
	"github.com/ipld/go-ipld-prime/schema"
	"github.com/ucan-wg/go-ucan/pkg/policy/literal"
)

const (
	argsSchema = "type Args { String : Any }"
	argsName   = "Args"
)

var (
	once sync.Once
	ts   *schema.TypeSystem
	err  error
)

func argsType() schema.Type {
	once.Do(func() {
		ts, err = ipld.LoadSchemaBytes([]byte(argsSchema))
	})
	if err != nil {
		panic(err)
	}

	return ts.TypeByName(argsName)
}

var ErrUnsupported = fmt.Errorf("failure adding unsupported type to meta")

// Args are the Command's argumennts when an invocation Token is processed
// by the executor.
//
// This type must be compatible with the IPLD type represented by the IPLD
// schema { String : Any }.
type Args struct {
	Keys   []string
	Values map[string]ipld.Node
}

// New returns a pointer to an initialized Args value.
func New() *Args {
	return &Args{
		Values: map[string]ipld.Node{},
	}
}

// FromIPLD unwraps an Args instance from an ipld.Node.
func FromIPLD(node ipld.Node) (*Args, error) {
	var err error

	defer func() {
		err = handlePanic(recover())
	}()

	obj := bindnode.Unwrap(node)

	args, ok := obj.(*Args)
	if !ok {
		err = fmt.Errorf("failed to convert to Args")
	}

	return args, err
}

// Add inserts a key/value pair in the Args set.
//
// Accepted types for val are: bool, string, int, int8, int16,
// int32, int64, uint, uint8, uint16, uint32, float32, float64, []byte,
// []any, map[string]any, ipld.Node and nil.
func (m *Args) Add(key string, val any) error {
	if _, ok := m.Values[key]; ok {
		return fmt.Errorf("duplicate key %q", key)
	}

	node, err := anyNode(val)
	if err != nil {
		return err
	}

	m.Values[key] = node
	m.Keys = append(m.Keys, key)

	return nil
}

// Include merges the provided arguments into the existing arguments.
//
// If duplicate keys are encountered, the new value is silently dropped
// without causing an error.
func (m *Args) Include(other *Args) {
	for _, key := range other.Keys {
		if _, ok := m.Values[key]; ok {
			// don't overwrite
			continue
		}
		m.Values[key] = other.Values[key]
		m.Keys = append(m.Keys, key)
	}
}

// ToIPLD wraps an instance of an Args with an ipld.Node.
func (m *Args) ToIPLD() (ipld.Node, error) {
	var err error

	defer func() {
		err = handlePanic(recover())
	}()

	return bindnode.Wrap(m, argsType()), err
}

func anyNode(val any) (ipld.Node, error) {
	var err error

	defer func() {
		err = handlePanic(recover())
	}()

	if val == nil {
		return literal.Null(), nil
	}

	if cast, ok := val.(ipld.Node); ok {
		return cast, nil
	}

	if cast, ok := val.(cid.Cid); ok {
		return literal.LinkCid(cast), err
	}

	var rv reflect.Value

	rv.Kind()

	if cast, ok := val.(reflect.Value); ok {
		rv = cast
	} else {
		rv = reflect.ValueOf(val)
	}

	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		rv = rv.Elem()
	}

	switch rv.Kind() {
	case reflect.Slice:
		if rv.Type().Elem().Kind() == reflect.Uint8 {
			return literal.Bytes(val.([]byte)), nil
		}

		l := make([]reflect.Value, rv.Len())

		for i := 0; i < rv.Len(); i++ {
			l[i] = rv.Index(i)
		}

		return literal.List(l)
	case reflect.Map:
		if rv.Type().Key().Kind() != reflect.String {
			return nil, fmt.Errorf("unsupported map key type: %s", rv.Type().Key().Name())
		}

		m := make(map[string]reflect.Value, rv.Len())
		it := rv.MapRange()

		for it.Next() {
			m[it.Key().String()] = it.Value()
		}

		return literal.Map(m)
	case reflect.String:
		return literal.String(rv.String()), nil
	case reflect.Bool:
		return literal.Bool(rv.Bool()), nil
		// reflect.Int64 may exceed the safe 53-bit limit of JavaScript
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return literal.Int(rv.Int()), nil
	// reflect.Uint64 can't be safely converted to int64
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
		return literal.Int(int64(rv.Uint())), nil
	case reflect.Float32, reflect.Float64:
		return literal.Float(rv.Float()), nil
	default:
		return nil, fmt.Errorf("unsupported Args type: %s", rv.Type().Name())
	}
}

func handlePanic(rec any) error {
	if err, ok := rec.(error); ok {
		return err
	}

	return fmt.Errorf("%v", rec)
}
