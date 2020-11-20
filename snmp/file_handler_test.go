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
}
