package container

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"iter"

	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/datamodel"

	"github.com/ucan-wg/go-ucan/token"
	"github.com/ucan-wg/go-ucan/token/delegation"
	"github.com/ucan-wg/go-ucan/token/invocation"
)

var ErrNotFound = fmt.Errorf("not found")
var ErrMultipleInvocations = fmt.Errorf("multiple invocations")

// Reader is a token container reader. It exposes the tokens conveniently decoded.
type Reader map[cid.Cid]token.Token

// GetToken returns an arbitrary decoded token, from its CID.
// If not found, ErrNotFound is returned.
func (ctn Reader) GetToken(cid cid.Cid) (token.Token, error) {
	tkn, ok := ctn[cid]
	if !ok {
		return nil, ErrNotFound
	}
	return tkn, nil
}

// GetDelegation is the same as GetToken but only return a delegation.Token, with the right type.
func (ctn Reader) GetDelegation(cid cid.Cid) (*delegation.Token, error) {
	tkn, err := ctn.GetToken(cid)
	if errors.Is(err, ErrNotFound) {
		return nil, delegation.ErrDelegationNotFound
	}
	if err != nil {
		return nil, err
	}
	if tkn, ok := tkn.(*delegation.Token); ok {
		return tkn, nil
	}
	return nil, delegation.ErrDelegationNotFound
}

// GetAllDelegations returns all the delegation.Token in the container.
func (ctn Reader) GetAllDelegations() iter.Seq2[cid.Cid, *delegation.Token] {
	return func(yield func(cid.Cid, *delegation.Token) bool) {
		for c, t := range ctn {
			if t, ok := t.(*delegation.Token); ok {
				if !yield(c, t) {
					return
				}
			}
		}
	}
}

// GetInvocation returns a single invocation.Token.
// If none are found, ErrNotFound is returned.
// If more than one invocation exist, ErrMultipleInvocations is returned.
func (ctn Reader) GetInvocation() (*invocation.Token, error) {
	var res *invocation.Token
	for _, t := range ctn {
		if inv, ok := t.(*invocation.Token); ok {
			if res != nil {
				return nil, ErrMultipleInvocations
			}
			res = inv
		}
	}
	if res == nil {
		return nil, ErrNotFound
	}
	return res, nil
}

// GetAllInvocations returns all the invocation.Token in the container.
func (ctn Reader) GetAllInvocations() iter.Seq2[cid.Cid, *invocation.Token] {
	return func(yield func(cid.Cid, *invocation.Token) bool) {
		for c, t := range ctn {
			if t, ok := t.(*invocation.Token); ok {
				if !yield(c, t) {
					return
				}
			}
		}
	}
}

// FromCbor decodes a DAG-CBOR encoded container.
func FromCbor(data []byte) (Reader, error) {
	return FromCborReader(bytes.NewReader(data))
}

// FromCborReader is the same as FromCbor, but with an io.Reader.
func FromCborReader(r io.Reader) (Reader, error) {
	n, err := ipld.DecodeStreaming(r, dagcbor.Decode)
	if err != nil {
		return nil, err
	}
	if n.Kind() != datamodel.Kind_Map {
		return nil, fmt.Errorf("invalid container format: expected map")
	}
	if n.Length() != 1 {
		return nil, fmt.Errorf("invalid container format: expected single version key")
	}

	// get the first (and only) key-value pair
	it := n.MapIterator()
	key, tokensNode, err := it.Next()
	if err != nil {
		return nil, err
	}

	version, err := key.AsString()
	if err != nil {
		return nil, fmt.Errorf("invalid container format: version must be string")
	}
	if version != currentContainerVersion {
		return nil, fmt.Errorf("unsupported container version: %s", version)
	}

	if tokensNode.Kind() != datamodel.Kind_List {
		return nil, fmt.Errorf("invalid container format: tokens must be a list")
	}

	ctn := make(Reader, tokensNode.Length())
	it2 := tokensNode.ListIterator()
	for !it2.Done() {
		_, val, err := it2.Next()
		if err != nil {
			return nil, err
		}
		data, err := val.AsBytes()
		if err != nil {
			return nil, err
		}
		err = ctn.addToken(data)
		if err != nil {
			return nil, err
		}
	}
	return ctn, nil
}

// FromCborBase64 decodes a base64 DAG-CBOR encoded container.
func FromCborBase64(data []byte) (Reader, error) {
	return FromCborBase64Reader(bytes.NewReader(data))
}

// FromCborBase64Reader is the same as FromCborBase64, but with an io.Reader.
func FromCborBase64Reader(r io.Reader) (Reader, error) {
	return FromCborReader(base64.NewDecoder(base64.StdEncoding, r))
}

// FromCar decodes a CAR file encoded container.
func FromCar(data []byte) (Reader, error) {
	return FromCarReader(bytes.NewReader(data))
}

// FromCarReader is the same as FromCar, but with an io.Reader.
func FromCarReader(r io.Reader) (Reader, error) {
	_, it, err := readCar(r)
	if err != nil {
		return nil, err
	}

	ctn := make(Reader)

	for block, err := range it {
		if err != nil {
			return nil, err
		}

		err = ctn.addToken(block.data)
		if err != nil {
			return nil, err
		}
	}

	return ctn, nil
}

// FromCarBase64 decodes a base64 CAR file encoded container.
func FromCarBase64(data []byte) (Reader, error) {
	return FromCarReader(bytes.NewReader(data))
}

// FromCarBase64Reader is the same as FromCarBase64, but with an io.Reader.
func FromCarBase64Reader(r io.Reader) (Reader, error) {
	return FromCarReader(base64.NewDecoder(base64.StdEncoding, r))
}

func (ctn Reader) addToken(data []byte) error {
	tkn, c, err := token.FromSealed(data)
	if err != nil {
		return err
	}
	ctn[c] = tkn
	return nil
}
