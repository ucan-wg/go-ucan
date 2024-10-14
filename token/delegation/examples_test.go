package delegation_test

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"github.com/libp2p/go-libp2p/core/crypto"

	"github.com/ucan-wg/go-ucan/did"
	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/pkg/policy"
	"github.com/ucan-wg/go-ucan/pkg/policy/literal"
	"github.com/ucan-wg/go-ucan/pkg/policy/selector"
	"github.com/ucan-wg/go-ucan/token/delegation"
	"github.com/ucan-wg/go-ucan/token/internal/envelope"
)

// The following example shows how to create a delegation.Token with
// distinct DIDs for issuer (iss), audience (aud) and subject (sub).
func ExampleNew() {
	issPriv, issPub, err := crypto.GenerateEd25519Key(rand.Reader)
	printThenPanicOnErr(err)

	issDid, err := did.FromPubKey(issPub)
	printThenPanicOnErr(err)
	fmt.Println("issDid:", issDid)

	audDid := did.MustParse(AudienceDID)
	subDid := did.MustParse(subjectDID)

	// The command defines the shape of the arguments that will be evaluated against the policy
	cmd := command.MustParse("/foo/bar")

	// The policy defines what is allowed to do.
	pol := policy.Policy{
		policy.Equal(selector.MustParse(".status"), literal.String("draft")),
		policy.All(selector.MustParse(".reviewer"),
			policy.MustLike(selector.MustParse(".email"), "*@example.com"),
		),
		policy.Any(selector.MustParse(".tags"), policy.Or(
			policy.Equal(selector.Identity, literal.String("news")),
			policy.Equal(selector.Identity, literal.String("press")),
		)),
	}

	tkn, err := delegation.New(issPriv, audDid, cmd, pol,
		delegation.WithSubject(subDid),
		delegation.WithExpirationAfter(time.Hour),
		delegation.WithMeta("foo", "bar"),
		delegation.WithMeta("baz", 123),
	)
	printThenPanicOnErr(err)

	// "Seal", meaning encode and wrap into a signed envelope.
	data, id, err := tkn.ToSealed(issPriv)
	printThenPanicOnErr(err)

	printCIDAndSealed(id, data)

	// Example output:
	//
	// issDid: did:key:z6MksKbqUiXRKVDHQJ2yezG83M6d68AQbz9rtajULF575X3s
	//
	// CID (base58BTC): zdpuAtXJQXZt123WNczSueoBrVcyKoJ2LH1aTmf41dZrisJJA
	//
	// DAG-CBOR (base64) out: glhAGCWszoibTPgkBSe5pk03wsB2orGzRKFvxLeqoDTNixxzXTDGKTj4ZfZrGOyCxf6rNW5zP8x2esFKV/akgy/nAaJhaEQ07QFxc3VjYW4vZGxnQDEuMC4wLXJjLjGoY2F1ZHg4ZGlkOmtleTp6Nk1rcTVZbWJKY1RyUEV4TkRpMjZpbXJUQ3BLaGVwakJGQlNIcXJCRE4yQXJQa3ZjY21kaC9mb28vYmFyY2V4cBpnDPLWY2lzc3g4ZGlkOmtleTp6Nk1rc0ticVVpWFJLVkRIUUoyeWV6RzgzTTZkNjhBUWJ6OXJ0YWpVTEY1NzVYM3NjcG9sg4NiPT1nLnN0YXR1c2VkcmFmdINjYWxsaS5yZXZpZXdlcoNkbGlrZWYuZW1haWxtKkBleGFtcGxlLmNvbYNjYW55ZS50YWdzgmJvcoKDYj09YS5kbmV3c4NiPT1hLmVwcmVzc2NzdWJ4OGRpZDprZXk6ejZNa3RBMXVCZENwcTR1SkJxRTlqak1pTHl4WkJnOWE2eGdQUEtKak1xc3M2WmMyZG1ldGGiY2Jhehh7Y2Zvb2NiYXJlbm9uY2VMgb9wlP/cdMKutRg+
	//
	// Converted to DAG-JSON out:
	// [
	//	{
	//		"/": {
	//			"bytes": "GCWszoibTPgkBSe5pk03wsB2orGzRKFvxLeqoDTNixxzXTDGKTj4ZfZrGOyCxf6rNW5zP8x2esFKV/akgy/nAQ"
	//		}
	//	},
	//	{
	//		"h": {
	//			"/": {
	//				"bytes": "NO0BcQ"
	//			}
	//		},
	//		"ucan/dlg@1.0.0-rc.1": {
	//			"aud": "did:key:z6Mkq5YmbJcTrPExNDi26imrTCpKhepjBFBSHqrBDN2ArPkv",
	//			"cmd": "/foo/bar",
	//			"exp": 1728901846,
	//			"iss": "did:key:z6MksKbqUiXRKVDHQJ2yezG83M6d68AQbz9rtajULF575X3s",
	//			"meta": {
	//				"baz": 123,
	//				"foo": "bar"
	//			},
	//			"nonce": {
	//				"/": {
	//					"bytes": "gb9wlP/cdMKutRg+"
	//				}
	//			},
	//			"pol": [
	//				[
	//					"==",
	//					".status",
	//					"draft"
	//				],
	//				[
	//					"all",
	//					".reviewer",
	//					[
	//						"like",
	//						".email",
	//						"*@example.com"
	//					]
	//				],
	//				[
	//					"any",
	//					".tags",
	//					[
	//						"or",
	//						[
	//							[
	//								"==",
	//								".",
	//								"news"
	//							],
	//							[
	//								"==",
	//								".",
	//								"press"
	//							]
	//						]
	//					]
	//				]
	//			],
	//			"sub": "did:key:z6MktA1uBdCpq4uJBqE9jjMiLyxZBg9a6xgPPKJjMqss6Zc2"
	//		}
	//	}
	// ]
}

// The following example shows how to create a UCAN root delegation.Token
// - a delegation.Token with the subject (sub) set to the value of issuer
// (iss).
func ExampleRoot() {
	issPriv, issPub, err := crypto.GenerateEd25519Key(rand.Reader)
	printThenPanicOnErr(err)

	issDid, err := did.FromPubKey(issPub)
	printThenPanicOnErr(err)
	fmt.Println("issDid:", issDid)

	audDid := did.MustParse(AudienceDID)

	// The command defines the shape of the arguments that will be evaluated against the policy
	cmd := command.MustParse("/foo/bar")

	// The policy defines what is allowed to do.
	pol := policy.Policy{
		policy.Equal(selector.MustParse(".status"), literal.String("draft")),
		policy.All(selector.MustParse(".reviewer"),
			policy.MustLike(selector.MustParse(".email"), "*@example.com"),
		),
		policy.Any(selector.MustParse(".tags"), policy.Or(
			policy.Equal(selector.Identity, literal.String("news")),
			policy.Equal(selector.Identity, literal.String("press")),
		)),
	}

	tkn, err := delegation.Root(issPriv, audDid, cmd, pol,
		delegation.WithExpirationAfter(time.Hour),
		delegation.WithMeta("foo", "bar"),
		delegation.WithMeta("baz", 123),
	)
	printThenPanicOnErr(err)

	// "Seal", meaning encode and wrap into a signed envelope.
	data, id, err := tkn.ToSealed(issPriv)
	printThenPanicOnErr(err)

	printCIDAndSealed(id, data)

	// Example output:
	//
	// issDid: did:key:z6MkshW2ADRrmfBuuBpKJiyNd7acLK1yjnxFJuBimwjQ4Bo5
	//
	// CID (base58BTC): zdpuAoBzE3kJK1qZC9EXH7h6iCwym1TqfxT9XzUfFNfcjcAKh
	//
	// DAG-CBOR (base64) out: glhADpyBSrTdRn2oZdJU26CjgFbaH7LbTDWyyAdgIwAW0p151XSdJwoBS2vTCp0+7sEkf4X2wl6N5IhxiKyQ8OkbCaJhaEQ07QFxc3VjYW4vZGxnQDEuMC4wLXJjLjGoY2F1ZHg4ZGlkOmtleTp6Nk1rcTVZbWJKY1RyUEV4TkRpMjZpbXJUQ3BLaGVwakJGQlNIcXJCRE4yQXJQa3ZjY21kaC9mb28vYmFyY2V4cBpnDPPCY2lzc3g4ZGlkOmtleTp6Nk1rc2hXMkFEUnJtZkJ1dUJwS0ppeU5kN2FjTEsxeWpueEZKdUJpbXdqUTRCbzVjcG9sg4NiPT1nLnN0YXR1c2VkcmFmdINjYWxsaS5yZXZpZXdlcoNkbGlrZWYuZW1haWxtKkBleGFtcGxlLmNvbYNjYW55ZS50YWdzgmJvcoKDYj09YS5kbmV3c4NiPT1hLmVwcmVzc2NzdWJ4OGRpZDprZXk6ejZNa3NoVzJBRFJybWZCdXVCcEtKaXlOZDdhY0xLMXlqbnhGSnVCaW13alE0Qm81ZG1ldGGiY2Jhehh7Y2Zvb2NiYXJlbm9uY2VMDmBGXMa/TCvhLHqu
	//
	// Converted to DAG-JSON out:
	// [
	//	{
	//		"/": {
	//			"bytes": "DpyBSrTdRn2oZdJU26CjgFbaH7LbTDWyyAdgIwAW0p151XSdJwoBS2vTCp0+7sEkf4X2wl6N5IhxiKyQ8OkbCQ"
	//		}
	//	},
	//	{
	//		"h": {
	//			"/": {
	//				"bytes": "NO0BcQ"
	//			}
	//		},
	//		"ucan/dlg@1.0.0-rc.1": {
	//			"aud": "did:key:z6Mkq5YmbJcTrPExNDi26imrTCpKhepjBFBSHqrBDN2ArPkv",
	//			"cmd": "/foo/bar",
	//			"exp": 1728902082,
	//			"iss": "did:key:z6MkshW2ADRrmfBuuBpKJiyNd7acLK1yjnxFJuBimwjQ4Bo5",
	//			"meta": {
	//				"baz": 123,
	//				"foo": "bar"
	//			},
	//			"nonce": {
	//				"/": {
	//					"bytes": "DmBGXMa/TCvhLHqu"
	//				}
	//			},
	//			"pol": [
	//				[
	//					"==",
	//					".status",
	//					"draft"
	//				],
	//				[
	//					"all",
	//					".reviewer",
	//					[
	//						"like",
	//						".email",
	//						"*@example.com"
	//					]
	//				],
	//				[
	//					"any",
	//					".tags",
	//					[
	//						"or",
	//						[
	//							[
	//								"==",
	//								".",
	//								"news"
	//							],
	//							[
	//								"==",
	//								".",
	//								"press"
	//							]
	//						]
	//					]
	//				]
	//			],
	//			"sub": "did:key:z6MkshW2ADRrmfBuuBpKJiyNd7acLK1yjnxFJuBimwjQ4Bo5"
	//		}
	//	}
	// ]
}

// The following example demonstrates how to get a delegation.Token from
// a DAG-CBOR []byte.
func ExampleToken_FromSealed() {
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
	//     ["like", ".email", "*@example.com"]],
	//   ["any", ".tags",
	//     ["or", [
	//       ["==", ".", "news"],
	//       ["==", ".", "press"]]]
	//     ]
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
