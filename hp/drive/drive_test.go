package drive

import (
	"regexp"
	"testing"

	"github.com/NETWAYS/go-check"
)

const affectedDrive = "VO0480JFDGT"
const affectedDriveFixed = "HPD8"

func matchRegexp(rx string, str string) bool {
	r := regexp.MustCompile(rx)

	return r.MatchString(str)
}

func TestPhysicalDrive_GetNagiosStatus(t *testing.T) {
	drive := &PhysicalDrive{
		ID:     "1.1",
		Model:  "OTHERDRIVE",
		FwRev:  "HPD1",
		Serial: "ABC123",
		Status: "ok",
		Hours:  1337,
	}

	var (
		status int
		info   string
	)

	// good
	status, info = drive.GetNagiosStatus()

	if check.OK != status {
		t.Fatalf("expected %v, got %v", check.OK, status)
	}

	if !matchRegexp(`\(1\.1 \) model=\w+ serial=ABC123 firmware=HPD1 hours=1337`, info) {
		t.Fatalf("%s did not match regex", info)
	}

	// failed
	drive.Status = "failed"
	status, info = drive.GetNagiosStatus()

	if check.Critical != status {
		t.Fatalf("expected %v, got %v", check.Critical, status)
	}

	if !matchRegexp(`\(1\.1 \) model=\w+ serial=ABC123 firmware=HPD1 hours=1337 - status: failed`, info) {
		t.Fatalf("%s did not match regex", info)
	}

	// affected
	drive.Status = "ok"
	drive.Model = affectedDrive
	status, info = drive.GetNagiosStatus()

	if check.Critical != status {
		t.Fatalf("expected %v, got %v", check.Critical, status)
	}

	if !matchRegexp(`\(1\.1 \) model=\w+ serial=ABC123 firmware=HPD1 hours=1337 - affected`, info) {
		t.Fatalf("%s did not match regex", info)
	}

	// affected but fixed
	drive.FwRev = affectedDriveFixed
	status, info = drive.GetNagiosStatus()
	if check.OK != status {
		t.Fatalf("expected %v, got %v", check.OK, status)
	}

	if !matchRegexp(`\(1\.1 \) model=\w+ serial=ABC123 firmware=\w+ hours=1337 - .*applied`, info) {
		t.Fatalf("%s did not match regex", info)
	}
}
