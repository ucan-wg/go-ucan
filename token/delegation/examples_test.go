package delegation_test

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/MetaMask/go-did-it/didtest"
	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/codec/dagjson"

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
	// issDid: did:key:z6Mkf4WtCwPDtamsZvBJA4eSVcE7vZuRPy5Skm4HaoQv81i1
	//
	// CID (base58BTC): zdpuB1bzMdGAtzBej9ZBJbW3ppsjP2Cf9KZ8xMjdhUjaf3vz1
	//
	// DAG-CBOR (base64) out: glhAQhZTTb4U1rin/oLFre618Ol5/leaP758T6EkHYfQaNmFYad7yVTer8bbM1Zp2LxrV3eEYeWkHeL3F0V3zWC/CaJhaEg0Ae0B7QETcXN1Y2FuL2RsZ0AxLjAuMC1yYy4xqWNhdWR4OGRpZDprZXk6ejZNa2pXRlU4SHRmTXF1V241TUVROXIyMnpXcXlDTjF2YXpHdzZnNnB6Y2l6aU5WY2NtZGgvZm9vL2JhcmNleHAaaItoW2Npc3N4OGRpZDprZXk6ejZNa2Y0V3RDd1BEdGFtc1p2QkpBNGVTVmNFN3ZadVJQeTVTa200SGFvUXY4MWkxY25iZhpoi1qHY3BvbIODYj09Zy5zdGF0dXNlZHJhZnSDY2FsbGkucmV2aWV3ZXKDZGxpa2VmLmVtYWlsbSpAZXhhbXBsZS5jb22DY2FueWUudGFnc4Jib3KCg2I9PWEuZG5ld3ODYj09YS5lcHJlc3Njc3VieDhkaWQ6a2V5Ono2TWtuVXoxbVNqNHB2UzZhVVVIZWtDSGRVUHY3SEJoRHlEQlpRMlczVnVqYzVxQ2RtZXRhomNiYXoYe2Nmb29jYmFyZW5vbmNlTHa1D4DJ78LvscA/hg==
	// Converted to DAG-JSON out:
	//	[
	//		{
	//			"/": {
	//				"bytes": "QhZTTb4U1rin/oLFre618Ol5/leaP758T6EkHYfQaNmFYad7yVTer8bbM1Zp2LxrV3eEYeWkHeL3F0V3zWC/CQ"
	//			}
	//		},
	//		{
	//			"h": {
	//				"/": {
	//					"bytes": "NAHtAe0BE3E"
	//				}
	//			},
	//			"ucan/dlg@1.0.0-rc.1": {
	//				"aud": "did:key:z6MkjWFU8HtfMquWn5MEQ9r22zWqyCN1vazGw6g6pzciziNV",
	//				"cmd": "/foo/bar",
	//				"exp": 1753966683,
	//				"iss": "did:key:z6Mkf4WtCwPDtamsZvBJA4eSVcE7vZuRPy5Skm4HaoQv81i1",
	//				"meta": {
	//					"baz": 123,
	//					"foo": "bar"
	//				},
	//				"nbf": 1753963143,
	//				"nonce": {
	//					"/": {
	//						"bytes": "drUPgMnvwu+xwD+G"
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
	//				"sub": "did:key:z6MknUz1mSj4pvS6aUUHekCHdUPv7HBhDyDBZQ2W3Vujc5qC"
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
	// CID (base58BTC): zdpuAkwYz8nY7uU8j3F6wVTfFY1VEoExwvUAYBEwRWfTozddE
	//
	// DAG-CBOR (base64) out: glhATIXb2wnBByq/llaQ4RLoWTxheAwamCNo2sKL8SQJbq0EVRvdfUQDKNpuMVkvtyUR6tUdZlKv1BcXjfGEaF2XAKJhaEg0Ae0B7QETcXN1Y2FuL2RsZ0AxLjAuMC1yYy4xqWNhdWR4OGRpZDprZXk6ejZNa2Y0V3RDd1BEdGFtc1p2QkpBNGVTVmNFN3ZadVJQeTVTa200SGFvUXY4MWkxY2NtZGgvZm9vL2JhcmNleHAaaItnfWNpc3N4OGRpZDprZXk6ejZNa25VejFtU2o0cHZTNmFVVUhla0NIZFVQdjdIQmhEeURCWlEyVzNWdWpjNXFDY25iZhpoi1mpY3BvbIODYj09Zy5zdGF0dXNlZHJhZnSDY2FsbGkucmV2aWV3ZXKDZGxpa2VmLmVtYWlsbSpAZXhhbXBsZS5jb22DY2FueWUudGFnc4Jib3KCg2I9PWEuZG5ld3ODYj09YS5lcHJlc3Njc3VieDhkaWQ6a2V5Ono2TWtuVXoxbVNqNHB2UzZhVVVIZWtDSGRVUHY3SEJoRHlEQlpRMlczVnVqYzVxQ2RtZXRhomNiYXoYe2Nmb29jYmFyZW5vbmNlTGxwbjHKcevt5dZ0Xg==
	//
	// Converted to DAG-JSON out:
	//	[
	//		{
	//			"/": {
	//				"bytes": "TIXb2wnBByq/llaQ4RLoWTxheAwamCNo2sKL8SQJbq0EVRvdfUQDKNpuMVkvtyUR6tUdZlKv1BcXjfGEaF2XAA"
	//			}
	//		},
	//		{
	//			"h": {
	//				"/": {
	//					"bytes": "NAHtAe0BE3E"
	//				}
	//			},
	//			"ucan/dlg@1.0.0-rc.1": {
	//				"aud": "did:key:z6Mkf4WtCwPDtamsZvBJA4eSVcE7vZuRPy5Skm4HaoQv81i1",
	//				"cmd": "/foo/bar",
	//				"exp": 1753966461,
	//				"iss": "did:key:z6MknUz1mSj4pvS6aUUHekCHdUPv7HBhDyDBZQ2W3Vujc5qC",
	//				"meta": {
	//					"baz": 123,
	//					"foo": "bar"
	//				},
	//				"nbf": 1753962921,
	//				"nonce": {
	//					"/": {
	//						"bytes": "bHBuMcpx6+3l1nRe"
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
	//				"sub": "did:key:z6MknUz1mSj4pvS6aUUHekCHdUPv7HBhDyDBZQ2W3Vujc5qC"
	//			}
	//		}
	//	]
}

// The following example demonstrates how to get a delegation.Token from
// a DAG-CBOR []byte.
func ExampleFromSealed() {
	const cborBase64 = "glhACBuW/rjVKyBPUVPxexsafwBe7y84k0yzywq3hQW2rs2TNmWA5wexAQ+jTkSQ07zhmQRA/wytBfqWkx24+sjlD6JhaEg0Ae0B7QETcXN1Y2FuL2RsZ0AxLjAuMC1yYy4xp2NhdWR4OGRpZDprZXk6ejZNa2pXRlU4SHRmTXF1V241TUVROXIyMnpXcXlDTjF2YXpHdzZnNnB6Y2l6aU5WY2NtZGgvZm9vL2JhcmNleHD2Y2lzc3g4ZGlkOmtleTp6Nk1rZjRXdEN3UER0YW1zWnZCSkE0ZVNWY0U3dlp1UlB5NVNrbTRIYW9RdjgxaTFjcG9sg4NiPT1nLnN0YXR1c2VkcmFmdINjYWxsaS5yZXZpZXdlcoNkbGlrZWYuZW1haWxtKkBleGFtcGxlLmNvbYNjYW55ZS50YWdzgmJvcoKDYj09YS5kbmV3c4NiPT1hLmVwcmVzc2NzdWJ4OGRpZDprZXk6ejZNa25VejFtU2o0cHZTNmFVVUhla0NIZFVQdjdIQmhEeURCWlEyVzNWdWpjNXFDZW5vbmNlTDVJDkO3LVTKnhxKKw=="

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
	// CID (base58BTC): zdpuAsgmtmC849BEApGCJm2fTSzwtHiAqQnuJCMBBBCbFiadd
	// Issuer (iss): did:key:z6Mkf4WtCwPDtamsZvBJA4eSVcE7vZuRPy5Skm4HaoQv81i1
	// Audience (aud): did:key:z6MkjWFU8HtfMquWn5MEQ9r22zWqyCN1vazGw6g6pzciziNV
	// Subject (sub): did:key:z6MknUz1mSj4pvS6aUUHekCHdUPv7HBhDyDBZQ2W3Vujc5qC
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
	// Nonce (nonce): 35490e43b72d54ca9e1c4a2b
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
