package command

import "github.com/teguh522/payslip/cmd/internal/pkg/helper"

type OvertimeCommand struct {
	EmployeeID  string
	PeriodID    string
	Date        helper.DateOnly
	Hours       float64
	Description string
	CreatedBy   string
	UpdatedBy   string
}

func NewOvertimeCommand(employeeID, periodID string, date helper.DateOnly, hours float64, description, createdBy, updatedBy string) *OvertimeCommand {
	return &OvertimeCommand{
		EmployeeID:  employeeID,
		PeriodID:    periodID,
		Date:        date,
		Hours:       hours,
		Description: description,
		CreatedBy:   createdBy,
		UpdatedBy:   updatedBy,
	}
}
