package snmp

import (
	"github.com/gosnmp/gosnmp"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestIsOid(t *testing.T) {
	assert.False(t, IsOid(""))
	assert.False(t, IsOid(".x"))
	assert.False(t, IsOid(".1.2.3.x"))
	assert.False(t, IsOid(".1.2.3..4"))
	assert.False(t, IsOid(".1.2.3.4."))

	assert.True(t, IsOid(".1.2.3.4"))
	assert.True(t, IsOid(".1.2.311.44"))
}

func TestIsOidPartOf(t *testing.T) {
	assert.False(t, IsOidPartOf("", ""))
	assert.False(t, IsOidPartOf(".1.2.3.4", ".1.1"))
	assert.True(t, IsOidPartOf(".1.2", ".1"))
	assert.True(t, IsOidPartOf(".1.2", ".1.2"))
}

func TestGetSubOid(t *testing.T) {
	assert.Equal(t, GetSubOid(".1.2.3.4", ".1.2.5"), "")
	assert.Equal(t, GetSubOid(".1.2.3.4", ".1.2.3"), "4")
	assert.Equal(t, GetSubOid(".1.2.3.4.5.6.7", ".1.2.3"), "4.5.6.7")
}

func getEnvDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}

	return value
}

func getSnmpClientFromEnv(t *testing.T) gosnmp.Handler {
	h := gosnmp.NewHandler()

	h.SetTarget(getEnvDefault("SNMP_HOST", "localhost"))
	h.SetCommunity(getEnvDefault("SNMP_COMMUNITY", "public"))

	version, err := VersionFromString(getEnvDefault("SNMP_VERSION", "2"))
	if err != nil {
		t.Fatal(err)
	}

	h.SetVersion(version)
	h.SetRetries(1)

	if err := h.Connect(); err != nil {
		t.Fatal(err)
	}

	return h
}
