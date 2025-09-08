package snmp

import (
	"os"
	"reflect"
	"testing"
)

func TestSnmpTable_Walk(t *testing.T) {
	if os.Getenv("NETWORK_TESTS_ENABLED") == "" {
		t.Skip("NETWORK_TESTS_ENABLED not set")
	}

	client := getSnmpClientFromEnv(t)

	table := &Table{
		Client: client,
		Oid:    ".1.3.6.1.2.1.2.2", // IF-MIB::ifTable
	}

	err := table.Walk()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// assert.NotEmpty(t, table.Columns, "expect at least the default columns")
	// assert.NotEmpty(t, table.Values, "expect at least some rows, even a basic container should have localhost")
}

func TestSortOIDs(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{
			input:    []string{"4.5.6", "1.2.3"},
			expected: []string{"1.2.3", "4.5.6"},
		},
		{
			input:    []string{"1.2.14", "1.2.10", "1.2.3"},
			expected: []string{"1.2.3", "1.2.10", "1.2.14"},
		},
		{
			input:    []string{"1.2.14", "1.2.10", "1.2.3.4.5.6"},
			expected: []string{"1.2.3.4.5.6", "1.2.10", "1.2.14"},
		},
	}

	for _, test := range tests {
		result := SortOIDs(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("SortOIDs(%v) = %v, want %v", test.input, result, test.expected)
		}
	}
}
