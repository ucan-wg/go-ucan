package selector

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/schema"
)

// Selector describes a UCAN policy selector, as specified here:
// https://github.com/ucan-wg/delegation/blob/4094d5878b58f5d35055a3b93fccda0b8329ebae/README.md#selectors
type Selector []Segment

func (s Selector) String() string {
	var str string
	for _, seg := range s {
		str += seg.String()
	}
	return str
}

type Segment interface {
	// Identity flags that this selector is the identity selector.
	Identity() bool
	// Optional flags that this selector is optional.
	Optional() bool
	// Iterator flags that this selector is an iterator segment.
	Iterator() bool
	// Slice flags that this segemnt targets a range of a slice.
	Slice() []int
	// Field is the name of a field in a struct/map.
	Field() string
	// Index is an index of a slice.
	Index() int
	// String returns the segment's string representation.
	String() string
}

var Identity = segment{".", true, false, false, nil, "", 0}

var (
	indexRegex = regexp.MustCompile(`^-?\d+$`)
	sliceRegex = regexp.MustCompile(`^((\-?\d+:\-?\d*)|(\-?\d*:\-?\d+))$`)
	fieldRegex = regexp.MustCompile(`^\.[a-zA-Z_]*?$`)
)

type segment struct {
	str      string
	identity bool
	optional bool
	iterator bool
	slice    []int
	field    string
	index    int
}

func (s segment) String() string {
	return s.str
}

func (s segment) Identity() bool {
	return s.identity
}

func (s segment) Optional() bool {
	return s.optional
}

func (s segment) Iterator() bool {
	return s.iterator
}

func (s segment) Slice() []int {
	return s.slice
}

func (s segment) Field() string {
	return s.field
}

func (s segment) Index() int {
	return s.index
}

func Parse(str string) (Selector, error) {
	if string(str[0]) != "." {
		return nil, NewParseError("selector must start with identity segment '.'", str, 0, string(str[0]))
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
				return nil, NewParseError("selector contains unsupported recursive descent segment: '..'", str, col, tok)
			}
			sel = append(sel, Identity)
		case "[]":
			sel = append(sel, segment{tok, false, opt, true, nil, "", 0})
		default:
			if strings.HasPrefix(seg, "[") && strings.HasSuffix(seg, "]") {
				lookup := seg[1 : len(seg)-1]

				if indexRegex.MatchString(lookup) { // index
					idx, err := strconv.Atoi(lookup)
					if err != nil {
						return nil, NewParseError("invalid index", str, col, tok)
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
							return nil, NewParseError("invalid slice index", str, col, tok)
						}
						rng = append(rng, i)
					}
					if splt[1] != "" {
						i, err := strconv.Atoi(splt[1])
						if err != nil {
							return nil, NewParseError("invalid slice index", str, col, tok)
						}
						rng = append(rng, i)
					}
					sel = append(sel, segment{str: tok, optional: opt, slice: rng})
				} else {
					return nil, NewParseError(fmt.Sprintf("invalid segment: %s", seg), str, col, tok)
				}
			} else if fieldRegex.MatchString(seg) {
				sel = append(sel, segment{str: tok, optional: opt, field: seg[1:]})
			} else {
				return nil, NewParseError(fmt.Sprintf("invalid segment: %s", seg), str, col, tok)
			}
		}
		col += len(tok)
	}
	return sel, nil
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

type ParseError interface {
	error
	Name() string
	Message() string
	Source() string
	Column() int
	Token() string
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

func NewParseError(message string, source string, column int, token string) error {
	return parseerr{message, source, column, token}
}

func MustParse(sel string) Selector {
	s, err := Parse(sel)
	if err != nil {
		panic(err)
	}
	return s
}

// Select uses a selector to extract an IPLD node or set of nodes from the
// passed subject node.
func Select(sel Selector, subject ipld.Node) (ipld.Node, []ipld.Node, error) {
	return resolve(sel, subject, nil)
}

func resolve(sel Selector, subject ipld.Node, at []string) (ipld.Node, []ipld.Node, error) {
	cur := subject
	for i, seg := range sel {
		if seg.Identity() {
			continue
		} else if seg.Iterator() {
			if cur != nil && cur.Kind() == datamodel.Kind_List {
				var many []ipld.Node
				it := cur.ListIterator()
				for {
					if it.Done() {
						break
					}

					k, v, err := it.Next()
					if err != nil {
						return nil, nil, err
					}

					key := fmt.Sprintf("%d", k)
					o, m, err := resolve(sel[i+1:], v, append(at[:], key))
					if err != nil {
						return nil, nil, err
					}

					if m != nil {
						many = append(many, m...)
					} else {
						many = append(many, o)
					}
				}
				return nil, many, nil
			} else if cur != nil && cur.Kind() == datamodel.Kind_Map {
				var many []ipld.Node
				it := cur.MapIterator()
				for {
					if it.Done() {
						break
					}

					k, v, err := it.Next()
					if err != nil {
						return nil, nil, err
					}

					key, _ := k.AsString()
					o, m, err := resolve(sel[i+1:], v, append(at[:], key))
					if err != nil {
						return nil, nil, err
					}

					if m != nil {
						many = append(many, m...)
					} else {
						many = append(many, o)
					}
				}
				return nil, many, nil
			} else if seg.Optional() {
				cur = nil
			} else {
				return nil, nil, NewResolutionError(fmt.Sprintf("can not iterate over kind: %s", kindString(cur)), at)
			}

		} else if seg.Field() != "" {
			at = append(at, seg.Field())
			if cur != nil && cur.Kind() == datamodel.Kind_Map {
				n, err := cur.LookupByString(seg.Field())
				if err != nil {
					if isMissing(err) {
						if seg.Optional() {
							cur = nil
						} else {
							return nil, nil, NewResolutionError(fmt.Sprintf("object has no field named: %s", seg.Field()), at)
						}
					} else {
						return nil, nil, err
					}
				}
				cur = n
			} else if seg.Optional() {
				cur = nil
			} else {
				return nil, nil, NewResolutionError(fmt.Sprintf("can not access field: %s on kind: %s", seg.Field(), kindString(cur)), at)
			}
		} else if seg.Slice() != nil {
			if cur != nil && cur.Kind() == datamodel.Kind_List {
				return nil, nil, NewResolutionError("list slice selection not yet implemented", at)
			} else if cur != nil && cur.Kind() == datamodel.Kind_Bytes {
				return nil, nil, NewResolutionError("bytes slice selection not yet implemented", at)
			} else if seg.Optional() {
				cur = nil
			} else {
				return nil, nil, NewResolutionError(fmt.Sprintf("can not index: %s on kind: %s", seg.Field(), kindString(cur)), at)
			}
		} else {
			at = append(at, fmt.Sprintf("%d", seg.Index()))
			if cur != nil && cur.Kind() == datamodel.Kind_List {
				n, err := cur.LookupByIndex(int64(seg.Index()))
				if err != nil {
					if isMissing(err) {
						if seg.Optional() {
							cur = nil
						} else {
							return nil, nil, NewResolutionError(fmt.Sprintf("index out of bounds: %d", seg.Index()), at)
						}
					} else {
						return nil, nil, err
					}
				}
				cur = n
			} else if seg.Optional() {
				cur = nil
			} else {
				return nil, nil, NewResolutionError(fmt.Sprintf("can not access field: %s on kind: %s", seg.Field(), kindString(cur)), at)
			}
		}
	}

	return cur, nil, nil
}

func kindString(n datamodel.Node) string {
	if n == nil {
		return "null"
	}
	return n.Kind().String()
}

func isMissing(err error) bool {
	if _, ok := err.(datamodel.ErrNotExists); ok {
		return true
	}
	if _, ok := err.(schema.ErrNoSuchField); ok {
		return true
	}
	if _, ok := err.(schema.ErrInvalidKey); ok {
		return true
	}
	return false
}

type ResolutionError interface {
	error
	Name() string
	Message() string
	At() []string
}

type resolutionerr struct {
	msg string
	at  []string
}

func (r resolutionerr) Name() string {
	return "ResolutionError"
}

func (r resolutionerr) Message() string {
	return fmt.Sprintf("can not resolve path: .%s", strings.Join(r.at, "."))
}

func (r resolutionerr) At() []string {
	return r.at
}

func (r resolutionerr) Error() string {
	return r.Message()
}

func NewResolutionError(message string, at []string) error {
	return resolutionerr{message, at}
}
