check_hp_disk_firmware
======================

![Go build](https://github.com/NETWAYS/check_hp_disk_firmware/workflows/Go/badge.svg?branch=master)

<!-- Note: Update `Readme` in main.go when changing this! -->

Icinga / Nagios check plugin to verify HPE controllers an SSD disks are not affected by certain vulnerabilities.

For controllers:
> HPE Smart Array SR Gen10 Controller Firmware Version 2.65 (or later) provided in the (HPE document a00097210) is
> required to prevent a potential data inconsistency on select RAID configurations with Smart Array Gen10 Firmware
> Version 1.98 through 2.62, based on the following scenarios. HPE strongly recommends performing this upgrade at the
> customer's earliest opportunity per the "Action Required" in the table located in the Resolution section.
> Neglecting to perform the recommended resolution could result in potential subsequent errors and potential data
> inconsistency.

For SSD disks:
> HPE SAS Solid State Drives - Critical Firmware Upgrade Required for Certain HPE SAS Solid State Drive Models to
> Prevent Drive Failure at 32,768 or 40,000 Hours of Operation

Please see support documents from HPE:
* [a00092491](https://support.hpe.com/hpesc/public/docDisplay?docLocale=en_US&docId=emr_na-a00092491en_us)
* [a00097382](https://support.hpe.com/hpesc/public/docDisplay?docLocale=en_US&docId=a00097382en_us)
* [a00097210](https://support.hpe.com/hpesc/public/docDisplay?docLocale=en_US&docId=a00097210en_us)

**IMPORTANT:** Read the documentation for HPE! The plugin and its documentation is a best effort to find and detect
affected hardware. There is ABSOLUTELY NO WARRANTY, see the license!

## Usage

Arguments:

      -H, --hostname string        SNMP host (default "localhost")
      -c, --community string       SNMP community (default "public")
      -P, --protocol string        SNMP protocol (default "2c")
          --timeout int            SNMP timeout in seconds (default 15)
          --snmpwalk-file string   Read output from snmpwalk
      -4, --ipv4                   Use IPv4
      -6, --ipv6                   Use IPv6
      -V, --version                Show version
          --debug                  Enable debug output

Simply run the command:

    $ ./check_hp_disk_firmware -H localhost -c public

## Installation

This is a golang project, either download the binary from the releases:

https://github.com/NETWAYS/check_hp_disk_firmware/releases

Also see the included [CheckCommand for Icinga 2](icinga2.conf).

You can download or build the project locally with go:

    $ go get github.com/NETWAYS/check_hp_cve
    
    $ git clone https://github.com/NETWAYS/check_hp_disk_firmware
    $ cd check_hp_disk_firmware/
    $ go build -o check_hp_disk_firmware .

## Example

    OK - All 2 controllers and 33 drives seem fine
    [OK] controller (0) model=p816i-a serial=XXX firmware=1.65 - firmware older than affected
    [OK] controller (4) model=p408e-p serial=XXX firmware=1.65 - firmware older than affected
    [OK] (0.9 ) model=MO003200JWFWR serial=XXX firmware=HPD2 hours=8086
    [OK] (0.11) model=EK000400GWEPE serial=XXX firmware=HPG0 hours=8086
    [OK] (0.12) model=EK000400GWEPE serial=XXX firmware=HPG0 hours=8086
    [OK] (0.14) model=MO003200JWFWR serial=XXX firmware=HPD2 hours=8086
    [OK] (4.0 ) model=MO3200JFFCL serial=XXX firmware=HPD8 hours=7568 - firmware update applied
    [OK] (4.1 ) model=MO3200JFFCL serial=XXX firmware=HPD8 hours=7568 - firmware update applied
    [OK] (4.2 ) model=MO3200JFFCL serial=XXX firmware=HPD8 hours=7568 - firmware update applied
    [OK] (4.3 ) model=MO3200JFFCL serial=XXX firmware=HPD8 hours=7568 - firmware update applied
    [OK] (4.4 ) model=MO3200JFFCL serial=XXX firmware=HPD8 hours=7568 - firmware update applied
    [OK] (4.5 ) model=MO3200JFFCL serial=XXX firmware=HPD8 hours=7568 - firmware update applied
    [OK] (4.6 ) model=MO3200JFFCL serial=XXX firmware=HPD8 hours=7568 - firmware update applied
    [OK] (4.24) model=MO3200JFFCL serial=XXX firmware=HPD8 hours=7568 - firmware update applied
    [OK] (4.25) model=MO3200JFFCL serial=XXX firmware=HPD8 hours=7568 - firmware update applied
    [OK] (4.26) model=MO3200JFFCL serial=XXX firmware=HPD8 hours=7568 - firmware update applied
    [OK] (4.27) model=MO3200JFFCL serial=XXX firmware=HPD8 hours=7568 - firmware update applied
    [OK] (4.28) model=MO3200JFFCL serial=XXX firmware=HPD8 hours=7568 - firmware update applied
    [OK] (4.29) model=MO3200JFFCL serial=XXX firmware=HPD8 hours=7568 - firmware update applied
    [OK] (4.30) model=MO3200JFFCL serial=XXX firmware=HPD8 hours=7568 - firmware update applied
    [OK] (4.31) model=MO3200JFFCL serial=XXX firmware=HPD8 hours=7568 - firmware update applied
    [OK] (4.50) model=MO3200JFFCL serial=XXX firmware=HPD8 hours=7568 - firmware update applied
    [OK] (4.51) model=MO003200JWFWR serial=XXX firmware=HPD2 hours=7568
    [OK] (4.52) model=MO3200JFFCL serial=XXX firmware=HPD8 hours=7568 - firmware update applied
    [OK] (4.53) model=MO3200JFFCL serial=XXX firmware=HPD8 hours=7568 - firmware update applied
    [OK] (4.54) model=MO3200JFFCL serial=XXX firmware=HPD8 hours=7568 - firmware update applied
    [OK] (4.55) model=MO3200JFFCL serial=XXX firmware=HPD8 hours=7568 - firmware update applied
    [OK] (4.56) model=MO3200JFFCL serial=XXX firmware=HPD8 hours=7568 - firmware update applied
    [OK] (4.75) model=MO3200JFFCL serial=XXX firmware=HPD8 hours=7568 - firmware update applied
    [OK] (4.76) model=MO3200JFFCL serial=XXX firmware=HPD8 hours=7568 - firmware update applied
    [OK] (4.77) model=MO3200JFFCL serial=XXX firmware=HPD8 hours=7568 - firmware update applied
    [OK] (4.78) model=MO3200JFFCL serial=XXX firmware=HPD8 hours=7568 - firmware update applied
    [OK] (4.79) model=MO3200JFFCL serial=XXX firmware=HPD8 hours=7568 - firmware update applied
    [OK] (4.80) model=MO3200JFFCL serial=XXX firmware=HPD8 hours=7568 - firmware update applied
    [OK] (4.81) model=MO3200JFFCL serial=XXX firmware=HPD8 hours=7568 - firmware update applied

## Limitations

* No SNMPv3 support is implemented

## Contribute

If you find bugs or want to add features, please open an issue or pull-request on GitHub.

You can help with problems by supplying the output of snmpwalk for the system you experience problems with:

    $ snmpwalk -c public -v2c -On HOST 1.3.6.1.4.1.232
    
Please make sure to either censor the output of any private information, or send an e-mail to support@netways.de,
so we can provide you with a secure upload link, that won't be shared with public.

## Technical Details

Supported hardware is split into modules: [hp/cntlr](hp/cntlr) [hp/phy_drv](hp/phy_drv)

Known models and affected firmware is documented in: [hp/cntlr/firmware_data.go](hp/cntlr/firmware_data.go) [hp/phy_drv/firmware_data.go](hp/phy_drv/firmware_data.go)

This data can be easily enhanced in the future. Make sure to document source documents and versions as well, and check
the accompanying firmware and status functions.

The check reads the `cpqDaCntlrTable` and `cpqDaPhyDrvTable` tables from SNMP, which should be available over the
IPMI agent or the locally installed HP tools, hooked into the SNMP daemon of the operating system.

## License

Copyright (C) 2020 Markus Frosch <markus.frosch@netways.de>

Copyright (C) 2020 NETWAYS <info@netways.de>

This program is free software; you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation; either version 2 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License along
with this program; if not, write to the Free Software Foundation, Inc.,
51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
