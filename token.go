// Package ucan implements User-Controlled Authorization Network tokens by
// fission:
// https://whitepaper.fission.codes/access-control/ucan/ucan-tokens
//
// From the paper:
// The UCAN format is designed as an authenticated digraph in some larger
// authorization space. The other way to view this is as a function from a set
// of authorizations (“UCAN proofs“) to a subset output (“UCAN capabilities”).
package ucan

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/crypto"
	mh "github.com/multiformats/go-multihash"
)

// ErrInvalidToken indicates an access token is invalid
var ErrInvalidToken = errors.New("invalid access token")

const (
	// UCANVersion is the current version of the UCAN spec
	UCANVersion = "0.4.0"
	// UCANVersionKey is the key used in version headers for the UCAN spec
	UCANVersionKey = "ucv"
	// PrfKey denotes "Proofs" in a UCAN. Stored in JWT Claims
	PrfKey = "prf"
	// FctKey denotes "Facts" in a UCAN. Stored in JWT Claims
	FctKey = "fct"
	// AttKey denotes "Attenuations" in a UCAN. Stored in JWT Claims
	AttKey = "att"
	// CapKey indicates a resource Capability. Used in an attenuation
	CapKey = "cap"
)

// UCAN is a JSON Web Token (JWT) that contains special keys
type UCAN struct {
	// Entire UCAN as a signed JWT string
	Raw string
	// the "inputs" to this token, a chain UCAN tokens with broader scopes &
	// deadlines than this token
	Proofs []Proof `json:"prf,omitempty"`
	// the "outputs" of this token, an array of heterogenous resources &
	// capabilities
	Attenuations Attenuations `json:"att,omitempty"`
	// Facts are facts, jack.
	Facts []Fact `json:"fct,omitempty"`
}

// CID calculates the cid of a UCAN using the default prefix
func (t *UCAN) CID() (cid.Cid, error) {
	pref := cid.Prefix{
		Version:  1,
		Codec:    cid.Raw,
		MhType:   mh.SHA2_256,
		MhLength: -1, // default length
	}

	return t.PrefixCID(pref)
}

// PrefixCID calculates the CID of a token with a supplied prefix
func (t *UCAN) PrefixCID(pref cid.Prefix) (cid.Cid, error) {
	return pref.Sum([]byte(t.Raw))
}

// Claims is the claims component of a UCAN token. UCAN claims are expressed
// as a standard JWT claims object with additional special fields
type Claims struct {
	*jwt.StandardClaims
	// the "inputs" to this token, a chain UCAN tokens with broader scopes &
	// deadlines than this token
	// Proofs are UCAN chains, leading back to a self-evident origin token
	Proofs []Proof `json:"prf,omitempty"`
	// the "outputs" of this token, an array of heterogenous resources &
	// capabilities
	Attenuations Attenuations `json:"att,omitempty"`
	// Facts are facts, jack.
	Facts []Fact `json:"fct,omitempty"`
}

// Fact is self-evident statement
type Fact struct {
	cidString string
	value     map[string]interface{}
}

// func (fct *Fact) MarshalJSON() (p[])

// func (fct *Fact) UnmarshalJSON(p []byte) error {
// 	var str string
// 	if json.Unmarshal(p, &str); err == nil {
// 	}
// }

// CIDBytesResolver is a small interface for turning a CID into the bytes
// they reference. In practice this may be backed by a network connection that
// can fetch CIDs, eg: IPFS.
type CIDBytesResolver interface {
	ResolveCIDBytes(ctx context.Context, id cid.Cid) ([]byte, error)
}

// Source creates tokens, and provides a verification key for all tokens it
// creates
//
// implementations of Source must conform to the assertion test defined in the
// spec subpackage
type Source interface {
	NewOriginUCAN(subjectDID string, att Attenuations, fct []Fact, notBefore, expires time.Time) (*UCAN, error)
	NewAttenuatedUCAN(parent *UCAN, subjectDID string, att Attenuations, fct []Fact, notBefore, expires time.Time) (*UCAN, error)
}

type pkSource struct {
	pk            crypto.PrivKey
	issuerDID     string
	signingMethod jwt.SigningMethod

	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey

	ap       AttenuationConstructor
	resolver CIDBytesResolver
	store    TokenStore
}

// assert pkSource implements tokens at compile time
var _ Source = (*pkSource)(nil)

// NewPrivKeySource creates an authentication interface backed by a single
// private key. Intended for a node running as remote, or providing a public API
func NewPrivKeySource(privKey crypto.PrivKey) (Source, error) {
	methodStr := ""
	keyType := privKey.Type()
	switch keyType {
	case crypto.RSA:
		methodStr = "RS256"
	case crypto.Ed25519:
		methodStr = "EdDSA"
	default:
		return nil, fmt.Errorf("unsupported key type for token creation: %q", keyType)
	}

	signingMethod := jwt.GetSigningMethod(methodStr)

	rawPrivBytes, err := privKey.Raw()
	if err != nil {
		return nil, err
	}
	signKey, err := x509.ParsePKCS1PrivateKey(rawPrivBytes)
	if err != nil {
		return nil, err
	}

	rawPubBytes, err := privKey.GetPublic().Raw()
	if err != nil {
		return nil, err
	}
	verifyKeyiface, err := x509.ParsePKIXPublicKey(rawPubBytes)
	if err != nil {
		return nil, err
	}
	verifyKey, ok := verifyKeyiface.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("public key is not an RSA key. got type: %T", verifyKeyiface)
	}

	issuerDID, err := DIDStringFromPublicKey(privKey.GetPublic())
	if err != nil {
		return nil, err
	}

	return &pkSource{
		pk:            privKey,
		signingMethod: signingMethod,
		verifyKey:     verifyKey,
		signKey:       signKey,
		issuerDID:     issuerDID,
	}, nil
}

func (a *pkSource) NewOriginUCAN(subjectDID string, att Attenuations, fct []Fact, nbf, exp time.Time) (*UCAN, error) {
	return a.newUCAN(subjectDID, nil, att, fct, nbf, exp)
}

func (a *pkSource) NewAttenuatedUCAN(parent *UCAN, subjectDID string, att Attenuations, fct []Fact, nbf, exp time.Time) (*UCAN, error) {
	if !parent.Attenuations.Contains(att) {
		return nil, fmt.Errorf("scope of ucan attenuations must be less than it's parent")
	}
	return a.newUCAN(subjectDID, append(parent.Proofs, Proof(parent.Raw)), att, fct, nbf, exp)
}

// CreateToken returns a new JWT token
func (a *pkSource) newUCAN(subjectDID string, prf []Proof, att Attenuations, fct []Fact, nbf, exp time.Time) (*UCAN, error) {
	// create a signer for rsa 256
	t := jwt.New(a.signingMethod)

	// if _, err := did.Parse(subjectDID); err != nil {
	// 	return nil, fmt.Errorf("invalid subject DID: %w", err)
	// }

	t.Header[UCANVersionKey] = UCANVersion

	var (
		nbfUnix int64
		expUnix int64
	)

	if !nbf.IsZero() {
		nbfUnix = nbf.Unix()
	}
	if !exp.IsZero() {
		expUnix = exp.Unix()
	}

	// set our claims
	t.Claims = &Claims{
		StandardClaims: &jwt.StandardClaims{
			Issuer:    a.issuerDID,
			Subject:   subjectDID,
			NotBefore: nbfUnix,
			// set the expire time
			// see http://tools.ietf.org/html/draft-ietf-oauth-json-web-token-20#section-4.1.4
			ExpiresAt: expUnix,
		},
		Attenuations: att,
		Facts:        fct,
		Proofs:       prf,
	}

	raw, err := t.SignedString(a.signKey)
	if err != nil {
		return nil, err
	}

	return &UCAN{
		Raw:          raw,
		Attenuations: att,
		Facts:        fct,
		Proofs:       prf,
	}, nil
}

type DIDPubKeyResolver interface {
	ResolveDIDKey(ctx context.Context, did string) (crypto.PubKey, error)
}

func DIDStringFromPublicKey(pub crypto.PubKey) (string, error) {
	rawPubBytes, err := pub.Raw()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("did:key:%s", base64.URLEncoding.EncodeToString(rawPubBytes)), nil
}

type StringDIDPubKeyResolver struct{}

func (StringDIDPubKeyResolver) ResolveDIDKey(ctx context.Context, didStr string) (crypto.PubKey, error) {
	// id, err := did.Parse(didStr)
	// if err != nil {
	// 	return nil, fmt.Errorf("invalid DID: %w", err)
	// }

	data, err := base64.URLEncoding.DecodeString(strings.TrimPrefix(didStr, "did:key:"))
	if err != nil {
		return nil, err
	}

	return crypto.UnmarshalRsaPublicKey(data)
}

type UCANParser struct {
	ap   AttenuationConstructor
	cidr CIDBytesResolver
	didr DIDPubKeyResolver
}

func NewUCANParser(ap AttenuationConstructor, didr DIDPubKeyResolver, cidr CIDBytesResolver) *UCANParser {
	return &UCANParser{
		ap:   ap,
		cidr: cidr,
		didr: didr,
	}
}

// ParseAndVerify will parse, validate and return a token
func (p *UCANParser) ParseAndVerify(ctx context.Context, raw string) (*UCAN, error) {
	return p.parseAndVerify(ctx, raw, nil)
}

func (p *UCANParser) parseAndVerify(ctx context.Context, raw string, child *UCAN) (*UCAN, error) {
	tok, err := jwt.Parse(raw, p.matchVerifyKeyFunc(ctx))
	if err != nil {
		return nil, err
	}

	mc, ok := tok.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("parser fail")
	}

	var att Attenuations
	if acci, ok := mc[AttKey].([]interface{}); ok {
		for i, a := range acci {
			if mapv, ok := a.(map[string]interface{}); ok {
				a, err := p.ap(mapv)
				if err != nil {
					return nil, err
				}
				att = append(att, a)
			} else {
				return nil, fmt.Errorf(`"att[%d]" is not an object`, i)
			}
		}
	} else {
		return nil, fmt.Errorf(`"att" key is not an array`)
	}

	var prf []Proof
	if prfi, ok := mc[PrfKey].([]interface{}); ok {
		for i, a := range prfi {
			if pStr, ok := a.(string); ok {
				prf = append(prf, Proof(pStr))
			} else {
				return nil, fmt.Errorf(`"prf[%d]" is not a string`, i)
			}
		}
	} else if mc[PrfKey] != nil {
		return nil, fmt.Errorf(`"prf" key is not an array`)
	}

	return &UCAN{
		Raw:          raw,
		Attenuations: att,
		Proofs:       prf,
	}, nil
}

func (p *UCANParser) matchVerifyKeyFunc(ctx context.Context) func(tok *jwt.Token) (interface{}, error) {
	return func(tok *jwt.Token) (interface{}, error) {
		mc, ok := tok.Claims.(jwt.MapClaims)
		if !ok {
			return nil, fmt.Errorf("parser fail")
		}

		iss, ok := mc["iss"].(string)
		if !ok {
			return nil, fmt.Errorf(`"iss" claims key is required`)
		}

		pubKey, err := p.didr.ResolveDIDKey(ctx, iss)
		if err != nil {
			return nil, err
		}

		rawPubBytes, err := pubKey.Raw()
		if err != nil {
			return nil, err
		}
		verifyKeyiface, err := x509.ParsePKIXPublicKey(rawPubBytes)
		if err != nil {
			return nil, err
		}
		verifyKey, ok := verifyKeyiface.(*rsa.PublicKey)
		if !ok {
			return nil, fmt.Errorf("public key is not an RSA key. got type: %T", verifyKeyiface)
		}

		return verifyKey, nil
	}
}
