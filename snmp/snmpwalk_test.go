package snmp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const testWalkLength = 242

func TestReadWalk(t *testing.T) {
	h, err := NewFileHandlerFromFile("testdata/if-mib.txt")
	assert.NoError(t, err)
	assert.Equal(t, testWalkLength, len(h.Data))
}
