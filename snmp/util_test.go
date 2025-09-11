package snmp

import (
	"os"
	"testing"

	"github.com/gosnmp/gosnmp"
)

func TestIsOid(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", false},
		{".x", false},
		{".1.2.3.x", false},
		{".1.2.3..4", false},
		{".1.2.3.4.", false},
		{".1.2.3.4", true},
		{".1.2.311.44", true},
	}

	for _, test := range tests {
		result := IsOid(test.input)
		if result != test.expected {
			t.Errorf("IsOid(%q) = %v, want %v", test.input, result, test.expected)
		}
	}
}

func TestIsOidPartOf(t *testing.T) {
	tests := []struct {
		oid      string
		prefix   string
		expected bool
	}{
		{"", "", false},
		{".1.2.3.4", ".1.1", false},
		{".1.2", ".1", true},
		{".1.2", ".1.2", true},
	}

	for _, test := range tests {
		result := IsOidPartOf(test.oid, test.prefix)
		if result != test.expected {
			t.Errorf("IsOidPartOf(%q, %q) = %v, want %v", test.oid, test.prefix, result, test.expected)
		}
	}
}

func TestGetSubOid(t *testing.T) {
	tests := []struct {
		oid      string
		prefix   string
		expected string
	}{
		{".1.2.3.4", ".1.2.5", ""},
		{".1.2.3.4", ".1.2.3", "4"},
		{".1.2.3.4.5.6.7", ".1.2.3", "4.5.6.7"},
	}

	for _, test := range tests {
		result := GetSubOid(test.oid, test.prefix)
		if result != test.expected {
			t.Errorf("GetSubOid(%q, %q) = %q, want %q", test.oid, test.prefix, result, test.expected)
		}
	}
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
