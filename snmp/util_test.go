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

func getSnmpClientFromEnv(t *testing.T) *gosnmp.GoSNMP {
	d := *gosnmp.Default
	c := d
	c.Target = getEnvDefault("SNMP_HOST", "localhost")
	c.Community = getEnvDefault("SNMP_COMMUNITY", "public")

	if err := SetVersion(&c, getEnvDefault("SNMP_VERSION", "2")); err != nil {
		t.Fatal(err)
	}

	c.Retries = 1

	if err := c.Connect(); err != nil {
		t.Fatal(err)
	}

	return &c
}
