package ilo

import (
	"fmt"

	"github.com/NETWAYS/check_hp_firmware/hp/mib"
	"github.com/NETWAYS/go-check"
	"github.com/gosnmp/gosnmp"
	"github.com/hashicorp/go-version"
)

type Ilo struct {
	ModelID     int
	Model       string
	RomRevision string
}

// GetIloInformation retrieves the iLO's Model and Rom Revision via SNMP
// and returns an Ilo struct.
func GetIloInformation(client gosnmp.Handler) (ilo *Ilo, err error) {
	oids := []string{
		mib.CpqSm2CntlrModel + ".0",
		mib.CpqSm2CntlrRomRevision + ".0",
	}

	ilo = &Ilo{}

	iloVariables, err := client.Get(oids)

	if err != nil {
		err = fmt.Errorf("could not get SNMP data for iLO: %w", err)
		return
	}

	// Since we only have two variable of different type we don't need to check their names
	for _, v := range iloVariables.Variables {
		switch v.Type {
		case gosnmp.OctetString: // CpqSm2CntlrRomRevision
			// Using Sprintf makes this work for (string) and ([]byte)
			ilo.RomRevision = fmt.Sprintf("%s", v.Value)
		case gosnmp.Integer: // CpqSm2CntlrModel
			modelID := v.Value.(int)
			ilo.ModelID = modelID

			if model, ok := mib.CpqSm2CntlrModelMap[modelID]; ok {
				ilo.Model = model
			}
		}
	}

	return
}

// GetNagiosStatus validates the iLO's data against the known models
// in this plugin.
func (ilo *Ilo) GetNagiosStatus(returnStateforPatch int) (state int, output string) {
	// nolint: ineffassign
	state = check.Unknown

	// Check if the SNMP id is an older model, then alert
	if ilo.ModelID <= OlderModels {
		state = check.Critical
		output = fmt.Sprintf("ILO model %s (%d) is pretty old and likely unsafe", ilo.Model, ilo.ModelID)

		return
	}

	// Check if we know fixed versions for the generation, other models are only reported
	modelInfo, found := FixedVersionMap[ilo.Model]
	if !found {
		state = check.OK
		output = fmt.Sprintf("Integrated Lights-Out model %s (%d) revision %s not known for any issues",
			ilo.Model, ilo.ModelID, ilo.RomRevision)

		return
	}

	output = fmt.Sprintf("Integrated Lights-Out %s revision %s ", modelInfo.Name, ilo.RomRevision)

	if !isNewerVersion(modelInfo.FixedRelease, ilo.RomRevision) {
		state = returnStateforPatch
		output += "- Patch available, should be at least " + modelInfo.FixedRelease
	} else {
		state = check.OK
		output += "- version newer than affected"
	}

	return
}

// isNewerVersion compares the current version against the required version
func isNewerVersion(required, current string) bool {
	currentVersion, cErr := version.NewVersion(current)
	requiredVersion, rErr := version.NewVersion(required)

	if cErr != nil || rErr != nil {
		return false
	}

	return currentVersion.GreaterThanOrEqual(requiredVersion)
}
