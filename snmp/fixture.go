package snmp

import (
	"github.com/soniah/gosnmp"
	"strconv"
)

type Fixture struct {
	PDUs []FixturePDU
}

type FixturePDU struct {
	Oid   string
	Type  string
	Value string
}

var netSnmpNameToType = map[string]gosnmp.Asn1BER{
	"OID": gosnmp.ObjectIdentifier,
	"INTEGER": gosnmp.Integer,
}

var netSnmpTypeToName = map[gosnmp.Asn1BER]string {
	
}

//func (*Fixture) Append(pdu gosnmp.SnmpPDU) error {
//
//}
//
//func MarshalPDU(pdu gosnmp.SnmpPDU) (string, error) {
//	return "", nil
//}
//
//func Unmarshal(json string) (*FixturePDU, error) {
//	return nil, nil
//}
//
//func TypeToString(pduType gosnmp.Asn1BER) string {
//
//}

func StringToType(s string) (gosnmp.Asn1BER, error) {
	var value gosnmp.Asn1BER
	switch s {
	case "OID":
		value = gosnmp.ObjectIdentifier
	case "INTEGER":
		value = gosnmp.Integer
	case "STRING":
		value = gosnmp.OctetString
	case "Timeticks":
		value = gosnmp.TimeTicks
	case "Hex-STRING":
	case "Gauge32":
		// TODO
	default:
		i, err := strconv.ParseUint(s, 0, 64)
		if err != nil {
			// TODO: more verbose error?
			return 0, err
		}
		value = gosnmp.Asn1BER(i)
	}

	return value, nil
}
