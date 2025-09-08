package ilo

import (
	"strings"
	"testing"

	"github.com/NETWAYS/check_hp_firmware/snmp"

	"github.com/NETWAYS/go-check"
)

func TestIlo_GetNagiosStatus(t *testing.T) {
	testcases := map[string]struct {
		ilo            Ilo
		expectedState  int
		expectedOutput string
	}{
		"too-old": {
			ilo: Ilo{
				Model:       "pciIntegratedLightsOutRemoteInsight3",
				ModelID:     9,
				RomRevision: "1.40",
			},
			expectedState:  check.Warning,
			expectedOutput: "Patch available",
		},
		"newer": {
			ilo: Ilo{
				Model:       "pciIntegratedLightsOutRemoteInsight3",
				ModelID:     9,
				RomRevision: "2.18",
			},
			expectedState:  check.OK,
			expectedOutput: "2.18 - version newer than affected",
		},
		"pretty-old": {
			ilo: Ilo{
				Model:       "pciIntegratedLightsOutRemoteInsight2",
				ModelID:     7,
				RomRevision: "2.18",
			},
			expectedState:  check.Critical,
			expectedOutput: "is pretty old and likely unsafe",
		},
		"unknown-new": {
			ilo: Ilo{
				Model:       "verynew",
				ModelID:     12,
				RomRevision: "2.18",
			},
			expectedState:  check.OK,
			expectedOutput: "verynew (12) revision 2.18 not known for any issues",
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			state, output := tc.ilo.GetNagiosStatus(1)
			if state != tc.expectedState {
				t.Fatalf("expected %v, got %v", tc.expectedState, state)
			}

			if !strings.Contains(output, tc.expectedOutput) {
				t.Fatalf("expected %v, got %v", tc.expectedOutput, output)
			}
		})
	}
}

func TestIsNewerVersion(t *testing.T) {
	tests := []struct {
		current  string
		required string
		expected bool
	}{
		{"1.0", "1.0", true},
		{"1.0", "1.1", true},
		{"1.0", "5", true},
		{"1.0", "10.1.0", true},
		{"1.0", "0.9", false},
		{"1.0", "0.9", false}, // Duplicate test case (can be removed if not intentional)
		{"1.0", "0.0", false},
		{"1.0", "0", false},
		{"1.0", "foobar", false},
		{"foobar", "1.0", false},
		{"xxx", "xxx", false},
	}

	for _, test := range tests {
		result := isNewerVersion(test.current, test.required)
		if result != test.expected {
			t.Errorf("isNewerVersion(%q, %q) = %v, want %v",
				test.current, test.required, result, test.expected)
		}
	}
}

func TestGetIloInformation_ilo5(t *testing.T) {
	c, _ := snmp.NewFileHandlerFromFile("../../testdata/ilo5.txt")

	i, err := GetIloInformation(c)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if i.ModelID != 11 {
		t.Errorf("Expected ModelID to be %d, got %d", 11, i.ModelID)
	}
	if i.Model != "pciIntegratedLightsOutRemoteInsight5" {
		t.Errorf("Expected Model to be %q, got %q", "pciIntegratedLightsOutRemoteInsight5", i.Model)
	}
	if i.RomRevision != "3.00" {
		t.Errorf("Expected RomRevision to be %q, got %q", "3.00", i.RomRevision)
	}
}

func TestGetIloInformation_ilo6(t *testing.T) {
	c, _ := snmp.NewFileHandlerFromFile("../../testdata/ilo6.txt")

	i, err := GetIloInformation(c)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if i.ModelID != 12 {
		t.Errorf("Expected ModelID to be %d, got %d", 12, i.ModelID)
	}
	if i.Model != "pciIntegratedLightsOutRemoteInsight6" {
		t.Errorf("Expected Model to be %q, got %q", "pciIntegratedLightsOutRemoteInsight6", i.Model)
	}
	if i.RomRevision != "1.55" {
		t.Errorf("Expected RomRevision to be %q, got %q", "1.55", i.RomRevision)
	}
}
