package invocation_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	"github.com/libp2p/go-libp2p/core/crypto"

	"github.com/ucan-wg/go-ucan/did"
	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/token/invocation"
)

func ExampleNew() {
	privKey, iss, sub, cmd, args, prf, meta, err := setupExampleNew()
	if err != nil {
		fmt.Println("failed to create setup:", err.Error())

		return
	}

	inv, err := invocation.New(iss, cmd, sub, prf,
		invocation.WithArgument("uri", args["uri"]),
		invocation.WithArgument("headers", args["headers"]),
		invocation.WithArgument("payload", args["payload"]),
		invocation.WithMeta("env", "development"),
		invocation.WithMeta("tags", meta["tags"]),
		invocation.WithExpirationIn(time.Minute),
		invocation.WithoutIssuedAt())
	if err != nil {
		fmt.Println("failed to create invocation:", err.Error())

		return
	}

	data, cid, err := inv.ToSealed(privKey)
	if err != nil {
		fmt.Println("failed to seal invocation:", err.Error())

		return
	}

	json, err := prettyDAGJSON(data)
	if err != nil {
		fmt.Println("failed to pretty DAG-JSON:", err.Error())

		return
	}

	fmt.Println("CID:", cid)
	fmt.Println("Token (pretty DAG-JSON):")
	fmt.Println(json)

	// Expected CID and DAG-JSON output:
	// CID: bafyreid2n5q45vk4osned7k5huocbe3mxbisonh5vujepqftc5ftr543ae
	// Token (pretty DAG-JSON):
	// [
	//   {
	// 	"/": {
	// 	  "bytes": "gvyL7kdSkgmaDpDU/Qj9ohRwxYLCHER52HFMSFEqQqEcQC9qr4JCPP1f/WybvGGuVzYiA0Hx4JO+ohNz8BxUAA"
	// 	}
	//   },
	//   {
	// 	"h": {
	// 	  "/": {
	// 		"bytes": "NO0BcQ"
	// 	  }
	// 	},
	// 	"ucan/inv@1.0.0-rc.1": {
	// 	  "args": {
	// 		"headers": {
	// 		  "Content-Type": "application/json"
	// 		},
	// 		"payload": {
	// 		  "body": "UCAN is great",
	// 		  "draft": true,
	// 		  "title": "UCAN for Fun and Profit",
	// 		  "topics": [
	// 			"authz",
	// 			"journal"
	// 		  ]
	// 		},
	// 		"uri": "https://example.com/blog/posts"
	// 	  },
	// 	  "cmd": "/crud/create",
	// 	  "exp": 1729788921,
	// 	  "iss": "did:key:z6MkhniGGyP88eZrq2dpMvUPdS2RQMhTUAWzcu6kVGUvEtCJ",
	// 	  "meta": {
	// 		"env": "development",
	// 		"tags": [
	// 		  "blog",
	// 		  "post",
	// 		  "pr#123"
	// 		]
	// 	  },
	// 	  "nonce": {
	// 		"/": {
	// 		  "bytes": "2xXPoZwWln1TfXIp"
	// 		}
	// 	  },
	// 	  "prf": [
	// 		{
	// 		  "/": "bafyreigx3qxd2cndpe66j2mdssj773ecv7tqd7wovcnz5raguw6lj7sjoe"
	// 		},
	// 		{
	// 		  "/": "bafyreib34ira254zdqgehz6f2bhwme2ja2re3ltcalejv4x4tkcveujvpa"
	// 		},
	// 		{
	// 		  "/": "bafyreibkb66tpo2ixqx3fe5hmekkbuasrod6olt5bwm5u5pi726mduuwlq"
	// 		}
	// 	  ],
	// 	  "sub": "did:key:z6MktWuvPvBe5UyHnDGuEdw8aJ5qrhhwLG6jy7cQYM6ckP6P"
	// 	}
	//   }
	// ]
}

func prettyDAGJSON(data []byte) (string, error) {
	var node ipld.Node

	node, err := ipld.Decode(data, dagcbor.Decode)
	if err != nil {
		return "", err
	}

	jsonData, err := ipld.Encode(node, dagjson.Encode)
	if err != nil {
		return "", err
	}

	var out bytes.Buffer
	if err := json.Indent(&out, jsonData, "", "  "); err != nil {
		return "", err
	}

	return out.String(), nil
}

func setupExampleNew() (privKey crypto.PrivKey, iss, sub did.DID, cmd command.Command, args map[string]any, prf []cid.Cid, meta map[string]any, errs error) {
	var err error

	privKey, iss, err = did.GenerateEd25519()
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("failed to generate Issuer identity: %w", err))
	}

	_, sub, err = did.GenerateEd25519()
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("failed to generate Subject identity: %w", err))
	}

	cmd, err = command.Parse("/crud/create")
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("failed to parse command: %w", err))
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	payload := map[string]any{
		"body":   "UCAN is great",
		"draft":  true,
		"title":  "UCAN for Fun and Profit",
		"topics": []string{"authz", "journal"},
	}

	args = map[string]any{
		// you can also directly pass IPLD values
		"uri":     basicnode.NewString("https://example.com/blog/posts"),
		"headers": headers,
		"payload": payload,
	}

	prf = make([]cid.Cid, 3)
	for i, v := range []string{
		"zdpuAzx4sBrBCabrZZqXgvK3NDzh7Mf5mKbG11aBkkMCdLtCp",
		"zdpuApTCXfoKh2sB1KaUaVSGofCBNPUnXoBb6WiCeitXEibZy",
		"zdpuAoFdXRPw4n6TLcncoDhq1Mr6FGbpjAiEtqSBrTSaYMKkf",
	} {
		prf[i], err = cid.Parse(v)
		if err != nil {
			errs = errors.Join(errs, fmt.Errorf("failed to parse proof cid: %w", err))
		}
	}

	meta = map[string]any{
		"env":  basicnode.NewString("development"),
		"tags": []string{"blog", "post", "pr#123"},
	}

	return // WARNING: named return values
}
