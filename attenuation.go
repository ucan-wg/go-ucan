package ucan

import (
	"encoding/json"
)

const CapKey = "cap"

// Attenuations is a list of attenuations
type Attenuations []Attenuation

// Contains is true if all attenuations in b are contained
func (att Attenuations) Contains(b Attenuations) bool {
LOOP:
	for _, x := range b {
		for _, y := range att {
			if y.Contains(x) {
				continue LOOP
			}
		}
		return false
	}
	return true
}

type AttenuationConstructor func(v map[string]interface{}) (Attenuation, error)

type Attenuation struct {
	Cap Capability
	Rsc Resource
}

func (a Attenuation) Contains(b Attenuation) bool {
	return a.Rsc.Type() == b.Rsc.Type() && a.Rsc.Contains(b.Rsc) && a.Cap.Contains(b.Cap)
}

func NewAttenuation(cap Capability, rsc Resource) Attenuation {
	return Attenuation{
		Rsc: rsc,
		Cap: cap,
	}
}

func (a Attenuation) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		a.Rsc.Type(): a.Rsc.Value(),
		CapKey:       a.Cap.String(),
	})
}

// ResourcePool is a pool of type strings to
var ResourcePool map[string]ResourceConstructor

type Resource interface {
	Type() string
	Value() string
	Contains(b Resource) bool
}

type ResourceConstructor func(typ, val string) Resource

type stringLengthRsc struct {
	t string
	v string
}

func NewStringLengthResource(typ, val string) Resource {
	return stringLengthRsc{
		t: typ,
		v: val,
	}
}

func (r stringLengthRsc) Type() string {
	return r.t
}

func (r stringLengthRsc) Value() string {
	return r.v
}

func (r stringLengthRsc) Contains(b Resource) bool {
	return len(r.Value()) < len(b.Value())
}

// Capability is the interface for an action users can perform
type Capability interface {
	String() string
	Contains(b Capability) bool
}

// NestedCapabilities is a basic implementation of the Capabilities interface
// based on a hierarchal list of strings
type NestedCapabilities struct {
	cap       string
	idx       int
	hierarchy *[]string
}

// assert at compile-time NestedCapabilities implements Capability
var _ Capability = (*NestedCapabilities)(nil)

// NewNestedCapabilities
func NewNestedCapabilities(strs ...string) NestedCapabilities {
	return NestedCapabilities{
		cap:       strs[0],
		idx:       0,
		hierarchy: &strs,
	}
}

func (nc NestedCapabilities) Cap(str string) Capability {
	idx := -1
	for i, c := range *nc.hierarchy {
		if c == str {
			idx = i
		}
	}

	return NestedCapabilities{
		cap:       str,
		idx:       idx,
		hierarchy: nc.hierarchy,
	}
}

func (nc NestedCapabilities) String() string {
	return nc.cap
}

func (nc NestedCapabilities) Contains(cap Capability) bool {
	str := cap.String()
	for i, c := range *nc.hierarchy {
		if c == str {
			if i > nc.idx {
				return false
			}
			return true
		}
	}
	return false
}
