package cntlr

import (
	"github.com/NETWAYS/check_hp_firmware/hp/mib"
	"github.com/NETWAYS/check_hp_firmware/snmp"
	"github.com/gosnmp/gosnmp"
	"io"
)

type CpqDaCntlrTable struct {
	Snmp *snmp.Table
}

func GetCpqDaCntlrTable(client gosnmp.Handler) (*CpqDaCntlrTable, error) {
	table := CpqDaCntlrTable{}
	table.Snmp = &snmp.Table{
		Client: client,
		Oid:    mib.CpqDaCntlrTable,
	}

	return &table, table.Snmp.Walk()
}

func LoadCpqDaCntlrTable(stream io.Reader) (*CpqDaCntlrTable, error) {
	table := CpqDaCntlrTable{}

	snmpTable, err := snmp.LoadTableFromWalkOutput(mib.CpqDaCntlrTable, stream)
	if err != nil {
		return nil, err
	}

	table.Snmp = snmpTable

	return &table, nil
}

func (t *CpqDaCntlrTable) ListIds() []string {
	return t.Snmp.GetSortedOIDs()
}

func (t *CpqDaCntlrTable) GetValue(id string, oid string) (interface{}, error) {
	return t.Snmp.GetValue(id, oid)
}

func (t *CpqDaCntlrTable) GetStringValue(id string, oid string) (string, error) {
	return t.Snmp.GetStringValue(id, oid)
}

func (t *CpqDaCntlrTable) GetUintValue(id string, oid string) (uint, error) {
	return t.Snmp.GetUintValue(id, oid)
}

func (t *CpqDaCntlrTable) GetIntValue(id string, oid string) (int, error) {
	return t.Snmp.GetIntValue(id, oid)
}
