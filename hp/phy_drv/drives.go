package phy_drv

import (
	"github.com/NETWAYS/check_hp_firmware/hp/mib"
	"github.com/NETWAYS/check_hp_firmware/snmp"
	"github.com/gosnmp/gosnmp"
	"io"
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
	return t.Snmp.GetValue(id, oid)
}

func (t *CpqDaPhyDrvTable) GetStringValue(id string, oid string) (string, error) {
	return t.Snmp.GetStringValue(id, oid)
}

func (t *CpqDaPhyDrvTable) GetUintValue(id string, oid string) (uint, error) {
	return t.Snmp.GetUintValue(id, oid)
}

func (t *CpqDaPhyDrvTable) GetIntValue(id string, oid string) (int, error) {
	return t.Snmp.GetIntValue(id, oid)
}
