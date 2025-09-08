package snmp

import (
	"testing"
)

const testWalkLength = 242

func TestReadWalk(t *testing.T) {
	h, err := NewFileHandlerFromFile("testdata/if-mib.txt")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if testWalkLength != len(h.Data) {
		t.Fatalf("expected %d, got %d", testWalkLength, len(h.Data))
	}

}

func TestIsValidOid(t *testing.T) {
	err := IsValidOid(".1.2.3.4.6.999999")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = IsValidOid("")
	if err == nil {
		t.Fatalf("expected error, got none")
	}

	err = IsValidOid("1.2.3.4")
	if err == nil {
		t.Fatalf("expected error, got none")
	}

	err = IsValidOid(".a.b.c.d")
	if err == nil {
		t.Fatalf("expected error, got none")
	}
}

func TestEnsureValidOid(t *testing.T) {
	oid, err := EnsureValidOid("1.2.3.4.6.999999")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if ".1.2.3.4.6.999999" != oid {
		t.Fatalf("expected %s, got %s", ".1.2.3.4.6.999999", oid)
	}

	oid, err = EnsureValidOid(".1.2.3.4.6.999999")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if ".1.2.3.4.6.999999" != oid {
		t.Fatalf("expected %s, got %s", ".1.2.3.4.6.999999", oid)
	}
}
