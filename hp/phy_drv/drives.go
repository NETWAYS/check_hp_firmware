package phy_drv

import (
	"fmt"
	"github.com/NETWAYS/check_hp_disk_firmware/hp/mib"
	"github.com/NETWAYS/check_hp_disk_firmware/snmp"
	"github.com/soniah/gosnmp"
	"io"
	"strings"
)

type CpqDaPhyDrvTable struct {
	Snmp *snmp.Table
}

func GetCpqDaPhyDrvTable(client *gosnmp.GoSNMP) (*CpqDaPhyDrvTable, error) {
	table := CpqDaPhyDrvTable{}
	table.Snmp = &snmp.Table{
		Client: client,
		Oid:    mib.CpqDaPhyDrvTable,
	}

	return &table, table.Snmp.Walk()
}

func LoadCpqDaPhyDrvTable(stream io.Reader) (*CpqDaPhyDrvTable, error) {
	table := CpqDaPhyDrvTable{}
	snmpTable, err := snmp.LoadTableFromWalkOutput(mib.CpqDaPhyDrvTable, stream)
	if err != nil {
		return nil, err
	}
	table.Snmp = snmpTable

	return &table, nil
}

func (t *CpqDaPhyDrvTable) ListIds() []string {
	return t.Snmp.GetSortedOIDs()
}

func (t *CpqDaPhyDrvTable) GetValue(id string, oid string) (interface{}, error) {
	parts := strings.Split(oid, ".")
	column := parts[len(parts)-1]

	drive, ok := t.Snmp.Values[id]
	if !ok {
		return nil, fmt.Errorf("could not find drive %s while looking for column %s", id, column)
	}

	value, ok := drive[column]
	if !ok {
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

func (t *CpqDaPhyDrvTable) GetUintValue(id string, oid string) (uint, error) {
	value, err := t.GetValue(id, oid)
	if err != nil {
		return 0, err
	}

	return value.(uint), nil
}

func (t *CpqDaPhyDrvTable) GetIntValue(id string, oid string) (int, error) {
	value, err := t.GetValue(id, oid)
	if err != nil {
		return 0, err
	}

	return value.(int), nil
}
