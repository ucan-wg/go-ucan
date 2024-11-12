// Package args provides the type that represents the Arguments passed to
// a command within an invocation.Token as well as a convenient Add method
// to incrementally build the underlying map.
package args

import (
	"fmt"
	"sort"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/fluent/qp"
	"github.com/ipld/go-ipld-prime/node/basicnode"

	"github.com/ucan-wg/go-ucan/pkg/policy/literal"
)

// Args are the Command's arguments when an invocation Token is processed
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

// Add inserts a key/value pair in the Args set.
//
// Accepted types for val are: bool, string, int, int8, int16,
// int32, int64, uint, uint8, uint16, uint32, float32, float64, []byte,
// []any, map[string]any, ipld.Node and nil.
func (a *Args) Add(key string, val any) error {
	if _, ok := a.Values[key]; ok {
		return fmt.Errorf("duplicate key %q", key)
	}

	node, err := literal.Any(val)
	if err != nil {
		return err
	}

	a.Values[key] = node
	a.Keys = append(a.Keys, key)

	return nil
}

// Include merges the provided arguments into the existing arguments.
//
// If duplicate keys are encountered, the new value is silently dropped
// without causing an error.
func (a *Args) Include(other *Args) {
	for _, key := range other.Keys {
		if _, ok := a.Values[key]; ok {
			// don't overwrite
			continue
		}
		a.Values[key] = other.Values[key]
		a.Keys = append(a.Keys, key)
	}
}

// ToIPLD wraps an instance of an Args with an ipld.Node.
func (a *Args) ToIPLD() (ipld.Node, error) {
	sort.Strings(a.Keys)
	return qp.BuildMap(basicnode.Prototype.Any, int64(len(a.Keys)), func(ma datamodel.MapAssembler) {
		for _, key := range a.Keys {
			qp.MapEntry(ma, key, qp.Node(a.Values[key]))
		}
	})
}

// Equals tells if two Args hold the same values.
func (a *Args) Equals(other *Args) bool {
	if len(a.Keys) != len(other.Keys) {
		return false
	}
	if len(a.Values) != len(other.Values) {
		return false
	}
	for _, key := range a.Keys {
		if !ipld.DeepEqual(a.Values[key], other.Values[key]) {
			return false
		}
	}
	return true
}
