package ilo

import (
	"github.com/NETWAYS/go-check"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIlo_GetNagiosStatus(t *testing.T) {
	ilo := Ilo{
		Model:       "pciIntegratedLightsOutRemoteInsight3",
		ModelID:     9,
		RomRevision: "1.40",
	}

	state, output := ilo.GetNagiosStatus()
	assert.Equal(t, check.Critical, state)
	assert.Contains(t, output, "too old")

	ilo.RomRevision = "2.18"
	state, output = ilo.GetNagiosStatus()
	assert.Equal(t, check.OK, state)
	assert.Contains(t, output, "2.18")

	ilo.Model = "pciIntegratedLightsOutRemoteInsight2"
	ilo.ModelID = 7
	state, output = ilo.GetNagiosStatus()
	assert.Equal(t, check.Critical, state)
	assert.Contains(t, output, "pretty old")

	ilo.Model = "someNewerModel"
	ilo.ModelID = 12
	state, output = ilo.GetNagiosStatus()
	assert.Equal(t, check.OK, state)
	assert.Contains(t, output, "not known")
}

func TestCompareVer(t *testing.T) {
	assert.True(t, CompareVersion("1.0", "1.0"))
	assert.True(t, CompareVersion("1.0", "1.1"))
	assert.False(t, CompareVersion("1.0", "0.9"))
}
