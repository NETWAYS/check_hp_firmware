package main

import (
	"fmt"
	"github.com/NETWAYS/check_hp_firmware/hp/cntlr"
	"github.com/NETWAYS/check_hp_firmware/hp/ilo"
	"github.com/NETWAYS/check_hp_firmware/hp/phy_drv"
	"github.com/NETWAYS/check_hp_firmware/snmp"
	"github.com/NETWAYS/go-check"
	"github.com/NETWAYS/go-check/result"
	"github.com/gosnmp/gosnmp"
	"time"
)

const Readme = `
Icinga / Nagios check plugin to verify HPE controllers an SSD disks or ilo are not affected by certain vulnerabilities.

**HPE Controllers**

	HPE Smart Array SR Gen10 Controller Firmware Version 2.65 (or later) provided in the (HPE document a00097210) is
	required to prevent a potential data inconsistency on select RAID configurations with Smart Array Gen10 Firmware
	Version 1.98 through 2.62, based on the following scenarios. HPE strongly recommends performing this upgrade at the
	customer's earliest opportunity per the "Action Required" in the table located in the Resolution section.
	Neglecting to perform the recommended resolution could result in potential subsequent errors and potential data
	inconsistency.

The check will alert you with a CRITICAL when the firmware is in the affected range with:

* "if you have RAID 1/10/ADM - update immediately!"
* "if you have RAID 5/6/50/60 - update immediately!"

And it will add a short note when "firmware older than affected" or "firmware has been updated". At the moment the
plugin does not verify configured logical drives, but we believe you should update in any case.

**HPE SSD SAS disks**

	HPE SAS Solid State Drives - Critical Firmware Upgrade Required for Certain HPE SAS Solid State Drive Models to
	Prevent Drive Failure at 32,768 or 40,000 Hours of Operation

The check will raise a CRITICAL when the drive needs to be updated with the note "affected by FW bug", and when
the drive is patched with "firmware update applied".

**HPE Integrated Lights-Out**
	Multiple security vulnerabilities have been identified in Integrated Lights-Out 3 (iLO 3),
	Integrated Lights-Out 4 (iLO 4), and Integrated Lights-Out 5 (iLO 5) firmware. The vulnerabilities could be remotely
	exploited to execute code, cause denial of service, and expose sensitive information. HPE has released updated
	firmware to mitigate these vulnerabilities.

	The check will raise a CRITICAL when the Integrated Lights-Out needs to be updated. Below you will find a list with
	the least version of each Integrated Lights-Out version:
	 - HPE Integrated Lights-Out 3 (iLO 3) firmware v1.93 or later.
	 - HPE Integrated Lights-Out 4 (iLO 4) firmware v2.75 or later
	 - HPE Integrated Lights-Out 5 (iLO 5) firmware v2.18 or later.


Please see support documents from HPE:
* https://support.hpe.com/hpesc/public/docDisplay?docLocale=en_US&docId=emr_na-a00092491en_us
* https://support.hpe.com/hpesc/public/docDisplay?docLocale=en_US&docId=a00097382en_us
* https://support.hpe.com/hpesc/public/docDisplay?docLocale=en_US&docId=a00097210en_us
* https://support.hpe.com/hpesc/public/docDisplay?docId=hpesbhf04012en_us

**IMPORTANT:** Read the documentation for HPE! The plugin and its documentation is a best effort to find and detect
affected hardware. There is ABSOLUTELY NO WARRANTY, see the license!

**Note:** This plugin was initially named check_hp_disk_firmware
`

// Check for HP Controller CVEs via SNMP
func main() {
	config := check.NewConfig()
	config.Name = "check_hp_firmware"
	config.Readme = Readme
	config.Version = buildVersion()
	config.Timeout = 15

	var (
		fs        = config.FlagSet
		host      = fs.StringP("hostname", "H", "localhost", "SNMP host")
		community = fs.StringP("community", "c", "public", "SNMP community")
		protocol  = fs.StringP("protocol", "P", "2c", "SNMP protocol")
		file      = fs.String("snmpwalk-file", "", "Read output from snmpwalk")
		ignoreIlo = fs.Bool("ignore-ilo-version", false, "Don't check the ILO version")
		ipv4      = fs.BoolP("ipv4", "4", false, "Use IPv4")
		ipv6      = fs.BoolP("ipv6", "6", false, "Use IPv6")
	)

	_ = fs.BoolP("ilo", "I", false, "Checks the version of iLo")
	_ = fs.MarkHidden("ilo")

	config.ParseArguments()

	var (
		client     gosnmp.Handler
		cntlrTable *cntlr.CpqDaCntlrTable
		driveTable *phy_drv.CpqDaPhyDrvTable
	)

	var err error
	if *file != "" {
		client, err = snmp.NewFileHandlerFromFile(*file)
		if err != nil {
			check.ExitError(err)
		}
	} else {
		client = gosnmp.NewHandler()
		client.SetTarget(*host)
		client.SetCommunity(*community)
		client.SetTimeout(time.Duration(config.Timeout) - 1*time.Second)
		client.SetRetries(1)

		version, err := snmp.VersionFromString(*protocol)
		if err != nil {
			check.ExitError(err)
		}

		client.SetVersion(version)
	}

	// Initialize connection
	if *ipv4 {
		err = client.ConnectIPv4()
	} else if *ipv6 {
		err = client.ConnectIPv6()
	} else {
		err = client.Connect()
	}

	if err != nil {
		check.ExitError(err)
	}

	defer func() {
		_ = client.Close()
	}()

	// Load controller data
	cntlrTable, err = cntlr.GetCpqDaCntlrTable(client)
	if err != nil {
		check.ExitError(err)
	}

	// Load drive data
	driveTable, err = phy_drv.GetCpqDaPhyDrvTable(client)
	if err != nil {
		check.ExitError(err)
	}

	if len(cntlrTable.Snmp.Values) == 0 {
		check.Exit(3, "No HP controller data found!")
	}

	controllers, err := cntlr.GetControllersFromTable(cntlrTable)
	if err != nil {
		check.ExitError(err)
	}

	if len(driveTable.Snmp.Values) == 0 {
		check.Exit(3, "No HP drive data found!")
	}

	drives, err := phy_drv.GetPhysicalDrivesFromTable(driveTable)
	if err != nil {
		check.ExitError(err)
	}

	overall := result.Overall{}

	// check the ILO Version unless set
	if !*ignoreIlo {
		iloData, err := ilo.GetIloInformation(client)
		if err != nil {
			check.ExitError(err)
		}

		overall.Add(iloData.GetNagiosStatus())
	}

	countControllers := 0

	for _, controller := range controllers {
		controllerStatus, desc := controller.GetNagiosStatus()

		overall.Add(controllerStatus, desc)

		countControllers += 1
	}

	countDrives := 0

	for _, drive := range drives {
		driveStatus, desc := drive.GetNagiosStatus()

		overall.Add(driveStatus, desc)

		countDrives += 1
	}

	var summary string

	status := overall.GetStatus()

	switch status {
	case check.OK:
		summary = fmt.Sprintf("All %d controllers and %d drives seem fine", countControllers, countDrives)
	case check.Warning:
		summary = fmt.Sprintf("Found %d warnings", overall.GetStatus())
	case check.Critical:
		summary = fmt.Sprintf("Found %d critical problems", overall.GetStatus())
	}

	overall.Summary = summary
	check.Exit(status, overall.GetOutput())
}
