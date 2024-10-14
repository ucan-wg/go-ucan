package policy

// https://github.com/ucan-wg/delegation/blob/4094d5878b58f5d35055a3b93fccda0b8329ebae/README.md#policy

import (
	"fmt"
	"strings"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagjson"

	"github.com/ucan-wg/go-ucan/pkg/policy/selector"
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

func (p Policy) String() string {
	if len(p) == 0 {
		return "[]"
	}
	childs := make([]string, len(p))
	for i, statement := range p {
		childs[i] = strings.ReplaceAll(statement.String(), "\n", "\n  ")
	}
	return fmt.Sprintf("[\n  %s\n]", strings.Join(childs, ",\n  "))
}

type Statement interface {
	Kind() string
	String() string
}

type equality struct {
	kind     string
	selector selector.Selector
	value    ipld.Node
}

func (e equality) Kind() string {
	return e.kind
}

func (e equality) String() string {
	child, err := ipld.Encode(e.value, dagjson.Encode)
	if err != nil {
		return "ERROR: INVALID VALUE"
	}
	return fmt.Sprintf(`["%s", "%s", %s]`, e.kind, e.selector, strings.ReplaceAll(string(child), "\n", "\n  "))
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

func (n negation) String() string {
	child := n.statement.String()
	return fmt.Sprintf(`["%s", "%s"]`, n.Kind(), strings.ReplaceAll(child, "\n", "\n  "))
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

func (c connective) String() string {
	childs := make([]string, len(c.statements))
	for i, statement := range c.statements {
		childs[i] = strings.ReplaceAll(statement.String(), "\n", "\n  ")
	}
	return fmt.Sprintf("[\"%s\", [\n  %s]]\n", c.kind, strings.Join(childs, ",\n  "))
}

func And(stmts ...Statement) Statement {
	return connective{kind: KindAnd, statements: stmts}
}

func Or(stmts ...Statement) Statement {
	return connective{kind: KindOr, statements: stmts}
}

type wildcard struct {
	selector selector.Selector
	pattern  glob
}

func (n wildcard) Kind() string {
	return KindLike
}

func (n wildcard) String() string {
	return fmt.Sprintf(`["%s", "%s", "%s"]`, n.Kind(), n.selector, n.pattern)
}

func Like(selector selector.Selector, pattern string) (Statement, error) {
	g, err := parseGlob(pattern)
	if err != nil {
		return nil, err
	}

	return wildcard{selector: selector, pattern: g}, nil
}

func MustLike(selector selector.Selector, pattern string) Statement {
	g, err := Like(selector, pattern)
	if err != nil {
		panic(err)
	}
	return g
}

type quantifier struct {
	kind      string
	selector  selector.Selector
	statement Statement
}

func (n quantifier) Kind() string {
	return n.kind
}

func (n quantifier) String() string {
	child := n.statement.String()
	return fmt.Sprintf("[\"%s\", \"%s\",\n  %s]", n.Kind(), n.selector, strings.ReplaceAll(child, "\n", "\n  "))
}

func All(selector selector.Selector, statement Statement) Statement {
	return quantifier{kind: KindAll, selector: selector, statement: statement}
}

func Any(selector selector.Selector, statement Statement) Statement {
	return quantifier{kind: KindAny, selector: selector, statement: statement}
}
