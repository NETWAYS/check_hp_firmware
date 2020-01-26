package hp

import (
	"fmt"
	"github.com/NETWAYS/check_hp_cve/snmp"
	"github.com/soniah/gosnmp"
	"io"
	"sort"
	"strings"
)

type CpqDaPhyDrvTable struct {
	Snmp *snmp.Table
}

func GetCpqDaPhyDrvTable(client *gosnmp.GoSNMP) (*CpqDaPhyDrvTable, error) {
	table := CpqDaPhyDrvTable{}
	table.Snmp = &snmp.Table{
		Client: client,
		Oid:    SnmpCpqDaPhyDrvTable,
	}

	return &table, table.Snmp.Walk()
}

func LoadCpqDaPhyDrvTable(stream io.Reader) (*CpqDaPhyDrvTable, error) {
	table := CpqDaPhyDrvTable{}
	snmpTable, err := snmp.LoadTableFromWalkOutput(SnmpCpqDaPhyDrvTable, stream)
	if err != nil {
		return nil, err
	}
	table.Snmp = snmpTable

	return &table, nil
}

func (t *CpqDaPhyDrvTable) ListIds() []string {
	ids := make([]string, 0, len(t.Snmp.Values))
	for k := range t.Snmp.Values {
		ids = append(ids, k)
	}

	// TODO: sort numerically
	sort.Strings(ids)

	return ids
}

func (t *CpqDaPhyDrvTable) GetValue(id string, oid string) (interface{}, error) {
	parts := strings.Split(oid, ".")
	column := parts[len(parts)-1]

	drive, ok := t.Snmp.Values[id]
	if ! ok {
		return nil, fmt.Errorf("could not find drive %s while looking for column %s", id, column)
	}

	value, ok := drive[column]
	if ! ok {
		return nil, fmt.Errorf("could not find column %s for drive %s", column, id)
	}

	return value.Value, nil
}

func (t *CpqDaPhyDrvTable) GetStringValue(id string, oid string) (string, error) {
	value, err := t.GetValue(id, oid)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", value), nil
}

func (t *CpqDaPhyDrvTable) GetUintValue(id string, oid string) (uint64, error) {
	value, err := t.GetValue(id, oid)
	if err != nil {
		return 0, err
	}

	return value.(uint64), nil
}

func (t *CpqDaPhyDrvTable) GetIntValue(id string, oid string) (int64, error) {
	value, err := t.GetValue(id, oid)
	if err != nil {
		return 0, err
	}

	return value.(int64), nil
}
