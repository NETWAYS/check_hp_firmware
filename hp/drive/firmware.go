package drive

import (
	"fmt"
	"regexp"
	"strconv"
)

type AffectedModel struct {
	ModelNo       string
	Description   string
	FixedFirmware string
}

type AffectedIndex map[string]*AffectedModel

var (
	AffectedModels = AffectedIndex{}
	versionRE      = regexp.MustCompile(`^([A-Z]+)([0-9]+)$`)
)

func init() {
	for _, drive := range AffectedModelList {
		AffectedModels[drive.ModelNo] = drive
	}
}

func SplitFirmware(firmware string) (prefix string, version int, err error) {
	match := versionRE.FindStringSubmatch(firmware)
	if match == nil {
		return "", 0, fmt.Errorf("could not parse firmware version: %s", firmware)
	}

	version, err = strconv.Atoi(match[2])
	if err != nil {
		return "", 0, fmt.Errorf("unable to parse version: %w", err)
	}

	return match[1], version, nil
}

func IsFirmwareFixed(model *AffectedModel, firmware string) (bool, error) {
	currentPrefix, currentVersion, err := SplitFirmware(firmware)
	if err != nil {
		return false, err
	}

	fixedPrefix, fixedVersion, err := SplitFirmware(model.FixedFirmware)
	if err != nil {
		return false, err
	}

	if currentPrefix != fixedPrefix {
		return false, fmt.Errorf("could not compare versions between: current=%s and fixed=%s",
			firmware, model.FixedFirmware)
	}

	return currentVersion >= fixedVersion, nil
}
