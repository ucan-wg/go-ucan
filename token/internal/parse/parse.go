package parse

import (
	"time"

	"github.com/ucan-wg/go-ucan/did"
)

func OptionalDID(s *string) (did.DID, error) {
	if s == nil {
		return did.Undef, nil
	}
	return did.Parse(*s)
}

func OptionalTimestamp(sec *int64) *time.Time {
	if sec == nil {
		return nil
	}
	t := time.Unix(*sec, 0)
	return &t
}
