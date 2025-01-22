package policy

// https://github.com/ucan-wg/delegation/blob/4094d5878b58f5d35055a3b93fccda0b8329ebae/README.md#policy

import (
	"fmt"
	"strings"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagjson"

	selpkg "github.com/ucan-wg/go-ucan/pkg/policy/selector"
)

const (
	KindEqual              = "=="   // implemented by equality
	KindNotEqual           = "!="   // implemented by equality
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

type Constructor func() (Statement, error)

func Construct(cstors ...Constructor) (Policy, error) {
	stmts, err := assemble(cstors)
	if err != nil {
		return nil, err
	}
	return stmts, nil
}

func MustConstruct(cstors ...Constructor) Policy {
	pol, err := Construct(cstors...)
	if err != nil {
		panic(err)
	}
	return pol
}

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
	selector selpkg.Selector
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

func Equal(selector string, value ipld.Node) Constructor {
	return func() (Statement, error) {
		sel, err := selpkg.Parse(selector)
		return equality{kind: KindEqual, selector: sel, value: value}, err
	}
}

func NotEqual(selector string, value ipld.Node) Constructor {
	return func() (Statement, error) {
		sel, err := selpkg.Parse(selector)
		return equality{kind: KindNotEqual, selector: sel, value: value}, err
	}
}

func GreaterThan(selector string, value ipld.Node) Constructor {
	return func() (Statement, error) {
		sel, err := selpkg.Parse(selector)
		return equality{kind: KindGreaterThan, selector: sel, value: value}, err
	}
}

func GreaterThanOrEqual(selector string, value ipld.Node) Constructor {
	return func() (Statement, error) {
		sel, err := selpkg.Parse(selector)
		return equality{kind: KindGreaterThanOrEqual, selector: sel, value: value}, err
	}
}

func LessThan(selector string, value ipld.Node) Constructor {
	return func() (Statement, error) {
		sel, err := selpkg.Parse(selector)
		return equality{kind: KindLessThan, selector: sel, value: value}, err
	}
}

func LessThanOrEqual(selector string, value ipld.Node) Constructor {
	return func() (Statement, error) {
		sel, err := selpkg.Parse(selector)
		return equality{kind: KindLessThanOrEqual, selector: sel, value: value}, err
	}
}

type negation struct {
	statement Statement
}

func (n negation) Kind() string {
	return KindNot
}

func (n negation) String() string {
	child := n.statement.String()
	return fmt.Sprintf(`["%s", %s]`, n.Kind(), strings.ReplaceAll(child, "\n", "\n  "))
}

func Not(cstor Constructor) Constructor {
	return func() (Statement, error) {
		stmt, err := cstor()
		return negation{statement: stmt}, err
	}
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
	return fmt.Sprintf("[\"%s\", [\n  %s\n]]", c.kind, strings.Join(childs, ",\n  "))
}

func And(cstors ...Constructor) Constructor {
	return func() (Statement, error) {
		stmts, err := assemble(cstors)
		if err != nil {
			return nil, err
		}
		return connective{kind: KindAnd, statements: stmts}, nil
	}
}

func Or(cstors ...Constructor) Constructor {
	return func() (Statement, error) {
		stmts, err := assemble(cstors)
		if err != nil {
			return nil, err
		}
		return connective{kind: KindOr, statements: stmts}, nil
	}
}

type wildcard struct {
	selector selpkg.Selector
	pattern  glob
}

func (n wildcard) Kind() string {
	return KindLike
}

func (n wildcard) String() string {
	return fmt.Sprintf(`["%s", "%s", "%s"]`, n.Kind(), n.selector, n.pattern)
}

func Like(selector string, pattern string) Constructor {
	return func() (Statement, error) {
		g, err := parseGlob(pattern)
		if err != nil {
			return nil, err
		}
		sel, err := selpkg.Parse(selector)
		return wildcard{selector: sel, pattern: g}, err
	}
}

type quantifier struct {
	kind      string
	selector  selpkg.Selector
	statement Statement
}

func (n quantifier) Kind() string {
	return n.kind
}

func (n quantifier) String() string {
	child := n.statement.String()
	return fmt.Sprintf("[\"%s\", \"%s\",\n  %s\n]", n.Kind(), n.selector, strings.ReplaceAll(child, "\n", "\n  "))
}

func All(selector string, cstor Constructor) Constructor {
	return func() (Statement, error) {
		stmt, err := cstor()
		if err != nil {
			return nil, err
		}
		sel, err := selpkg.Parse(selector)
		return quantifier{kind: KindAll, selector: sel, statement: stmt}, err
	}
}

func Any(selector string, cstor Constructor) Constructor {
	return func() (Statement, error) {
		stmt, err := cstor()
		if err != nil {
			return nil, err
		}
		sel, err := selpkg.Parse(selector)
		return quantifier{kind: KindAny, selector: sel, statement: stmt}, err
	}
}

func assemble(cstors []Constructor) ([]Statement, error) {
	stmts := make([]Statement, 0, len(cstors))
	for _, cstor := range cstors {
		stmt, err := cstor()
		if err != nil {
			return nil, err
		}
		stmts = append(stmts, stmt)
	}
	return stmts, nil
}
