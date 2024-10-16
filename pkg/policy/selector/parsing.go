package selector

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	identity   = Selector{segment{str: ".", identity: true}}
	indexRegex = regexp.MustCompile(`^-?\d+$`)
	sliceRegex = regexp.MustCompile(`^((\-?\d+:\-?\d*)|(\-?\d*:\-?\d+))$`)
	fieldRegex = regexp.MustCompile(`^\.[a-zA-Z_]*?$`)
)

func Parse(str string) (Selector, error) {
	if len(str) == 0 {
		return nil, newParseError("empty selector", str, 0, "")
	}
	if string(str[0]) != "." {
		return nil, newParseError("selector must start with identity segment '.'", str, 0, string(str[0]))
	}
	if str == "." {
		return identity, nil
	}

	col := 0
	var sel Selector
	for _, tok := range tokenize(str) {
		seg := tok
		opt := strings.HasSuffix(tok, "?")
		if opt {
			seg = tok[0 : len(tok)-1]
		}
		switch seg {
		case ".":
			if len(sel) > 0 && sel[len(sel)-1].Identity() {
				return nil, newParseError("selector contains unsupported recursive descent segment: '..'", str, col, tok)
			}
			sel = append(sel, segment{str: ".", identity: true})
		case "[]":
			sel = append(sel, segment{str: tok, optional: opt, iterator: true})
		default:
			if strings.HasPrefix(seg, "[") && strings.HasSuffix(seg, "]") {
				lookup := seg[1 : len(seg)-1]

				if indexRegex.MatchString(lookup) { // index
					idx, err := strconv.Atoi(lookup)
					if err != nil {
						return nil, newParseError("invalid index", str, col, tok)
					}
					sel = append(sel, segment{str: tok, optional: opt, index: idx})
				} else if strings.HasPrefix(lookup, "\"") && strings.HasSuffix(lookup, "\"") { // explicit field
					sel = append(sel, segment{str: tok, optional: opt, field: lookup[1 : len(lookup)-1]})
				} else if sliceRegex.MatchString(lookup) { // slice [3:5] or [:5] or [3:]
					var rng []int
					splt := strings.Split(lookup, ":")
					if splt[0] == "" {
						rng = append(rng, 0)
					} else {
						i, err := strconv.Atoi(splt[0])
						if err != nil {
							return nil, newParseError("invalid slice index", str, col, tok)
						}
						rng = append(rng, i)
					}
					if splt[1] != "" {
						i, err := strconv.Atoi(splt[1])
						if err != nil {
							return nil, newParseError("invalid slice index", str, col, tok)
						}
						rng = append(rng, i)
					}
					sel = append(sel, segment{str: tok, optional: opt, slice: rng})
				} else {
					return nil, newParseError(fmt.Sprintf("invalid segment: %s", seg), str, col, tok)
				}
			} else if fieldRegex.MatchString(seg) {
				sel = append(sel, segment{str: tok, optional: opt, field: seg[1:]})
			} else {
				return nil, newParseError(fmt.Sprintf("invalid segment: %s", seg), str, col, tok)
			}
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

type parseerr struct {
	msg string
	src string
	col int
	tok string
}

func (p parseerr) Name() string {
	return "ParseError"
}

func (p parseerr) Message() string {
	return p.msg
}

func (p parseerr) Column() int {
	return p.col
}

func (p parseerr) Error() string {
	return p.msg
}

func (p parseerr) Source() string {
	return p.src
}

func (p parseerr) Token() string {
	return p.tok
}

func newParseError(message string, source string, column int, token string) error {
	return parseerr{message, source, column, token}
}
