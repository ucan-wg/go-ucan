package ucan_test

import (
	"context"
	"fmt"
	"time"

	"github.com/qri-io/ucan"
)

func ExampleWalkthrough() {
	source, err := ucan.NewPrivKeyUCANSource(keyOne)
	panicIfError(err)

	subjectDID, err := ucan.DIDStringFromPublicKey(keyOne.GetPublic())
	panicIfError(err)

	caps := ucan.NewNestedCapabilities("SUPER_USER", "OVERWRITE", "SOFT_DELETE", "REVISE", "CREATE")
	att := ucan.Attenuations{
		{caps.Cap("SUPER_USER"), ucan.NewStringLengthResource("api", "*")},
		{caps.Cap("SUPER_USER"), ucan.NewStringLengthResource("dataset", "b5:world_bank_population:*")},
	}
	zero := time.Time{}

	// create a root UCAN
	origin, err := source.NewOriginUCAN(subjectDID, att, nil, zero, zero)
	panicIfError(err)

	id, err := origin.CID()
	panicIfError(err)

	fmt.Printf("cid of root UCAN: %s\n", id.String())

	att = ucan.Attenuations{
		{caps.Cap("SUPER_USER"), ucan.NewStringLengthResource("dataset", "third:resource")},
	}

	if _, err = source.NewAttenuatedUCAN(origin, subjectDID, att, nil, zero, zero); err != nil {
		fmt.Println(err)
	}

	att = ucan.Attenuations{
		{caps.Cap("OVERWRITE"), ucan.NewStringLengthResource("dataset", "b5:world_bank_population:*")},
	}

	derivedToken, err := source.NewAttenuatedUCAN(origin, subjectDID, att, nil, zero, zero)
	panicIfError(err)

	id, err = derivedToken.CID()
	panicIfError(err)

	fmt.Printf("cid of derived UCAN: %s\n", id.String())

	p := exampleParser()
	_, err = p.ParseAndVerify(context.Background(), origin.Raw)
	panicIfError(err)

	// Output:
	// cid of root UCAN: bafkreidhsvhlctwylgeibl2eeapdvbl3qm3mbqcqhxhvy4grmr25ji77hu
	// scope of ucan attenuations must be less than it's parent
	// cid of derived UCAN: bafkreifglbwtr27fbzmv3uardlygvggr722fckusfvfyfsonwkroca7efu
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func exampleParser() *ucan.UCANParser {
	caps := ucan.NewNestedCapabilities("SUPER_USER", "OVERWRITE", "SOFT_DELETE", "REVISE", "CREATE")

	ac := func(m map[string]interface{}) (ucan.Attenuation, error) {
		var (
			cap string
			rsc ucan.Resource
		)
		for key, vali := range m {
			val, ok := vali.(string)
			if !ok {
				return ucan.Attenuation{}, fmt.Errorf(`expected attenuation value to be a string`)
			}

			if key == ucan.CapKey {
				cap = val
			} else {
				rsc = ucan.NewStringLengthResource(key, val)
			}
		}

		return ucan.Attenuation{
			Rsc: rsc,
			Cap: caps.Cap(cap),
		}, nil
	}

	store := ucan.NewMemTokenStore()
	return ucan.NewUCANParser(ac, ucan.StringDIDPubKeyResolver{}, store.(ucan.CIDBytesResolver))
}
