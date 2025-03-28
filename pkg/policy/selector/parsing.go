package selector

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/ucan-wg/go-ucan/pkg/policy/limits"
)

var (
	indexRegex = regexp.MustCompile(`^-?\d+$`)
	sliceRegex = regexp.MustCompile(`^((\-?\d+:\-?\d*)|(\-?\d*:\-?\d+))$`)

	// Field name requirements:
	// - Must start with ASCII letter, Unicode letter, or underscore
	// - Can contain:
	//   - ASCII letters (a-z, A-Z)
	//   - ASCII digits (0-9)
	//   - Unicode letters (\p{L})
	//   - Dollar sign ($)
	//   - Underscore (_)
	//   - Hyphen (-)
	fieldRegex = regexp.MustCompile(`^\.[a-zA-Z_\p{L}][a-zA-Z0-9$_\p{L}\-]*$`)
)

func Parse(str string) (Selector, error) {
	if len(str) == 0 {
		return nil, newParseError("empty selector", str, 0, "")
	}
	if string(str[0]) != "." {
		return nil, newParseError("selector must start with identity segment '.'", str, 0, string(str[0]))
	}
	if str == "." {
		return Selector{segment{str: ".", identity: true}}, nil
	}
	if str == ".?" {
		return Selector{segment{str: ".?", identity: true, optional: true}}, nil
	}

	col := 0
	var sel Selector
	for _, tok := range tokenize(str) {
		seg := tok
		opt := strings.HasSuffix(tok, "?")
		if opt {
			seg = strings.TrimRight(tok, "?")
		}
		switch {
		case seg == ".":
			if len(sel) > 0 && sel[len(sel)-1].Identity() {
				return nil, newParseError("selector contains unsupported recursive descent segment: '..'", str, col, tok)
			}
			sel = append(sel, segment{str: ".", identity: true})

		case seg == "[]":
			sel = append(sel, segment{str: tok, optional: opt, iterator: true})

		case strings.HasPrefix(seg, "[") && strings.HasSuffix(seg, "]"):
			lookup := seg[1 : len(seg)-1]

			switch {
			// index, [123]
			case indexRegex.MatchString(lookup):
				idx, err := strconv.Atoi(lookup)
				if err != nil {
					return nil, newParseError("invalid index", str, col, tok)
				}
				if int64(idx) > limits.MaxInt53 || int64(idx) < limits.MinInt53 {
					return nil, newParseError(fmt.Sprintf("index %d exceeds safe integer bounds", idx), str, col, tok)
				}
				sel = append(sel, segment{str: tok, optional: opt, index: idx})

			// explicit field, ["abcd"]
			case strings.HasPrefix(lookup, "\"") && strings.HasSuffix(lookup, "\""):
				fieldName := lookup[1 : len(lookup)-1]
				if strings.Contains(fieldName, ":") {
					return nil, newParseError(fmt.Sprintf("invalid segment: %s", seg), str, col, tok)
				}
				sel = append(sel, segment{str: tok, optional: opt, field: fieldName})

			// slice [3:5] or [:5] or [3:], also negative numbers
			case sliceRegex.MatchString(lookup):
				var rng [2]int64
				splt := strings.Split(lookup, ":")
				if splt[0] == "" {
					rng[0] = math.MinInt
				} else {
					i, err := strconv.ParseInt(splt[0], 10, 0)
					if err != nil {
						return nil, newParseError("invalid slice index", str, col, tok)
					}
					if i > limits.MaxInt53 || i < limits.MinInt53 {
						return nil, newParseError(fmt.Sprintf("slice index %d exceeds safe integer bounds", i), str, col, tok)
					}
					rng[0] = i
				}
				if splt[1] == "" {
					rng[1] = math.MaxInt
				} else {
					i, err := strconv.ParseInt(splt[1], 10, 0)
					if err != nil {
						return nil, newParseError("invalid slice index", str, col, tok)
					}
					if i > limits.MaxInt53 || i < limits.MinInt53 {
						return nil, newParseError(fmt.Sprintf("slice index %d exceeds safe integer bounds", i), str, col, tok)
					}
					rng[1] = i
				}
				sel = append(sel, segment{str: tok, optional: opt, slice: rng[:]})

			default:
				return nil, newParseError(fmt.Sprintf("invalid segment: %s", seg), str, col, tok)
			}

		case fieldRegex.MatchString(seg):
			sel = append(sel, segment{str: tok, optional: opt, field: seg[1:]})
		default:
			return nil, newParseError(fmt.Sprintf("invalid segment: %s", seg), str, col, tok)
		}
		col += len(tok)
	}
	return sel, nil
}

func MustParse(sel string) Selector {
	s, err := Parse(sel)
	if err != nil {
		panic(err)
	}
	return s
}

func tokenize(str string) []string {
	var toks []string
	col := 0
	ofs := 0
	ctx := ""

	for col < len(str) {
		char := string(str[col])

		if char == "\"" && string(str[col-1]) != "\\" {
			col++
			if ctx == "\"" {
				ctx = ""
			} else {
				ctx = "\""
			}
			continue
		}

		if ctx == "\"" {
			col++
			continue
		}

		if char == "." || char == "[" {
			if ofs < col {
				toks = append(toks, str[ofs:col])
			}
			ofs = col
		}
		col++
	}

	if ofs < col && ctx != "\"" {
		toks = append(toks, str[ofs:col])
	}

	return toks
}

type parseErr struct {
	msg string
	src string
	col int
	tok string
}

func (p parseErr) Name() string {
	return "ParseError"
}

func (p parseErr) Message() string {
	return p.msg
}

func (p parseErr) Column() int {
	return p.col
}

func (p parseErr) Error() string {
	return p.msg
}

func (p parseErr) Source() string {
	return p.src
}

func (p parseErr) Token() string {
	return p.tok
}

func newParseError(message string, source string, column int, token string) error {
	return parseErr{message, source, column, token}
}
