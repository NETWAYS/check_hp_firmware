package snmp

import (
	"testing"

	"github.com/gosnmp/gosnmp"
)

func TestNewFileHandlerFromFile(t *testing.T) {
	h, err := NewFileHandlerFromFile("testdata/if-mib.txt")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if testWalkLength != len(h.Data) {
		t.Fatalf("expected %d, got %d", testWalkLength, len(h.Data))
	}
}

func TestFileHandler_Walk(t *testing.T) {
	h, err := NewFileHandlerFromFile("testdata/if-mib.txt")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var counter int

	err = h.Walk(".1.3.6.1.2.1.2.2.1.1", func(pdu gosnmp.SnmpPDU) error {
		counter++
		return nil
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if counter != 11 {
		t.Fatalf("expected %d, got %d", 11, counter)
	}

	// without leading dot
	counter = 0
	err = h.Walk("1.3.6.1.2.1.2.2.1.1", func(pdu gosnmp.SnmpPDU) error {
		counter++
		return nil
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if counter != 11 {
		t.Fatalf("expected %d, got %d", 11, counter)
	}
}

func TestFileHandler_Get(t *testing.T) {
	h, err := NewFileHandlerFromFile("testdata/if-mib.txt")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	oid := ".1.3.6.1.2.1.2.2.1.3.1"

	p, err := h.Get([]string{oid})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(p.Variables) != 1 {
		t.Fatalf("expected %d, got %d", 1, len(p.Variables))
	}

	pdu := p.Variables[0]

	if oid != pdu.Name {
		t.Fatalf("expected %s, got %s", oid, pdu.Name)
	}

	if gosnmp.Integer != pdu.Type {
		t.Fatalf("expected %s, got %s", gosnmp.Integer, pdu.Type)
	}

	if 24 != pdu.Value {
		t.Fatalf("expected %d, got %d", 24, pdu.Value)
	}

	// without leading dot
	oid = "1.3.6.1.2.1.2.2.1.3.1"

	p, err = h.Get([]string{oid})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(p.Variables) != 1 {
		t.Fatalf("expected %d, got %d", 1, len(p.Variables))
	}
}
