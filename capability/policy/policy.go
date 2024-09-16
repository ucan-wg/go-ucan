package policy

// https://github.com/ucan-wg/delegation/blob/4094d5878b58f5d35055a3b93fccda0b8329ebae/README.md#policy

import (
	"github.com/ipld/go-ipld-prime"

	"github.com/ucan-wg/go-ucan/capability/policy/selector"
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
	return equality{kind: KindEqual, selector: selector, value: value}
}

func GreaterThan(selector selector.Selector, value ipld.Node) Statement {
	return equality{kind: KindGreaterThan, selector: selector, value: value}
}

func GreaterThanOrEqual(selector selector.Selector, value ipld.Node) Statement {
	return equality{kind: KindGreaterThanOrEqual, selector: selector, value: value}
}

func LessThan(selector selector.Selector, value ipld.Node) Statement {
	return equality{kind: KindLessThan, selector: selector, value: value}
}

func LessThanOrEqual(selector selector.Selector, value ipld.Node) Statement {
	return equality{kind: KindLessThanOrEqual, selector: selector, value: value}
}

type negation struct {
	statement Statement
}

func (n negation) Kind() string {
	return KindNot
}

func Not(stmt Statement) Statement {
	return negation{statement: stmt}
}

type connective struct {
	kind       string
	statements []Statement
}

func (c connective) Kind() string {
	return c.kind
}

func And(stmts ...Statement) Statement {
	return connective{kind: KindAnd, statements: stmts}
}

func Or(stmts ...Statement) Statement {
	return connective{kind: KindOr, statements: stmts}
}

type wildcard struct {
	selector selector.Selector
	pattern  string
}

func (n wildcard) Kind() string {
	return KindLike
}

func Like(selector selector.Selector, pattern string) (Statement, error) {
	return wildcard{selector: selector, pattern: pattern}, nil
}

type quantifier struct {
	kind      string
	selector  selector.Selector
	statement Statement
}

func (n quantifier) Kind() string {
	return n.kind
}

func All(selector selector.Selector, statement Statement) Statement {
	return quantifier{kind: KindAll, selector: selector, statement: statement}
}

func Any(selector selector.Selector, statement Statement) Statement {
	return quantifier{kind: KindAny, selector: selector, statement: statement}
}
