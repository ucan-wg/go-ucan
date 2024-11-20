// Package args provides the type that represents the Arguments passed to
// a command within an invocation.Token as well as a convenient Add method
// to incrementally build the underlying map.
package args

import (
	"fmt"
	"iter"
	"sort"
	"strings"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/fluent/qp"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	"github.com/ipld/go-ipld-prime/printer"

	"github.com/ucan-wg/go-ucan/pkg/policy/literal"
)

// Args are the Command's arguments when an invocation Token is processed by the executor.
// This also serves as a way to construct the underlying IPLD data with minimum allocations
// and transformations, while hiding the IPLD complexity from the caller.
type Args struct {
	// This type must be compatible with the IPLD type represented by the IPLD
	// schema { String : Any }.

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
// Accepted types for val are any CBOR compatible type, or directly IPLD values.
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

type Iterator interface {
	Iter() iter.Seq2[string, ipld.Node]
}

// Include merges the provided arguments into the existing arguments.
//
// If duplicate keys are encountered, the new value is silently dropped
// without causing an error.
func (a *Args) Include(other Iterator) {
	for key, value := range other.Iter() {
		if _, ok := a.Values[key]; ok {
			// don't overwrite
			continue
		}
		a.Values[key] = value
		a.Keys = append(a.Keys, key)
	}
}

// Iter iterates over the args key/values
func (a *Args) Iter() iter.Seq2[string, ipld.Node] {
	return func(yield func(string, ipld.Node) bool) {
		for _, key := range a.Keys {
			if !yield(key, a.Values[key]) {
				return
			}
		}
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

func (a *Args) String() string {
	sort.Strings(a.Keys)

	buf := strings.Builder{}
	buf.WriteString("{")

	for _, key := range a.Keys {
		buf.WriteString("\n\t")
		buf.WriteString(key)
		buf.WriteString(": ")
		buf.WriteString(strings.ReplaceAll(printer.Sprint(a.Values[key]), "\n", "\n\t"))
		buf.WriteString(",")
	}

	if len(a.Keys) > 0 {
		buf.WriteString("\n")
	}
	buf.WriteString("}")

	return buf.String()
}

// ReadOnly returns a read-only version of Args.
func (a *Args) ReadOnly() ReadOnly {
	return ReadOnly{args: a}
}

// Clone makes a deep copy.
func (a *Args) Clone() *Args {
	res := &Args{
		Keys:   make([]string, len(a.Keys)),
		Values: make(map[string]ipld.Node, len(a.Values)),
	}
	copy(res.Keys, a.Keys)
	for k, v := range a.Values {
		res.Values[k] = v
	}
	return res
}
