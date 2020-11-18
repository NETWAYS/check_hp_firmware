package snmp

import (
	"fmt"
	"github.com/gosnmp/gosnmp"
	"github.com/mcuadros/go-version"
	log "github.com/sirupsen/logrus"
	"sort"
	"strings"
)

type Table struct {
	Client  gosnmp.Handler
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

	if err := t.Client.Walk(t.Oid, t.addWalkValue); err != nil {
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

func (t *Table) GetValue(id string, oid string) (interface{}, error) {
	parts := strings.Split(oid, ".")
	column := parts[len(parts)-1]

	drive, ok := t.Values[id]
	if !ok {
		return nil, fmt.Errorf("could not find row %s while looking for column %s", id, column)
	}

	value, ok := drive[column]
	if !ok {
		return nil, fmt.Errorf("could not find column %s for row %s", column, id)
	}

	return value.Value, nil
}

func (t *Table) GetStringValue(id string, oid string) (string, error) {
	value, err := t.GetValue(id, oid)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", value), nil
}

func (t *Table) GetUintValue(id string, oid string) (uint, error) {
	value, err := t.GetValue(id, oid)
	if err != nil {
		return 0, err
	}

	return value.(uint), nil
}

func (t *Table) GetIntValue(id string, oid string) (int, error) {
	value, err := t.GetValue(id, oid)
	if err != nil {
		return 0, err
	}

	return value.(int), nil
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
