// Package literal holds a collection of functions to create IPLD types to use in policies, selector and args.
package literal

import (
	"fmt"
	"reflect"

	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/fluent/qp"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/ipld/go-ipld-prime/node/basicnode"
)

var Bool = basicnode.NewBool
var Int = basicnode.NewInt
var Float = basicnode.NewFloat
var String = basicnode.NewString
var Bytes = basicnode.NewBytes
var Link = basicnode.NewLink

func LinkCid(cid cid.Cid) ipld.Node {
	return Link(cidlink.Link{Cid: cid})
}

func Null() ipld.Node {
	nb := basicnode.Prototype.Any.NewBuilder()
	nb.AssignNull()
	return nb.Build()
}

// Map creates an IPLD node from a map[string]any
func Map[T any](m map[string]T) (ipld.Node, error) {
	return qp.BuildMap(basicnode.Prototype.Any, int64(len(m)), func(ma datamodel.MapAssembler) {
		for k, v := range m {
			qp.MapEntry(ma, k, anyAssemble(v))
		}
	})
}

// List creates an IPLD node from a []any
func List[T any](l []T) (ipld.Node, error) {
	return qp.BuildList(basicnode.Prototype.Any, int64(len(l)), func(la datamodel.ListAssembler) {
		for _, val := range l {
			qp.ListEntry(la, anyAssemble(val))
		}
	})
}

func anyAssemble(val any) qp.Assemble {
	var rt reflect.Type
	var rv reflect.Value

	// support for recursive calls, staying in reflection land
	if cast, ok := val.(reflect.Value); ok {
		rt = cast.Type()
		rv = cast
	} else {
		rt = reflect.TypeOf(val)
		rv = reflect.ValueOf(val)
	}

	// we need to dereference in some cases, to get the real value type
	if rt.Kind() == reflect.Ptr || rt.Kind() == reflect.Interface {
		rv = rv.Elem()
		rt = rv.Type()
	}

	switch rt.Kind() {
	case reflect.Array, reflect.Slice:
		return qp.List(int64(rv.Len()), func(la datamodel.ListAssembler) {
			for i := range rv.Len() {
				qp.ListEntry(la, anyAssemble(rv.Index(i)))
			}
		})
	case reflect.Map:
		if rt.Key().Kind() != reflect.String {
			break
		}
		it := rv.MapRange()
		return qp.Map(int64(rv.Len()), func(ma datamodel.MapAssembler) {
			for it.Next() {
				qp.MapEntry(ma, it.Key().String(), anyAssemble(it.Value()))
			}
		})
	case reflect.Bool:
		return qp.Bool(rv.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return qp.Int(rv.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return qp.Int(int64(rv.Uint()))
	case reflect.Float32, reflect.Float64:
		return qp.Float(rv.Float())
	case reflect.String:
		return qp.String(rv.String())
	default:
	}

	panic(fmt.Sprintf("unsupported type %T", val))
}
