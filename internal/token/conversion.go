package token

import (
	"github.com/ipld/go-ipld-prime/datamodel"
)

func ToIPLDMapStringAny(m map[string]datamodel.Node) Map__String__Any {
	keys := make([]string, len(m))
	i := 0

	for k := range m {
		keys[i] = k
		i++
	}

	return Map__String__Any{
		Keys:   keys,
		Values: m,
	}
}

func FromIPLDMapStringAny(m Map__String__Any) map[string]datamodel.Node {
	return m.Values
}
