package main

const Project = "check_hp_disk_firmware"
const Version = "1.0.1"

var GitCommit string

func buildVersion() string {
	version := Version
	if GitCommit != "" {
		version += " - " + GitCommit
	}
	return version
}
