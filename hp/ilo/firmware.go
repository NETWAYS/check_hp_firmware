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
	oidModel := []string{mib.CpqSm2CntlrModel + ".0"}
	oidRev := []string{mib.CpqSm2CntlrRomRevision + ".0"}

	ilo = &Ilo{}

	iloModel, err := client.Get(oidModel)
	if err != nil {
		err = fmt.Errorf("could not get model for Ilo: %s", oidModel[0])
		return
	}

	ilo.ModelID = iloModel.Variables[0].Value.(int)
	if model, ok := mib.CpqSm2CntlrModelMap[ilo.ModelID]; ok {
		ilo.Model = model
	}

	iloRev, err := client.Get(oidRev)
	if err != nil {
		err = fmt.Errorf("could not get revision for Ilo: %s", oidRev[0])
		return
	}

	ilo.RomRevision = iloRev.Variables[0].Value.(string)

	return
}

// GetNagiosStatus validates the iLO's data against the known models
// in this plugin.
func (ilo *Ilo) GetNagiosStatus() (state int, output string) {
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
		state = check.Critical
		output += "- version too old, should be at least " + modelInfo.FixedRelease
	} else {
		state = check.OK
		output += "- version newer than affected"
	}

	return
}

// isNewerVersion compares the current version against the required version
func isNewerVersion(required, current string) bool {
	v, err := version.NewVersion(current)
	if err != nil {
		return false
	}

	c, err := version.NewConstraint(">=" + required)
	if err != nil {
		// TODO remove panic
		panic(err)
	}

	return c.Check(v)
}
