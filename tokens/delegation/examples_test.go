package delegation_test

import (
	"bytes"
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
	"github.com/ucan-wg/go-ucan/tokens/delegation"
	"github.com/ucan-wg/go-ucan/tokens/internal/envelope"
)

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
	panicAndPrintIfErr(err)
	data, id, err := tkn.ToSealed(issuerPrivKey)
	panicAndPrintIfErr(err)

	printCIDAndSealed(id, data)

	//Output:
	//CID: zdpuAw26pFuvZa2Z9YAtpZZnWN6VmnRFr7Z8LVY5c7RVWoxGY
	//DAG-CBOR (hex) out: 8258409a702481f8c0c78480e69cc99ad687449b5318da45d72ea8a9be32846a0cd87d845066c1613e2b5438cc2818028e17461e38a99b4d0bf224794ac421dec20405a261684434ed0171737563616e2f646c6740312e302e302d72632e31a86361756478386469643a6b65793a7a364d6b7135596d624a6354725045784e44693236696d725443704b6865706a4246425348717242444e324172506b7663636d64682f666f6f2f62617263657870f66369737378386469643a6b65793a7a364d6b707a6e326e335a4754325661714d4753514333747a6d7a563454533953373169467344584531576e6f4e483263706f6c8383623d3d672e7374617475736564726166748363616c6c692e726576696577657283646c696b65662e656d61696c6d2a406578616d706c652e636f6d8363616e79652e7461677382626f728283623d3d612e646e65777383623d3d612e6570726573736373756278386469643a6b65793a7a364d6b74413175426443707134754a427145396a6a4d694c79785a4267396136786750504b4a6a4d717373365a6332646d657461a0656e6f6e63654c000102030405060708090a0b
	//Converted to DAG-JSON out:
	//[
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
	//]
}

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
	panicAndPrintIfErr(err)
	data, id, err := tkn.ToSealed(issuerPrivKey)
	panicAndPrintIfErr(err)

	printCIDAndSealed(id, data)

	//Output:
	//CID: zdpuAnbsR3e6DK8hBk5WA7KwbHYN6CKY4a3Bv1GNehvFYShQ8
	//DAG-CBOR (hex) out: 825840ebb01205ccc5ff09483f411210d9fee193502ae923713373f9fa3b2b6b586ba3949b4ad62020c92640d69bb949790b7e2af480f98e1cb474d06c0af72ebee60ea261684434ed0171737563616e2f646c6740312e302e302d72632e31a86361756478386469643a6b65793a7a364d6b7135596d624a6354725045784e44693236696d725443704b6865706a4246425348717242444e324172506b7663636d64682f666f6f2f62617263657870f66369737378386469643a6b65793a7a364d6b707a6e326e335a4754325661714d4753514333747a6d7a563454533953373169467344584531576e6f4e483263706f6c8383623d3d672e7374617475736564726166748363616c6c692e726576696577657283646c696b65662e656d61696c6d2a406578616d706c652e636f6d8363616e79652e7461677382626f728283623d3d612e646e65777383623d3d612e6570726573736373756278386469643a6b65793a7a364d6b707a6e326e335a4754325661714d4753514333747a6d7a563454533953373169467344584531576e6f4e4832646d657461a0656e6f6e63654c000102030405060708090a0b
	//Converted to DAG-JSON out:
	//[
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
	//]
}

func ExampleToken_FromSealed() {
	cborBytes := exampleCBORData()
	fmt.Println("DAG-CBOR (hex) in:", hex.EncodeToString(cborBytes))

	tkn, err := delegation.FromSealed(cborBytes)
	panicAndPrintIfErr(err)

	fmt.Println("CID:", envelope.CIDToBase58BTC(tkn.CID()))
	fmt.Println("Issuer (iss):", tkn.Issuer().String())
	fmt.Println("Audience (aud):", tkn.Audience().String())
	fmt.Println("Subject (sub):", tkn.Subject().String())
	fmt.Println("Command (cmd):", tkn.Command().String())
	fmt.Println("Policy (pol): TODO")
	fmt.Println("Nonce (nonce):", hex.EncodeToString(tkn.Nonce()))
	fmt.Println("Meta (meta): TODO")
	fmt.Println("NotBefore (nbf):", tkn.NotBefore())
	fmt.Println("Expiration (exp):", tkn.Expiration())

	//Output:
	//DAG-CBOR (hex) in: 825840ebb01205ccc5ff09483f411210d9fee193502ae923713373f9fa3b2b6b586ba3949b4ad62020c92640d69bb949790b7e2af480f98e1cb474d06c0af72ebee60ea261684434ed0171737563616e2f646c6740312e302e302d72632e31a86361756478386469643a6b65793a7a364d6b7135596d624a6354725045784e44693236696d725443704b6865706a4246425348717242444e324172506b7663636d64682f666f6f2f62617263657870f66369737378386469643a6b65793a7a364d6b707a6e326e335a4754325661714d4753514333747a6d7a563454533953373169467344584531576e6f4e483263706f6c8383623d3d672e7374617475736564726166748363616c6c692e726576696577657283646c696b65662e656d61696c6d2a406578616d706c652e636f6d8363616e79652e7461677382626f728283623d3d612e646e65777383623d3d612e6570726573736373756278386469643a6b65793a7a364d6b707a6e326e335a4754325661714d4753514333747a6d7a563454533953373169467344584531576e6f4e4832646d657461a0656e6f6e63654c000102030405060708090a0b
	//CID: zdpuAnbsR3e6DK8hBk5WA7KwbHYN6CKY4a3Bv1GNehvFYShQ8
	//Issuer (iss): did:key:z6Mkpzn2n3ZGT2VaqMGSQC3tzmzV4TS9S71iFsDXE1WnoNH2
	//Audience (aud): did:key:z6Mkq5YmbJcTrPExNDi26imrTCpKhepjBFBSHqrBDN2ArPkv
	//Subject (sub): did:key:z6Mkpzn2n3ZGT2VaqMGSQC3tzmzV4TS9S71iFsDXE1WnoNH2
	//Command (cmd): /foo/bar
	//Policy (pol): TODO
	//Nonce (nonce): 000102030405060708090a0b
	//Meta (meta): TODO
	//NotBefore (nbf): <nil>
	//Expiration (exp): <nil>
}

func exampleCBORData() []byte {
	data, err := hex.DecodeString("825840ebb01205ccc5ff09483f411210d9fee193502ae923713373f9fa3b2b6b586ba3949b4ad62020c92640d69bb949790b7e2af480f98e1cb474d06c0af72ebee60ea261684434ed0171737563616e2f646c6740312e302e302d72632e31a86361756478386469643a6b65793a7a364d6b7135596d624a6354725045784e44693236696d725443704b6865706a4246425348717242444e324172506b7663636d64682f666f6f2f62617263657870f66369737378386469643a6b65793a7a364d6b707a6e326e335a4754325661714d4753514333747a6d7a563454533953373169467344584531576e6f4e483263706f6c8383623d3d672e7374617475736564726166748363616c6c692e726576696577657283646c696b65662e656d61696c6d2a406578616d706c652e636f6d8363616e79652e7461677382626f728283623d3d612e646e65777383623d3d612e6570726573736373756278386469643a6b65793a7a364d6b707a6e326e335a4754325661714d4753514333747a6d7a563454533953373169467344584531576e6f4e4832646d657461a0656e6f6e63654c000102030405060708090a0b")
	panicAndPrintIfErr(err)

	return data
}

func exampleDID(didStr string) did.DID {
	id, err := did.Parse(didStr)
	panicAndPrintIfErr(err)

	return id
}

func exampleCommand(cmdStr string) command.Command {
	cmd, err := command.Parse(cmdStr)
	panicAndPrintIfErr(err)

	return cmd
}

func examplePolicy(policyJSON string) policy.Policy {
	pol, err := policy.FromDagJson(policyJSON)
	panicAndPrintIfErr(err)

	return pol
}

func examplePrivKey(privKeyCfg string) crypto.PrivKey {
	privKeyMar, err := crypto.ConfigDecodeKey(privKeyCfg)
	panicAndPrintIfErr(err)

	privKey, err := crypto.UnmarshalPrivateKey(privKeyMar)
	panicAndPrintIfErr(err)

	return privKey
}

func panicAndPrintIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func printCIDAndSealed(id cid.Cid, data []byte) {
	fmt.Println("CID:", envelope.CIDToBase58BTC(id))
	fmt.Println("DAG-CBOR (hex) out:", hex.EncodeToString(data))
	fmt.Println("Converted to DAG-JSON out:")

	node, err := ipld.Decode(data, dagcbor.Decode)
	panicAndPrintIfErr(err)

	rawJSON, err := ipld.Encode(node, dagjson.Encode)
	panicAndPrintIfErr(err)

	prettyJSON := &bytes.Buffer{}
	err = json.Indent(prettyJSON, rawJSON, "", "\t")
	panicAndPrintIfErr(err)

	fmt.Println(prettyJSON.String())
}
