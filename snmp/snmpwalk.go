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

var (
	commentedOrEmptyLine = regexp.MustCompile(`^\s*(//|#|$)`)
	timeTickValue        = regexp.MustCompile(`^\(\d+\)`)
)

type WalkData map[string]*gosnmp.SnmpPDU

func ReadWalk(r io.Reader) (pduList WalkData, err error) {
	var (
		s     = bufio.NewScanner(r)
		pdu   *gosnmp.SnmpPDU
		lines uint
	)

	pduList = WalkData{}

	for s.Scan() {
		line := s.Text()
		lines++

		if commentedOrEmptyLine.MatchString(line) {
			continue
		}

		pdu, err = ParseWalkLine(line)
		if err != nil {
			err = fmt.Errorf("could not parse line %d: %w", lines, err)
			return
		}

		if pdu != nil && pdu.Name != "" {
			pduList[pdu.Name] = pdu
		}
	}

	return
}

func ParseWalkLine(line string) (pdu *gosnmp.SnmpPDU, err error) {
	parts := strings.SplitN(line, " = ", 2)

	if len(parts) != 2 || !IsOid(parts[0]) {
		// TODO: This can be the case for wrapped Hex-STRING lines, we are ignoring it for now...
		//err = fmt.Errorf("not a key = value line")
		return
	}

	if len(parts[1]) > 17 && parts[1][:17] == "No more variables" {
		// end of walk
		return
	}

	pdu = &gosnmp.SnmpPDU{
		Name: parts[0],
	}

	if parts[1] == `""` {
		// snmpwalk just lists "", which basically means null
		pdu.Type = gosnmp.Null
		return
	}

	split := strings.SplitN(parts[1], ": ", 2)
	if len(split) != 2 {
		err = fmt.Errorf("could not split type from value: %v", parts[1])
		return
	}

	netSnmpType, bareValue := split[0], split[1]

	// TODO: compare list with net-snmp source code!
	switch netSnmpType {
	case "OID":
		pdu.Type = gosnmp.ObjectIdentifier
		pdu.Value = bareValue
	case "INTEGER":
		pdu.Type = gosnmp.Integer

		i64, err := strconv.ParseInt(bareValue, 10, 32)
		if err == nil {
			pdu.Value = int(i64)
		}
	case "STRING":
		pdu.Type = gosnmp.OctetString
		pdu.Value = strings.Trim(bareValue, `"`)
	case "Hex-STRING":
		var bytes []byte

		hexString := strings.Join(strings.Split(bareValue, " "), "")
		bytes, err = hex.DecodeString(hexString)
		pdu.Value = string(bytes)
	case "Counter32":
		pdu.Type = gosnmp.Counter32

		i64, err := strconv.ParseUint(bareValue, 10, 32)
		if err == nil {
			pdu.Value = uint(i64)
		}
	case "Counter64":
		pdu.Type = gosnmp.Counter64
		pdu.Value, err = strconv.ParseUint(bareValue, 10, 64)
	case "Gauge32":
		pdu.Type = gosnmp.Gauge32

		i64, err := strconv.ParseUint(bareValue, 10, 32)
		if err == nil {
			pdu.Value = uint(i64)
		}
	case "Timeticks":
		pdu.Type = gosnmp.TimeTicks
		pdu.Value = timeTickValue.FindString(bareValue)
	case "NULL":
		pdu.Type = gosnmp.Null
		pdu.Value = []byte{}
	case "IpAddress":
		pdu.Type = gosnmp.IPAddress
		pdu.Value = bareValue
	default:
		err = fmt.Errorf("can not parse net-mib type %s of oid %s", netSnmpType, pdu.Name)
		return
	}

	if err != nil {
		err = fmt.Errorf("could not parse %s from oid %s: %w", netSnmpType, pdu.Name, err)
		return
	}

	return
}
