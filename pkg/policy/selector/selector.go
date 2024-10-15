package selector

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	"github.com/ipld/go-ipld-prime/schema"
)

// Selector describes a UCAN policy selector, as specified here:
// https://github.com/ucan-wg/delegation/blob/4094d5878b58f5d35055a3b93fccda0b8329ebae/README.md#selectors
type Selector []segment

// Select perform the selection described by the selector on the input IPLD DAG.
// If no error, Select returns either one ipld.Node or a []ipld.Node.
func (s Selector) Select(subject ipld.Node) (ipld.Node, []ipld.Node, error) {
	return resolve(s, subject, nil)
}

func (s Selector) String() string {
	var res strings.Builder
	for _, seg := range s {
		res.WriteString(seg.String())
	}
	return res.String()
}

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

// String returns the segment's string representation.
func (s segment) String() string {
	return s.str
}

// Identity flags that this selector is the identity selector.
func (s segment) Identity() bool {
	return s.identity
}

// Optional flags that this selector is optional.
func (s segment) Optional() bool {
	return s.optional
}

// Iterator flags that this selector is an iterator segment.
func (s segment) Iterator() bool {
	return s.iterator
}

// Slice flags that this segment targets a range of a slice.
func (s segment) Slice() []int {
	return s.slice
}

// Field is the name of a field in a struct/map.
func (s segment) Field() string {
	return s.field
}

// Index is an index of a slice.
func (s segment) Index() int {
	return s.index
}

func resolve(sel Selector, subject ipld.Node, at []string) (ipld.Node, []ipld.Node, error) {
	cur := subject
	for i, seg := range sel {
		if seg.Identity() {
			continue
		}

		// 1st level: handle the different segment types (iterator, field, slice, index)
		// 2nd level: handle different node kinds (list, map, string, bytes)
		switch {
		case seg.Iterator():
			if cur == nil || cur.Kind() == datamodel.Kind_Null {
				if seg.Optional() {
					// build empty list
					nb := basicnode.Prototype.List.NewBuilder()
					assembler, err := nb.BeginList(0)
					if err != nil {
						return nil, nil, err
					}

					if err = assembler.Finish(); err != nil {
						return nil, nil, err
					}

					return nb.Build(), nil, nil
				} else {
					return nil, nil, newResolutionError(fmt.Sprintf("can not iterate over kind: %s", kindString(cur)), at)
				}
			} else {
				var many []ipld.Node
				switch cur.Kind() {
				case datamodel.Kind_List:
					it := cur.ListIterator()
					for !it.Done() {
						_, v, err := it.Next()
						if err != nil {
							return nil, nil, err
						}

						// check if there are more iterator segments
						if len(sel) > i+1 && sel[i+1].Iterator() {
							if v.Kind() == datamodel.Kind_List {
								// recursively resolve the remaining selector segments
								var o ipld.Node
								var m []ipld.Node
								o, m, err = resolve(sel[i+1:], v, at)
								if err != nil {
									// if the segment is optional and an error occurs, skip the current iteration.
									if seg.Optional() {
										continue
									} else {
										return nil, nil, err
									}
								}
								if m != nil {
									many = append(many, m...)
								} else if o != nil {
									many = append(many, o)
								}
							} else {
								// if the current value is not a list and the next segment is optional, skip the current iteration
								if sel[i+1].Optional() {
									continue
								} else {
									return nil, nil, newResolutionError(fmt.Sprintf("can not iterate over kind: %s", kindString(v)), at)
								}
							}
						} else {
							// if there are no more iterator segments, append the current value to the result
							many = append(many, v)
						}
					}
				case datamodel.Kind_Map:
					it := cur.MapIterator()
					for !it.Done() {
						_, v, err := it.Next()
						if err != nil {
							return nil, nil, err
						}

						if len(sel) > i+1 && sel[i+1].Iterator() {
							if v.Kind() == datamodel.Kind_List {
								var o ipld.Node
								var m []ipld.Node
								o, m, err = resolve(sel[i+1:], v, at)
								if err != nil {
									if seg.Optional() {
										continue
									} else {
										return nil, nil, err
									}
								}
								if m != nil {
									many = append(many, m...)
								} else if o != nil {
									many = append(many, o)
								}
							} else {
								if sel[i+1].Optional() {
									continue
								} else {
									return nil, nil, newResolutionError(fmt.Sprintf("can not iterate over kind: %s", kindString(v)), at)
								}
							}
						} else {
							many = append(many, v)
						}
					}
				default:
					return nil, nil, newResolutionError(fmt.Sprintf("can not iterate over kind: %s", kindString(cur)), at)
				}

				return nil, many, nil
			}

		case seg.Field() != "":
			at = append(at, seg.Field())
			if cur == nil {
				if seg.Optional() {
					cur = nil
				} else {
					return nil, nil, newResolutionError(fmt.Sprintf("can not access field: %s on kind: %s", seg.Field(), kindString(cur)), at)
				}
			} else {
				switch cur.Kind() {
				case datamodel.Kind_Map:
					n, err := cur.LookupByString(seg.Field())
					if err != nil {
						if isMissing(err) {
							if seg.Optional() {
								cur = nil
							} else {
								return nil, nil, newResolutionError(fmt.Sprintf("object has no field named: %s", seg.Field()), at)
							}
						} else {
							return nil, nil, err
						}
					} else {
						cur = n
					}
				case datamodel.Kind_List:
					var many []ipld.Node
					it := cur.ListIterator()
					for !it.Done() {
						_, v, err := it.Next()
						if err != nil {
							return nil, nil, err
						}
						if v.Kind() == datamodel.Kind_Map {
							n, err := v.LookupByString(seg.Field())
							if err == nil {
								many = append(many, n)
							}
						}
					}
					if len(many) > 0 {
						cur = nil
						return nil, many, nil
					} else if seg.Optional() {
						cur = nil
					} else {
						return nil, nil, newResolutionError(fmt.Sprintf("no elements in list have field named: %s", seg.Field()), at)
					}
				default:
					if seg.Optional() {
						cur = nil
					} else {
						return nil, nil, newResolutionError(fmt.Sprintf("can not access field: %s on kind: %s", seg.Field(), kindString(cur)), at)
					}
				}
			}

		case seg.Slice() != nil:
			if cur == nil {
				if seg.Optional() {
					cur = nil
				} else {
					return nil, nil, newResolutionError(fmt.Sprintf("can not slice on kind: %s", kindString(cur)), at)
				}
			} else {
				slice := seg.Slice()
				var start, end, length int64
				switch cur.Kind() {
				case datamodel.Kind_List:
					length = cur.Length()
					start, end = resolveSliceIndices(slice, length)
				case datamodel.Kind_Bytes:
					b, _ := cur.AsBytes()
					length = int64(len(b))
					start, end = resolveSliceIndices(slice, length)
				case datamodel.Kind_String:
					str, _ := cur.AsString()
					length = int64(len(str))
					start, end = resolveSliceIndices(slice, length)
				default:
					return nil, nil, newResolutionError(fmt.Sprintf("can not slice on kind: %s", kindString(cur)), at)
				}

				if start < 0 || end < start || end > length {
					if seg.Optional() {
						cur = nil
					} else {
						return nil, nil, newResolutionError(fmt.Sprintf("slice out of bounds: [%d:%d]", start, end), at)
					}
				} else {
					switch cur.Kind() {
					case datamodel.Kind_List:
						if end > cur.Length() {
							end = cur.Length()
						}
						nb := basicnode.Prototype.List.NewBuilder()
						assembler, _ := nb.BeginList(int64(end - start))
						for i := start; i < end; i++ {
							item, _ := cur.LookupByIndex(int64(i))
							assembler.AssembleValue().AssignNode(item)
						}
						assembler.Finish()
						cur = nb.Build()
					case datamodel.Kind_Bytes:
						b, _ := cur.AsBytes()
						l := int64(len(b))
						if end > l {
							end = l
						}
						cur = basicnode.NewBytes(b[start:end])
					case datamodel.Kind_String:
						str, _ := cur.AsString()
						l := int64(len(str))
						if end > l {
							end = l
						}
						cur = basicnode.NewString(str[start:end])
					}
				}
			}

		default:
			at = append(at, fmt.Sprintf("%d", seg.Index()))
			if cur == nil {
				if seg.Optional() {
					cur = nil
				} else {
					return nil, nil, newResolutionError(fmt.Sprintf("can not access index: %d on kind: %s", seg.Index(), kindString(cur)), at)
				}
			} else {
				idx := seg.Index()
				switch cur.Kind() {
				case datamodel.Kind_List:
					if idx < 0 {
						idx = int(cur.Length()) + idx
					}
					if idx < 0 || idx >= int(cur.Length()) {
						if seg.Optional() {
							cur = nil
						} else {
							return nil, nil, newResolutionError(fmt.Sprintf("index out of bounds: %d", seg.Index()), at)
						}
					} else {
						cur, _ = cur.LookupByIndex(int64(idx))
					}
				case datamodel.Kind_String:
					str, _ := cur.AsString()
					if idx < 0 {
						idx = len(str) + idx
					}
					if idx < 0 || idx >= len(str) {
						if seg.Optional() {
							cur = nil
						} else {
							return nil, nil, newResolutionError(fmt.Sprintf("index out of bounds: %d", seg.Index()), at)
						}
					} else {
						cur = basicnode.NewString(string(str[idx]))
					}
				case datamodel.Kind_Bytes:
					b, _ := cur.AsBytes()
					if idx < 0 {
						idx = len(b) + idx
					}
					if idx < 0 || idx >= len(b) {
						if seg.Optional() {
							cur = nil
						} else {
							return nil, nil, newResolutionError(fmt.Sprintf("index out of bounds: %d", seg.Index()), at)
						}
					} else {
						cur = basicnode.NewInt(int64(b[idx]))
					}
				default:
					return nil, nil, newResolutionError(fmt.Sprintf("can not access index: %d on kind: %s", seg.Index(), kindString(cur)), at)
				}
			}
		}
	}

	return cur, nil, nil
}

// resolveSliceIndices resolves the start and end indices for slicing a list or byte array.
//
// It takes the slice indices from the selector segment and the length of the list or byte array,
// and returns the resolved start and end indices. Negative indices are supported.
//
// Parameters:
//   - slice: The slice indices from the selector segment.
//   - length: The length of the list or byte array being sliced.
//
// Returns:
//   - start: The resolved start index for slicing.
//   - end: The resolved end index for slicing.
func resolveSliceIndices(slice []int, length int64) (int64, int64) {
	start, end := int64(0), length
	if len(slice) > 0 {
		start = int64(slice[0])
		if start < 0 {
			start = length + start
			if start < 0 {
				start = 0
			}
		}
	}
	if len(slice) > 1 {
		end = int64(slice[1])
		if end <= 0 {
			end = length + end
			if end < start {
				end = start
			}
		}
	}

	return start, end
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

func newResolutionError(message string, at []string) error {
	return resolutionerr{message, at}
}
