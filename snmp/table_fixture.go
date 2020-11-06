package snmp

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"github.com/gosnmp/gosnmp"
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

	if len(parts) != 2 || !IsOid(parts[0]) {
		// unknown line
		return nil
	}

	oid := parts[0]

	// check if we are interested in the oid
	subOid := GetSubOid(oid, t.Oid)
	if subOid == "" {
		// other data in walk, ignoring it
		return nil
	}

	var netSnmpType string
	var bareValue string

	if parts[1] == "\"\"" {
		// snmpwalk just lists "", which basically means null
		netSnmpType = "NULL"
	} else if len(parts[1]) > 17 && parts[1][:17] == "No more variables" {
		// end of walk
		return nil
	} else {
		split := strings.SplitN(parts[1], ": ", 2)
		if len(split) != 2 {
			return fmt.Errorf("could not split type from value: %v", parts[1])
		}

		netSnmpType = split[0]
		bareValue = split[1]
	}

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
		i64, err := strconv.ParseInt(bareValue, 10, 32)
		if err == nil {
			value = int(i64)
		}
	case "STRING":
		snmpType = gosnmp.OctetString
		value = strings.Trim(bareValue, "\"")
	case "Hex-STRING":
		hexString := strings.Join(strings.Split(bareValue, " "), "")
		var bytes []byte
		bytes, err = hex.DecodeString(hexString)
		value = string(bytes)
	case "Counter32":
		snmpType = gosnmp.Counter32
		i64, err := strconv.ParseUint(bareValue, 10, 32)
		if err == nil {
			value = uint(i64)
		}
	case "Counter64":
		snmpType = gosnmp.Counter64
		value, err = strconv.ParseUint(bareValue, 10, 64)
	case "Gauge32":
		snmpType = gosnmp.Gauge32
		i64, err := strconv.ParseUint(bareValue, 10, 32)
		if err == nil {
			value = uint(i64)
		}
	case "Timeticks":
		snmpType = gosnmp.TimeTicks
		re := regexp.MustCompile(`^\(\d+\)`)
		value = re.FindString(bareValue)
	case "NULL":
		snmpType = gosnmp.Null
		value = []byte{}
	case "IpAddress":
		snmpType = gosnmp.IPAddress
		value = bareValue
	default:
		return fmt.Errorf("can not parse net-mib type %s of oid %s", netSnmpType, oid)
	}

	if err != nil {
		return fmt.Errorf("could not parse %s from oid %s: %s", netSnmpType, oid, err)
	}

	data := gosnmp.SnmpPDU{
		Name:  oid,
		Type:  snmpType,
		Value: value,
	}

	return t.addWalkValue(data)
}
