package cntlr

// Source: https://support.hpe.com/hpesc/public/docDisplay?docLocale=en_US&docId=a00097210en_us
// Version: 3
// Effective date: 2020-04-06

// Note: Always validate IsAffected() when changing values here!!
const VersionAffectedRaid1 = "2.62"

var VersionAffectedRaid5 = []string{"1.98", "1.99", "2.02", "2.03"}

const VersionFixed = "2.65"

//AffectedModelList list of model numbers for controllers
var AffectedModelList = []*AffectedModel{
	{"e208i-p"},
	{"e208i-a"},
	{"e208i-c"},
	{"e208e-p"},
	{"p204i-b"},
	{"p204i-c"},
	{"p408i-p"},
	{"p408i-a"},
	{"p408e-p"},
	{"p408i-c"},
	{"p408e-m"},
	{"p416ie-m"},
	{"p816i-a"},
	{"p408i-sb"},
}
