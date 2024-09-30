package tokens

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec"
	"github.com/ipld/go-ipld-prime/codec/cbor"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"github.com/ipld/go-ipld-prime/codec/json"
	"github.com/ipld/go-ipld-prime/codec/raw"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ucan-wg/go-ucan/tokens/internal/envelope"
)

// EnvelopeInfo describes the fields of an Envelope enclosing a UCAN token.
type EnvelopeInfo struct {
	Signature    []byte
	Tag          string
	VarsigHeader []byte
}

// InspectEnvelope accepts arbitrary data and attempts to decode it as a
// UCAN token's Envelope.
func Inspect(data []byte) (EnvelopeInfo, error) {
	undef := EnvelopeInfo{}

	node, err := decodeAny(data)
	if err != nil {
		return undef, err
	}

	info, err := envelope.Inspect(node)
	if err != nil {
		return undef, err
	}

	iterator := info.SigPayloadNode.MapIterator()
	foundVarsigHeader := false
	foundTokenPayload := false
	tag := ""
	i := 0

	for !iterator.Done() {
		k, _, err := iterator.Next()
		if err != nil {
			return undef, err
		}

		key, err := k.AsString()
		if err != nil {
			return undef, err
		}

		if key == envelope.VarsigHeaderKey {
			foundVarsigHeader = true
			i++

			continue
		}

		if strings.HasPrefix(key, envelope.UCANTagPrefix) {
			tag = key
			foundTokenPayload = true
			i++
		}
	}

	if i != 2 {
		return undef, fmt.Errorf("expected two and only two fields in SigPayload: %d", i)
	}

	if !foundVarsigHeader {
		return undef, errors.New("failed to find VarsigHeader field")
	}

	if !foundTokenPayload {
		return undef, errors.New("failed to find TokenPayload field")
	}

	return EnvelopeInfo{
		Signature:    info.Signature,
		Tag:          tag,
		VarsigHeader: info.VarsigHeader,
	}, nil
}

func decodeAny(data []byte) (datamodel.Node, error) {
	for _, decoder := range []codec.Decoder{
		dagcbor.Decode,
		dagjson.Decode,
		cbor.Decode,
		json.Decode,
		raw.Decode,
	} {
		if node, err := ipld.Decode(data, decoder); err == nil {
			return node, nil
		}
	}

	return nil, errors.New("failed to decode (any) the provided data")
}
