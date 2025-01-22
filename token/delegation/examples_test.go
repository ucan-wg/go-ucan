package delegation_test

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/codec/dagjson"

	"github.com/ucan-wg/go-ucan/did/didtest"
	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/pkg/policy"
	"github.com/ucan-wg/go-ucan/pkg/policy/literal"
	"github.com/ucan-wg/go-ucan/token/delegation"
	"github.com/ucan-wg/go-ucan/token/internal/envelope"
)

// The following example shows how to create a delegation.Token with
// distinct DIDs for issuer (iss), audience (aud) and subject (sub).
func ExampleNew() {
	fmt.Println("issDid:", didtest.PersonaBob.DID().String())

	// The command defines the shape of the arguments that will be evaluated against the policy
	cmd := command.MustParse("/foo/bar")

	// The policy defines what is allowed to do.
	pol := policy.MustConstruct(
		policy.Equal(".status", literal.String("draft")),
		policy.All(".reviewer",
			policy.Like(".email", "*@example.com"),
		),
		policy.Any(".tags", policy.Or(
			policy.Equal(".", literal.String("news")),
			policy.Equal(".", literal.String("press")),
		)),
	)

	tkn, err := delegation.New(didtest.PersonaBob.DID(), didtest.PersonaCarol.DID(), cmd, pol, didtest.PersonaAlice.DID(),
		delegation.WithExpirationIn(time.Hour),
		delegation.WithNotBeforeIn(time.Minute),
		delegation.WithMeta("foo", "bar"),
		delegation.WithMeta("baz", 123),
	)
	printThenPanicOnErr(err)

	// "Seal", meaning encode and wrap into a signed envelope.
	data, id, err := tkn.ToSealed(didtest.PersonaBob.PrivKey())
	printThenPanicOnErr(err)

	printCIDAndSealed(id, data)

	// Example output:
	//
	// issDid: did:key:z6MkvJPmEZZYbgiw1ouT1oouTsTFBHJSts9ophVsNgcRmYxU
	//
	// CID (base58BTC): zdpuAsqfZkgg2jgZyob23sq1J9xwtf9PHgt1PsskVCMq7Vvxk
	//
	// DAG-CBOR (base64) out: lhAOnjc0bPptlI5MxRBrIK3YmAP1CxKfXOPkz6MHt/UJCx2gCN+6gXZX2N+BIJvmy8XmAO5sT2GYimiV7HlJH1AA6JhaEQ07QFxc3VjYW4vZGxnQDEuMC4wLXJjLjGpY2F1ZHg4ZGlkOmtleTp6Nk1rZ3VwY2hoNUh3dUhhaFM3WXN5RThiTHVhMU1yOHAyaUtOUmh5dlN2UkFzOW5jY21kaC9mb28vYmFyY2V4cBpnROP/Y2lzc3g4ZGlkOmtleTp6Nk1rdkpQbUVaWlliZ2l3MW91VDFvb3VUc1RGQkhKU3RzOW9waFZzTmdjUm1ZeFVjbmJmGmdE1itjcG9sg4NiPT1nLnN0YXR1c2VkcmFmdINjYWxsaS5yZXZpZXdlcoNkbGlrZWYuZW1haWxtKkBleGFtcGxlLmNvbYNjYW55ZS50YWdzgmJvcoKDYj09YS5kbmV3c4NiPT1hLmVwcmVzc2NzdWJ4OGRpZDprZXk6ejZNa3V1a2syc2tEWExRbjdOSzNFaDlqTW5kWWZ2REJ4eGt0Z3BpZEpBcWI3TTNwZG1ldGGiY2Jhehh7Y2Zvb2NiYXJlbm9uY2VMv+Diy6GExIuM1eX4
	// Converted to DAG-JSON out:
	//	[
	//		{
	//			"/": {
	//				"bytes": "5rvl8uKmDVGvAVSt4m/0MGiXl9dZwljJJ9m2qHCoIB617l26UvMxyH5uvN9hM7ozfVATiq4mLhoGgm9IGnEEAg"
	//			}
	//		},
	//		{
	//			"h": {
	//				"/": {
	//					"bytes": "NO0BcQ"
	//				}
	//			},
	//			"ucan/dlg@1.0.0-rc.1": {
	//				"aud": "did:key:z6Mkq5YmbJcTrPExNDi26imrTCpKhepjBFBSHqrBDN2ArPkv",
	//				"cmd": "/foo/bar",
	//				"exp": 1728933098,
	//				"iss": "did:key:z6MkhVFznPeR572rTK51UjoTNpnF8cxuWfPm9oBMPr7y8ABe",
	//				"meta": {
	//					"baz": 123,
	//					"foo": "bar"
	//				},
	//				"nbf": 1728929558,
	//				"nonce": {
	//					"/": {
	//						"bytes": "u0HMgJ5Y+M84I/66"
	//					}
	//				},
	//				"pol": [
	//					[
	//						"==",
	//						".status",
	//						"draft"
	//					],
	//					[
	//						"all",
	//						".reviewer",
	//						[
	//							"like",
	//							".email",
	//							"*@example.com"
	//						]
	//					],
	//					[
	//						"any",
	//						".tags",
	//						[
	//							"or",
	//							[
	//								[
	//									"==",
	//									".",
	//									"news"
	//								],
	//								[
	//									"==",
	//									".",
	//									"press"
	//								]
	//							]
	//						]
	//					]
	//				],
	//				"sub": "did:key:z6MktA1uBdCpq4uJBqE9jjMiLyxZBg9a6xgPPKJjMqss6Zc2"
	//			}
	//		}
	//	]
}

// The following example shows how to create a UCAN root delegation.Token
// - a delegation.Token with the subject (sub) set to the value of issuer
// (iss).
func ExampleRoot() {
	// The command defines the shape of the arguments that will be evaluated against the policy
	cmd := command.MustParse("/foo/bar")

	// The policy defines what is allowed to do.
	pol := policy.MustConstruct(
		policy.Equal(".status", literal.String("draft")),
		policy.All(".reviewer",
			policy.Like(".email", "*@example.com"),
		),
		policy.Any(".tags", policy.Or(
			policy.Equal(".", literal.String("news")),
			policy.Equal(".", literal.String("press")),
		)),
	)

	tkn, err := delegation.Root(didtest.PersonaAlice.DID(), didtest.PersonaBob.DID(), cmd, pol,
		delegation.WithExpirationIn(time.Hour),
		delegation.WithNotBeforeIn(time.Minute),
		delegation.WithMeta("foo", "bar"),
		delegation.WithMeta("baz", 123),
	)
	printThenPanicOnErr(err)

	// "Seal", meaning encode and wrap into a signed envelope.
	data, id, err := tkn.ToSealed(didtest.PersonaAlice.PrivKey())
	printThenPanicOnErr(err)

	printCIDAndSealed(id, data)

	// Example output:
	//
	// issDid: did:key:z6MknWJqz17Y4AfsXSJUFKomuBR4GTkViM7kJYutzTMkCyFF
	//
	// CID (base58BTC): zdpuAkwYz8nY7uU8j3F6wVTfFY1VEoExwvUAYBEwRWfTozddE
	//
	// DAG-CBOR (base64) out: glhAVpW67FJ+myNi+azvnw2jivuiqXTuMrDZI2Qdaa8jE1Oi3mkjnm7DyqSQGADcomcuDslMWKmJ+OIyvbPG5PtSA6JhaEQ07QFxc3VjYW4vZGxnQDEuMC4wLXJjLjGpY2F1ZHg4ZGlkOmtleTp6Nk1rdkpQbUVaWlliZ2l3MW91VDFvb3VUc1RGQkhKU3RzOW9waFZzTmdjUm1ZeFVjY21kaC9mb28vYmFyY2V4cBpnROVoY2lzc3g4ZGlkOmtleTp6Nk1rdXVrazJza0RYTFFuN05LM0VoOWpNbmRZZnZEQnh4a3RncGlkSkFxYjdNM3BjbmJmGmdE15RjcG9sg4NiPT1nLnN0YXR1c2VkcmFmdINjYWxsaS5yZXZpZXdlcoNkbGlrZWYuZW1haWxtKkBleGFtcGxlLmNvbYNjYW55ZS50YWdzgmJvcoKDYj09YS5kbmV3c4NiPT1hLmVwcmVzc2NzdWJ4OGRpZDprZXk6ejZNa3V1a2syc2tEWExRbjdOSzNFaDlqTW5kWWZ2REJ4eGt0Z3BpZEpBcWI3TTNwZG1ldGGiY2Jhehh7Y2Zvb2NiYXJlbm9uY2VMwzDc03WBciJIGPWG
	//
	// Converted to DAG-JSON out:
	//	[
	//		{
	//			"/": {
	//				"bytes": "VpW67FJ+myNi+azvnw2jivuiqXTuMrDZI2Qdaa8jE1Oi3mkjnm7DyqSQGADcomcuDslMWKmJ+OIyvbPG5PtSAw"
	//			}
	//		},
	//		{
	//			"h": {
	//				"/": {
	//					"bytes": "NO0BcQ"
	//				}
	//			},
	//			"ucan/dlg@1.0.0-rc.1": {
	//				"aud": "did:key:z6MkvJPmEZZYbgiw1ouT1oouTsTFBHJSts9ophVsNgcRmYxU",
	//				"cmd": "/foo/bar",
	//				"exp": 1732568424,
	//				"iss": "did:key:z6Mkuukk2skDXLQn7NK3Eh9jMndYfvDBxxktgpidJAqb7M3p",
	//				"meta": {
	//					"baz": 123,
	//					"foo": "bar"
	//				},
	//				"nbf": 1732564884,
	//				"nonce": {
	//					"/": {
	//						"bytes": "wzDc03WBciJIGPWG"
	//					}
	//				},
	//				"pol": [
	//					[
	//						"==",
	//						".status",
	//						"draft"
	//					],
	//					[
	//						"all",
	//						".reviewer",
	//						[
	//							"like",
	//							".email",
	//							"*@example.com"
	//						]
	//					],
	//					[
	//						"any",
	//						".tags",
	//						[
	//							"or",
	//							[
	//								[
	//									"==",
	//									".",
	//									"news"
	//								],
	//								[
	//									"==",
	//									".",
	//									"press"
	//								]
	//							]
	//						]
	//					]
	//				],
	//				"sub": "did:key:z6Mkuukk2skDXLQn7NK3Eh9jMndYfvDBxxktgpidJAqb7M3p"
	//			}
	//		}
	//	]
}

// The following example demonstrates how to get a delegation.Token from
// a DAG-CBOR []byte.
func ExampleFromSealed() {
	const cborBase64 = "glhAmnAkgfjAx4SA5pzJmtaHRJtTGNpF1y6oqb4yhGoM2H2EUGbBYT4rVDjMKBgCjhdGHjipm00L8iR5SsQh3sIEBaJhaEQ07QFxc3VjYW4vZGxnQDEuMC4wLXJjLjGoY2F1ZHg4ZGlkOmtleTp6Nk1rcTVZbWJKY1RyUEV4TkRpMjZpbXJUQ3BLaGVwakJGQlNIcXJCRE4yQXJQa3ZjY21kaC9mb28vYmFyY2V4cPZjaXNzeDhkaWQ6a2V5Ono2TWtwem4ybjNaR1QyVmFxTUdTUUMzdHptelY0VFM5UzcxaUZzRFhFMVdub05IMmNwb2yDg2I9PWcuc3RhdHVzZWRyYWZ0g2NhbGxpLnJldmlld2Vyg2RsaWtlZi5lbWFpbG0qQGV4YW1wbGUuY29tg2NhbnllLnRhZ3OCYm9ygoNiPT1hLmRuZXdzg2I9PWEuZXByZXNzY3N1Yng4ZGlkOmtleTp6Nk1rdEExdUJkQ3BxNHVKQnFFOWpqTWlMeXhaQmc5YTZ4Z1BQS0pqTXFzczZaYzJkbWV0YaBlbm9uY2VMAAECAwQFBgcICQoL"

	cborBytes, err := base64.StdEncoding.DecodeString(cborBase64)
	printThenPanicOnErr(err)

	tkn, c, err := delegation.FromSealed(cborBytes)
	printThenPanicOnErr(err)

	fmt.Println("CID (base58BTC):", envelope.CIDToBase58BTC(c))
	fmt.Println("Issuer (iss):", tkn.Issuer().String())
	fmt.Println("Audience (aud):", tkn.Audience().String())
	fmt.Println("Subject (sub):", tkn.Subject().String())
	fmt.Println("Command (cmd):", tkn.Command().String())
	fmt.Println("Policy (pol):", tkn.Policy().String())
	fmt.Println("Nonce (nonce):", hex.EncodeToString(tkn.Nonce()))
	fmt.Println("Meta (meta):", tkn.Meta().String())
	fmt.Println("NotBefore (nbf):", tkn.NotBefore())
	fmt.Println("Expiration (exp):", tkn.Expiration())

	// Output:
	// CID (base58BTC): zdpuAw26pFuvZa2Z9YAtpZZnWN6VmnRFr7Z8LVY5c7RVWoxGY
	// Issuer (iss): did:key:z6Mkpzn2n3ZGT2VaqMGSQC3tzmzV4TS9S71iFsDXE1WnoNH2
	// Audience (aud): did:key:z6Mkq5YmbJcTrPExNDi26imrTCpKhepjBFBSHqrBDN2ArPkv
	// Subject (sub): did:key:z6MktA1uBdCpq4uJBqE9jjMiLyxZBg9a6xgPPKJjMqss6Zc2
	// Command (cmd): /foo/bar
	// Policy (pol): [
	//   ["==", ".status", "draft"],
	//   ["all", ".reviewer",
	//     ["like", ".email", "*@example.com"]
	//   ],
	//   ["any", ".tags",
	//     ["or", [
	//       ["==", ".", "news"],
	//       ["==", ".", "press"]
	//     ]]
	//   ]
	// ]
	// Nonce (nonce): 000102030405060708090a0b
	// Meta (meta): {}
	// NotBefore (nbf): <nil>
	// Expiration (exp): <nil>
}

func printCIDAndSealed(id cid.Cid, data []byte) {
	fmt.Println("CID (base58BTC):", envelope.CIDToBase58BTC(id))
	fmt.Println("DAG-CBOR (base64) out:", base64.StdEncoding.EncodeToString(data))
	fmt.Println("Converted to DAG-JSON out:")

	node, err := ipld.Decode(data, dagcbor.Decode)
	printThenPanicOnErr(err)

	rawJSON, err := ipld.Encode(node, dagjson.Encode)
	printThenPanicOnErr(err)

	prettyJSON := &bytes.Buffer{}
	err = json.Indent(prettyJSON, rawJSON, "", "\t")
	printThenPanicOnErr(err)

	fmt.Println(prettyJSON.String())
}

func printThenPanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
