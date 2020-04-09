package phy_drv

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const testModelA = "VO0480JFDGT"
const testModelB = "EK0800JVYPN"

func TestSplitFirmware(t *testing.T) {
	prefix, version, err := SplitFirmware("HPD5")
	assert.Equal(t, "HPD", prefix)
	assert.Equal(t, 5, version)
	assert.Nil(t, err)

	prefix, version, err = SplitFirmware("HPD10")
	assert.Equal(t, "HPD", prefix)
	assert.Equal(t, 10, version)
	assert.Nil(t, err)

	prefix, version, err = SplitFirmware("1HPD5")
	assert.Error(t, err)
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
		assert.Equal(t, expect, ok)
		assert.NoError(t, err)
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
		assert.Equal(t, expect, ok)
		assert.NoError(t, err)
	}

	otherModel := &AffectedModel{"ABC", "nothing", "HPX11"}
	_, err := IsFirmwareFixed(otherModel, "HPD5")
	assert.Error(t, err)
}
