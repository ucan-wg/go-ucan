package policy

import (
	"cmp"
	"fmt"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/storacha-network/go-ucanto/core/policy/literal"
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
			n, err := selectNode(s.Selector(), node)
			if err != nil {
				if _, ok := err.(datamodel.ErrNotExists); ok {
					return false, nil
				}
				return false, fmt.Errorf("selecting node: %w", err)
			}
			return isDeepEqual(s.Value(), n)
		}
	case Kind_GreaterThan:
		if s, ok := statement.(InequalityStatement); ok {
			n, err := selectNode(s.Selector(), node)
			if err != nil {
				if _, ok := err.(datamodel.ErrNotExists); ok {
					return false, nil
				}
				return false, fmt.Errorf("selecting node: %w", err)
			}
			return isOrdered(s, n, gt)
		}
	case Kind_GreaterThanOrEqual:
		if s, ok := statement.(InequalityStatement); ok {
			n, err := selectNode(s.Selector(), node)
			if err != nil {
				if _, ok := err.(datamodel.ErrNotExists); ok {
					return false, nil
				}
				return false, fmt.Errorf("selecting node: %w", err)
			}
			return isOrdered(s, n, gte)
		}
	case Kind_LessThan:
		if s, ok := statement.(InequalityStatement); ok {
			n, err := selectNode(s.Selector(), node)
			if err != nil {
				if _, ok := err.(datamodel.ErrNotExists); ok {
					return false, nil
				}
				return false, fmt.Errorf("selecting node: %w", err)
			}
			return isOrdered(s, n, lt)
		}
	case Kind_LessThanOrEqual:
		if s, ok := statement.(InequalityStatement); ok {
			n, err := selectNode(s.Selector(), node)
			if err != nil {
				if _, ok := err.(datamodel.ErrNotExists); ok {
					return false, nil
				}
				return false, fmt.Errorf("selecting node: %w", err)
			}
			return isOrdered(s, n, lte)
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
	return false, fmt.Errorf("statement kind not implemented: %s", statement.Kind())
}

func selectNode(sel selector.Selector, node ipld.Node) (child ipld.Node, err error) {
	if sel.Identity() {
		child = node
	} else if sel.Field() != "" {
		child, err = node.LookupByString(sel.Field())
	} else {
		child, err = node.LookupByIndex(int64(sel.Index()))
	}
	return
}

func isOrdered(stmt InequalityStatement, node ipld.Node, satisfies func(order int) bool) (bool, error) {
	if stmt.Value().Kind() == literal.Kind_Int && node.Kind() == ipld.Kind_Int {
		a, err := node.AsInt()
		if err != nil {
			return false, fmt.Errorf("extracting node int: %w", err)
		}
		b, err := stmt.Value().AsInt()
		if err != nil {
			return false, fmt.Errorf("extracting selector int: %w", err)
		}
		return satisfies(cmp.Compare(a, b)), nil
	}

	if stmt.Value().Kind() == literal.Kind_Float && node.Kind() == ipld.Kind_Float {
		a, err := node.AsFloat()
		if err != nil {
			return false, fmt.Errorf("extracting node float: %w", err)
		}
		b, err := stmt.Value().AsFloat()
		if err != nil {
			return false, fmt.Errorf("extracting selector float: %w", err)
		}
		return satisfies(cmp.Compare(a, b)), nil
	}

	return false, fmt.Errorf("selector type %s is not compatible with node type %s: kind mismatch: need int or float", stmt.Value().Kind(), node.Kind())
}

func isDeepEqual(value literal.Literal, node ipld.Node) (bool, error) {
	switch value.Kind() {
	case literal.Kind_String:
		if node.Kind() != ipld.Kind_String {
			return false, nil
		}
		a, err := node.AsString()
		if err != nil {
			return false, fmt.Errorf("extracting node string: %w", err)
		}
		b, err := value.AsString()
		if err != nil {
			return false, fmt.Errorf("extracting selector string: %w", err)
		}
		return a == b, nil
	case literal.Kind_Int:
		if node.Kind() != ipld.Kind_Int {
			return false, nil
		}
		a, err := node.AsInt()
		if err != nil {
			return false, fmt.Errorf("extracting node int: %w", err)
		}
		b, err := value.AsInt()
		if err != nil {
			return false, fmt.Errorf("extracting selector int: %w", err)
		}
		return a == b, nil
	case literal.Kind_Float:
		if node.Kind() != ipld.Kind_Float {
			return false, nil
		}
		a, err := node.AsFloat()
		if err != nil {
			return false, fmt.Errorf("extracting node float: %w", err)
		}
		b, err := value.AsFloat()
		if err != nil {
			return false, fmt.Errorf("extracting selector float: %w", err)
		}
		return a == b, nil
	case literal.Kind_IPLD:
		v, err := value.AsNode()
		if err != nil {
			return false, fmt.Errorf("extracting selector node: %w", err)
		}
		if v.Kind() != node.Kind() {
			return false, nil
		}
		// TODO: should be easy enough to do the basic types, map, struct and list
		// might be harder.
		switch v.Kind() {
		case ipld.Kind_Bool:
			a, err := node.AsBool()
			if err != nil {
				return false, fmt.Errorf("extracting node boolean: %w", err)
			}
			b, err := v.AsBool()
			if err != nil {
				return false, fmt.Errorf("extracting selector node boolean: %w", err)
			}
			return a == b, nil
		case ipld.Kind_Link:
			a, err := node.AsLink()
			if err != nil {
				return false, fmt.Errorf("extracting node link: %w", err)
			}
			b, err := v.AsLink()
			if err != nil {
				return false, fmt.Errorf("extracting selector node link: %w", err)
			}
			return a.Binary() == b.Binary(), nil
		}
		return false, fmt.Errorf("unsupported IPLD kind: %s", v.Kind())
	}
	return false, fmt.Errorf("unknown literal kind: %s", value.Kind())
}

func gt(order int) bool  { return order == 1 }
func gte(order int) bool { return order == 0 || order == 1 }
func lt(order int) bool  { return order == -1 }
func lte(order int) bool { return order == 0 || order == -1 }
