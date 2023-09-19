package did

import (
	"testing"
)

func TestParseDIDKey(t *testing.T) {
	str := "did:key:z6Mkod5Jr3yd5SC7UDueqK4dAAw5xYJYjksy722tA9Boxc4z"
	d, err := Parse(str)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if d.String() != str {
		t.Fatalf("expected %v to equal %v", d.String(), str)
	}
}

func TestDecodeDIDKey(t *testing.T) {
	str := "did:key:z6Mkod5Jr3yd5SC7UDueqK4dAAw5xYJYjksy722tA9Boxc4z"
	d0, err := Parse(str)
	if err != nil {
		t.Fatalf("%v", err)
	}
	d1, err := Decode(d0.Bytes())
	if err != nil {
		t.Fatalf("%v", err)
	}
	if d1.String() != str {
		t.Fatalf("expected %v to equal %v", d1.String(), str)
	}
}

func TestParseDIDWeb(t *testing.T) {
	str := "did:web:up.web3.storage"
	d, err := Parse(str)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if d.String() != str {
		t.Fatalf("expected %v to equal %v", d.String(), str)
	}
}

func TestDecodeDIDWeb(t *testing.T) {
	str := "did:web:up.web3.storage"
	d0, err := Parse(str)
	if err != nil {
		t.Fatalf("%v", err)
	}
	d1, err := Decode(d0.Bytes())
	if err != nil {
		t.Fatalf("%v", err)
	}
	if d1.String() != str {
		t.Fatalf("expected %v to equal %v", d1.String(), str)
	}
}

func TestEquivalence(t *testing.T) {
	u0 := DID{}
	u1 := Undef
	if u0 != u1 {
		t.Fatalf("undef DID not equivalent")
	}

	d0, err := Parse("did:key:z6Mkod5Jr3yd5SC7UDueqK4dAAw5xYJYjksy722tA9Boxc4z")
	if err != nil {
		t.Fatalf("%v", err)
	}

	d1, err := Parse("did:key:z6Mkod5Jr3yd5SC7UDueqK4dAAw5xYJYjksy722tA9Boxc4z")
	if err != nil {
		t.Fatalf("%v", err)
	}

	if d0 != d1 {
		t.Fatalf("two equivalent DID not equivalent")
	}
}
