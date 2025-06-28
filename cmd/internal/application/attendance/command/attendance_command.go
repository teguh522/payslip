package command

import (
	"github.com/google/uuid"
	"github.com/teguh522/payslip/cmd/internal/pkg/helper"
)

type AttendanceCommand struct {
	Date       helper.DateOnly
	CheckIn    string
	CheckOut   string
	CreatedBy  string
	UpdatedBy  string
	EmployeeID uuid.UUID
	PeriodID   uuid.UUID
}

func NewAttendanceCommand(date helper.DateOnly, checkIn, checkOut, createdBy, updatedBy string, employeeID, periodID uuid.UUID) *AttendanceCommand {
	return &AttendanceCommand{
		Date:       date,
		CheckIn:    checkIn,
		CheckOut:   checkOut,
		CreatedBy:  createdBy,
		UpdatedBy:  updatedBy,
		EmployeeID: employeeID,
		PeriodID:   periodID,
	}
}
