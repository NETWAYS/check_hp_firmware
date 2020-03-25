package hp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const testModel = "VO0480JFDGT"

func TestIsAffected(t *testing.T) {
	assert.False(t, IsAffected("UNKNOWN"))
	assert.True(t, IsAffected(testModel))
}

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
	tests := map[string]bool{
		"HPD5":  false,
		"HPD7":  false,
		"HPD8":  true,
		"HPD9":  true,
		"HPD10": true,
	}

	model := AffectedModels[testModel]

	for fw, expect := range tests {
		ok, err := IsFirmwareFixed(model, fw)
		assert.Equal(t, expect, ok)
		assert.NoError(t, err)
	}

	otherModel := &AffectedModel{"ABC", "nothing", "HPX11"}
	_, err := IsFirmwareFixed(otherModel, "HPD5")
	assert.Error(t, err)
}
