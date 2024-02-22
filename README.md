# check_hp_firmware

Monitoring check plugin to verify HPE controllers an SSD disks are not affected by certain vulnerabilities.

*Note:* This plugin was initially named `check_hp_disk_firmware`.

Current Limitations:

* No SNMPv3 support is implemented

## HPE Controllers

> HPE Smart Array SR Gen10 Controller Firmware Version 2.65 (or later) provided in the (HPE document a00097210) is
> required to prevent a potential data inconsistency on select RAID configurations with Smart Array Gen10 Firmware
> Version 1.98 through 2.62, based on the following scenarios. HPE strongly recommends performing this upgrade at the
> customer's earliest opportunity per the "Action Required" in the table located in the Resolution section.
> Neglecting to perform the recommended resolution could result in potential subsequent errors and potential data
> inconsistency.

The check will alert you with a CRITICAL when the firmware is in the affected range with:

* `if you have RAID 1/10/ADM - update immediately!`
* `if you have RAID 5/6/50/60 - update immediately!`

And it will add a short note when `firmware older than affected` or `firmware has been updated`. At the moment the
plugin does not verify configured logical drives, but we believe you should update in any case.

## HPE SSD SAS disks

> HPE SAS Solid State Drives - Critical Firmware Upgrade Required for Certain HPE SAS Solid State Drive Models to
> Prevent Drive Failure at 32,768 or 40,000 Hours of Operation

The check will raise a CRITICAL when the drive needs to be updated with the note `affected by FW bug`, and when
the drive is patched with `firmware update applied`.

## HPE Integrated Lights-Out

Multiple security vulnerabilities have been identified in Integrated Lights-Out 3 (iLO 3),
Integrated Lights-Out 4 (iLO 4), and Integrated Lights-Out 5 (iLO 5) firmware. The vulnerabilities could be remotely
exploited to execute code, cause denial of service, and expose sensitive information. HPE has released updated
firmware to mitigate these vulnerabilities.

The check will raise a CRITICAL when the Integrated Lights-Out needs to be updated. Below you will find a list with
the least version of each Integrated Lights-Out version:

- HPE iLO 6 v1.56 or later
- HPE iLO 5 v3.01 or later
- HPE iLO 4 v2.82 or later

**IMPORTANT:** Always read the latest HPE Security Bulletins. https://support.hpe.com/connect/s/securitybulletinlibrary

The plugin and its documentation is a best effort to find and detect affected hardware. There is no warranty, see the license.

## Usage

Arguments:

```
-H, --hostname string        SNMP host (default "localhost")
-c, --community string       SNMP community (default "public")
-P, --protocol string        SNMP protocol (default "2c")
    --timeout int            SNMP timeout in seconds (default 15)
    --snmpwalk-file string   Read output from snmpwalk
-I, --ignore-ilo-version     Don't check the ILO version
-4, --ipv4                   Use IPv4
-6, --ipv6                   Use IPv6
-V, --version                Show version
    --debug                  Enable debug output
```

Simply run the command:

```bash
check_hp_firmware -H localhost -c public
```

# Installation

This is a Golang project, either download the binary from the releases:

https://github.com/NETWAYS/check_hp_firmware/releases

Also see the included [CheckCommand for Icinga 2](icinga2.conf).

You can download or build the project locally with go:

```bash
git clone https://github.com/NETWAYS/check_hp_firmware
cd check_hp_firmware/
make build
```

## Example

    OK - All 2 controllers and 33 drives seem fine
    [OK] Integrated Lights-Out 5 revision 2.18 - version newer than affected
    [OK] controller (0) model=p816i-a serial=XXX firmware=1.65 - firmware older than affected
    [OK] controller (4) model=p408e-p serial=XXX firmware=1.65 - firmware older than affected
    [OK] (0.9 ) model=MO003200JWFWR serial=XXX firmware=HPD2 hours=8086
    [OK] (0.11) model=EK000400GWEPE serial=XXX firmware=HPG0 hours=8086
    [OK] (0.14) model=MO003200JWFWR serial=XXX firmware=HPD2 hours=8086
    [OK] (4.0 ) model=MO3200JFFCL serial=XXX firmware=HPD8 hours=7568 - firmware update applied
    [OK] (4.31) model=MO3200JFFCL serial=XXX firmware=HPD8 hours=7568 - firmware update applied
    [OK] (4.50) model=MO3200JFFCL serial=XXX firmware=HPD8 hours=7568 - firmware update applied
    [OK] (4.51) model=MO003200JWFWR serial=XXX firmware=HPD2 hours=7568
    [OK] (4.52) model=MO3200JFFCL serial=XXX firmware=HPD8 hours=7568 - firmware update applied
    [OK] (4.78) model=MO3200JFFCL serial=XXX firmware=HPD8 hours=7568 - firmware update applied
    [OK] (4.81) model=MO3200JFFCL serial=XXX firmware=HPD8 hours=7568 - firmware update applied


# Contribute

If you find bugs or want to add features, please open an issue or pull-request on GitHub.

You can help with problems by supplying the output of snmpwalk for the system you experience problems with:

    $ snmpwalk -c public -v2c -On HOST 1.3.6.1.4.1.232

Please make sure to either censor the output of any private information, or send an e-mail to support@netways.de,
so we can provide you with a secure upload link, that won't be shared with public.

## Technical Details

Supported hardware is split into modules: [hp/cntlr](hp/cntlr) [hp/drive](hp/drive) [hp/ilo](hp/ilo)

Known models and affected firmware is documented in: [hp/cntlr/firmware_data.go](hp/cntlr/firmware_data.go) [hp/phy_drv/firmware_data.go](hp/phy_drv/firmware_data.go) [hp/ilo/firmware_data.go](hp/ilo/firmware_data.go)

This data can be easily enhanced in the future. Make sure to document source documents and versions as well, and check
the accompanying firmware and status functions.

The check reads the `cpqDaCntlrTable` and `cpqDaPhyDrvTable` tables from SNMP, which should be available over the
IPMI agent or the locally installed HP tools, hooked into the SNMP daemon of the operating system.

# License

Copyright (C) 2020 NETWAYS <info@netways.de>

Copyright (C) 2020 Markus Frosch <markus.frosch@netways.de>

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
