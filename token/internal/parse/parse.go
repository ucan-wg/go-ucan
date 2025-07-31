package parse

import (
	"fmt"
	"time"

	"github.com/MetaMask/go-did-it"

	"github.com/ucan-wg/go-ucan/pkg/policy/limits"
)

func OptionalDID(s *string) (did.DID, error) {
	if s == nil {
		return nil, nil
	}
	return did.Parse(*s)
}

func OptionalTimestamp(sec *int64) (*time.Time, error) {
	if sec == nil {
		return nil, nil
	}

	if *sec > limits.MaxInt53 || *sec < limits.MinInt53 {
		return nil, fmt.Errorf("timestamp value %d exceeds safe integer bounds", *sec)
	}

	t := time.Unix(*sec, 0)
	return &t, nil
}
