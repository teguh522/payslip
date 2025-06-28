package command

import (
	"github.com/google/uuid"
	"github.com/teguh522/payslip/cmd/internal/pkg/helper"
)

type AttendanceCommand struct {
	Date       helper.DateOnly
	CheckIn    string
	CreatedBy  string
	UpdatedBy  string
	EmployeeID uuid.UUID
	PeriodID   uuid.UUID
}

func NewAttendanceCommand(date helper.DateOnly, checkIn, createdBy, updatedBy string, employeeID, periodID uuid.UUID) *AttendanceCommand {
	return &AttendanceCommand{
		Date:       date,
		CheckIn:    checkIn,
		CreatedBy:  createdBy,
		UpdatedBy:  updatedBy,
		EmployeeID: employeeID,
		PeriodID:   periodID,
	}
}

type AttendanceCheckOutCommand struct {
	Date       helper.DateOnly
	Checkout   string
	UpdatedBy  string
	EmployeeID uuid.UUID
	PeriodID   uuid.UUID
}

func NewAttendanceCheckOutCommand(date helper.DateOnly, checkout, updatedBy string, employeeID, periodID uuid.UUID) *AttendanceCheckOutCommand {
	return &AttendanceCheckOutCommand{
		Date:       date,
		Checkout:   checkout,
		UpdatedBy:  updatedBy,
		EmployeeID: employeeID,
		PeriodID:   periodID,
	}
}
