package policy

import (
	"fmt"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/must"
	"github.com/ipld/go-ipld-prime/node/basicnode"

	"github.com/ucan-wg/go-ucan/pkg/policy/selector"
)

func FromIPLD(node datamodel.Node) (Policy, error) {
	return statementsFromIPLD("/", node)
}

func FromDagJson(json string) (Policy, error) {
	nodes, err := ipld.Decode([]byte(json), dagjson.Decode)
	if err != nil {
		return nil, err
	}
	return FromIPLD(nodes)
}

func statementFromIPLD(path string, node datamodel.Node) (Statement, error) {
	// sanity checks
	if node.Kind() != datamodel.Kind_List {
		return nil, ErrNotATuple(path)
	}
	if node.Length() != 2 && node.Length() != 3 {
		return nil, ErrUnrecognizedShape(path)
	}

	// extract operator
	opNode, _ := node.LookupByIndex(0)
	if opNode.Kind() != datamodel.Kind_String {
		return nil, ErrNotAString(path)
	}
	op := must.String(opNode)

	arg2AsSelector := func(op string) (selector.Selector, error) {
		nd, _ := node.LookupByIndex(1)
		if nd.Kind() != datamodel.Kind_String {
			return nil, ErrNotAString(combinePath(path, op, 1))
		}
		sel, err := selector.Parse(must.String(nd))
		if err != nil {
			return nil, ErrInvalidSelector(combinePath(path, op, 1), err)
		}
		return sel, nil
	}

	switch node.Length() {
	case 2:
		switch op {
		case KindNot:
			arg2, _ := node.LookupByIndex(1)
			statement, err := statementFromIPLD(combinePath(path, op, 1), arg2)
			if err != nil {
				return nil, err
			}
			return Not(statement), nil

		case KindAnd, KindOr:
			arg2, _ := node.LookupByIndex(1)
			statement, err := statementsFromIPLD(combinePath(path, op, 1), arg2)
			if err != nil {
				return nil, err
			}
			return connective{kind: op, statements: statement}, nil

		default:
			return nil, ErrUnrecognizedOperator(path, op)
		}
	case 3:
		switch op {
		case KindEqual, KindLessThan, KindLessThanOrEqual, KindGreaterThan, KindGreaterThanOrEqual:
			sel, err := arg2AsSelector(op)
			if err != nil {
				return nil, err
			}
			arg3, _ := node.LookupByIndex(2)
			return equality{kind: op, selector: sel, value: arg3}, nil

		case KindLike:
			sel, err := arg2AsSelector(op)
			if err != nil {
				return nil, err
			}
			pattern, _ := node.LookupByIndex(2)
			if pattern.Kind() != datamodel.Kind_String {
				return nil, ErrNotAString(combinePath(path, op, 2))
			}
			res, err := Like(sel, must.String(pattern))
			if err != nil {
				return nil, ErrInvalidPattern(combinePath(path, op, 2), err)
			}
			return res, nil

		case KindAll, KindAny:
			sel, err := arg2AsSelector(op)
			if err != nil {
				return nil, err
			}
			statementsNode, _ := node.LookupByIndex(2)
			statement, err := statementFromIPLD(combinePath(path, op, 1), statementsNode)
			if err != nil {
				return nil, err
			}
			return quantifier{kind: op, selector: sel, statement: statement}, nil

		default:
			return nil, ErrUnrecognizedOperator(path, op)
		}

	default:
		return nil, ErrUnrecognizedShape(path)
	}
}

func statementsFromIPLD(path string, node datamodel.Node) ([]Statement, error) {
	// sanity checks
	if node.Kind() != datamodel.Kind_List {
		return nil, ErrNotATuple(path)
	}
	if node.Length() == 0 {
		return nil, nil
	}

	res := make([]Statement, node.Length())

	for i := int64(0); i < node.Length(); i++ {
		nd, _ := node.LookupByIndex(i)
		statement, err := statementFromIPLD(fmt.Sprintf("%s%d/", path, i), nd)
		if err != nil {
			return nil, err
		}
		res[i] = statement
	}

	return res, nil
}

func (p Policy) ToIPLD() (datamodel.Node, error) {
	return statementsToIPLD(p)
}

func statementsToIPLD(statements []Statement) (datamodel.Node, error) {
	list := basicnode.Prototype.List.NewBuilder()
	// can't error, we have the right builder.
	listBuilder, _ := list.BeginList(int64(len(statements)))
	for _, argStatement := range statements {
		node, err := statementToIPLD(argStatement)
		if err != nil {
			return nil, err
		}
		err = listBuilder.AssembleValue().AssignNode(node)
		if err != nil {
			return nil, err
		}
	}
	err := listBuilder.Finish()
	if err != nil {
		return nil, err
	}
	return list.Build(), nil
}

func statementToIPLD(statement Statement) (datamodel.Node, error) {
	list := basicnode.Prototype.List.NewBuilder()

	length := int64(3)
	switch statement.(type) {
	case negation, connective:
		length = 2
	}

	// can't error, we have the right builder.
	listBuilder, _ := list.BeginList(length)

	switch statement := statement.(type) {
	case equality:
		err := listBuilder.AssembleValue().AssignString(statement.kind)
		if err != nil {
			return nil, err
		}
		err = listBuilder.AssembleValue().AssignString(statement.selector.String())
		if err != nil {
			return nil, err
		}
		err = listBuilder.AssembleValue().AssignNode(statement.value)
		if err != nil {
			return nil, err
		}

	case negation:
		err := listBuilder.AssembleValue().AssignString(statement.Kind())
		if err != nil {
			return nil, err
		}
		node, err := statementToIPLD(statement.statement)
		if err != nil {
			return nil, err
		}
		err = listBuilder.AssembleValue().AssignNode(node)
		if err != nil {
			return nil, err
		}

	case connective:
		err := listBuilder.AssembleValue().AssignString(statement.kind)
		if err != nil {
			return nil, err
		}
		args, err := statementsToIPLD(statement.statements)
		if err != nil {
			return nil, err
		}
		err = listBuilder.AssembleValue().AssignNode(args)
		if err != nil {
			return nil, err
		}

	case wildcard:
		err := listBuilder.AssembleValue().AssignString(statement.Kind())
		if err != nil {
			return nil, err
		}
		err = listBuilder.AssembleValue().AssignString(statement.selector.String())
		if err != nil {
			return nil, err
		}
		err = listBuilder.AssembleValue().AssignString(string(statement.pattern))
		if err != nil {
			return nil, err
		}

	case quantifier:
		err := listBuilder.AssembleValue().AssignString(statement.kind)
		if err != nil {
			return nil, err
		}
		err = listBuilder.AssembleValue().AssignString(statement.selector.String())
		if err != nil {
			return nil, err
		}
		args, err := statementToIPLD(statement.statement)
		if err != nil {
			return nil, err
		}
		err = listBuilder.AssembleValue().AssignNode(args)
		if err != nil {
			return nil, err
		}
	}

	err := listBuilder.Finish()
	if err != nil {
		return nil, err
	}

	return list.Build(), nil
}

func combinePath(prev string, operator string, index int) string {
	return fmt.Sprintf("%s%d-%s/", prev, index, operator)
}
