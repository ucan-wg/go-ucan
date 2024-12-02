package limits

import (
	"testing"

	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/fluent/qp"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	"github.com/stretchr/testify/require"
)

func TestValidateIntegerBoundsIPLD(t *testing.T) {
	buildMap := func() datamodel.Node {
		nb := basicnode.Prototype.Any.NewBuilder()
		qp.Map(1, func(ma datamodel.MapAssembler) {
			qp.MapEntry(ma, "foo", qp.Int(MaxInt53+1))
		})(nb)
		return nb.Build()
	}

	buildList := func() datamodel.Node {
		nb := basicnode.Prototype.Any.NewBuilder()
		qp.List(1, func(la datamodel.ListAssembler) {
			qp.ListEntry(la, qp.Int(MinInt53-1))
		})(nb)
		return nb.Build()
	}

	tests := []struct {
		name    string
		input   datamodel.Node
		wantErr bool
	}{
		{
			name:    "valid int",
			input:   basicnode.NewInt(42),
			wantErr: false,
		},
		{
			name:    "max safe int",
			input:   basicnode.NewInt(MaxInt53),
			wantErr: false,
		},
		{
			name:    "min safe int",
			input:   basicnode.NewInt(MinInt53),
			wantErr: false,
		},
		{
			name:    "above MaxInt53",
			input:   basicnode.NewInt(MaxInt53 + 1),
			wantErr: true,
		},
		{
			name:    "below MinInt53",
			input:   basicnode.NewInt(MinInt53 - 1),
			wantErr: true,
		},
		{
			name:    "nested map with invalid int",
			input:   buildMap(),
			wantErr: true,
		},
		{
			name:    "nested list with invalid int",
			input:   buildList(),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateIntegerBoundsIPLD(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), "exceeds safe bounds")
			} else {
				require.NoError(t, err)
			}
		})
	}
}
