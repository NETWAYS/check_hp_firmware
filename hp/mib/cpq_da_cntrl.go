package mib

//noinspection GoUnusedConst,SpellCheckingInspection
const (
	CpqDaCntlrTable                 = `.1.3.6.1.4.1.232.3.2.2.1`
	CpqDaCntlrEntry                 = `.1.3.6.1.4.1.232.3.2.2.1.1`
	CpqDaCntlrIndex                 = `.1.3.6.1.4.1.232.3.2.2.1.1.1`
	CpqDaCntlrModel                 = `.1.3.6.1.4.1.232.3.2.2.1.1.2`
	CpqDaCntlrFWRev                 = `.1.3.6.1.4.1.232.3.2.2.1.1.3`
	CpqDaCntlrStndIntr              = `.1.3.6.1.4.1.232.3.2.2.1.1.4`
	CpqDaCntlrSlot                  = `.1.3.6.1.4.1.232.3.2.2.1.1.5`
	CpqDaCntlrCondition             = `.1.3.6.1.4.1.232.3.2.2.1.1.6`
	CpqDaCntlrProductRev            = `.1.3.6.1.4.1.232.3.2.2.1.1.7`
	CpqDaCntlrPartnerSlot           = `.1.3.6.1.4.1.232.3.2.2.1.1.8`
	CpqDaCntlrCurrentRole           = `.1.3.6.1.4.1.232.3.2.2.1.1.9`
	CpqDaCntlrBoardStatus           = `.1.3.6.1.4.1.232.3.2.2.1.1.10`
	CpqDaCntlrPartnerBoardStatus    = `.1.3.6.1.4.1.232.3.2.2.1.1.11`
	CpqDaCntlrBoardCondition        = `.1.3.6.1.4.1.232.3.2.2.1.1.12`
	CpqDaCntlrPartnerBoardCondition = `.1.3.6.1.4.1.232.3.2.2.1.1.13`
	CpqDaCntlrDriveOwnership        = `.1.3.6.1.4.1.232.3.2.2.1.1.14`
	CpqDaCntlrSerialNumber          = `.1.3.6.1.4.1.232.3.2.2.1.1.15`
	CpqDaCntlrRedundancyType        = `.1.3.6.1.4.1.232.3.2.2.1.1.16`
	CpqDaCntlrRedundancyError       = `.1.3.6.1.4.1.232.3.2.2.1.1.17`
	CpqDaCntlrAccessModuleStatus    = `.1.3.6.1.4.1.232.3.2.2.1.1.18`
	CpqDaCntlrDaughterBoardType     = `.1.3.6.1.4.1.232.3.2.2.1.1.19`
	CpqDaCntlrHwLocation            = `.1.3.6.1.4.1.232.3.2.2.1.1.20`
	CpqDaCntlrNumberOfBuses         = `.1.3.6.1.4.1.232.3.2.2.1.1.21`
	CpqDaCntlrBlinkTime             = `.1.3.6.1.4.1.232.3.2.2.1.1.22`
	CpqDaCntlrRebuildPriority       = `.1.3.6.1.4.1.232.3.2.2.1.1.23`
	CpqDaCntlrExpandPriority        = `.1.3.6.1.4.1.232.3.2.2.1.1.24`
	CpqDaCntlrNumberOfInternalPorts = `.1.3.6.1.4.1.232.3.2.2.1.1.25`
	CpqDaCntlrNumberOfExternalPorts = `.1.3.6.1.4.1.232.3.2.2.1.1.26`
	CpqDaCntlrDriveWriteCacheState  = `.1.3.6.1.4.1.232.3.2.2.1.1.27`
)

var CpqDaCntlrBoardStatusMap = StringMap{
	1: "other",
	2: "ok",
	3: "generalFailure",
	4: "cableProblem",
	5: "poweredOff",
}

var CpqDaCntlrModelMap = StringMap{
	1:   "other",
	2:   "ida",
	3:   "idaExpansion",
	4:   "ida-2",
	5:   "smart",
	6:   "smart-2e",
	7:   "smart-2p",
	8:   "smart-2sl",
	9:   "smart-3100es",
	10:  "smart-3200",
	11:  "smart-2dh",
	12:  "smart-221",
	13:  "sa-4250es",
	14:  "sa-4200",
	15:  "sa-integrated",
	16:  "sa-431",
	17:  "sa-5300",
	18:  "raidLc2",
	19:  "sa-5i",
	20:  "sa-532",
	21:  "sa-5312",
	22:  "sa-641",
	23:  "sa-642",
	24:  "sa-6400",
	25:  "sa-6400em",
	26:  "sa-6i",
	27:  "sa-generic",
	29:  "sa-p600",
	30:  "sa-p400",
	31:  "sa-e200",
	32:  "sa-e200i",
	33:  "sa-p400i",
	34:  "sa-p800",
	35:  "sa-e500",
	36:  "sa-p700m",
	37:  "sa-p212",
	38:  "sa-p410",
	39:  "sa-p410i",
	40:  "sa-p411",
	41:  "sa-b110i",
	42:  "sa-p712m",
	43:  "sa-p711m",
	44:  "sa-p812",
	45:  "sw-1210m",
	46:  "sa-p220i",
	47:  "sa-p222",
	48:  "sa-p420",
	49:  "sa-p420i",
	50:  "sa-p421",
	51:  "sa-b320i",
	52:  "sa-p822",
	53:  "sa-p721m",
	54:  "sa-b120i",
	55:  "hps-1224",
	56:  "hps-1228",
	57:  "hps-1228m",
	58:  "sa-p822se",
	59:  "hps-1224e",
	60:  "hps-1228e",
	61:  "hps-1228em",
	62:  "sa-p230i",
	63:  "sa-p430i",
	64:  "sa-p430",
	65:  "sa-p431",
	66:  "sa-p731m",
	67:  "sa-p830i",
	68:  "sa-p830",
	69:  "sa-p831",
	70:  "sa-p530",
	71:  "sa-p531",
	72:  "sa-p244br",
	73:  "sa-p246br",
	74:  "sa-p440",
	75:  "sa-p440ar",
	76:  "sa-p441",
	77:  "sa-p741m",
	78:  "sa-p840",
	79:  "sa-p841",
	80:  "sh-h240ar",
	81:  "sh-h244br",
	82:  "sh-h240",
	83:  "sh-h241",
	84:  "sa-b140i",
	85:  "sh-generic",
	86:  "sa-p240nr",
	87:  "sh-h240nr",
	88:  "sa-p840ar",
	89:  "sa-p542d",
	90:  "s100i",
	91:  "e208i-p",  // affected by a00097210
	92:  "e208i-a",  // affected by a00097210
	93:  "e208i-c",  // affected by a00097210
	94:  "e208e-p",  // affected by a00097210
	95:  "p204i-b",  // affected by a00097210
	96:  "p204i-c",  // affected by a00097210
	97:  "p408i-p",  // affected by a00097210
	98:  "p408i-a",  // affected by a00097210
	99:  "p408e-p",  // affected by a00097210
	100: "p408i-c",  // affected by a00097210
	101: "p408e-m",  // affected by a00097210
	102: "p416ie-m", // affected by a00097210
	103: "p816i-a",  // affected by a00097210
	104: "p408i-sb", // affected by a00097210
}
