package ilo

import (
	"github.com/NETWAYS/go-check"
	"github.com/stretchr/testify/assert"
	"testing"
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
			expectedState:  check.Critical,
			expectedOutput: "too old",
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
			state, output := tc.ilo.GetNagiosStatus()
			assert.Equal(t, state, tc.expectedState)
			assert.Contains(t, output, tc.expectedOutput)
		})
	}

}

func TestCompareVer(t *testing.T) {
	// Compare required Version with current Version
	assert.True(t, CompareVersion("1.0", "1.0"))
	assert.True(t, CompareVersion("1.0", "1.1"))
	assert.True(t, CompareVersion("1.0", "5"))
	assert.True(t, CompareVersion("1.0", "10.1.0"))

	assert.False(t, CompareVersion("1.0", "0.9"))
	assert.False(t, CompareVersion("1.0", "0.9"))
	assert.False(t, CompareVersion("1.0", "0.0"))
	assert.False(t, CompareVersion("1.0", "0"))
}
