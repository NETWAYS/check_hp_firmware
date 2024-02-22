package drive

import (
	"github.com/NETWAYS/check_hp_firmware/hp/mib"
	"github.com/NETWAYS/check_hp_firmware/snmp"
	"github.com/gosnmp/gosnmp"
)

type CpqDaPhyDrvTable struct {
	Snmp *snmp.Table
}

func GetCpqDaPhyDrvTable(client gosnmp.Handler) (*CpqDaPhyDrvTable, error) {
	table := CpqDaPhyDrvTable{}
	table.Snmp = &snmp.Table{
		Client: client,
		Oid:    mib.CpqDaPhyDrvTable,
	}

	return &table, table.Snmp.Walk()
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
