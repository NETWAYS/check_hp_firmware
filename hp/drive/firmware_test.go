package drive

import (
	"testing"
)

const testModelA = "VO0480JFDGT"
const testModelB = "EK0800JVYPN"

func TestSplitFirmware(t *testing.T) {
	prefix, version, err := SplitFirmware("HPD5")

	if "HPD" != prefix {
		t.Fatalf("expected %s, got %s", "HPD", prefix)
	}

	if 5 != version {
		t.Fatalf("expected %d, got %d", 5, version)
	}

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	prefix, version, err = SplitFirmware("HPD10")

	if "HPD" != prefix {
		t.Fatalf("expected %s, got %s", "HPD", prefix)
	}

	if 10 != version {
		t.Fatalf("expected %d, got %d", 10, version)
	}

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	_, _, err = SplitFirmware("1HPD5")

	if err == nil {
		t.Fatalf("expected error, got none")
	}
}

func TestIsFirmwareFixed(t *testing.T) {
	testsA := map[string]bool{
		"HPD5":  false,
		"HPD7":  false,
		"HPD8":  true,
		"HPD9":  true,
		"HPD10": true,
	}

	modelA := AffectedModels[testModelA]

	for fw, expect := range testsA {
		ok, err := IsFirmwareFixed(modelA, fw)

		if ok != expect {
			t.Fatalf("expected %v, got %v", ok, expect)
		}

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	}

	testsB := map[string]bool{
		"HPD5": false,
		"HPD6": false,
		"HPD7": true,
		"HPD8": true,
	}

	modelB := AffectedModels[testModelB]

	for fw, expect := range testsB {
		ok, err := IsFirmwareFixed(modelB, fw)

		if ok != expect {
			t.Fatalf("expected %v, got %v", ok, expect)
		}

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	}

	otherModel := &AffectedModel{"ABC", "nothing", "HPX11"}
	_, err := IsFirmwareFixed(otherModel, "HPD5")

	if err == nil {
		t.Fatalf("expected error, got none")
	}
}
