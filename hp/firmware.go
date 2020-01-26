package hp

import (
	"fmt"
	"strconv"
)

// Source https://support.hpe.com/hpsc/doc/public/display?docId=emr_na-a00092491en_us
// Version: 2
// Effective date: 2019-12-12

const FirmwarePrefix = "HPD"
const FirmwareFixed = 8

//AffectedModels list of model numbers for drives that are affected with their description as value
var AffectedModels = map[string]string{
	"VO0480JFDGT":   "HP 480 GB 12 Gbit SAS 2.5\" RI PLP SC SSD",
	"VO0960JFDGU":   "HP 960GB 12 Gbit SAS 2.5\" RI PLP SC SSD",
	"VO1920JFDGV":   "HP 1,92 TB 12 Gbit SAS 2.5\" RI PLP SC SSD",
	"VO3840JFDHA":   "HP 3,84 TB 12 Gbit SAS 2.5\" RI PLP SC SSD",
	"MO0400JFFCF":   "HP 400 GB 12 Gbit SAS 2.5\" MU PLP SC SSD S2",
	"MO0800JFFCH":   "HP 800 GB 12 Gbit SAS 2.5\" MU PLP SC SSD S2",
	"MO1600JFFCK":   "HP 1,6 TB 12 Gbit SAS 2.5\" MU PLP SC SSD S2",
	"MO3200JFFCL":   "HP 3,2 TB 12 Gbit SAS 2.5\" MU PLP SC SSD S2",
	"VO000480JWDAR": "HPE 480 GB SAS SFF RI SC DS SSD",
	"VO000960JWDAT": "HPE 960 GB SAS SFF RI SC DS SSD",
	"VO001920JWDAU": "HPE 1,92 TB SAS RI SFF SC DS SSD",
	"VO003840JWDAV": "HPE 3,84 TB SAS RI SFF SC DS SSD",
	"VO007680JWCNK": "HPE 7,68 TB SAS 12G RI SFF SC DS SSD",
	"VO015300JWCNL": "HPE 15,3 TB SAS 12G RI SFF SC DS SSD",
	"VK000960JWSSQ": "HPE 960 GB SAS RI SFF SC DS SSD",
	"VK001920JWSSR": "HPE 1,92 TB SAS RI SFF SC DS SSD",
	"VK003840JWSST": "HPE 3,84 TB SAS RI SFF SC DS SSD",
	// duplicate
	//"VK003840JWSST": "HPE 3,84 TB SAS RI LFF SCC DS SPL SSD",
	"VK007680JWSSU": "HPE 7,68 TB SAS RI SFF SC DS SSD",
	"VO015300JWSSV": "HPE 15,3 TB SAS RI SFF SC DS SSD",
}

func IsFirmwareFixed(firmware string) (bool, error) {
	if len(firmware) < 4 || firmware[:3] != FirmwarePrefix {
		return false, fmt.Errorf("can not check firmware %s - only built for %s pattern",
			firmware, FirmwarePrefix)
	}

	version, err := strconv.ParseUint(firmware[3:], 10, 64)
	if err != nil {
		return false, fmt.Errorf("can not parse version number from: %s", firmware)
	}

	return version >= FirmwareFixed, nil
}
