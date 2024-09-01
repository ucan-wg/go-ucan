package policy

import "fmt"

type errWithPath struct {
	path string
	msg  string
}

func (e errWithPath) Error() string {
	return fmt.Sprintf("IPLD path '%s': %s", e.path, e.msg)
}

func ErrInvalidSelector(path string, err error) error {
	return errWithPath{path: path, msg: fmt.Sprintf("invalid selector: %s", err)}
}

func ErrInvalidPattern(path string, err error) error {
	return errWithPath{path: path, msg: fmt.Sprintf("invalid pattern: %s", err)}
}

func ErrNotAString(path string) error {
	return errWithPath{path: path, msg: ""}
}

func ErrUnrecognizedOperator(path string, op string) error {
	return errWithPath{path: path, msg: fmt.Sprintf("unrecognized operator '%s'", safeStr(op))}
}

func ErrUnrecognizedShape(path string) error {
	return errWithPath{path: path, msg: "unrecognized shape"}
}

func ErrNotATuple(path string) error {
	return errWithPath{path: path, msg: "not a tuple"}
}

func ErrEmptyList(path string) error {
	return errWithPath{path: path, msg: "empty list"}
}

func safeStr(str string) string {
	if len(str) > 10 {
		return str[:10]
	}
	return str
}
