package cntlr

import (
	"testing"

	"github.com/NETWAYS/go-check"
	"github.com/stretchr/testify/assert"
)

func TestIlo_GetNagiosStatus(t *testing.T) {
	testcases := map[string]struct {
		controller     Controller
		expectedState  int
		expectedOutput string
	}{
		"status-ok": {
			controller: Controller{
				ID:     "id123",
				Model:  "model",
				FwRev:  "revision",
				Serial: "12345",
				Status: "ok",
			},
			expectedState:  check.OK,
			expectedOutput: "controller (id123) model=model serial=12345 firmware=revision",
		},
		"status-not-ok-not-affected": {
			controller: Controller{
				ID:     "id123",
				Model:  "model",
				FwRev:  "revision",
				Serial: "12345",
				Status: "not-ok",
			},
			expectedState:  check.Critical,
			expectedOutput: "controller (id123) model=model serial=12345 firmware=revision",
		},
		"status-not-ok-affected": {
			controller: Controller{
				ID:     "id123",
				Model:  "e208i-p",
				FwRev:  "1.98",
				Serial: "12345",
				Status: "ok",
			},
			expectedState:  check.Critical,
			expectedOutput: "controller (id123) model=e208i-p serial=12345 firmware=1.98 - if you have RAID 5/6/50/60 - update immediately!",
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			state, output := tc.controller.GetNagiosStatus()
			assert.Equal(t, state, tc.expectedState)
			assert.Contains(t, output, tc.expectedOutput)
		})
	}
}
