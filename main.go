package main

import (
	"fmt"
	"time"

	"github.com/NETWAYS/check_hp_firmware/hp/cntlr"
	"github.com/NETWAYS/check_hp_firmware/hp/ilo"
	"github.com/NETWAYS/check_hp_firmware/hp/phy_drv"
	"github.com/NETWAYS/check_hp_firmware/snmp"
	"github.com/NETWAYS/go-check"
	"github.com/NETWAYS/go-check/result"
	"github.com/gosnmp/gosnmp"
)

var (
	// These get filled at build time with the proper vaules
	version = "development"
	commit  = "HEAD"
	date    = "latest"
)

func buildVersion() string {
	result := version

	if commit != "" {
		result = fmt.Sprintf("%s\ncommit: %s", result, commit)
	}

	if date != "" {
		result = fmt.Sprintf("%s\ndate: %s", result, date)
	}

	return result
}

const Readme = "Monitoring plugin to verify HPE controllers an SSD disks or iLO are not affected by certain vulnerabilities."

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
	// nolint: gocritic
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
		check.ExitRaw(3, "No HP controller data found!")
	}

	controllers, err := cntlr.GetControllersFromTable(cntlrTable)
	if err != nil {
		check.ExitError(err)
	}

	if len(driveTable.Snmp.Values) == 0 {
		check.ExitRaw(3, "No HP drive data found!")
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
	check.ExitRaw(status, overall.GetOutput())
}
