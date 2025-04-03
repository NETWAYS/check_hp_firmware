package snmp

import (
	"fmt"

	"github.com/gosnmp/gosnmp"
)

func IsOid(oid string) bool {
	if oid == "" || oid[:1] != "." || len(oid) < 2 {
		return false
	}

	lastChar := rune(oid[0])

	for _, char := range oid[1:] {
		if char == '.' {
			if lastChar == '.' {
				return false
			}
		} else if !(char >= '0' && char <= '9') { //nolint:staticcheck
			return false
		}

		lastChar = char
	}

	return lastChar != '.'
}

// IsOidPartOf tests if an OID is equal of below another OID
func IsOidPartOf(oid string, baseOid string) bool {
	if !IsOid(oid) || !IsOid(baseOid) {
		return false
	}

	lenBase := len(baseOid)
	if oid[:lenBase] == baseOid {
		if len(oid) == lenBase || oid[lenBase:lenBase+1] == "." {
			return true
		}
	}

	return false
}

func GetSubOid(oid string, baseOid string) string {
	if !IsOid(oid) || !IsOid(baseOid) || !IsOidPartOf(oid, baseOid) {
		return ""
	}

	l := len(baseOid)

	return oid[l+1:]
}

func VersionFromString(version string) (v gosnmp.SnmpVersion, err error) {
	switch version {
	case "1":
		v = gosnmp.Version2c
	case "2", "2c":
		v = gosnmp.Version2c
	case "3":
		v = gosnmp.Version3
		// TODO: support v3?
		err = fmt.Errorf("SNMPv3 config not implemented")
	default:
		err = fmt.Errorf("unknown SNMP version: %s", version)
	}

	return
}
