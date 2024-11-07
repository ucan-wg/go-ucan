package policy

import (
	"cmp"
	"fmt"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/must"
)

// MatchTrace, if set, will print tracing statements to stdout of the policy matching resolution.
var MatchTrace = false

// Match determines if the IPLD node satisfies the policy.
func (p Policy) Match(node datamodel.Node) bool {
	for _, stmt := range p {
		res, _ := matchStatement(stmt, node)
		switch res {
		case matchResultNoData, matchResultFalse:
			return false
		case matchResultOptionalNoData, matchResultTrue:
			// continue
		}
	}
	return true
}

// PartialMatch returns false IIF one non-optional Statement has the corresponding data and doesn't match.
// If the data is missing or the non-optional Statement is matching, true is returned.
//
// This allows performing the policy checking in multiple steps, and find immediately if a Statement already failed.
// A final call to Match is necessary to make sure that the policy is fully matched, with no missing data
// (apart from optional values).
//
// The first Statement failing to match is returned as well.
func (p Policy) PartialMatch(node datamodel.Node) (bool, Statement) {
	for _, stmt := range p {
		res, leaf := matchStatement(stmt, node)
		switch res {
		case matchResultFalse:
			return false, leaf
		case matchResultNoData, matchResultOptionalNoData, matchResultTrue:
			// continue
		}
	}
	return true, nil
}

type matchResult int8

const (
	matchResultTrue           matchResult = iota // statement has data and resolve to true
	matchResultFalse                             // statement has data and resolve to false
	matchResultNoData                            // statement has no data
	matchResultOptionalNoData                    // statement has no data and is optional
)

// matchStatement evaluate the policy against the given ipld.Node and returns:
// - matchResultTrue: if the selector matched and the statement evaluated to true.
// - matchResultFalse: if the selector matched and the statement evaluated to false.
// - matchResultNoData: if the selector didn't match the expected data.
// For matchResultTrue and matchResultNoData, the leaf-most (innermost) statement failing to be true is returned,
// as well as the corresponding root-most encompassing statement.
func matchStatement(cur Statement, node ipld.Node) (output matchResult, leafMost Statement) {
	var boolToRes = func(v bool) (matchResult, Statement) {
		if v {
			return matchResultTrue, nil
		} else {
			return matchResultFalse, cur
		}
	}
	if MatchTrace {
		defer func() {
			fmt.Printf("match %v --> %v\n", cur, matchResToStr(output))
		}()
	}

	switch cur.Kind() {
	case KindEqual:
		if s, ok := cur.(equality); ok {
			res, err := s.selector.Select(node)
			if err != nil {
				return matchResultNoData, cur
			}
			if res == nil { // optional selector didn't match
				return matchResultOptionalNoData, nil
			}
			return boolToRes(datamodel.DeepEqual(s.value, res))
		}
	case KindGreaterThan:
		if s, ok := cur.(equality); ok {
			res, err := s.selector.Select(node)
			if err != nil {
				return matchResultNoData, cur
			}
			if res == nil { // optional selector didn't match
				return matchResultOptionalNoData, nil
			}
			return boolToRes(isOrdered(s.value, res, gt))
		}
	case KindGreaterThanOrEqual:
		if s, ok := cur.(equality); ok {
			res, err := s.selector.Select(node)
			if err != nil {
				return matchResultNoData, cur
			}
			if res == nil { // optional selector didn't match
				return matchResultOptionalNoData, nil
			}
			return boolToRes(isOrdered(s.value, res, gte))
		}
	case KindLessThan:
		if s, ok := cur.(equality); ok {
			res, err := s.selector.Select(node)
			if err != nil {
				return matchResultNoData, cur
			}
			if res == nil { // optional selector didn't match
				return matchResultOptionalNoData, nil
			}
			return boolToRes(isOrdered(s.value, res, lt))
		}
	case KindLessThanOrEqual:
		if s, ok := cur.(equality); ok {
			res, err := s.selector.Select(node)
			if err != nil {
				return matchResultNoData, cur
			}
			if res == nil { // optional selector didn't match
				return matchResultOptionalNoData, nil
			}
			return boolToRes(isOrdered(s.value, res, lte))
		}
	case KindNot:
		if s, ok := cur.(negation); ok {
			res, leaf := matchStatement(s.statement, node)
			switch res {
			case matchResultNoData, matchResultOptionalNoData:
				return res, leaf
			case matchResultTrue:
				return matchResultFalse, leaf
			case matchResultFalse:
				return matchResultTrue, leaf
			}
		}
	case KindAnd:
		if s, ok := cur.(connective); ok {
			for _, cs := range s.statements {
				res, leaf := matchStatement(cs, node)
				switch res {
				case matchResultNoData, matchResultOptionalNoData:
					return res, leaf
				case matchResultTrue:
					// continue
				case matchResultFalse:
					return matchResultFalse, leaf
				}
			}
			return matchResultTrue, nil
		}
	case KindOr:
		if s, ok := cur.(connective); ok {
			if len(s.statements) == 0 {
				return matchResultTrue, nil
			}
			for _, cs := range s.statements {
				res, leaf := matchStatement(cs, node)
				switch res {
				case matchResultNoData, matchResultOptionalNoData:
					return res, leaf
				case matchResultTrue:
					return matchResultTrue, leaf
				case matchResultFalse:
					// continue
				}
			}
			return matchResultFalse, cur
		}
	case KindLike:
		if s, ok := cur.(wildcard); ok {
			res, err := s.selector.Select(node)
			if err != nil {
				return matchResultNoData, cur
			}
			if res == nil { // optional selector didn't match
				return matchResultOptionalNoData, nil
			}
			v, err := res.AsString()
			if err != nil {
				return matchResultFalse, cur // not a string
			}
			return boolToRes(s.pattern.Match(v))
		}
	case KindAll:
		if s, ok := cur.(quantifier); ok {
			res, err := s.selector.Select(node)
			if err != nil {
				return matchResultNoData, cur
			}
			if res == nil {
				return matchResultOptionalNoData, nil
			}
			it := res.ListIterator()
			if it == nil {
				return matchResultFalse, cur // not a list
			}
			for !it.Done() {
				_, v, err := it.Next()
				if err != nil {
					panic("should never happen")
				}
				matchRes, leaf := matchStatement(s.statement, v)
				switch matchRes {
				case matchResultNoData, matchResultOptionalNoData:
					return matchRes, leaf
				case matchResultTrue:
					// continue
				case matchResultFalse:
					return matchResultFalse, leaf
				}
			}
			return matchResultTrue, nil
		}
	case KindAny:
		if s, ok := cur.(quantifier); ok {
			res, err := s.selector.Select(node)
			if err != nil {
				return matchResultNoData, cur
			}
			if res == nil {
				return matchResultOptionalNoData, nil
			}
			it := res.ListIterator()
			if it == nil {
				return matchResultFalse, cur // not a list
			}
			for !it.Done() {
				_, v, err := it.Next()
				if err != nil {
					panic("should never happen")
				}
				matchRes, leaf := matchStatement(s.statement, v)
				switch matchRes {
				case matchResultNoData, matchResultOptionalNoData:
					return matchRes, leaf
				case matchResultTrue:
					return matchResultTrue, nil
				case matchResultFalse:
					// continue
				}
			}
			return matchResultFalse, cur
		}
	}
	panic(fmt.Errorf("unimplemented statement kind: %s", cur.Kind()))
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

func matchResToStr(res matchResult) string {
	switch res {
	case matchResultTrue:
		return "True"
	case matchResultFalse:
		return "False"
	case matchResultNoData:
		return "NoData"
	case matchResultOptionalNoData:
		return "OptionalNoData"
	default:
		panic("invalid matchResult")
	}
}
