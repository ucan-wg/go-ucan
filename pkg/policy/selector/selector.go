package selector

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/fluent/qp"
	"github.com/ipld/go-ipld-prime/node/basicnode"
)

// Selector describes a UCAN policy selector, as specified here:
// https://github.com/ucan-wg/delegation/blob/4094d5878b58f5d35055a3b93fccda0b8329ebae/README.md#selectors
type Selector []segment

// Select perform the selection described by the selector on the input IPLD DAG.
// Select can return:
//   - exactly one matched IPLD node
//   - a resolutionerr error if not being able to resolve to a node
//   - nil and no errors, if the selector couldn't match on an optional segment (with ?).
func (s Selector) Select(subject ipld.Node) (ipld.Node, error) {
	return resolve(s, subject, nil)
}

func (s Selector) String() string {
	var res strings.Builder
	for _, seg := range s {
		res.WriteString(seg.String())
	}
	return res.String()
}

type segment struct {
	str      string
	identity bool
	optional bool
	iterator bool
	slice    []int64 // either 0-length or 2-length
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
func (s segment) Slice() []int64 {
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

func resolve(sel Selector, subject ipld.Node, at []string) (ipld.Node, error) {
	errIfNotOptional := func(s segment, err error) error {
		if !s.Optional() {
			return err
		}
		return nil
	}

	cur := subject
	for _, seg := range sel {
		// 1st level: handle the different segment types (iterator, field, slice, index)
		// 2nd level: handle different node kinds (list, map, string, bytes)
		switch {
		case seg.Identity():
			continue

		case seg.Iterator():
			switch {
			case cur == nil || cur.Kind() == datamodel.Kind_Null:
				if seg.Optional() {
					// build an empty list
					n, _ := qp.BuildList(basicnode.Prototype.Any, 0, func(_ datamodel.ListAssembler) {})
					return n, nil
				}
				return nil, newResolutionError(fmt.Sprintf("can not iterate over kind: %s", kindString(cur)), at)

			case cur.Kind() == datamodel.Kind_List:
				// iterators are no-op on list
				continue

			case cur.Kind() == datamodel.Kind_Map:
				// iterators on maps collect the values
				nd, err := qp.BuildList(basicnode.Prototype.Any, cur.Length(), func(l datamodel.ListAssembler) {
					it := cur.MapIterator()
					for !it.Done() {
						_, v, err := it.Next()
						if err != nil {
							// recovered by BuildList
							// Error is bubbled up, but should never occur as we already checked the type,
							// and are using the iterator correctly.
							// This is verified with fuzzing.
							panic(err)
						}
						qp.ListEntry(l, qp.Node(v))
					}
				})
				if err != nil {
					panic("should never happen")
				}
				return nd, nil

			default:
				return nil, newResolutionError(fmt.Sprintf("can not iterate over kind: %s", kindString(cur)), at)
			}

		case seg.Field() != "":
			at = append(at, seg.Field())
			switch {
			case cur == nil:
				err := newResolutionError(fmt.Sprintf("can not access field: %s on kind: %s", seg.Field(), kindString(cur)), at)
				return nil, errIfNotOptional(seg, err)

			case cur.Kind() == datamodel.Kind_Map:
				n, err := cur.LookupByString(seg.Field())
				if err != nil {
					// the only possible error is missing field as we already check the type
					if seg.Optional() {
						cur = nil
					} else {
						return nil, newResolutionError(fmt.Sprintf("object has no field named: %s", seg.Field()), at)
					}
				} else {
					cur = n
				}

			default:
				err := newResolutionError(fmt.Sprintf("can not access field: %s on kind: %s", seg.Field(), kindString(cur)), at)
				return nil, errIfNotOptional(seg, err)
			}

		case len(seg.Slice()) > 0:
			if cur == nil {
				err := newResolutionError(fmt.Sprintf("can not slice on kind: %s", kindString(cur)), at)
				return nil, errIfNotOptional(seg, err)
			}

			slice := seg.Slice()

			switch cur.Kind() {
			case datamodel.Kind_List:
				start, end := resolveSliceIndices(slice, cur.Length())
				sliced, err := qp.BuildList(basicnode.Prototype.Any, end-start, func(l datamodel.ListAssembler) {
					for i := start; i < end; i++ {
						item, err := cur.LookupByIndex(i)
						if err != nil {
							// recovered by BuildList
							// Error is bubbled up, but should never occur as we already checked the type and boundaries
							// This is verified with fuzzing.
							panic(err)
						}
						qp.ListEntry(l, qp.Node(item))
					}
				})
				if err != nil {
					panic("should never happen")
				}
				cur = sliced

			case datamodel.Kind_Bytes:
				b, _ := cur.AsBytes()
				start, end := resolveSliceIndices(slice, int64(len(b)))
				cur = basicnode.NewBytes(b[start:end])

			case datamodel.Kind_String:
				str, _ := cur.AsString()
				runes := []rune(str)
				start, end := resolveSliceIndices(slice, int64(len(runes)))
				cur = basicnode.NewString(string(runes[start:end]))

			default:
				return nil, newResolutionError(fmt.Sprintf("can not slice on kind: %s", kindString(cur)), at)
			}

		default: // Index()
			at = append(at, strconv.Itoa(seg.Index()))

			if cur == nil {
				err := newResolutionError(fmt.Sprintf("can not access index: %d on kind: %s", seg.Index(), kindString(cur)), at)
				return nil, errIfNotOptional(seg, err)
			}

			idx := seg.Index()
			switch cur.Kind() {
			case datamodel.Kind_List:
				if idx < 0 {
					idx = int(cur.Length()) + idx
				}
				if idx < 0 || idx >= int(cur.Length()) {
					err := newResolutionError(fmt.Sprintf("index out of bounds: %d", seg.Index()), at)
					return nil, errIfNotOptional(seg, err)
				}
				cur, _ = cur.LookupByIndex(int64(idx))

			case datamodel.Kind_Bytes:
				b, _ := cur.AsBytes()
				if idx < 0 {
					idx = len(b) + idx
				}
				if idx < 0 || idx >= len(b) {
					err := newResolutionError(fmt.Sprintf("index %d out of bounds for bytes of length %d", seg.Index(), len(b)), at)
					return nil, errIfNotOptional(seg, err)
				}
				cur = basicnode.NewInt(int64(b[idx]))

			default:
				return nil, newResolutionError(fmt.Sprintf("can not access index: %d on kind: %s", seg.Index(), kindString(cur)), at)
			}
		}
	}

	// segment exhausted, we return where we are
	return cur, nil
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
//   - end: The resolved **excluded** end index for slicing.
func resolveSliceIndices(slice []int64, length int64) (start int64, end int64) {
	if len(slice) != 2 {
		panic("should always be 2-length")
	}

	start, end = slice[0], slice[1]

	// adjust boundaries
	switch {
	case slice[0] == math.MinInt:
		start = 0
	case slice[0] < 0:
		start = length + slice[0]
	}
	switch {
	case slice[1] == math.MaxInt:
		end = length
	case slice[1] < 0:
		end = length + slice[1]
	}

	// backward iteration is not allowed, shortcut to an empty result
	if start >= end {
		start, end = 0, 0
	}
	// clamp out of bound
	if start < 0 {
		start = 0
	}
	if start > length {
		start = length
	}
	if end > length {
		end = length
	}

	return start, end
}

func kindString(n datamodel.Node) string {
	if n == nil {
		return "null"
	}
	return n.Kind().String()
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
