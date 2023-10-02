package phy_drv

// Source A: https://support.hpe.com/hpsc/doc/public/display?docId=emr_na-a00092491en_us
// Version: 2
// Effective date: 2019-12-12

const FixedA = "HPD8"

// Source B: https://support.hpe.com/hpesc/public/docDisplay?docLocale=en_US&docId=a00097382en_us
// Version: 1
// Last Updated: 2020-03-20

const FixedB = "HPD7"

// AffectedModelList list of model numbers for drives that are affected with their description as value
var AffectedModelList = []*AffectedModel{
	// Drives from bulletin A
	{"VO0480JFDGT", "HP 480 GB 12 Gbit SAS 2.5\" RI PLP SC SSD", FixedA},
	{"VO0960JFDGU", "HP 960GB 12 Gbit SAS 2.5\" RI PLP SC SSD", FixedA},
	{"VO1920JFDGV", "HP 1,92 TB 12 Gbit SAS 2.5\" RI PLP SC SSD", FixedA},
	{"VO3840JFDHA", "HP 3,84 TB 12 Gbit SAS 2.5\" RI PLP SC SSD", FixedA},
	{"MO0400JFFCF", "HP 400 GB 12 Gbit SAS 2.5\" MU PLP SC SSD S2", FixedA},
	{"MO0800JFFCH", "HP 800 GB 12 Gbit SAS 2.5\" MU PLP SC SSD S2", FixedA},
	{"MO1600JFFCK", "HP 1,6 TB 12 Gbit SAS 2.5\" MU PLP SC SSD S2", FixedA},
	{"MO3200JFFCL", "HP 3,2 TB 12 Gbit SAS 2.5\" MU PLP SC SSD S2", FixedA},
	{"VO000480JWDAR", "HPE 480 GB SAS SFF RI SC DS SSD", FixedA},
	{"VO000960JWDAT", "HPE 960 GB SAS SFF RI SC DS SSD", FixedA},
	{"VO001920JWDAU", "HPE 1,92 TB SAS RI SFF SC DS SSD", FixedA},
	{"VO003840JWDAV", "HPE 3,84 TB SAS RI SFF SC DS SSD", FixedA},
	{"VO007680JWCNK", "HPE 7,68 TB SAS 12G RI SFF SC DS SSD", FixedA},
	{"VO015300JWCNL", "HPE 15,3 TB SAS 12G RI SFF SC DS SSD", FixedA},
	{"VK000960JWSSQ", "HPE 960 GB SAS RI SFF SC DS SSD", FixedA},
	{"VK001920JWSSR", "HPE 1,92 TB SAS RI SFF SC DS SSD", FixedA},
	{"VK003840JWSST", "HPE 3,84 TB SAS RI SFF SC DS SSD", FixedA},
	// duplicate
	// {"VK003840JWSST", "HPE 3,84 TB SAS RI LFF SCC DS SPL SSD", FixedA},
	{"VK007680JWSSU", "HPE 7,68 TB SAS RI SFF SC DS SSD", FixedA},
	{"VO015300JWSSV", "HPE 15,3 TB SAS RI SFF SC DS SSD", FixedA},

	// Drives from bulletin B
	{"EK0800JVYPN", "HPE 800GB 12G SAS WI-1 SFF SC SSD", FixedB},
	{"EO1600JVYPP", "HPE 1.6TB 12G SAS WI-1 SFF SC SSD", FixedB},
	{"MK0800JVYPQ", "HPE 800GB 12G SAS MU-1 SFF SC SSD", FixedB},
	{"MO1600JVYPR", "HPE 1.6TB 12G SAS MU-1 SFF SC SSD", FixedB},
}
