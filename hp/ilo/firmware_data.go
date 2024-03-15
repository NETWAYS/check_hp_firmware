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
// HPE iLO 6 v1.56 Release Notes https://support.hpe.com/hpesc/public/docDisplay?docLocale=en_US&docId=sd00003963en_us
// HPE iLO 5 v3.01 Release Notes https://support.hpe.com/hpesc/public/docDisplay?docLocale=en_US&docId=sd00003959en_us
// HPE iLO 4 v2.82 Release Notes https://support.hpe.com/hpesc/public/docDisplay?docId=c03334036en_us&page=index.html
var FixedVersionMap = map[string]PlatformInfo{
	"pciIntegratedLightsOutRemoteInsight3": {"3", "1.93"},
	"pciIntegratedLightsOutRemoteInsight4": {"4", "2.82"},
	"pciIntegratedLightsOutRemoteInsight5": {"5", "3.01"},
	"pciIntegratedLightsOutRemoteInsight6": {"6", "1.56"},
}
