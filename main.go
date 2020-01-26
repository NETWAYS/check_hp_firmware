package main

import (
	"github.com/NETWAYS/check_hp_cve/hp"
	"github.com/NETWAYS/check_hp_cve/nagios"
	"github.com/NETWAYS/check_hp_cve/snmp"
	log "github.com/sirupsen/logrus"
	"github.com/soniah/gosnmp"
	flag "github.com/spf13/pflag"
	"os"
	"time"
)

// Check for HP PhysicalDrive CVEs via SNMP
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

	debug := flagSet.Bool("debug", false, "Enable debug output")

	var err error
	err = flagSet.Parse(os.Args[1:])
	if err != nil {
		if err != flag.ErrHelp {
			nagios.ExitError(err)
		}
		os.Exit(3)
	}

	if *debug {
		log.SetLevel(log.DebugLevel)
	}

	var client *gosnmp.GoSNMP
	var table *hp.CpqDaPhyDrvTable

	if *file != "" {
		var fh *os.File
		fh, err = os.Open(*file)
		if err != nil {
			nagios.ExitError(err)
		}
		defer fh.Close()

		table, err = hp.LoadCpqDaPhyDrvTable(fh)
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

		table, err = hp.GetCpqDaPhyDrvTable(client)
	}
	if err != nil {
		nagios.ExitError(err)
	}

	ids := table.ListIds()
	if len(ids) == 0 {
		nagios.Exit(3, "No HP drive data found!")
	}

	drives, err := hp.GetPhysicalDrivesFromTable(table)
	if err != nil {
		nagios.ExitError(err)
	}

	// TODO: check if drives found?

	overall := nagios.Overall{}

	for _, drive := range drives {
		driveStatus, desc := drive.GetNagiosStatus()
		overall.Add(driveStatus, desc)
	}

	status := overall.GetStatus()
	var summary string
	switch status {
	case nagios.OK:
		summary = "All drives seem fine"
	case nagios.Warning:
		summary = "Found warnings for drives"
	case nagios.Critical:
		summary = "Found critical problems on drives"
	}
	overall.Summary = summary
	nagios.Exit(status, overall.GetOutput())
}
