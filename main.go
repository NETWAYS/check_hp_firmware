package main

import (
	"fmt"
	"github.com/NETWAYS/check_hp_firmware/hp/cntlr"
	"github.com/NETWAYS/check_hp_firmware/hp/phy_drv"
	"github.com/NETWAYS/check_hp_firmware/nagios"
	"github.com/NETWAYS/check_hp_firmware/snmp"
	log "github.com/sirupsen/logrus"
	"github.com/soniah/gosnmp"
	flag "github.com/spf13/pflag"
	"io"
	"os"
	"time"
)

const Readme = `
Icinga / Nagios check plugin to verify HPE controllers an SSD disks are not affected by certain vulnerabilities.

For controllers:

  HPE Smart Array SR Gen10 Controller Firmware Version 2.65 (or later) provided in the (HPE document a00097210) is
  required to prevent a potential data inconsistency on select RAID configurations with Smart Array Gen10 Firmware
  Version 1.98 through 2.62, based on the following scenarios. HPE strongly recommends performing this upgrade at the
  customer's earliest opportunity per the "Action Required" in the table located in the Resolution section.
  Neglecting to perform the recommended resolution could result in potential subsequent errors and potential data
  inconsistency.

For SSD disks:

  HPE SAS Solid State Drives - Critical Firmware Upgrade Required for Certain HPE SAS Solid State Drive Models to
  Prevent Drive Failure at 32,768 or 40,000 Hours of Operation

Please see support documents from HPE:
* https://support.hpe.com/hpesc/public/docDisplay?docLocale=en_US&docId=emr_na-a00092491en_us
* https://support.hpe.com/hpesc/public/docDisplay?docLocale=en_US&docId=a00097382en_us
* https://support.hpe.com/hpesc/public/docDisplay?docLocale=en_US&docId=a00097210en_us

**IMPORTANT:** Read the documentation for HPE! The plugin and its documentation is a best effort to find and detect
affected hardware. There is ABSOLUTELY NO WARRANTY, see the license!

**Note:** This plugin was initially named check_hp_disk_firmware
`

// Check for HP Controller CVEs via SNMP
func main() {
	flagSet := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	flagSet.SortFlags = false

	host := flagSet.StringP("hostname", "H", "localhost", "SNMP host")
	community := flagSet.StringP("community", "c", "public", "SNMP community")
	protocol := flagSet.StringP("protocol", "P", "2c", "SNMP protocol")
	timeout := flagSet.Int64("timeout", 15, "SNMP timeout in seconds")

	file := flagSet.String("snmpwalk-file", "", "Read output from snmpwalk")

	ipv4 := flagSet.BoolP("ipv4", "4", false, "Use IPv4")
	ipv6 := flagSet.BoolP("ipv6", "6", false, "Use IPv6")

	version := flagSet.BoolP("version", "V", false, "Show version")

	debug := flagSet.Bool("debug", false, "Enable debug output")

	flagSet.Usage = func() {
		fmt.Printf("Usage: %s [-H <hostname>] [-c <community>]\n", os.Args[0])
		fmt.Println(Readme)
		fmt.Printf("Version: %s\n", buildVersion())
		fmt.Println()
		fmt.Println("Arguments:")
		flagSet.PrintDefaults()
	}

	var err error
	err = flagSet.Parse(os.Args[1:])
	if err != nil {
		if err != flag.ErrHelp {
			nagios.ExitError(err)
		}
		os.Exit(3)
	}

	if *version {
		fmt.Printf("%s version %s\n", Project, buildVersion())
		os.Exit(nagios.Unknown)
	}

	if *debug {
		log.SetLevel(log.DebugLevel)
	} else {
		defer nagios.CatchPanic()
	}

	var client *gosnmp.GoSNMP
	var cntlrTable *cntlr.CpqDaCntlrTable
	var driveTable *phy_drv.CpqDaPhyDrvTable

	if *file != "" {
		var fh *os.File
		fh, err = os.Open(*file)
		if err != nil {
			nagios.ExitError(err)
		}
		defer fh.Close()

		cntlrTable, err = cntlr.LoadCpqDaCntlrTable(fh)
		if err != nil {
			nagios.ExitError(err)
		}

		// jump back to start
		_, err = fh.Seek(0, io.SeekStart)
		if err != nil {
			nagios.ExitError(err)
		}

		driveTable, err = phy_drv.LoadCpqDaPhyDrvTable(fh)
		if err != nil {
			nagios.ExitError(err)
		}
	} else {
		defaultClient := *gosnmp.Default
		client = &defaultClient
		client.Target = *host
		client.Community = *community
		client.Timeout = time.Duration(*timeout) * time.Second
		client.Retries = 1

		if err := snmp.SetVersion(client, *protocol); err != nil {
			nagios.ExitError(err)
		}

		if *ipv4 {
			err = client.ConnectIPv4()
		} else if *ipv6 {
			err = client.ConnectIPv6()
		} else {
			err = client.Connect()
		}
		if err != nil {
			nagios.ExitError(err)
		}
		defer client.Conn.Close()

		cntlrTable, err = cntlr.GetCpqDaCntlrTable(client)
		if err != nil {
			nagios.ExitError(err)
		}

		driveTable, err = phy_drv.GetCpqDaPhyDrvTable(client)
		if err != nil {
			nagios.ExitError(err)
		}
	}

	if len(cntlrTable.Snmp.Values) == 0 {
		nagios.Exit(3, "No HP controller data found!")
	}

	controllers, err := cntlr.GetControllersFromTable(cntlrTable)
	if err != nil {
		nagios.ExitError(err)
	}

	if len(driveTable.Snmp.Values) == 0 {
		nagios.Exit(3, "No HP drive data found!")
	}

	drives, err := phy_drv.GetPhysicalDrivesFromTable(driveTable)
	if err != nil {
		nagios.ExitError(err)
	}

	overall := nagios.Overall{}

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

	status := overall.GetStatus()
	var summary string
	switch status {
	case nagios.OK:
		summary = fmt.Sprintf("All %d controllers and %d drives seem fine", countControllers, countDrives)
	case nagios.Warning:
		summary = fmt.Sprintf("Found %d warnings", overall.Warnings)
	case nagios.Critical:
		summary = fmt.Sprintf("Found %d critical problems", overall.Criticals)
	}
	overall.Summary = summary
	nagios.Exit(status, overall.GetOutput())
}
