package snmp

import (
	"os"
	"testing"
)

const fixtureHp = "../testdata/cpqDaPhyDrvTable.txt"
const tableOidHp = ".1.3.6.1.4.1.232.3.2.5.1"

func TestLoadFromWalkOutput(t *testing.T) {
	file, err := os.Open(fixtureHp)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	table, err := LoadTableFromWalkOutput(tableOidHp, file)
	if err != nil {
		t.Fatal(err)
	}

	if len(table.Values) == 0 {
		t.Fatal("Rows are empty!")
	}
}
