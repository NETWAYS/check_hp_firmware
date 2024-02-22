package cntlr

import (
	"fmt"
	"strings"

	"github.com/NETWAYS/check_hp_firmware/hp/mib"
	"github.com/NETWAYS/go-check"
)

type Controller struct {
	ID     string
	Model  string
	FwRev  string
	Serial string
	Status string
}

func NewControllerFromTable(t *CpqDaCntlrTable, id string) (*Controller, error) {
	if _, ok := t.Snmp.Values[id]; !ok {
		return nil, fmt.Errorf("could not find controller %s in table", id)
	}

	var err error

	controller := &Controller{}
	controller.ID = id

	modelI, err := t.GetIntValue(id, mib.CpqDaCntlrModel)
	if err != nil {
		return nil, fmt.Errorf("could not get model for controller %s: %w", id, err)
	}

	if model, ok := mib.CpqDaCntlrModelMap[modelI]; ok {
		controller.Model = model
	} else {
		return nil, fmt.Errorf("unknown model int for controller: %s: %d", id, modelI)
	}

	controller.FwRev, err = t.GetStringValue(id, mib.CpqDaCntlrFWRev)
	if err != nil {
		return nil, fmt.Errorf("could not get fwrev for controller %s: %w", id, err)
	}

	controller.Serial, err = t.GetStringValue(id, mib.CpqDaCntlrSerialNumber)
	if err != nil {
		return nil, fmt.Errorf("could not get serial for controller %s: %w", id, err)
	}

	statusI, err := t.GetIntValue(id, mib.CpqDaCntlrBoardStatus)
	if err != nil {
		return nil, fmt.Errorf("could not get status for controller %s: %w", id, err)
	}

	if status, ok := mib.CpqDaCntlrBoardStatusMap[statusI]; ok {
		controller.Status = status
	} else {
		return nil, fmt.Errorf("invalid status for controller: %s: %d", id, statusI)
	}

	return controller, nil
}

func GetControllersFromTable(t *CpqDaCntlrTable) ([]*Controller, error) {
	controllers := make([]*Controller, 0, len(t.ListIds()))

	for _, id := range t.ListIds() {
		controller, err := NewControllerFromTable(t, id)
		if err != nil {
			return nil, err
		}

		controllers = append(controllers, controller)
	}

	return controllers, nil
}

// GetNagiosStatus validates the Controller's data against the known models
// in this plugin.
func (d *Controller) GetNagiosStatus() (int, string) {
	description := fmt.Sprintf("controller (%s) model=%s serial=%s firmware=%s",
		d.ID, d.Model, strings.TrimSpace(d.Serial), d.FwRev)

	if d.Status != "ok" {
		return check.Critical, description + " - status: " + d.Status
	}

	if _, affected := AffectedModels[d.Model]; affected {
		status, info := IsAffected(d.FwRev)
		return status, description + " - " + info
	}

	return check.OK, description
}
