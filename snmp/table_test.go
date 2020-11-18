package snmp

import (
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func init() {
	if os.Getenv("SNMP_DEBUG") != "" {
		log.SetLevel(log.DebugLevel)
	}
}

func TestSnmpTable_Walk(t *testing.T) {
	if os.Getenv("NETWORK_TESTS_ENABLED") == "" {
		return
	}

	client := getSnmpClientFromEnv(t)

	table := &Table{
		Client: client,
		Oid:    ".1.3.6.1.2.1.2.2", // IF-MIB::ifTable
	}

	// TODO
	assert.NoError(t, table.Walk())
}

func TestSortOIDS(t *testing.T) {
	assert.Equal(t,
		[]string{"1.2.3", "4.5.6"},
		SortOIDs([]string{"4.5.6", "1.2.3"}))

	assert.Equal(t,
		[]string{"1.2.3", "1.2.10", "1.2.14"},
		SortOIDs([]string{"1.2.14", "1.2.10", "1.2.3"}))

	assert.Equal(t,
		[]string{"1.2.3.4.5.6", "1.2.10", "1.2.14"},
		SortOIDs([]string{"1.2.14", "1.2.10", "1.2.3.4.5.6"}))
}
