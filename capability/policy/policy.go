package policy

// https://github.com/ucan-wg/delegation/blob/4094d5878b58f5d35055a3b93fccda0b8329ebae/README.md#policy

import (
	"github.com/gobwas/glob"
	"github.com/ipld/go-ipld-prime"

	"github.com/ucan-wg/go-ucan/v1/capability/policy/selector"
)

const (
	KindEqual              = "=="   // implemented by equality
	KindGreaterThan        = ">"    // implemented by equality
	KindGreaterThanOrEqual = ">="   // implemented by equality
	KindLessThan           = "<"    // implemented by equality
	KindLessThanOrEqual    = "<="   // implemented by equality
	KindNot                = "not"  // implemented by negation
	KindAnd                = "and"  // implemented by connective
	KindOr                 = "or"   // implemented by connective
	KindLike               = "like" // implemented by wildcard
	KindAll                = "all"  // implemented by quantifier
	KindAny                = "any"  // implemented by quantifier
)

type Policy []Statement

type Statement interface {
	Kind() string
}

type equality struct {
	kind     string
	selector selector.Selector
	value    ipld.Node
}

func (e equality) Kind() string {
	return e.kind
}

func Equal(selector selector.Selector, value ipld.Node) Statement {
	return equality{KindEqual, selector, value}
}

func GreaterThan(selector selector.Selector, value ipld.Node) Statement {
	return equality{KindGreaterThan, selector, value}
}

func GreaterThanOrEqual(selector selector.Selector, value ipld.Node) Statement {
	return equality{KindGreaterThanOrEqual, selector, value}
}

func LessThan(selector selector.Selector, value ipld.Node) Statement {
	return equality{KindLessThan, selector, value}
}

func LessThanOrEqual(selector selector.Selector, value ipld.Node) Statement {
	return equality{KindLessThanOrEqual, selector, value}
}

type negation struct {
	statement Statement
}

func (n negation) Kind() string {
	return KindNot
}

func Not(stmt Statement) Statement {
	return negation{stmt}
}

type connective struct {
	kind       string
	statements []Statement
}

func (c connective) Kind() string {
	return c.kind
}

func And(stmts ...Statement) Statement {
	return connective{KindAnd, stmts}
}

func Or(stmts ...Statement) Statement {
	return connective{KindOr, stmts}
}

type wildcard struct {
	selector selector.Selector
	pattern  string
	glob     glob.Glob // not serialized
}

func (n wildcard) Kind() string {
	return KindLike
}

func Like(selector selector.Selector, pattern string) (Statement, error) {
	g, err := glob.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return wildcard{selector, pattern, g}, nil
}

type quantifier struct {
	kind       string
	selector   selector.Selector
	statements []Statement
}

func (n quantifier) Kind() string {
	return n.kind
}

func All(selector selector.Selector, policy ...Statement) Statement {
	return quantifier{KindAll, selector, policy}
}

func Any(selector selector.Selector, policy ...Statement) Statement {
	return quantifier{KindAny, selector, policy}
}
