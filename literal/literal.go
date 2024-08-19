package literal

import (
	"fmt"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/node/basicnode"
)

var ErrType = fmt.Errorf("literal is not this type")

const (
	Kind_IPLD   = "ipld"
	Kind_Int    = "int"
	Kind_Float  = "float"
	Kind_String = "string"
)

type Literal interface {
	Kind() string // ipld | integer | float | string
	AsNode() (ipld.Node, error)
	AsInt() (int64, error)
	AsFloat() (float64, error)
	AsString() (string, error)
}

type literal struct{}

func (l literal) AsFloat() (float64, error) {
	return 0, ErrType
}

func (l literal) AsInt() (int64, error) {
	return 0, ErrType
}

func (l literal) AsNode() (datamodel.Node, error) {
	return nil, ErrType
}

func (l literal) AsString() (string, error) {
	return "", ErrType
}

type node struct {
	literal
	value ipld.Node
}

func (l node) AsNode() (datamodel.Node, error) {
	return l.value, nil
}

func (l node) Kind() string {
	return Kind_IPLD
}

func Node(n ipld.Node) Literal {
	return node{value: n}
}

func Link(cid ipld.Link) Literal {
	nb := basicnode.Prototype.Link.NewBuilder()
	nb.AssignLink(cid)
	return node{value: nb.Build()}
}

func Bool(val bool) Literal {
	nb := basicnode.Prototype.Bool.NewBuilder()
	nb.AssignBool(val)
	return node{value: nb.Build()}
}

type nint struct {
	literal
	value int64
}

func (l nint) AsInt() (int64, error) {
	return l.value, nil
}

func (l nint) Kind() string {
	return Kind_Int
}

func Int(num int64) Literal {
	return nint{value: num}
}

type nfloat struct {
	literal
	value float64
}

func (l nfloat) AsFloat() (float64, error) {
	return l.value, nil
}

func (l nfloat) Kind() string {
	return Kind_Float
}

func Float(num float64) Literal {
	return nfloat{value: num}
}

type str struct {
	literal
	value string
}

func (l str) AsString() (string, error) {
	return l.value, nil
}

func (l str) Kind() string {
	return Kind_String
}

func String(s string) Literal {
	return str{value: s}
}
