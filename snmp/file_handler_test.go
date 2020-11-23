package snmp

import (
	"github.com/gosnmp/gosnmp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFileHandlerFromFile(t *testing.T) {
	h, err := NewFileHandlerFromFile("testdata/if-mib.txt")
	assert.NoError(t, err)
	assert.Equal(t, testWalkLength, len(h.Data))
}

func TestFileHandler_Walk(t *testing.T) {
	h, err := NewFileHandlerFromFile("testdata/if-mib.txt")
	assert.NoError(t, err)

	var counter int

	err = h.Walk(".1.3.6.1.2.1.2.2.1.1", func(pdu gosnmp.SnmpPDU) error {
		counter++
		return nil
	})

	assert.NoError(t, err)
	assert.Equal(t, counter, 11)

	// without leading dot
	counter = 0
	err = h.Walk("1.3.6.1.2.1.2.2.1.1", func(pdu gosnmp.SnmpPDU) error {
		counter++
		return nil
	})

	assert.NoError(t, err)
	assert.Equal(t, counter, 11)
}

func TestFileHandler_Get(t *testing.T) {
	h, err := NewFileHandlerFromFile("testdata/if-mib.txt")
	assert.NoError(t, err)

	oid := ".1.3.6.1.2.1.2.2.1.3.1"

	p, err := h.Get([]string{oid})
	assert.NoError(t, err)

	assert.Len(t, p.Variables, 1)

	pdu := p.Variables[0]
	assert.Equal(t, oid, pdu.Name)
	assert.Equal(t, gosnmp.Integer, pdu.Type)
	assert.Equal(t, 24, pdu.Value)

	// without leading dot
	oid = "1.3.6.1.2.1.2.2.1.3.1"

	p, err = h.Get([]string{oid})
	assert.NoError(t, err)

	assert.Len(t, p.Variables, 1)
}
