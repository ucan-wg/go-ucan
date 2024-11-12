package invocation_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/stretchr/testify/require"
	"github.com/ucan-wg/go-ucan/did"
	"github.com/ucan-wg/go-ucan/pkg/args"
	"github.com/ucan-wg/go-ucan/pkg/command"
	"github.com/ucan-wg/go-ucan/pkg/policy"
	"github.com/ucan-wg/go-ucan/token/delegation"
	"github.com/ucan-wg/go-ucan/token/invocation"
	"github.com/ucan-wg/go-ucan/token/invocation/invocationtest"

	"github.com/stretchr/testify/assert"
)

const (
	rootPrivKeyCfg = "CAESQNMjVXyblHUrtmvdzjfkdrUlLkxh/cHZMFirki7dJLFjW7cCs74VxjO8Wh04f8Xg0uqyKw7N0wTUBiPy2h4hGPQ="
	rootDIDStr     = "did:key:z6MkkdH136Kee5akqx1DiK3hE3WGx4SKmHo11tYw2cWmjjZV"
	rootTknCIDStr  = "bafyreibk6lkgd32zldpqqqsdn7diyosokelrhder5lic4ujadtxl5blkei"

	rootOnlyTknCIDStr = "bafyreidyknlfnsah63jujdkhe5vvjil2yoznlgoaezeq62qqhghaxzpfya"

	dlg0PrivKeyCfg = "CAESQFIb3aD0lZzidUzTtwdpHtCyx1VJxe+Uq4x/S+XQFDDgz/NZIi/TR3rhgUn550RSBOSxNmw0QnR0FOPmAB7SXAg="
	dlg0DIDStr     = "did:key:z6MktT1f5LXQ2MFSUwwTTY9DDU2QdBzZA11V5bjou4YDQY6K"
	dlg0TknCIDStr  = "bafyreihocbstcdvgyeczjoijiyo2bdppyplm2aglqumwtutyapd2zlp2bi"

	expiredTknDagJson    = `[{"/":{"bytes":"U4IH1Q52UrdOyOHkHtXSkH0uf5ouk10k/LTOMz3UvP2k1kqv9/rbGXUwhQCy6JP3s8hc+U3h/lBXYFYzIlULAw"}},{"h":{"/":{"bytes":"NO0BcQ"}},"ucan/dlg@1.0.0-rc.1":{"aud":"did:key:z6MkqTkYrRV6WpFCWRcEHBXnTV3cVqzMtVZREhJKAAswrrGq","cmd":"/seg0","exp":1731439015,"iss":"did:key:z6MktT1f5LXQ2MFSUwwTTY9DDU2QdBzZA11V5bjou4YDQY6K","nonce":{"/":{"bytes":"AAECAwQFBgcICQoL"}},"pol":[],"sub":"did:key:z6MkkdH136Kee5akqx1DiK3hE3WGx4SKmHo11tYw2cWmjjZV"}}]`
	expiredDlg0TknCIDStr = "bafyreigezbrohmjmzv5nxdw2nhjabnrvs52vcv4uvbuscnbyr2jzep2vru"

	inactiveDlg0TknCIDStr = "bafyreiffflxvtfiv722i3lomrqcogc6nncxai3voxeltvrphpuwecdfbgq"

	dlg1PrivKeyCfg = "CAESQPHkXp4OatPxpJ1veMAxEYzP4rd3/sPUz9PRQuJaDKTco5DJTdJC5iXxjCe1VDYAmRwlJOBvBdbsvS3qFgV6zoI="
	dlg1DIDStr     = "did:key:z6MkqTkYrRV6WpFCWRcEHBXnTV3cVqzMtVZREhJKAAswrrGq"
	dlg1TknCIDStr  = "bafyreihjbkvom3tdivzolh6aieuwamz4x6bu3dh667bxytdc5vzo37obo4"

	// dlg2PrivKeyCfg = "CAESQJryJWe3uGt+7BTH2doqsN+2H83rNXv7Yw0tMoKE+lBKTqByESll668G1Jiy9gW/8hm9/jVcrFD9Y1BWyokfNBE="
	// dlg2DIDStr     = "did:key:z6MkjkBk9hzMhcr6rYMXHPEsXy1LAWK5dHniNdbaEPojGXr8"

	// dlg3PrivKeyCfg = "CAESQPrLMlX2p+Dgz70YCyVXbHAJfVMT7lLUYAbuZ1rBKQmBLiD7WJ4Spc4VFZAsJ7HUnkneJWNTk/FFaN2z3pb/OZI="
	// dlg3DIDStr     = "did:key:z6MkhZKy9X4sLtbH1fGQmPVMjXmBAEQ3vAN6DRSfebSBCqpu"

	invPrivKeyCfg = "CAESQHJW8WYTZDRzxjLBjrFN35raIGvVsPoXAJB/5X+J8miboVWVLZFyQmxCAIXOMpwLqWW7R2I98qsCGvxgEJZ5qgY="
	invDIDStr     = "did:key:z6MkqK3NgTnZZo77iptQdU9proJn1ozMmcTSKR98t8sZzJAq"
	invTknCIDStr  = ""

	missingPrivKeyCfg = "CAESQMjRvrEIjpPYRQKmkAGw/pV0XgE958rYa4vlnKJjl1zz/sdnGnyV1xKLJk8D39edyjhHWyqcpgFnozQK62SG16k="
	missingDIDStr     = "bafyreigwypmw6eul6vadi6g6lnfbsfo2zck7gfzsbjoroqs3djhnzzc7mm"
	missingTknCIDStr  = "did:key:z6MkwboxFsH3kEuehBZ5fLkRmxi68yv1u38swA4r9Jm2VRma"
)

func TestToken_ExecutionAllowed(t *testing.T) {
	t.Parallel()

	t.Run("passes - only root", func(t *testing.T) {
		t.Parallel()

		args := invocationtest.EmptyArguments
		prf := invocationtest.Proof(t, rootOnlyTknCIDStr)
		testPasses(t, []string{"seg0"}, args, prf)
	})

	t.Run("passes - valid chain", func(t *testing.T) {
		t.Parallel()

		args := invocationtest.EmptyArguments
		prf := invocationtest.Proof(t, dlg1TknCIDStr, dlg0TknCIDStr, rootTknCIDStr)
		testPasses(t, []string{"seg0"}, args, prf)
	})

	t.Run("fails - no proof", func(t *testing.T) {
		t.Parallel()

		args := args.New()
		testFails(t, invocation.ErrNoProof, []string{"seg0"}, args, []cid.Cid{})
	})

	t.Run("fails - missing referenced delegation", func(t *testing.T) {
		t.Parallel()

		args := invocationtest.EmptyArguments
		prf := invocationtest.Proof(t, missingDIDStr, rootTknCIDStr)
		testFails(t, invocation.ErrMissingDelegation, []string{"seg0"}, args, prf)
	})

	t.Run("fails - referenced delegation expired", func(t *testing.T) {
		t.Parallel()

		args := invocationtest.EmptyArguments
		prf := invocationtest.Proof(t, dlg1TknCIDStr, expiredDlg0TknCIDStr, rootTknCIDStr)
		testFails(t, invocation.ErrDelegationExpired, []string{"seg0"}, args, prf)
	})

	t.Run("fails - referenced delegation inactive", func(t *testing.T) {
		t.Parallel()

		args := invocationtest.EmptyArguments
		prf := invocationtest.Proof(t, dlg1TknCIDStr, inactiveDlg0TknCIDStr, rootTknCIDStr)
		testFails(t, invocation.ErrDelegationInactive, []string{"seg0"}, args, prf)
	})

	t.Run("fails - last (or only) delegation not root", func(t *testing.T) {
		t.Parallel()

		args := args.New()
		prf := invocationtest.Proof(t, dlg1TknCIDStr)
		testFails(t, invocation.ErrLastNotRoot, []string{"seg0"}, args, prf)
	})

	t.Run("fails - broken chain", func(t *testing.T) {
		t.Parallel()

		args := invocationtest.EmptyArguments
		prf := invocationtest.Proof(t, dlg1TknCIDStr, rootTknCIDStr)
		testFails(t, invocation.ErrBrokenChain, []string{"seg0"}, args, prf)
	})

	t.Run("fails - first not issued to invoker", func(t *testing.T) {
		t.Parallel()

		args := invocationtest.EmptyArguments
		prf := invocationtest.Proof(t, dlg0TknCIDStr, rootTknCIDStr)
		testFails(t, invocation.ErrNotIssuedToInvoker, []string{"seg0"}, args, prf)
	})
}

func test(t *testing.T, cmdSegs []string, args *args.Args, prf []cid.Cid, opts ...invocation.Option) (bool, error) {
	ldr := newTestDelegationLoader(t)
	tkn := testInvocation(t, invPrivKeyCfg, rootDIDStr, rootDIDStr, cmdSegs, args, prf, opts...)

	return tkn.ExecutionAllowed(ldr)
}

func testFails(t *testing.T, expErr error, cmdSegs []string, args *args.Args, prf []cid.Cid, opts ...invocation.Option) {
	t.Helper()

	ok, err := test(t, cmdSegs, args, prf, opts...)
	require.ErrorIs(t, err, expErr)
	assert.False(t, ok)
}

func testPasses(t *testing.T, cmdSegs []string, args *args.Args, prf []cid.Cid, opts ...invocation.Option) {
	ok, err := test(t, cmdSegs, args, prf, opts...)
	require.NoError(t, err)
	assert.True(t, ok)
}

var _ invocation.DelegationLoader = (*testDelegationLoader)(nil)

type testDelegationLoader struct {
	tkns map[cid.Cid]*delegation.Token
}

func newTestDelegationLoader(t *testing.T) *testDelegationLoader {
	t.Helper()

	cmdSegs := []string{"seg0"}
	ldr := &testDelegationLoader{
		tkns: map[cid.Cid]*delegation.Token{},
	}

	// aud, err := did.Parse(dlg0DIDStr)
	// require.NoError(t, err)
	// cmd := command.New(cmdSegs...)

	// pol, err := policy.FromDagJson("[]")
	// require.NoError(t, err)

	// rootDlg, err := delegation.Root(rootPrivKey, aud, cmd, pol)
	// require.NoError(t, err)
	// _, rootCID, err := rootDlg.ToSealed(rootPrivKey)
	// require.NoError(t, err)

	// Do not add this one to the loader
	_, missingCID := testDelegation(t, missingPrivKeyCfg, dlg1DIDStr, rootDIDStr, cmdSegs, "[]")
	t.Log("Missing CID", missingCID)

	rootTkn, rootDlgCID := testDelegation(t, rootPrivKeyCfg, dlg0DIDStr, rootDIDStr, cmdSegs, "[]")
	t.Log("Root chain CID:", rootDlgCID)
	ldr.tkns[rootDlgCID] = rootTkn

	rootOnlyTkn, rootOnlyDlgCID := testDelegation(t, rootPrivKeyCfg, invDIDStr, rootDIDStr, cmdSegs, "[]")
	t.Log("Root only chain CID:", rootOnlyDlgCID)
	ldr.tkns[rootOnlyDlgCID] = rootOnlyTkn

	dlg0Tkn, dlg0CID := testDelegation(t, dlg0PrivKeyCfg, dlg1DIDStr, rootDIDStr, cmdSegs, "[]")
	t.Log("Dlg0 CID:", dlg0CID)
	ldr.tkns[dlg0CID] = dlg0Tkn

	// exp := time.Now().Add(time.Second)

	// expiredDlg0Tkn, expiredDlg0CID := testDelegation(t, dlg0PrivKeyCfg, dlg1DIDStr, rootDIDStr, cmdSegs, "[]", delegation.WithExpiration(exp))
	// t.Log("Expired Dlg0 CID:", expiredDlg0CID)

	// dlg0PrivKeyEnc, err := crypto.ConfigDecodeKey(dlg0PrivKeyCfg)
	// require.NoError(t, err)
	// dlg0PrivKey, err := crypto.UnmarshalPrivateKey(dlg0PrivKeyEnc)
	// require.NoError(t, err)

	// dlg0DagJson, err := expiredDlg0Tkn.ToDagJson(dlg0PrivKey)
	// require.NoError(t, err)

	// t.Log("Expired token DAGJSON:", string(dlg0DagJson))

	expiredDlg0Tkn, err := delegation.FromDagJson([]byte(expiredTknDagJson))
	require.NoError(t, err)
	expiredDlg0TknCID, err := cid.Parse(expiredDlg0TknCIDStr)
	ldr.tkns[expiredDlg0TknCID] = expiredDlg0Tkn

	nbf, err := time.Parse(time.RFC3339, "2500-01-01T00:00:00Z")
	require.NoError(t, err)

	inactiveDlg0Tkn, inactiveDlg0CID := testDelegation(t, dlg0PrivKeyCfg, dlg1DIDStr, rootDIDStr, cmdSegs, "[]", delegation.WithNotBefore(nbf))
	t.Log("Inactive Dlg0 CID:", inactiveDlg0CID)
	ldr.tkns[inactiveDlg0CID] = inactiveDlg0Tkn

	dlg1Tkn, dlg1CID := testDelegation(t, dlg1PrivKeyCfg, invDIDStr, rootDIDStr, cmdSegs, "[]")
	t.Log("Dlg1 CID:", dlg1CID)
	ldr.tkns[dlg1CID] = dlg1Tkn

	// for i := 0; i < 5; i++ {
	// 	privKey, id, err := did.GenerateEd25519()
	// 	require.NoError(t, err)

	// 	privKeyEnc, err := crypto.MarshalPrivateKey(privKey)
	// 	require.NoError(t, err)

	// 	t.Log("PrivKey:", crypto.ConfigEncodeKey(privKeyEnc), "DID:", id)

	// }

	t.Log("Delegation loader length:", len(ldr.tkns))

	return ldr
}

var ErrUnknownDelegationTokenCID = errors.New("unknown delegation token CID")

func (l *testDelegationLoader) GetDelegation(c cid.Cid) (*delegation.Token, error) {
	tkn, ok := l.tkns[c]
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrUnknownDelegationTokenCID, c.String())
	}

	return tkn, nil
}

func testDelegation(t *testing.T, privKeyCfg string, audStr string, subStr string, cmdSegs []string, polDAGJSON string, opts ...delegation.Option) (*delegation.Token, cid.Cid) {
	t.Helper()

	privKey, aud, sub, cmd := parseTokenArgs(t, privKeyCfg, audStr, subStr, cmdSegs)

	pol, err := policy.FromDagJson(polDAGJSON)
	require.NoError(t, err)

	opts = append(
		opts,
		delegation.WithNonce([]byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b}),
		delegation.WithSubject(sub),
	)

	tkn, err := delegation.New(privKey, aud, cmd, pol, opts...)
	require.NoError(t, err)

	_, id, err := tkn.ToSealed(privKey)
	require.NoError(t, err)

	return tkn, id
}

func testInvocation(t *testing.T, privKeyCfg string, audStr string, subStr string, cmdSegs []string, args *args.Args, prf []cid.Cid, opts ...invocation.Option) *invocation.Token {
	t.Helper()

	privKey, _, sub, cmd := parseTokenArgs(t, privKeyCfg, audStr, subStr, cmdSegs)

	iss, err := did.FromPrivKey(privKey)
	require.NoError(t, err)

	tkn, err := invocation.New(iss, sub, cmd, prf)
	require.NoError(t, err)

	return tkn
}

func testProof(t *testing.T, cidStrs ...string) []cid.Cid {
	var prf []cid.Cid

	for cidStr := range cidStrs {
		id, err := cid.Parse(cidStr)
		require.NoError(t, err)

		prf = append(prf, id)
	}

	return prf
}

func parseTokenArgs(
	t *testing.T,
	privKeyCfg string,
	audStr string,
	subStr string,
	cmdSegs []string,
) (
	privKey crypto.PrivKey,
	aud did.DID,
	sub did.DID,
	cmd command.Command,
) {
	t.Helper()

	var err error

	privKeyEnc, err := crypto.ConfigDecodeKey(privKeyCfg)
	require.NoError(t, err)
	privKey, err = crypto.UnmarshalPrivateKey(privKeyEnc)
	require.NoError(t, err)

	aud, err = did.Parse(audStr)
	require.NoError(t, err)

	sub, err = did.Parse(subStr)
	require.NoError(t, err)

	cmd = command.New(cmdSegs...)

	return
}
