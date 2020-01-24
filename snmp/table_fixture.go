package snmp

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"github.com/soniah/gosnmp"
	"io"
	"regexp"
	"strconv"
	"strings"
)

func LoadTableFromWalkOutput(oid string, stream io.Reader) (*Table, error) {
	t := Table{}
	t.Oid = oid

	t.Reset()

	scanner := bufio.NewScanner(stream)
	for scanner.Scan() {
		err := t.addSnmpWalkLine(scanner.Text())
		if err != nil {
			return nil, err
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &t, nil
}

func (t *Table) addSnmpWalkLine(line string) error {
	parts := strings.SplitN(line, " = ", 2)

	if len(parts) != 2 || ! IsOid(parts[0]) {
		// unknown line
		return nil
	}

	oid := parts[0]
	split := strings.SplitN(parts[1], ": ", 2)
	if len(split) != 2 {
		return fmt.Errorf("could not split type from value: %v", parts[1])
	}

	netSnmpType := split[0]
	bareValue := split[1]

	var snmpType gosnmp.Asn1BER
	var value interface{}
	var err error

	// TODO: compare list with net-snmp source code!
	switch netSnmpType {
	case "OID":
		snmpType = gosnmp.ObjectIdentifier
		value = bareValue
	case "INTEGER":
		snmpType = gosnmp.Integer
		value, err = strconv.ParseInt(bareValue, 10, 64)
	case "STRING":
		snmpType = gosnmp.OctetString
		value = strings.Trim(bareValue, "\"")
	case "Hex-STRING":
		value, err = hex.DecodeString(bareValue)
		if err != nil {}
	case "Counter32":
		snmpType = gosnmp.Counter32
		value, err = strconv.ParseUint(bareValue, 10, 32)
	case "Counter64":
		snmpType = gosnmp.Counter64
		value, err = strconv.ParseUint(bareValue, 10, 64)
	case "Gauge32":
		snmpType = gosnmp.Gauge32
		value, err = strconv.ParseUint(bareValue, 10, 32)
	case "Timeticks":
		snmpType = gosnmp.TimeTicks
		re := regexp.MustCompile(`^\(\d+\)`)
		value = re.FindString(bareValue)
	default:
		return fmt.Errorf("can not parse net-snmp type %s of oid %s", netSnmpType, oid)
	}

	if err != nil {
		return fmt.Errorf("could not parse %s from oid %s: %s", netSnmpType, oid, err)
	}

	// TODO: split type
	// TODO: convert value to proper byte

	data := gosnmp.SnmpPDU{
		Name:  oid,
		Type:  snmpType,
		Value: value,
	}

	return t.addWalkValue(data)
}
