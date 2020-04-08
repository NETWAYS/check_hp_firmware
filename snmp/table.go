package snmp

import (
	"fmt"
	"github.com/mcuadros/go-version"
	log "github.com/sirupsen/logrus"
	"github.com/soniah/gosnmp"
	"sort"
	"strings"
)

type Table struct {
	Client  *gosnmp.GoSNMP
	Oid     string
	Columns IndexedIds
	Values  TableRows
}

type IndexedIds map[string]string
type TableRows map[string]TableColumns
type TableColumns map[string]gosnmp.SnmpPDU

func (t *Table) Reset() {
	t.Columns = IndexedIds{}
	t.Values = TableRows{}
}

func (t *Table) Walk() error {
	t.Reset()

	log.WithFields(log.Fields{
		"oid": t.Oid,
	}).Debug("Starting walk for table")

	err := t.Client.Walk(t.Oid, t.addWalkValue)
	if err != nil {
		return err
	}

	if len(t.Values) == 0 {
		return fmt.Errorf("no data retrieved in walk for table: %s", t.Oid)
	}

	log.WithFields(log.Fields{
		"tableValues": t.Values,
	}).Debug("read table data from SNMP walk")

	return nil
}

// addWalkValue parses the PDU and stored result in an indexed way
//
// The OID part below the table is something like:
// 1.X.Y.Y
//
// 1    entry OID, just a construct to represent the row
// X    value OID
// Y.Y  actual index part (can be a longer chain)
//
// TODO: this might not apply to all tables
func (t *Table) addWalkValue(data gosnmp.SnmpPDU) error {
	subOid := GetSubOid(data.Name, t.Oid)
	if subOid == "" {
		// other data in walk, ignoring it
		return nil
	}
	parts := strings.Split(subOid, ".")

	if len(parts) < 3 {
		return fmt.Errorf("could not identify entry, column and id in oid: %s", data.Name)
	}

	entry := parts[0]
	column := parts[1]
	id := strings.Join(parts[2:], ".")

	log.WithFields(log.Fields{
		"oid":    data.Name,
		"entry":  entry,
		"column": column,
		"id":     id,
	}).Debug("Reading PDU data")

	if _, ok := t.Values[id]; !ok {
		t.Values[id] = TableColumns{}
	}

	// store data in indexed tree
	t.Values[id][column] = data

	// keep list of existing columns
	if _, ok := t.Columns[column]; !ok {
		t.Columns[column] = column
	}

	return nil
}

func (t *Table) GetSortedOIDs() []string {
	var ids []string
	for k := range t.Values {
		ids = append(ids, k)
	}

	return SortOIDs(ids)
}

func SortOIDs(list []string) []string {
	sort.Slice(list, func(i, j int) bool {
		return version.Compare(list[i], list[j], "<")
	})

	return list
}
