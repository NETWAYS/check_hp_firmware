package snmp

import (
	"fmt"
	"io"
	"os"

	"github.com/gosnmp/gosnmp"
)

// Provides a GoSNMP like data interface, but with data from a snmpwalk output
//
// Anyone can generate an output by running:
//
//	snmpwalk -c public -v2c -On HOST .1.3.6.1 >snmp-data.txt
//
// Warning: This does not implement all functions of gosnmp.Handler
type FileHandler struct {
	Data WalkData

	gosnmp.Handler
}

// Read data from a io.Reader and parse it for PDUs
//
// They will be stored in Data for later use.
func (h *FileHandler) ReadFromWalk(r io.Reader) error {
	if h.Data == nil {
		h.Data = WalkData{}
	}

	pduList, err := ReadWalk(r)
	if err != nil {
		return err
	}

	for k, v := range pduList {
		h.Data[k] = v
	}

	return nil
}

// Simulate Connect behavior by returning no error
func (h *FileHandler) Connect() error {
	return nil
}

// Simulate ConnectIPv4 behavior by returning no error
func (h *FileHandler) ConnectIPv4() error {
	return nil
}

// Simulate ConnectIPv6 behavior by returning no error
func (h *FileHandler) ConnectIPv6() error {
	return nil
}

// Simulate Close behavior by returning no error
func (h *FileHandler) Close() error {
	return nil
}

// Simulating Get() behavior by searching read in data
func (h *FileHandler) Get(oids []string) (result *gosnmp.SnmpPacket, err error) {
	result = &gosnmp.SnmpPacket{
		Version: gosnmp.Version2c,
	}

	for _, oid := range oids {
		oid, err = EnsureValidOid(oid)
		if err != nil {
			return
		}

		if pdu, ok := h.Data[oid]; ok {
			result.Variables = append(result.Variables, *pdu)
		}
	}

	return
}

// Not yet implemented
func (h *FileHandler) GetBulk(oids []string, nonRepeaters uint8,
	maxRepetitions uint32) (result *gosnmp.SnmpPacket, err error) {
	panic("not implemented")
}

// Not yet implemented
func (h *FileHandler) GetNext(oids []string) (result *gosnmp.SnmpPacket, err error) {
	panic("not implemented")
}

// Simulating Walk() behavior by searching read in data
func (h *FileHandler) Walk(rootOid string, walkFn gosnmp.WalkFunc) (err error) {
	rootOid, err = EnsureValidOid(rootOid)
	if err != nil {
		return
	}

	for oid, pdu := range h.Data {
		if !IsOidPartOf(oid, rootOid) {
			continue
		}

		err = walkFn(*pdu)
		if err != nil {
			return
		}
	}

	return
}

// Not yet implemented
func (h *FileHandler) WalkAll(rootOid string) (results []gosnmp.SnmpPDU, err error) {
	panic("not implemented")
}

// Not yet implemented
func (h *FileHandler) BulkWalk(rootOid string, walkFn gosnmp.WalkFunc) error {
	panic("not implemented")
}

// Not yet implemented
func (h *FileHandler) BulkWalkAll(rootOid string) (results []gosnmp.SnmpPDU, err error) {
	panic("not implemented")
}

// Create a new file handler and initialize it with data from an io.Reader
func NewFileHandler(r io.Reader) (h *FileHandler, err error) {
	h = &FileHandler{}
	err = h.ReadFromWalk(r)

	return
}

// Create a new file handler and initialize it with data by reading from a file
func NewFileHandlerFromFile(filePath string) (h *FileHandler, err error) {
	fh, err := os.Open(filePath)
	if err != nil {
		err = fmt.Errorf("could not open SNMP data file: %s - %w", filePath, err)
		return
	}

	defer fh.Close()

	return NewFileHandler(fh)
}
