package didkey

import (
	"testing"
)

func TestID(t *testing.T) {
	keyStrED := "did:key:z6MkpTHR8VNsBxYAAWHut2Geadd9jSwuBV8xRoAnwWsdvktH"
	id, err := Parse(keyStrED)
	if err != nil {
		t.Fatal(err)
	}

	if id.String() != keyStrED {
		t.Errorf("string mismatch.\nwant: %q\ngot:  %q", keyStrED, id.String())
	}

	keyStrRSA := "did:key:zDQw3CC91XtvcBm5kY3FGBFgJarujMEKcarucbThDbGWVo3cAQPgvXuH5g9kxe9PNVHhAQpRQzTZQ3Pchra5pqKt6SgwTdgHXJtcWwp2pqkAmuNqhUj22naAg526RvU8u5ZE3tidrxh8zajWsrXriFwkVtZfTbEBGKewUMRVhttA3hc8Rpa6gL5HRqZ44Uq1fdKzdsBMTfG5ohhRzesiwoEtRyb3w4SPhdVYrD5Rd2KJPyTUqCvuuFwLAJJcDaPkp36RnRnFbbtZEydPtdMscnJTmj5nD9kpSmkTHU2riMVLcGnJWRHkHWZPnfbtuXcQVEwiBSjpJVRcCFG6LDbgTqt4s6BYxJqT5zriEaigzvFuj9V4s7QSJCs1Sj92tgDmE6YyTSR7UN9di2oHDCzk"

	id, err = Parse(keyStrRSA)
	if err != nil {
		t.Fatal(err)
	}

	if id.String() != keyStrRSA {
		t.Errorf("string mismatch.\nwant: %q\ngot:  %q", keyStrRSA, id.String())
	}
}
