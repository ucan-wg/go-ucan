package parse

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ucan-wg/go-ucan/pkg/policy/limits"
)

func TestOptionalTimestamp(t *testing.T) {
	tests := []struct {
		name    string
		input   *int64
		wantErr bool
	}{
		{
			name:    "nil timestamp",
			input:   nil,
			wantErr: false,
		},
		{
			name:    "valid timestamp",
			input:   int64Ptr(1625097600),
			wantErr: false,
		},
		{
			name:    "max safe integer",
			input:   int64Ptr(limits.MaxInt53),
			wantErr: false,
		},
		{
			name:    "exceeds max safe integer",
			input:   int64Ptr(limits.MaxInt53 + 1),
			wantErr: true,
		},
		{
			name:    "below min safe integer",
			input:   int64Ptr(limits.MinInt53 - 1),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := OptionalTimestamp(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), "exceeds safe integer bounds")
				require.Nil(t, result)
			} else {
				require.NoError(t, err)
				if tt.input == nil {
					require.Nil(t, result)
				} else {
					require.NotNil(t, result)
				}
			}
		})
	}
}

func int64Ptr(i int64) *int64 {
	return &i
}
