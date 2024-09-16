package policy

import (
	"cmp"
	"fmt"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/must"

	"github.com/ucan-wg/go-ucan/capability/policy/selector"
)

// Match determines if the IPLD node matches the policy document.
func Match(policy Policy, node ipld.Node) bool {
	for _, stmt := range policy {
		ok := matchStatement(stmt, node)
		if !ok {
			return false
		}
	}
	return true
}

func matchStatement(statement Statement, node ipld.Node) bool {
	switch statement.Kind() {
	case KindEqual:
		if s, ok := statement.(equality); ok {
			one, many, err := selector.Select(s.selector, node)
			if err != nil {
				return false
			}
			if one != nil {
				return datamodel.DeepEqual(s.value, one)
			}
			if many != nil {
				for _, n := range many {
					if eq := datamodel.DeepEqual(s.value, n); eq {
						return true
					}
				}
			}

			return false
		}
	case KindGreaterThan:
		if s, ok := statement.(equality); ok {
			one, _, err := selector.Select(s.selector, node)
			if err != nil || one == nil {
				return false
			}
			return isOrdered(s.value, one, gt)
		}
	case KindGreaterThanOrEqual:
		if s, ok := statement.(equality); ok {
			one, _, err := selector.Select(s.selector, node)
			if err != nil || one == nil {
				return false
			}
			return isOrdered(s.value, one, gte)
		}
	case KindLessThan:
		if s, ok := statement.(equality); ok {
			one, _, err := selector.Select(s.selector, node)
			if err != nil || one == nil {
				return false
			}
			return isOrdered(s.value, one, lt)
		}
	case KindLessThanOrEqual:
		if s, ok := statement.(equality); ok {
			one, _, err := selector.Select(s.selector, node)
			if err != nil || one == nil {
				return false
			}
			return isOrdered(s.value, one, lte)
		}
	case KindNot:
		if s, ok := statement.(negation); ok {
			return !matchStatement(s.statement, node)
		}
	case KindAnd:
		if s, ok := statement.(connective); ok {
			for _, cs := range s.statements {
				r := matchStatement(cs, node)
				if !r {
					return false
				}
			}
			return true
		}
	case KindOr:
		if s, ok := statement.(connective); ok {
			if len(s.statements) == 0 {
				return true
			}
			for _, cs := range s.statements {
				r := matchStatement(cs, node)
				if r {
					return true
				}
			}
			return false
		}
	case KindLike:
		if s, ok := statement.(wildcard); ok {
			one, _, err := selector.Select(s.selector, node)
			if err != nil || one == nil {
				return false
			}
			v, err := one.AsString()
			if err != nil {
				return false
			}
			return globMatch(s.pattern, v)
		}
	case KindAll:
		if s, ok := statement.(quantifier); ok {
			_, many, err := selector.Select(s.selector, node)
			if err != nil || many == nil {
				return false
			}
			for _, n := range many {
				ok := matchStatement(s.statement, n)
				if !ok {
					return false
				}
			}
			return true
		}
	case KindAny:
		if s, ok := statement.(quantifier); ok {
			one, many, err := selector.Select(s.selector, node)
			if err != nil {
				return false
			}
			if one != nil {
				ok := matchStatement(s.statement, one)
				if ok {
					return true
				}
			}
			if many != nil {
				for _, n := range many {
					ok := matchStatement(s.statement, n)
					if ok {
						return true
					}
				}
			}

			return false
		}
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
