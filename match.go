package policy

import (
	"cmp"
	"fmt"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/must"
	"github.com/storacha-network/go-ucanto/core/policy/selector"
)

// Match determines if the IPLD node matches the policy document.
func Match(policy Policy, node ipld.Node) bool {
	for _, stmt := range policy {
		ok := matchStatement(stmt, node)
		if !ok {
			return ok
		}
	}
	return true
}

func matchStatement(statement Statement, node ipld.Node) bool {
	switch statement.Kind() {
	case Kind_Equal:
		if s, ok := statement.(EqualityStatement); ok {
			one, _, err := selector.Select(s.Selector(), node)
			if err != nil || one == nil {
				return false
			}
			return datamodel.DeepEqual(s.Value(), one)
		}
	case Kind_GreaterThan:
		if s, ok := statement.(InequalityStatement); ok {
			one, _, err := selector.Select(s.Selector(), node)
			if err != nil || one == nil {
				return false
			}
			return isOrdered(s.Value(), one, gt)
		}
	case Kind_GreaterThanOrEqual:
		if s, ok := statement.(InequalityStatement); ok {
			one, _, err := selector.Select(s.Selector(), node)
			if err != nil || one == nil {
				return false
			}
			return isOrdered(s.Value(), one, gte)
		}
	case Kind_LessThan:
		if s, ok := statement.(InequalityStatement); ok {
			one, _, err := selector.Select(s.Selector(), node)
			if err != nil || one == nil {
				return false
			}
			return isOrdered(s.Value(), one, lt)
		}
	case Kind_LessThanOrEqual:
		if s, ok := statement.(InequalityStatement); ok {
			one, _, err := selector.Select(s.Selector(), node)
			if err != nil || one == nil {
				return false
			}
			return isOrdered(s.Value(), one, lte)
		}
	case Kind_Not:
		if s, ok := statement.(NegationStatement); ok {
			return !matchStatement(s.Value(), node)
		}
	case Kind_And:
		if s, ok := statement.(ConjunctionStatement); ok {
			for _, cs := range s.Value() {
				r := matchStatement(cs, node)
				if !r {
					return false
				}
			}
			return true
		}
	case Kind_Or:
		if s, ok := statement.(DisjunctionStatement); ok {
			if len(s.Value()) == 0 {
				return true
			}
			for _, cs := range s.Value() {
				r := matchStatement(cs, node)
				if r {
					return true
				}
			}
			return false
		}
	case Kind_Like:
		if s, ok := statement.(WildcardStatement); ok {
			one, _, err := selector.Select(s.Selector(), node)
			if err != nil || one == nil {
				return false
			}
			v, err := one.AsString()
			if err != nil {
				return false
			}
			return s.Value().Match(v)
		}
	case Kind_All:
	case Kind_Any:
	}
	panic(fmt.Errorf("unimplemented statement kind: %s", statement.Kind()))
}

func isOrdered(expected ipld.Node, actual ipld.Node, satisfies func(order int) bool) bool {
	if expected.Kind() == ipld.Kind_Int && actual.Kind() == ipld.Kind_Int {
		a := must.Int(actual)
		b := must.Int(expected)
		return satisfies(cmp.Compare(a, b))
	}

	if expected.Kind() == ipld.Kind_Float && actual.Kind() == ipld.Kind_Float {
		a, err := actual.AsFloat()
		if err != nil {
			panic(fmt.Errorf("extracting node float: %w", err))
		}
		b, err := expected.AsFloat()
		if err != nil {
			panic(fmt.Errorf("extracting selector float: %w", err))
		}
		return satisfies(cmp.Compare(a, b))
	}

	return false
}

func gt(order int) bool  { return order == 1 }
func gte(order int) bool { return order == 0 || order == 1 }
func lt(order int) bool  { return order == -1 }
func lte(order int) bool { return order == 0 || order == -1 }
