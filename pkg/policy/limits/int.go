package limits

import (
	"fmt"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/must"
)

const (
	// MaxInt53 represents the maximum safe integer in JavaScript (2^53 - 1)
	MaxInt53 = 9007199254740991
	// MinInt53 represents the minimum safe integer in JavaScript (-2^53 + 1)
	MinInt53 = -9007199254740991
)

func ValidateIntegerBoundsIPLD(node ipld.Node) error {
	switch node.Kind() {
	case ipld.Kind_Int:
		val := must.Int(node)
		if val > MaxInt53 || val < MinInt53 {
			return fmt.Errorf("integer value %d exceeds safe bounds", val)
		}
	case ipld.Kind_List:
		it := node.ListIterator()
		for !it.Done() {
			_, v, err := it.Next()
			if err != nil {
				return err
			}
			if err := ValidateIntegerBoundsIPLD(v); err != nil {
				return err
			}
		}
	case ipld.Kind_Map:
		it := node.MapIterator()
		for !it.Done() {
			_, v, err := it.Next()
			if err != nil {
				return err
			}
			if err := ValidateIntegerBoundsIPLD(v); err != nil {
				return err
			}
		}
	}

	return nil
}
