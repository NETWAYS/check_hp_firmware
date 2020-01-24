package snmp

import (
	"github.com/soniah/gosnmp"
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

	switch getEnvDefault("SNMP_VERSION", "2") {
	case "1":
		c.Version = gosnmp.Version1
	case "2":
		c.Version = gosnmp.Version2c
	case "3":
		c.Version = gosnmp.Version3
		// TODO: support v3?
		t.Fatal("SNMPv3 config not implemented")
	}

	c.Retries = 1

	if err := c.Connect(); err != nil {
		t.Fatal(err)
	}

	return &c
}
