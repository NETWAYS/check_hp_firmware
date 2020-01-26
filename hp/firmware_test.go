package hp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsFirmwareFixed(t *testing.T) {
	tests := map[string]bool{
		"HPD5":  false,
		"HPD7":  false,
		"HPD8":  true,
		"HPD9":  true,
		"HPD10": true,
	}

	for fw, expect := range tests {
		ok, err := IsFirmwareFixed(fw)
		assert.Equal(t, expect, ok)
		assert.NoError(t, err)
	}
}
