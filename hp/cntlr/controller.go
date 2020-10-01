package cntlr

import (
	"fmt"
	"github.com/NETWAYS/check_hp_firmware/hp/mib"
	"github.com/NETWAYS/check_hp_firmware/nagios"
	"strings"
)

type Controller struct {
	Id     string
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
	controller.Id = id

	modelI, err := t.GetIntValue(id, mib.CpqDaCntlrModel)
	if err != nil {
		return nil, fmt.Errorf("could not get model for controller %s: %s", id, err)
	}
	if model, ok := mib.CpqDaCntlrModelMap[modelI]; ok {
		controller.Model = model
	} else {
		return nil, fmt.Errorf("unknown model int for controller: %s: %d", id, modelI)
	}

	controller.FwRev, err = t.GetStringValue(id, mib.CpqDaCntlrFWRev)
	if err != nil {
		return nil, fmt.Errorf("could not get fwrev for controller %s: %s", id, err)
	}

	controller.Serial, err = t.GetStringValue(id, mib.CpqDaCntlrSerialNumber)
	if err != nil {
		return nil, fmt.Errorf("could not get serial for controller %s: %s", id, err)
	}

	statusI, err := t.GetIntValue(id, mib.CpqDaCntlrBoardStatus)
	if err != nil {
		return nil, fmt.Errorf("could not get status for controller %s: %s", id, err)
	}
	if status, ok := mib.CpqDaCntlrBoardStatusMap[statusI]; ok {
		controller.Status = status
	} else {
		return nil, fmt.Errorf("invalid status for controller: %s: %d", id, statusI)
	}

	return controller, nil
}

func GetControllersFromTable(t *CpqDaCntlrTable) ([]*Controller, error) {
	ids := t.ListIds()
	var controllers []*Controller

	for _, id := range ids {
		controller, err := NewControllerFromTable(t, id)
		if err != nil {
			return nil, err
		}
		controllers = append(controllers, controller)
	}

	return controllers, nil
}

func (d *Controller) GetNagiosStatus() (int, string) {
	description := fmt.Sprintf("controller (%s) model=%s serial=%s firmware=%s",
		d.Id, d.Model, strings.TrimSpace(d.Serial), d.FwRev)

	if d.Status != "ok" {
		return nagios.Critical, description + " - status: " + d.Status
	}

	if _, affected := AffectedModels[d.Model]; affected {
		status, info := IsAffected(d.FwRev)
		return status, description + " - " + info
	}

	return nagios.OK, description
}
