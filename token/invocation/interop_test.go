package invocation_test

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"github.com/ipld/go-ipld-prime/schema"
	"github.com/stretchr/testify/require"

	"github.com/ucan-wg/go-ucan/token/delegation"
	"github.com/ucan-wg/go-ucan/token/internal/envelope"
	"github.com/ucan-wg/go-ucan/token/invocation"
)

// This comes from https://github.com/ucan-wg/spec/blob/main/fixtures/1.0.0/invocation.json
//
//go:embed testdata/interop_invocation.json
var interopInvocation []byte

//go:embed testdata/interop.ipldsch
var schemaBytes []byte

func fixturesType(t *testing.T) schema.Type {
	ts, err := ipld.LoadSchemaBytes(schemaBytes)
	require.NoError(t, err)
	return ts.TypeByName("Fixtures")
}

type ErrorModel struct {
	Name string
}

type VectorModel struct {
	Name        string
	Description string
	Invocation  []byte
	Proofs      [][]byte
	Error       *ErrorModel
}

type FixturesModel struct {
	Version  string
	Comments string
	Valid    []VectorModel
	Invalid  []VectorModel
}

func TestInterop(t *testing.T) {
	var fixtures FixturesModel
	_, err := ipld.Unmarshal(interopInvocation, dagjson.Decode, &fixtures, fixturesType(t))
	require.NoError(t, err)

	for _, vector := range fixtures.Valid {
		t.Run("valid "+vector.Name, func(t *testing.T) {
			err := decodeAndValidate(vector)
			require.NoError(t, err, "execution should have been allowed for invocation with %s", vector.Description)
		})
	}

	for _, vector := range fixtures.Invalid {
		t.Run("invalid "+vector.Name, func(t *testing.T) {
			err := decodeAndValidate(vector)
			require.Error(t, err, "execution should not have been allowed for invocation because %s", vector.Description)
		})
	}
}

type mapDelegationLoader struct {
	data map[cid.Cid]*delegation.Token
}

func (ml *mapDelegationLoader) GetDelegation(cid cid.Cid) (*delegation.Token, error) {
	d, ok := ml.data[cid]
	if !ok {
		return nil, fmt.Errorf("delegation not found: %s", cid)
	}
	return d, nil
}

func decodeAndValidate(vector VectorModel) error {
	inv, err := invocation.Decode(vector.Invocation, dagcbor.Decode)
	if err != nil {
		return err
	}
	proofs, err := decodeProofs(vector.Proofs)
	if err != nil {
		return err
	}
	err = inv.ExecutionAllowed(&mapDelegationLoader{proofs})
	if err != nil {
		return err
	}
	return nil
}

func decodeProofs(vectorProofs [][]byte) (map[cid.Cid]*delegation.Token, error) {
	proofs := map[cid.Cid]*delegation.Token{}
	for _, p := range vectorProofs {
		dlg, err := delegation.Decode(p, dagcbor.Decode)
		if err != nil {
			return nil, err
		}
		c, err := envelope.CIDFromBytes(p)
		if err != nil {
			return nil, err
		}
		proofs[c] = dlg
	}
	return proofs, nil
}
