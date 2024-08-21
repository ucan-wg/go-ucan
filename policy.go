package policy

// https://github.com/ucan-wg/delegation/blob/4094d5878b58f5d35055a3b93fccda0b8329ebae/README.md#policy

import (
	"github.com/gobwas/glob"
	"github.com/ipld/go-ipld-prime"
	"github.com/storacha-network/go-ucanto/core/policy/selector"
)

const (
	Kind_Equal              = "=="
	Kind_GreaterThan        = ">"
	Kind_GreaterThanOrEqual = ">="
	Kind_LessThan           = "<"
	Kind_LessThanOrEqual    = "<="
	Kind_Not                = "not"
	Kind_And                = "and"
	Kind_Or                 = "or"
	Kind_Like               = "like"
	Kind_All                = "all"
	Kind_Any                = "any"
)

type Policy = []Statement

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
	return equality{Kind_Equal, selector, value}
}

func GreaterThan(selector selector.Selector, value ipld.Node) InequalityStatement {
	return equality{Kind_GreaterThan, selector, value}
}

func GreaterThanOrEqual(selector selector.Selector, value ipld.Node) InequalityStatement {
	return equality{Kind_GreaterThanOrEqual, selector, value}
}

func LessThan(selector selector.Selector, value ipld.Node) InequalityStatement {
	return equality{Kind_LessThan, selector, value}
}

func LessThanOrEqual(selector selector.Selector, value ipld.Node) InequalityStatement {
	return equality{Kind_LessThanOrEqual, selector, value}
}

type negation struct {
	statement Statement
}

func (n negation) Kind() string {
	return Kind_Not
}

func (n negation) Value() Statement {
	return n.statement
}

func Not(stmt Statement) NegationStatement {
	return negation{stmt}
}

type conjunction struct {
	statements []Statement
}

func (n conjunction) Kind() string {
	return Kind_And
}

func (n conjunction) Value() []Statement {
	return n.statements
}

func And(stmts ...Statement) ConjunctionStatement {
	return conjunction{stmts}
}

type disjunction struct {
	statements []Statement
}

func (n disjunction) Kind() string {
	return Kind_Or
}

func (n disjunction) Value() []Statement {
	return n.statements
}

func Or(stmts ...Statement) DisjunctionStatement {
	return disjunction{stmts}
}

type wildcard struct {
	selector selector.Selector
	glob     glob.Glob
}

func (n wildcard) Kind() string {
	return Kind_Like
}

func (n wildcard) Selector() selector.Selector {
	return n.selector
}

func (n wildcard) Value() glob.Glob {
	return n.glob
}

func Like(selector selector.Selector, glob glob.Glob) WildcardStatement {
	return wildcard{selector, glob}
}

type quantifier struct {
	kind     string
	selector selector.Selector
	policy   Policy
}

func (n quantifier) Kind() string {
	return n.kind
}

func (n quantifier) Selector() selector.Selector {
	return n.selector
}

func (n quantifier) Value() Policy {
	return n.policy
}

func All(selector selector.Selector, policy Policy) QuantifierStatement {
	return quantifier{Kind_All, selector, policy}
}

func Any(selector selector.Selector, policy Policy) QuantifierStatement {
	return quantifier{Kind_Any, selector, policy}
}
