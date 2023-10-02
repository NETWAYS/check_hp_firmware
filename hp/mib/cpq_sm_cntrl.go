package mib

const (
	CpqSm2Cntlr                               = `.1.3.6.1.4.1.232.9.2.2`
	CpqSm2CntlrRomDate                        = `.1.3.6.1.4.1.232.9.2.2.1`
	CpqSm2CntlrRomRevision                    = `.1.3.6.1.4.1.232.9.2.2.2`
	CpqSm2CntlrVideoStatus                    = `.1.3.6.1.4.1.232.9.2.2.3`
	CpqSm2CntlrBatteryEnabled                 = `.1.3.6.1.4.1.232.9.2.2.4`
	CpqSm2CntlrBatteryStatus                  = `.1.3.6.1.4.1.232.9.2.2.5`
	CpqSm2CntlrBatteryPercentCharged          = `.1.3.6.1.4.1.232.9.2.2.6`
	CpqSm2CntlrAlertStatus                    = `.1.3.6.1.4.1.232.9.2.2.7`
	CpqSm2CntlrPendingAlerts                  = `.1.3.6.1.4.1.232.9.2.2.8`
	CpqSm2CntlrSelfTestErrors                 = `.1.3.6.1.4.1.232.9.2.2.9`
	CpqSm2CntlrAgentLocation                  = `.1.3.6.1.4.1.232.9.2.2.10`
	CpqSm2CntlrLastDataUpdate                 = `.1.3.6.1.4.1.232.9.2.2.11`
	CpqSm2CntlrDataStatus                     = `.1.3.6.1.4.1.232.9.2.2.12`
	CpqSm2CntlrColdReboot                     = `.1.3.6.1.4.1.232.9.2.2.13`
	CpqSm2CntlrBadLoginAttemptsThresh         = `.1.3.6.1.4.1.232.9.2.2.14`
	CpqSm2CntlrBoardSerialNumber              = `.1.3.6.1.4.1.232.9.2.2.15`
	CpqSm2CntlrRemoteSessionStatus            = `.1.3.6.1.4.1.232.9.2.2.16`
	CpqSm2CntlrInterfaceStatus                = `.1.3.6.1.4.1.232.9.2.2.17`
	CpqSm2CntlrSystemId                       = `.1.3.6.1.4.1.232.9.2.2.18`
	CpqSm2CntlrKeyboardCableStatus            = `.1.3.6.1.4.1.232.9.2.2.19`
	CpqSm2ServerIpAddress                     = `.1.3.6.1.4.1.232.9.2.2.20`
	CpqSm2CntlrModel                          = `.1.3.6.1.4.1.232.9.2.2.21`
	CpqSm2CntlrSelfTestErrorMask              = `.1.3.6.1.4.1.232.9.2.2.22`
	CpqSm2CntlrMouseCableStatus               = `.1.3.6.1.4.1.232.9.2.2.23`
	CpqSm2CntlrVirtualPowerCableStatus        = `.1.3.6.1.4.1.232.9.2.2.24`
	CpqSm2CntlrExternalPowerCableStatus       = `.1.3.6.1.4.1.232.9.2.2.25`
	CpqSm2CntlrHostGUID                       = `.1.3.6.1.4.1.232.9.2.2.26`
	CpqSm2CntlriLOSecurityOverrideSwitchState = `.1.3.6.1.4.1.232.9.2.2.27`
	CpqSm2CntlrHardwareVer                    = `.1.3.6.1.4.1.232.9.2.2.28`
	CpqSm2CntlrAction                         = `.1.3.6.1.4.1.232.9.2.2.29`
	CpqSm2CntlrLicenseActive                  = `.1.3.6.1.4.1.232.9.2.2.30`
	CpqSm2CntlrLicenseKey                     = `.1.3.6.1.4.1.232.9.2.2.31`
)

var CpqSm2CntlrModelMap = StringMap{
	1:  "other",
	2:  "eisaRemoteInsightBoard",               // EISA Remote Insight
	3:  "pciRemoteInsightBoard",                // PCI Remote Insight
	4:  "pciLightsOutRemoteInsightBoard",       // Remote Insight Lights-Out Edition
	5:  "pciIntegratedLightsOutRemoteInsight",  // Integrated Remote Insight Lights-Out Edition.
	6:  "pciLightsOutRemoteInsightBoardII",     // Remote Insight Lights-Out Edition version II
	7:  "pciIntegratedLightsOutRemoteInsight2", // Integrated Lights-Out 2 Edition
	8:  "pciLightsOut100series",                // Lights-Out 100 Edition for 100 Series of ProLiant servers
	9:  "pciIntegratedLightsOutRemoteInsight3", // Integrated Lights-Out 3 Edition
	10: "pciIntegratedLightsOutRemoteInsight4", // Integrated Lights-Out 4 Edition
	11: "pciIntegratedLightsOutRemoteInsight5", // Integrated Lights-Out 5 Edition
}
