package ilo

type PlatformInfo struct {
	Name         string
	FixedRelease string
}

// OlderModels or lower ids are too old to be safe and should be alerted
const OlderModels = 8

// FixedVersionMap lists the version for which the platform has been fixed by generation
// See mib.CpqSm2CntlrModelMap - not all models are mentioned here
//
// For vendor details see HPESBHF04012 https://support.hpe.com/hpesc/public/docDisplay?docId=hpesbhf04012en_us
var FixedVersionMap = map[string]PlatformInfo{
	"pciIntegratedLightsOutRemoteInsight3": {"3", "1.93"},
	"pciIntegratedLightsOutRemoteInsight4": {"4", "2.75"},
	"pciIntegratedLightsOutRemoteInsight5": {"5", "2.18"},
}
