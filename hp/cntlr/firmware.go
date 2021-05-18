package cntlr

import (
	"github.com/NETWAYS/go-check"
	"github.com/mcuadros/go-version"
)

type VersionInfo struct {
	Affected string
	Fixed    string
}

type AffectedModel struct {
	Model string
}

type AffectedIndex map[string]*AffectedModel

var AffectedModels = AffectedIndex{}

func init() {
	for _, controller := range AffectedModelList {
		AffectedModels[controller.Model] = controller
	}
}

// Note: we can't validate against existing logical drives at the moment
func IsAffected(firmware string) (int, string) {
	if version.Compare(firmware, VersionFixed, ">=") {
		return check.OK, "firmware has been updated"
	}

	if version.Compare(firmware, VersionAffectedRaid1, ">=") {
		return check.Critical, "if you have RAID 1/10/ADM - update immediately!"
	}

	for _, v := range VersionAffectedRaid5 {
		if v == firmware {
			return check.Critical, "if you have RAID 5/6/50/60 - update immediately!"
		}
	}

	return check.OK, "firmware older than affected"
}
