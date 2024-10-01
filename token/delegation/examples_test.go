package delegation_test

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"github.com/libp2p/go-libp2p/core/crypto"

	"github.com/ucan-wg/go-ucan/did"
	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/pkg/policy"
	"github.com/ucan-wg/go-ucan/token/delegation"
	"github.com/ucan-wg/go-ucan/token/internal/envelope"
)

// The following example shows how to create a delegation.Token with
// distinct DIDs for issuer (iss), audience (aud) and subject (sub).
func ExampleNew() {
	issuerPrivKey := examplePrivKey(issuerPrivKeyCfg)
	audienceDID := exampleDID(AudienceDID)
	command := exampleCommand(subJectCmd)
	policy := examplePolicy(subjectPol)
	subjectDID := exampleDID(subjectDID)

	// Don't do this in your code - a nonce should be a cryptographically
	// strong random slice of bytes to ensure the integrity of your private
	// key.  For this example, a fixed nonce is required to obtain the fixed
	// printed output (below).  If unsure of what value to supply for the
	// nonce, don't pass the WithNonce option and one will be generated
	// when the token is created.
	nonce := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b}

	tkn, err := delegation.New(
		issuerPrivKey,
		audienceDID,
		command,
		policy,
		delegation.WithSubject(subjectDID),
		delegation.WithNonce(nonce),
	)
	printThenPanicOnErr(err)
	data, id, err := tkn.ToSealed(issuerPrivKey)
	printThenPanicOnErr(err)

	printCIDAndSealed(id, data)

	// Output:
	// CID (base58BTC): zdpuAw26pFuvZa2Z9YAtpZZnWN6VmnRFr7Z8LVY5c7RVWoxGY
	// DAG-CBOR (base64) out: glhAmnAkgfjAx4SA5pzJmtaHRJtTGNpF1y6oqb4yhGoM2H2EUGbBYT4rVDjMKBgCjhdGHjipm00L8iR5SsQh3sIEBaJhaEQ07QFxc3VjYW4vZGxnQDEuMC4wLXJjLjGoY2F1ZHg4ZGlkOmtleTp6Nk1rcTVZbWJKY1RyUEV4TkRpMjZpbXJUQ3BLaGVwakJGQlNIcXJCRE4yQXJQa3ZjY21kaC9mb28vYmFyY2V4cPZjaXNzeDhkaWQ6a2V5Ono2TWtwem4ybjNaR1QyVmFxTUdTUUMzdHptelY0VFM5UzcxaUZzRFhFMVdub05IMmNwb2yDg2I9PWcuc3RhdHVzZWRyYWZ0g2NhbGxpLnJldmlld2Vyg2RsaWtlZi5lbWFpbG0qQGV4YW1wbGUuY29tg2NhbnllLnRhZ3OCYm9ygoNiPT1hLmRuZXdzg2I9PWEuZXByZXNzY3N1Yng4ZGlkOmtleTp6Nk1rdEExdUJkQ3BxNHVKQnFFOWpqTWlMeXhaQmc5YTZ4Z1BQS0pqTXFzczZaYzJkbWV0YaBlbm9uY2VMAAECAwQFBgcICQoL
	// Converted to DAG-JSON out:
	// [
	//	{
	//		"/": {
	//			"bytes": "mnAkgfjAx4SA5pzJmtaHRJtTGNpF1y6oqb4yhGoM2H2EUGbBYT4rVDjMKBgCjhdGHjipm00L8iR5SsQh3sIEBQ"
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
	//			"exp": null,
	//			"iss": "did:key:z6Mkpzn2n3ZGT2VaqMGSQC3tzmzV4TS9S71iFsDXE1WnoNH2",
	//			"meta": {},
	//			"nonce": {
	//				"/": {
	//					"bytes": "AAECAwQFBgcICQoL"
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
	issuerPrivKey := examplePrivKey(issuerPrivKeyCfg)
	audienceDID := exampleDID(AudienceDID)
	command := exampleCommand(subJectCmd)
	policy := examplePolicy(subjectPol)

	// Don't do this in your code - a nonce should be a cryptographically
	// strong random slice of bytes to ensure the integrity of your private
	// key.  For this example, a fixed nonce is required to obtain the fixed
	// printed output (below).  If unsure of what value to supply for the
	// nonce, don't pass the WithNonce option and one will be generated
	// when the token is created.
	nonce := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b}

	tkn, err := delegation.Root(
		issuerPrivKey,
		audienceDID,
		command,
		policy,
		delegation.WithNonce(nonce),
	)
	printThenPanicOnErr(err)
	data, id, err := tkn.ToSealed(issuerPrivKey)
	printThenPanicOnErr(err)

	printCIDAndSealed(id, data)

	// Output:
	// CID (base58BTC): zdpuAnbsR3e6DK8hBk5WA7KwbHYN6CKY4a3Bv1GNehvFYShQ8
	// DAG-CBOR (base64) out: glhA67ASBczF/wlIP0ESENn+4ZNQKukjcTNz+fo7K2tYa6OUm0rWICDJJkDWm7lJeQt+KvSA+Y4ctHTQbAr3Lr7mDqJhaEQ07QFxc3VjYW4vZGxnQDEuMC4wLXJjLjGoY2F1ZHg4ZGlkOmtleTp6Nk1rcTVZbWJKY1RyUEV4TkRpMjZpbXJUQ3BLaGVwakJGQlNIcXJCRE4yQXJQa3ZjY21kaC9mb28vYmFyY2V4cPZjaXNzeDhkaWQ6a2V5Ono2TWtwem4ybjNaR1QyVmFxTUdTUUMzdHptelY0VFM5UzcxaUZzRFhFMVdub05IMmNwb2yDg2I9PWcuc3RhdHVzZWRyYWZ0g2NhbGxpLnJldmlld2Vyg2RsaWtlZi5lbWFpbG0qQGV4YW1wbGUuY29tg2NhbnllLnRhZ3OCYm9ygoNiPT1hLmRuZXdzg2I9PWEuZXByZXNzY3N1Yng4ZGlkOmtleTp6Nk1rcHpuMm4zWkdUMlZhcU1HU1FDM3R6bXpWNFRTOVM3MWlGc0RYRTFXbm9OSDJkbWV0YaBlbm9uY2VMAAECAwQFBgcICQoL
	// Converted to DAG-JSON out:
	// [
	//	{
	//		"/": {
	//			"bytes": "67ASBczF/wlIP0ESENn+4ZNQKukjcTNz+fo7K2tYa6OUm0rWICDJJkDWm7lJeQt+KvSA+Y4ctHTQbAr3Lr7mDg"
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
	//			"exp": null,
	//			"iss": "did:key:z6Mkpzn2n3ZGT2VaqMGSQC3tzmzV4TS9S71iFsDXE1WnoNH2",
	//			"meta": {},
	//			"nonce": {
	//				"/": {
	//					"bytes": "AAECAwQFBgcICQoL"
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
	//			"sub": "did:key:z6Mkpzn2n3ZGT2VaqMGSQC3tzmzV4TS9S71iFsDXE1WnoNH2"
	//		}
	//	}
	// ]
}

// The following example demonstrates how to get a delegation.Token from
// a DAG-CBOR []byte.
func ExampleToken_FromSealed() {
	cborBytes := exampleCBORData()
	fmt.Println("DAG-CBOR (base64) in:", base64.StdEncoding.EncodeToString(cborBytes))

	tkn, err := delegation.FromSealed(cborBytes)
	printThenPanicOnErr(err)

	fmt.Println("CID (base58BTC):", envelope.CIDToBase58BTC(tkn.CID()))
	fmt.Println("Issuer (iss):", tkn.Issuer().String())
	fmt.Println("Audience (aud):", tkn.Audience().String())
	fmt.Println("Subject (sub):", tkn.Subject().String())
	fmt.Println("Command (cmd):", tkn.Command().String())
	fmt.Println("Policy (pol): TODO")
	fmt.Println("Nonce (nonce):", hex.EncodeToString(tkn.Nonce()))
	fmt.Println("Meta (meta): TODO")
	fmt.Println("NotBefore (nbf):", tkn.NotBefore())
	fmt.Println("Expiration (exp):", tkn.Expiration())

	// Output:
	// DAG-CBOR (base64) in: glhAmnAkgfjAx4SA5pzJmtaHRJtTGNpF1y6oqb4yhGoM2H2EUGbBYT4rVDjMKBgCjhdGHjipm00L8iR5SsQh3sIEBaJhaEQ07QFxc3VjYW4vZGxnQDEuMC4wLXJjLjGoY2F1ZHg4ZGlkOmtleTp6Nk1rcTVZbWJKY1RyUEV4TkRpMjZpbXJUQ3BLaGVwakJGQlNIcXJCRE4yQXJQa3ZjY21kaC9mb28vYmFyY2V4cPZjaXNzeDhkaWQ6a2V5Ono2TWtwem4ybjNaR1QyVmFxTUdTUUMzdHptelY0VFM5UzcxaUZzRFhFMVdub05IMmNwb2yDg2I9PWcuc3RhdHVzZWRyYWZ0g2NhbGxpLnJldmlld2Vyg2RsaWtlZi5lbWFpbG0qQGV4YW1wbGUuY29tg2NhbnllLnRhZ3OCYm9ygoNiPT1hLmRuZXdzg2I9PWEuZXByZXNzY3N1Yng4ZGlkOmtleTp6Nk1rdEExdUJkQ3BxNHVKQnFFOWpqTWlMeXhaQmc5YTZ4Z1BQS0pqTXFzczZaYzJkbWV0YaBlbm9uY2VMAAECAwQFBgcICQoL
	// CID (base58BTC): zdpuAw26pFuvZa2Z9YAtpZZnWN6VmnRFr7Z8LVY5c7RVWoxGY
	// Issuer (iss): did:key:z6Mkpzn2n3ZGT2VaqMGSQC3tzmzV4TS9S71iFsDXE1WnoNH2
	// Audience (aud): did:key:z6Mkq5YmbJcTrPExNDi26imrTCpKhepjBFBSHqrBDN2ArPkv
	// Subject (sub): did:key:z6MktA1uBdCpq4uJBqE9jjMiLyxZBg9a6xgPPKJjMqss6Zc2
	// Command (cmd): /foo/bar
	// Policy (pol): TODO
	// Nonce (nonce): 000102030405060708090a0b
	// Meta (meta): TODO
	// NotBefore (nbf): <nil>
	// Expiration (exp): <nil>
}

func exampleCBORData() []byte {
	data, err := base64.StdEncoding.DecodeString("glhAmnAkgfjAx4SA5pzJmtaHRJtTGNpF1y6oqb4yhGoM2H2EUGbBYT4rVDjMKBgCjhdGHjipm00L8iR5SsQh3sIEBaJhaEQ07QFxc3VjYW4vZGxnQDEuMC4wLXJjLjGoY2F1ZHg4ZGlkOmtleTp6Nk1rcTVZbWJKY1RyUEV4TkRpMjZpbXJUQ3BLaGVwakJGQlNIcXJCRE4yQXJQa3ZjY21kaC9mb28vYmFyY2V4cPZjaXNzeDhkaWQ6a2V5Ono2TWtwem4ybjNaR1QyVmFxTUdTUUMzdHptelY0VFM5UzcxaUZzRFhFMVdub05IMmNwb2yDg2I9PWcuc3RhdHVzZWRyYWZ0g2NhbGxpLnJldmlld2Vyg2RsaWtlZi5lbWFpbG0qQGV4YW1wbGUuY29tg2NhbnllLnRhZ3OCYm9ygoNiPT1hLmRuZXdzg2I9PWEuZXByZXNzY3N1Yng4ZGlkOmtleTp6Nk1rdEExdUJkQ3BxNHVKQnFFOWpqTWlMeXhaQmc5YTZ4Z1BQS0pqTXFzczZaYzJkbWV0YaBlbm9uY2VMAAECAwQFBgcICQoL")
	printThenPanicOnErr(err)

	return data
}

func exampleDID(didStr string) did.DID {
	id, err := did.Parse(didStr)
	printThenPanicOnErr(err)

	return id
}

func exampleCommand(cmdStr string) command.Command {
	cmd, err := command.Parse(cmdStr)
	printThenPanicOnErr(err)

	return cmd
}

func examplePolicy(policyJSON string) policy.Policy {
	pol, err := policy.FromDagJson(policyJSON)
	printThenPanicOnErr(err)

	return pol
}

func examplePrivKey(privKeyCfg string) crypto.PrivKey {
	privKeyMar, err := crypto.ConfigDecodeKey(privKeyCfg)
	printThenPanicOnErr(err)

	privKey, err := crypto.UnmarshalPrivateKey(privKeyMar)
	printThenPanicOnErr(err)

	return privKey
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
