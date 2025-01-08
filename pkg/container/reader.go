package container

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"iter"
	"strings"

	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/cbor"
	"github.com/ipld/go-ipld-prime/datamodel"

	"github.com/ucan-wg/go-ucan/token"
	"github.com/ucan-wg/go-ucan/token/delegation"
	"github.com/ucan-wg/go-ucan/token/invocation"
)

var ErrNotFound = fmt.Errorf("not found")
var ErrMultipleInvocations = fmt.Errorf("multiple invocations")

// Reader is a token container reader. It exposes the tokens conveniently decoded.
type Reader map[cid.Cid]token.Token

// FromBytes decodes a container from a []byte
func FromBytes(data []byte) (Reader, error) {
	return FromReader(bytes.NewReader(data))
}

// FromString decodes a container from a string
func FromString(s string) (Reader, error) {
	return FromReader(strings.NewReader(s))
}

// FromReader decodes a container from an io.Reader.
func FromReader(r io.Reader) (Reader, error) {
	payload, err := decodePayload(r)
	if err != nil {
		return nil, err
	}

	n, err := ipld.DecodeStreaming(payload, cbor.Decode)
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
	if version != containerVersionTag {
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
// If more than one invocation exists, ErrMultipleInvocations is returned.
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

func (ctn Reader) addToken(data []byte) error {
	tkn, c, err := token.FromSealed(data)
	if err != nil {
		return err
	}
	ctn[c] = tkn
	return nil
}
