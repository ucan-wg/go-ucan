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

// // Filter performs a recursive filtering of the Statement, and prunes what doesn't match the given path
// func (p Policy) Filter(path ...string) Policy {
// 	var filtered Policy
//
// 	for _, stmt := range p {
// 		newChild, remain := filter(stmt, path)
// 		if newChild != nil && len(remain) == 0 {
// 			filtered = append(filtered, newChild)
// 		}
// 	}
//
// 	return filtered
// }

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

// // filter performs a recursive filtering of the Statement, and prunes what doesn't match the given path
// func filter(stmt Statement, path []string) (Statement, []string) {
// 	// For each kind, we do some of the following if it applies:
// 	// - test the path against the selector, consuming segments
// 	// - for terminal statements (equality, wildcard), require all the segments to have been consumed
// 	// - recursively filter child (negation, quantifier) or children (connective) statements with the remaining path
// 	switch stmt.(type) {
// 	case equality:
// 		match, remain := stmt.(equality).selector.MatchPath(path...)
// 		if match && len(remain) == 0 {
// 			return stmt, remain
// 		}
// 		return nil, nil
// 	case negation:
// 		newChild, remain := filter(stmt.(negation).statement, path)
// 		if newChild != nil && len(remain) == 0 {
// 			return negation{
// 				statement: newChild,
// 			}, nil
// 		}
// 		return nil, nil
// 	case connective:
// 		var newChildren []Statement
// 		for _, child := range stmt.(connective).statements {
// 			newChild, remain := filter(child, path)
// 			if newChild != nil && len(remain) == 0 {
// 				newChildren = append(newChildren, newChild)
// 			}
// 		}
// 		if len(newChildren) == 0 {
// 			return nil, nil
// 		}
// 		return connective{
// 			kind:       stmt.(connective).kind,
// 			statements: newChildren,
// 		}, nil
// 	case wildcard:
// 		match, remain := stmt.(wildcard).selector.MatchPath(path...)
// 		if match && len(remain) == 0 {
// 			return stmt, remain
// 		}
// 		return nil, nil
// 	case quantifier:
// 		match, remain := stmt.(quantifier).selector.MatchPath(path...)
// 		if match && len(remain) == 0 {
// 			return stmt, remain
// 		}
// 		if !match {
// 			return nil, nil
// 		}
// 		newChild, remain := filter(stmt.(quantifier).statement, remain)
// 		if newChild != nil && len(remain) == 0 {
// 			return quantifier{
// 				kind:      stmt.(quantifier).kind,
// 				selector:  stmt.(quantifier).selector,
// 				statement: newChild,
// 			}, nil
// 		}
// 		return nil, nil
// 	default:
// 		panic(fmt.Errorf("unimplemented statement kind: %s", stmt.Kind()))
// 	}
// }

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
