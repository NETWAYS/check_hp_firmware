package phy_drv

import (
	"fmt"
	"github.com/NETWAYS/check_hp_firmware/hp/mib"
	"github.com/NETWAYS/check_hp_firmware/nagios"
)

type PhysicalDrive struct {
	Id     string
	Model  string
	FwRev  string
	Serial string
	Status string
	Hours  uint
}

func NewPhysicalDriveFromTable(t *CpqDaPhyDrvTable, id string) (*PhysicalDrive, error) {
	_, ok := t.Snmp.Values[id]
	if !ok {
		return nil, fmt.Errorf("could not find drive %s in table", id)
	}

	var err error
	drive := &PhysicalDrive{}
	drive.Id = id

	drive.Model, err = t.GetStringValue(id, mib.CpqDaPhyDrvModel)
	if err != nil {
		return nil, fmt.Errorf("could not get model for drive %s: %s", id, err)
	}

	drive.FwRev, err = t.GetStringValue(id, mib.CpqDaPhyDrvFWRev)
	if err != nil {
		return nil, fmt.Errorf("could not get fwrev for drive %s: %s", id, err)
	}

	drive.Serial, err = t.GetStringValue(id, mib.CpqDaPhyDrvSerialNum)
	if err != nil {
		return nil, fmt.Errorf("could not get serial for drive %s: %s", id, err)
	}

	statusI, err := t.GetIntValue(id, mib.CpqDaPhyDrvStatus)
	if err != nil {
		return nil, fmt.Errorf("could not get status for drive %s: %s", id, err)
	}
	if status, ok := mib.CpqDaPhyDrvStatusMap[statusI]; ok {
		drive.Status = status
	} else {
		return nil, fmt.Errorf("invalid status for drive: %s: %s", id, status)
	}

	drive.Hours, err = t.GetUintValue(id, mib.CpqDaPhyDrvRefHours)
	if err != nil {
		return nil, fmt.Errorf("could not get hours for drive %s: %s", id, err)
	}

	return drive, nil
}

func GetPhysicalDrivesFromTable(t *CpqDaPhyDrvTable) ([]*PhysicalDrive, error) {
	ids := t.ListIds()
	var drives []*PhysicalDrive

	for _, id := range ids {
		drive, err := NewPhysicalDriveFromTable(t, id)
		if err != nil {
			return nil, err
		}
		drives = append(drives, drive)
	}

	return drives, nil
}

func (d *PhysicalDrive) GetNagiosStatus() (int, string) {
	description := fmt.Sprintf("physical drive (%-4s) model=%s serial=%s firmware=%s hours=%d",
		d.Id, d.Model, d.Serial, d.FwRev, d.Hours)

	if d.Status != "ok" {
		return nagios.Critical, description + " - status: " + d.Status
	}

	if model, affected := AffectedModels[d.Model]; affected {
		ok, err := IsFirmwareFixed(model, d.FwRev)
		if err != nil {
			return nagios.Unknown, description + " - error: " + err.Error()
		} else if ok {
			return nagios.OK, description + " - firmware update applied"
		} else {
			return nagios.Critical, description + " - affected by FW bug"
		}
	}

	return nagios.OK, description
}
