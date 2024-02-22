package snmp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testWalkLength = 242

func TestReadWalk(t *testing.T) {
	h, err := NewFileHandlerFromFile("testdata/if-mib.txt")
	assert.NoError(t, err)
	assert.Equal(t, testWalkLength, len(h.Data))
}

func TestIsValidOid(t *testing.T) {
	assert.NoError(t, IsValidOid(".1.2.3.4.6.999999"))
	assert.Error(t, IsValidOid(""))
	assert.Error(t, IsValidOid("1.2.3.4"))
	assert.Error(t, IsValidOid(".a.b.c.d"))
}

func TestEnsureValidOid(t *testing.T) {
	oid, err := EnsureValidOid("1.2.3.4.6.999999")
	assert.NoError(t, err)
	assert.Equal(t, ".1.2.3.4.6.999999", oid)

	oid, err = EnsureValidOid(".1.2.3.4.6.999999")
	assert.NoError(t, err)
	assert.Equal(t, ".1.2.3.4.6.999999", oid)
}
