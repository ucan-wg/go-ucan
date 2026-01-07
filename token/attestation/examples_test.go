package attestation_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/MetaMask/go-did-it"
	didkeyctl "github.com/MetaMask/go-did-it/controller/did-key"
	"github.com/MetaMask/go-did-it/crypto/ed25519"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"github.com/ipld/go-ipld-prime/node/basicnode"

	"github.com/ucan-wg/go-ucan/token/attestation"
)

func ExampleNew() {
	privKey, iss, _, _, err := setupExampleNew()
	if err != nil {
		fmt.Println("failed to create setup:", err.Error())
		return
	}

	att, err := attestation.New(iss,
		attestation.WithClaim("claim1", "UCAN is great"),
		attestation.WithMeta("env", "development"),
		attestation.WithExpirationIn(time.Minute),
		attestation.WithoutIssuedAt())
	if err != nil {
		fmt.Println("failed to create attestation:", err.Error())
		return
	}

	// foo, _ := att.ToDagJson(privKey)
	// os.WriteFile("testdata/new.dagjson", foo, 0666)
	// fmt.Println(base64.StdEncoding.EncodeToString(privKey.ToBytes()))

	data, cid, err := att.ToSealed(privKey)
	if err != nil {
		fmt.Println("failed to seal attestation:", err.Error())
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
	// CID: bafyreibm5vo6gk75oreefkg6xkrrfb4d5dgkccgmutirjgtzi5j45svjm4
	// Token (pretty DAG-JSON):
	// [
	//  {
	//    "/": {
	//      "bytes": "ApuXUsUYhqostO2zfKZK50GW0gXYPtrlpoVA8EwGFdyYahQecOizVpl+9wy64aqk2rMP4Q0UEUKCTV0ONMdPAw"
	//    }
	//  },
	//  {
	//    "h": {
	//      "/": {
	//        "bytes": "NAHtAe0BE3E"
	//      }
	//    },
	//    "ucan/att@tbd": {
	//      "claims": {
	//        "claim1": "UCAN is great"
	//      },
	//      "exp": 1767790971,
	//      "iss": "did:key:z6Mkm4RzzBDfSHqmwV9dp5jFsLkVgKRYp1PhSj7VybCcLHC4",
	//      "meta": {
	//        "env": "development"
	//      },
	//      "nonce": {
	//        "/": {
	//          "bytes": "NjS8QPvft97jbtUG"
	//        }
	//      }
	//    }
	//  }
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

func setupExampleNew() (privKey ed25519.PrivateKey, iss did.DID, claims map[string]any, meta map[string]any, errs error) {
	var err error

	_, privKey, err = ed25519.GenerateKeyPair()
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("failed to generate Ed25519 keypair: %w", err))
		return
	}
	iss = didkeyctl.FromPrivateKey(privKey)

	claims = map[string]any{
		"claim1": "UCAN is great",
	}

	meta = map[string]any{
		"env": basicnode.NewString("development"),
	}

	return // WARNING: named return values
}
