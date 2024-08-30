package delegation

import (
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
)

func (p *PayloadModel) EncodeDagCbor() ([]byte, error) {
	return ipld.Marshal(dagcbor.Encode, p, PayloadType())
}

func (p *PayloadModel) EncodeDagJson() ([]byte, error) {
	return ipld.Marshal(dagjson.Encode, p, PayloadType())
}

func DecodeDagCbor(data []byte) (*PayloadModel, error) {
	var p PayloadModel
	_, err := ipld.Unmarshal(data, dagcbor.Decode, &p, PayloadType())
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func DecodeDagJson(data []byte) (*PayloadModel, error) {
	var p PayloadModel
	_, err := ipld.Unmarshal(data, dagjson.Decode, &p, PayloadType())
	if err != nil {
		return nil, err
	}
	return &p, nil
}
