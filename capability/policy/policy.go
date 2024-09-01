package policy

// https://github.com/ucan-wg/delegation/blob/4094d5878b58f5d35055a3b93fccda0b8329ebae/README.md#policy

import (
	"github.com/gobwas/glob"
	"github.com/ipld/go-ipld-prime"

	"github.com/ucan-wg/go-ucan/v1/capability/policy/selector"
)

const (
	KindEqual              = "=="
	KindGreaterThan        = ">"
	KindGreaterThanOrEqual = ">="
	KindLessThan           = "<"
	KindLessThanOrEqual    = "<="
	KindNot                = "not"
	KindAnd                = "and"
	KindOr                 = "or"
	KindLike               = "like"
	KindAll                = "all"
	KindAny                = "any"
)

type Policy []Statement

type Statement interface {
	Kind() string
}

type EqualityStatement interface {
	Statement
	Selector() selector.Selector
	Value() ipld.Node
}

type InequalityStatement interface {
	Statement
	Selector() selector.Selector
	Value() ipld.Node
}

type WildcardStatement interface {
	Statement
	Selector() selector.Selector
	Value() glob.Glob
}

type ConnectiveStatement interface {
	Statement
}

type NegationStatement interface {
	ConnectiveStatement
	Value() Statement
}

type ConjunctionStatement interface {
	ConnectiveStatement
	Value() []Statement
}

type DisjunctionStatement interface {
	ConnectiveStatement
	Value() []Statement
}

type QuantifierStatement interface {
	Statement
	Selector() selector.Selector
	Value() Policy
}

type equality struct {
	kind     string
	selector selector.Selector
	value    ipld.Node
}

func (e equality) Kind() string {
	return e.kind
}

func (e equality) Value() ipld.Node {
	return e.value
}

func (e equality) Selector() selector.Selector {
	return e.selector
}

func Equal(selector selector.Selector, value ipld.Node) EqualityStatement {
	return equality{KindEqual, selector, value}
}

func GreaterThan(selector selector.Selector, value ipld.Node) InequalityStatement {
	return equality{KindGreaterThan, selector, value}
}

func GreaterThanOrEqual(selector selector.Selector, value ipld.Node) InequalityStatement {
	return equality{KindGreaterThanOrEqual, selector, value}
}

func LessThan(selector selector.Selector, value ipld.Node) InequalityStatement {
	return equality{KindLessThan, selector, value}
}

func LessThanOrEqual(selector selector.Selector, value ipld.Node) InequalityStatement {
	return equality{KindLessThanOrEqual, selector, value}
}

type negation struct {
	statement Statement
}

func (n negation) Kind() string {
	return KindNot
}

func (n negation) Value() Statement {
	return n.statement
}

func Not(stmt Statement) NegationStatement {
	return negation{stmt}
}

type connective struct {
	kind       string
	statements []Statement
}

func (c connective) Kind() string {
	return c.kind
}

func (c connective) Value() []Statement {
	return c.statements
}

func And(stmts ...Statement) ConjunctionStatement {
	return connective{KindAnd, stmts}
}

func Or(stmts ...Statement) DisjunctionStatement {
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

func (n wildcard) Selector() selector.Selector {
	return n.selector
}

func (n wildcard) Value() glob.Glob {
	return n.glob
}

func Like(selector selector.Selector, pattern string) (WildcardStatement, error) {
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

func (n quantifier) Selector() selector.Selector {
	return n.selector
}

func (n quantifier) Value() Policy {
	return n.statements
}

func All(selector selector.Selector, policy ...Statement) QuantifierStatement {
	return quantifier{KindAll, selector, policy}
}

func Any(selector selector.Selector, policy ...Statement) QuantifierStatement {
	return quantifier{KindAny, selector, policy}
}
