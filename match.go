package policy

import (
	"cmp"
	"fmt"

	"github.com/ipld/go-ipld-prime"
	"github.com/storacha-network/go-ucanto/core/policy/selector"
)

// Match determines if the IPLD node matches the policy document.
func Match(policy Policy, node ipld.Node) (bool, error) {
	for _, stmt := range policy {
		ok, err := matchStatement(stmt, node)
		if err != nil || !ok {
			return ok, err
		}
	}
	return true, nil
}

func matchStatement(statement Statement, node ipld.Node) (bool, error) {
	switch statement.Kind() {
	case Kind_Equal:
		if s, ok := statement.(EqualityStatement); ok {
			one, _, err := selector.Select(s.Selector(), node)
			if err != nil || one == nil {
				return false, nil
			}
			return isDeepEqual(s.Value(), one)
		}
	case Kind_GreaterThan:
		if s, ok := statement.(InequalityStatement); ok {
			one, _, err := selector.Select(s.Selector(), node)
			if err != nil || one == nil {
				return false, nil
			}
			return isOrdered(s.Value(), one, gt)
		}
	case Kind_GreaterThanOrEqual:
		if s, ok := statement.(InequalityStatement); ok {
			one, _, err := selector.Select(s.Selector(), node)
			if err != nil || one == nil {
				return false, nil
			}
			return isOrdered(s.Value(), one, gte)
		}
	case Kind_LessThan:
		if s, ok := statement.(InequalityStatement); ok {
			one, _, err := selector.Select(s.Selector(), node)
			if err != nil || one == nil {
				return false, nil
			}
			return isOrdered(s.Value(), one, lt)
		}
	case Kind_LessThanOrEqual:
		if s, ok := statement.(InequalityStatement); ok {
			one, _, err := selector.Select(s.Selector(), node)
			if err != nil || one == nil {
				return false, nil
			}
			return isOrdered(s.Value(), one, lte)
		}
	case Kind_Negation:
		if s, ok := statement.(NegationStatement); ok {
			r, err := matchStatement(s.Value(), node)
			if err != nil {
				return false, err
			}
			return !r, err
		}
	case Kind_Conjunction:
		if s, ok := statement.(ConjunctionStatement); ok {
			for _, cs := range s.Value() {
				r, err := matchStatement(cs, node)
				if err != nil {
					return false, err
				}
				if !r {
					return false, nil
				}
			}
			return true, nil
		}
	case Kind_Disjunction:
		if s, ok := statement.(DisjunctionStatement); ok {
			if len(s.Value()) == 0 {
				return true, nil
			}
			for _, cs := range s.Value() {
				r, err := matchStatement(cs, node)
				if err != nil {
					return false, err
				}
				if r {
					return true, nil
				}
			}
			return false, nil
		}
	case Kind_Wildcard:
	case Kind_Universal:
	case Kind_Existential:
	}
	return false, fmt.Errorf("unimplemented statement kind: %s", statement.Kind())
}

func isOrdered(expected ipld.Node, actual ipld.Node, satisfies func(order int) bool) (bool, error) {
	if expected.Kind() == ipld.Kind_Int && actual.Kind() == ipld.Kind_Int {
		a, err := actual.AsInt()
		if err != nil {
			return false, fmt.Errorf("extracting node int: %w", err)
		}
		b, err := expected.AsInt()
		if err != nil {
			return false, fmt.Errorf("extracting selector int: %w", err)
		}
		return satisfies(cmp.Compare(a, b)), nil
	}

	if expected.Kind() == ipld.Kind_Float && actual.Kind() == ipld.Kind_Float {
		a, err := actual.AsFloat()
		if err != nil {
			return false, fmt.Errorf("extracting node float: %w", err)
		}
		b, err := expected.AsFloat()
		if err != nil {
			return false, fmt.Errorf("extracting selector float: %w", err)
		}
		return satisfies(cmp.Compare(a, b)), nil
	}

	return false, fmt.Errorf("unsupported IPLD kinds in ordered comparison: %s %s", expected.Kind(), actual.Kind())
}

func isDeepEqual(expected ipld.Node, actual ipld.Node) (bool, error) {
	if expected.Kind() != actual.Kind() {
		return false, nil
	}
	// TODO: should be easy enough to do the basic types, map, struct and list
	// might be harder.
	switch expected.Kind() {
	case ipld.Kind_String:
		a, err := actual.AsString()
		if err != nil {
			return false, fmt.Errorf("extracting node string: %w", err)
		}
		b, err := expected.AsString()
		if err != nil {
			return false, fmt.Errorf("extracting selector string: %w", err)
		}
		return a == b, nil
	case ipld.Kind_Int:
		if actual.Kind() != ipld.Kind_Int {
			return false, nil
		}
		a, err := actual.AsInt()
		if err != nil {
			return false, fmt.Errorf("extracting node int: %w", err)
		}
		b, err := expected.AsInt()
		if err != nil {
			return false, fmt.Errorf("extracting selector int: %w", err)
		}
		return a == b, nil
	case ipld.Kind_Float:
		if actual.Kind() != ipld.Kind_Float {
			return false, nil
		}
		a, err := actual.AsFloat()
		if err != nil {
			return false, fmt.Errorf("extracting node float: %w", err)
		}
		b, err := expected.AsFloat()
		if err != nil {
			return false, fmt.Errorf("extracting selector float: %w", err)
		}
		return a == b, nil
	case ipld.Kind_Bool:
		a, err := actual.AsBool()
		if err != nil {
			return false, fmt.Errorf("extracting node boolean: %w", err)
		}
		b, err := expected.AsBool()
		if err != nil {
			return false, fmt.Errorf("extracting selector node boolean: %w", err)
		}
		return a == b, nil
	case ipld.Kind_Link:
		a, err := actual.AsLink()
		if err != nil {
			return false, fmt.Errorf("extracting node link: %w", err)
		}
		b, err := expected.AsLink()
		if err != nil {
			return false, fmt.Errorf("extracting selector node link: %w", err)
		}
		return a.Binary() == b.Binary(), nil
	}
	return false, fmt.Errorf("unsupported IPLD kind in equality comparison: %s", expected.Kind())
}

func gt(order int) bool  { return order == 1 }
func gte(order int) bool { return order == 0 || order == 1 }
func lt(order int) bool  { return order == -1 }
func lte(order int) bool { return order == 0 || order == -1 }
