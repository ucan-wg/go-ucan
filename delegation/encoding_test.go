package delegation_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/ucan-wg/go-ucan/delegation"
)

func TestEncodingRoundTrip(t *testing.T) {
	const delegationJson = `
[
  {
    "/": {
      "bytes": "QWr0Pk+sSWE1nszuBMQzggbHX4ofJb8QRdwrLJK/AGCx2p4s/xaCRieomfstDjsV4ezBzX1HARvcoNgdwDQ8Aw"
    }
  },
  {
    "h": {
      "/": {
        "bytes": "NO0BcQ"
      }
    },
    "ucan/dlg@1.0.0-rc.1": {
      "aud": "did:key:z6Mkpzn2n3ZGT2VaqMGSQC3tzmzV4TS9S71iFsDXE1WnoNH2",
      "cmd": "/foo/bar",
      "exp": -62135596800,
      "iss": "did:key:z6Mkpzn2n3ZGT2VaqMGSQC3tzmzV4TS9S71iFsDXE1WnoNH2",
      "meta": {
        "bar": "barr",
        "foo": "fooo"
      },
      "nbf": -62135596800,
      "nonce": {
        "/": {
          "bytes": "X93ORvN1QIXrKPyEP5m5XoVK9VLX9nX8VV/+HlWrp9c"
        }
      },
      "pol": [
        [
          "==",
          ".status",
          "draft"
        ],
        [
          "all",
          ".reviewer",
          [
            "like",
            ".email",
            "*@example.com"
          ]
        ],
        [
          "any",
          ".tags",
          [
            "or",
            [
              [
                "==",
                ".",
                "news"
              ],
              [
                "==",
                ".",
                "press"
              ]
            ]
          ]
        ]
      ],
      "sub": "did:key:z6MktA1uBdCpq4uJBqE9jjMiLyxZBg9a6xgPPKJjMqss6Zc2"
    }
  }
]
`
	// format:    dagJson   -->   Delegation   -->   dagCbor   -->   Delegation   -->   dagJson
	// function:        FromDagJson()       ToDagCbor()    FromDagCbor()       ToDagJson()

	p1, err := delegation.FromDagJson([]byte(delegationJson))
	require.NoError(t, err)

	cborBytes, err := p1.ToDagCbor()
	require.NoError(t, err)
	fmt.Println("cborBytes length", len(cborBytes))
	fmt.Println("cbor", string(cborBytes))

	p2, err := delegation.FromDagCbor(cborBytes)
	require.NoError(t, err)
	fmt.Println("read Cbor", p2)

	readJson, err := p2.ToDagJson()
	require.NoError(t, err)
	fmt.Println("readJson length", len(readJson))
	fmt.Println("json: ", string(readJson))

	require.JSONEq(t, delegationJson, string(readJson))
}
