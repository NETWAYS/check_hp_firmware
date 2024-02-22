package cntlr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testInfo struct {
	rc   int
	info string
}

func TestIsAffected(t *testing.T) {
	versions := map[string]testInfo{
		"1.65": {0, "firmware older than affected"},
		"1.98": {2, "RAID 5"},
		"1.99": {2, "RAID 5"},
		"2.02": {2, "RAID 5"},
		"2.03": {2, "RAID 5"},
		"2.62": {2, "RAID 1"},
		"2.65": {0, "updated"},
	}

	for fw, expect := range versions {
		rc, info := IsAffected(fw)
		assert.Equal(t, rc, expect.rc)
		assert.Contains(t, info, expect.info)
	}
}
