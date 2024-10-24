package policy

import (
	"cmp"
	"fmt"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/must"
)

// Match determines if the IPLD node satisfies the policy.
func (p Policy) Match(node datamodel.Node) bool {
	for _, stmt := range p {
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
			res, err := s.selector.Select(node)
			if err != nil {
				return false
			}
			return datamodel.DeepEqual(s.value, res)
		}
	case KindGreaterThan:
		if s, ok := statement.(equality); ok {
			res, err := s.selector.Select(node)
			if err != nil {
				return false
			}
			return isOrdered(s.value, res, gt)
		}
	case KindGreaterThanOrEqual:
		if s, ok := statement.(equality); ok {
			res, err := s.selector.Select(node)
			if err != nil {
				return false
			}
			return isOrdered(s.value, res, gte)
		}
	case KindLessThan:
		if s, ok := statement.(equality); ok {
			res, err := s.selector.Select(node)
			if err != nil {
				return false
			}
			return isOrdered(s.value, res, lt)
		}
	case KindLessThanOrEqual:
		if s, ok := statement.(equality); ok {
			res, err := s.selector.Select(node)
			if err != nil {
				return false
			}
			return isOrdered(s.value, res, lte)
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
			res, err := s.selector.Select(node)
			if err != nil {
				return false
			}
			v, err := res.AsString()
			if err != nil {
				return false // not a string
			}
			return s.pattern.Match(v)
		}
	case KindAll:
		if s, ok := statement.(quantifier); ok {
			res, err := s.selector.Select(node)
			if err != nil {
				return false
			}
			it := res.ListIterator()
			if it == nil {
				return false // not a list
			}
			for !it.Done() {
				_, v, err := it.Next()
				if err != nil {
					return false
				}
				ok := matchStatement(s.statement, v)
				if !ok {
					return false
				}
			}
			return true
		}
	case KindAny:
		if s, ok := statement.(quantifier); ok {
			res, err := s.selector.Select(node)
			if err != nil {
				return false
			}
			it := res.ListIterator()
			if it == nil {
				return false // not a list
			}
			for !it.Done() {
				_, v, err := it.Next()
				if err != nil {
					return false
				}
				ok := matchStatement(s.statement, v)
				if ok {
					return true
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
