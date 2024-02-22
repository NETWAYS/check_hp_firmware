package drive

import (
	"fmt"

	"github.com/NETWAYS/check_hp_firmware/hp/mib"
	"github.com/NETWAYS/go-check"
)

type PhysicalDrive struct {
	ID     string
	Model  string
	FwRev  string
	Serial string
	Status string
	Hours  uint
}

func NewPhysicalDriveFromTable(t *CpqDaPhyDrvTable, id string) (*PhysicalDrive, error) {
	if _, ok := t.Snmp.Values[id]; !ok {
		return nil, fmt.Errorf("could not find drive %s in table", id)
	}

	var err error

	drive := &PhysicalDrive{
		ID: id,
	}

	drive.Model, err = t.GetStringValue(id, mib.CpqDaPhyDrvModel)
	if err != nil {
		return nil, fmt.Errorf("could not get model for drive %s: %w", id, err)
	}

	drive.FwRev, err = t.GetStringValue(id, mib.CpqDaPhyDrvFWRev)
	if err != nil {
		return nil, fmt.Errorf("could not get fwrev for drive %s: %w", id, err)
	}

	drive.Serial, err = t.GetStringValue(id, mib.CpqDaPhyDrvSerialNum)
	if err != nil {
		return nil, fmt.Errorf("could not get serial for drive %s: %w", id, err)
	}

	statusI, err := t.GetIntValue(id, mib.CpqDaPhyDrvStatus)
	if err != nil {
		return nil, fmt.Errorf("could not get status for drive %s: %w", id, err)
	}

	if status, ok := mib.CpqDaPhyDrvStatusMap[statusI]; ok {
		drive.Status = status
	} else {
		return nil, fmt.Errorf("invalid status for drive: %s: %s", id, status)
	}

	drive.Hours, err = t.GetUintValue(id, mib.CpqDaPhyDrvRefHours)
	if err != nil {
		return nil, fmt.Errorf("could not get hours for drive %s: %w", id, err)
	}

	return drive, nil
}

func GetPhysicalDrivesFromTable(t *CpqDaPhyDrvTable) ([]*PhysicalDrive, error) {
	drives := make([]*PhysicalDrive, 0, len(t.ListIds()))

	for _, id := range t.ListIds() {
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
		d.ID, d.Model, d.Serial, d.FwRev, d.Hours)

	if d.Status != "ok" {
		return check.Critical, description + " - status: " + d.Status
	}

	if model, affected := AffectedModels[d.Model]; affected {
		ok, err := IsFirmwareFixed(model, d.FwRev)
		// nolint: gocritic
		if err != nil {
			return check.Unknown, description + " - error: " + err.Error()
		} else if ok {
			return check.OK, description + " - firmware update applied"
		} else {
			return check.Critical, description + " - affected by FW bug"
		}
	}

	return check.OK, description
}
