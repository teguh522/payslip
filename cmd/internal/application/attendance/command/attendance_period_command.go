package command

import "github.com/teguh522/payslip/cmd/internal/pkg/helper"

type CreateAttendancePeriodCommand struct {
	StartDate helper.DateOnly
	EndDate   helper.DateOnly
	CreatedBy string
	UpdatedBy string
}

func NewCreateAttendancePeriodCommand(startDate, endDate helper.DateOnly, createdBy, updatedBy string) *CreateAttendancePeriodCommand {
	return &CreateAttendancePeriodCommand{
		StartDate: startDate,
		EndDate:   endDate,
		CreatedBy: createdBy,
		UpdatedBy: updatedBy,
	}
}
