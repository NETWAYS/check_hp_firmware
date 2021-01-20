package ilo

import (
	"fmt"
	"github.com/NETWAYS/check_hp_firmware/hp/mib"
	"github.com/NETWAYS/check_hp_firmware/nagios"
	"github.com/gosnmp/gosnmp"
	"github.com/hashicorp/go-version"
)

type Ilo struct {
	Model       string
	RomRevision string
}

func GetIloInformation(client gosnmp.Handler) (int, string)  {
	oidModel := []string{mib.CpqSm2CntlrModel + ".0"}
	oidRev := []string{ mib.CpqSm2CntlrRomRevision + ".0"}

	ilo := &Ilo{}
	parseErr := ""

	iloModel, err := client.Get(oidModel)
	if err != nil {
		return nagios.Critical, parseErr + "could not get model for Ilo"
	}

	iloRev, err := client.Get(oidRev)
	if err != nil {
		return nagios.Critical, parseErr + "could not get revision for Ilo"
	} else {
		ilo.RomRevision = iloRev.Variables[0].Value.(string)
	}

	if iloModel, ok := mib.CpqSm2CntlrModelMap[iloModel.Variables[0].Value.(int)]; ok {
		ilo.Model = iloModel
	} else {
		return nagios.Critical, parseErr + "unknown Ilo model"
	}

	description := fmt.Sprintf("Integrated Lights-Out=%s Revision=%s ", ilo.Model, ilo.RomRevision)

	if ilo.Model == "3" {
		if ( ! CompareVer("1.93", iloRev.Variables[0].Value.(string))) {
			return nagios.Critical, description +
				fmt.Sprintf("The Revision: %s does not satisfies constraints 1.93. Update Firmware immediately!",
				ilo.RomRevision)
		}
	} else if ilo.Model == "4" {
		if ( ! CompareVer("2.75", iloRev.Variables[0].Value.(string))) {
			return nagios.Critical, description +
				fmt.Sprintf("The Revision: %s does not satisfies constraints 2.75 Update Firmware immediately!",
				ilo.RomRevision)
		}
	} else if ilo.Model == "5" {
		if ( ! CompareVer("2.18", iloRev.Variables[0].Value.(string))) {
			return nagios.Critical, description +
				fmt.Sprintf("The Revision: %s does not satisfies constraints 2.18 Update Firmware immediately!",
				ilo.RomRevision)
		}
	} else {
		return nagios.Critical, description + fmt.Sprintf("the Ilo Version is to old")
	}

	return nagios.OK, description + fmt.Sprintf("The Revision:%s satisfies constraints", ilo.RomRevision)
}

func CompareVer(constr, vers string) (ret bool) {
	v, err := version.NewVersion(vers)
	if err != nil{
		return false
	}

	c, err := version.NewConstraint(">=" + constr)
	if err != nil {
		return false
	}

	if c.Check(v) {
		return true
	}

	return false
}
