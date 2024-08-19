package selector

import (
	"fmt"
	"strconv"
	"strings"
)

// Selector describes a UCAN policy selector, as specified here:
// https://github.com/ucan-wg/delegation/blob/4094d5878b58f5d35055a3b93fccda0b8329ebae/README.md#selectors
type Selector interface {
	// Identity flags that this selector is the identity selector.
	Identity() bool
	// Optional flags that this selector is optional.
	Optional() bool
	// Field is the name of a field in a struct/map.
	Field() string
	// Index is an index of a slice.
	Index() int
	// String returns the selector's string representation.
	String() string
}

type selector struct {
	str      string
	identity bool
	optional bool
	field    string
	index    int
}

func (s selector) Field() string {
	return s.field
}

func (s selector) Identity() bool {
	return s.identity
}

func (s selector) Index() int {
	return s.index
}

func (s selector) Optional() bool {
	return s.optional
}

func (s selector) String() string {
	return s.str
}

// TODO: probably regex or better parser
func Parse(sel string) (Selector, error) {
	s := sel
	if s == "." {
		return selector{sel, true, false, "", 0}, nil
	}

	optional := strings.HasSuffix(s, "?")
	if optional {
		s = s[0 : len(s)-1]
	}

	dotted := strings.HasPrefix(s, ".")
	if dotted {
		s = s[1:]
	}

	// collection values
	if s == "[]" {
		return nil, fmt.Errorf("unsupported selector: %s", sel)
	}

	if strings.HasPrefix(s, "[") && strings.HasSuffix(s, "]") {
		s = s[1 : len(s)-1]

		// explicit field selector
		if strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"") {
			return selector{sel, false, optional, s[1 : len(s)-1], 0}, nil
		}

		// collection range
		if strings.Contains(s, ":") {
			return nil, fmt.Errorf("unsupported selector: %s", sel)
		}

		// index selector
		idx, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("parsing index selector value: %s", err)
		}

		return selector{sel, false, optional, "", idx}, nil
	}

	if !dotted {
		return nil, fmt.Errorf("invalid selector: %s", sel)
	}

	// dotted field selector
	return selector{sel, false, optional, s, 0}, nil
}

func MustParse(sel string) Selector {
	s, err := Parse(sel)
	if err != nil {
		panic(err)
	}
	return s
}
