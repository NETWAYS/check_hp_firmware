package ilo

import (
	"testing"

	"github.com/NETWAYS/check_hp_firmware/snmp"

	"github.com/NETWAYS/go-check"
	"github.com/stretchr/testify/assert"
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
			assert.Equal(t, state, tc.expectedState)
			assert.Contains(t, output, tc.expectedOutput)
		})
	}
}

func TestIsNewerVersion(t *testing.T) {
	// Compare required Version with current Version
	assert.True(t, isNewerVersion("1.0", "1.0"))
	assert.True(t, isNewerVersion("1.0", "1.1"))
	assert.True(t, isNewerVersion("1.0", "5"))
	assert.True(t, isNewerVersion("1.0", "10.1.0"))

	assert.False(t, isNewerVersion("1.0", "0.9"))
	assert.False(t, isNewerVersion("1.0", "0.9"))
	assert.False(t, isNewerVersion("1.0", "0.0"))
	assert.False(t, isNewerVersion("1.0", "0"))

	assert.False(t, isNewerVersion("1.0", "foobar"))
	assert.False(t, isNewerVersion("foobar", "1.0"))
	assert.False(t, isNewerVersion("xxx", "xxx"))
}

func TestGetIloInformation_ilo5(t *testing.T) {
	c, _ := snmp.NewFileHandlerFromFile("../../testdata/ilo5.txt")

	i, err := GetIloInformation(c)

	assert.NoError(t, err)

	assert.Equal(t, 11, i.ModelID)
	assert.Equal(t, "pciIntegratedLightsOutRemoteInsight5", i.Model)
	assert.Equal(t, "3.00", i.RomRevision)
}

func TestGetIloInformation_ilo6(t *testing.T) {
	c, _ := snmp.NewFileHandlerFromFile("../../testdata/ilo6.txt")

	i, err := GetIloInformation(c)

	assert.NoError(t, err)

	assert.Equal(t, 12, i.ModelID)
	assert.Equal(t, "pciIntegratedLightsOutRemoteInsight6", i.Model)
	assert.Equal(t, "1.55", i.RomRevision)
}
