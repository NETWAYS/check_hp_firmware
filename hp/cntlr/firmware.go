package cntlr

import (
	"github.com/NETWAYS/go-check"
	"github.com/hashicorp/go-version"
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

// IsAffected validates the given version against known affected versions.
// Note: we can't validate against existing logical drives at the moment
func IsAffected(firmware string) (int, string) {
	firmwareVersion, _ := version.NewVersion(firmware)
	fixedVersion, _ := version.NewVersion(VersionFixed)

	if firmwareVersion.GreaterThanOrEqual(fixedVersion) {
		return check.OK, "firmware has been updated"
	}

	affectedRaid1, _ := version.NewVersion(VersionAffectedRaid1)

	if firmwareVersion.GreaterThanOrEqual(affectedRaid1) {
		return check.Critical, "if you have RAID 1/10/ADM - update immediately!"
	}

	for _, v := range VersionAffectedRaid5 {
		affectedRaid5, _ := version.NewVersion(v)
		if firmwareVersion.Equal(affectedRaid5) {
			return check.Critical, "if you have RAID 5/6/50/60 - update immediately!"
		}
	}

	return check.OK, "firmware older than affected"
}
