package command

import "time"

type CreateAttendancePeriodCommand struct {
	StartDate time.Time
	EndDate   time.Time
	CreatedBy string
	UpdatedBy string
}

func NewCreateAttendancePeriodCommand(startDate, endDate time.Time, createdBy, updatedBy string) *CreateAttendancePeriodCommand {
	return &CreateAttendancePeriodCommand{
		StartDate: startDate,
		EndDate:   endDate,
		CreatedBy: createdBy,
		UpdatedBy: updatedBy,
	}
}
