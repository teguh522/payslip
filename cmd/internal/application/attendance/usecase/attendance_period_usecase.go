package usecase

import (
	"context"

	"github.com/teguh522/payslip/cmd/internal/application/attendance/command"
	"github.com/teguh522/payslip/cmd/internal/application/attendance/dto"
	"github.com/teguh522/payslip/cmd/internal/domain/attendance/entity"
	"github.com/teguh522/payslip/cmd/internal/domain/attendance/repository"
)

type CreateAttendancePeriodUseCase struct {
	attendancePeriodRepo repository.AttendancePeriodRepository
}

func NewCreateAttendancePeriodUseCase(attendancePeriodRepo repository.AttendancePeriodRepository) *CreateAttendancePeriodUseCase {
	return &CreateAttendancePeriodUseCase{
		attendancePeriodRepo: attendancePeriodRepo,
	}
}

func (uc *CreateAttendancePeriodUseCase) Execute(ctx context.Context, cmd command.CreateAttendancePeriodCommand) (*dto.CreateAttendancePeriodResponse, error) {
	attendancePeriod, err := entity.NewAttendancePeriod(cmd.StartDate, cmd.EndDate, cmd.CreatedBy, cmd.UpdatedBy)
	if err != nil {
		return nil, err
	}

	err = uc.attendancePeriodRepo.CreateAttendancePeriod(ctx, attendancePeriod)
	if err != nil {
		return nil, err
	}

	return &dto.CreateAttendancePeriodResponse{
		ID:        attendancePeriod.ID,
		StartDate: attendancePeriod.StartDate,
		EndDate:   attendancePeriod.EndDate,
		Status:    attendancePeriod.Status,
		CreatedBy: attendancePeriod.CreatedBy,
		UpdatedBy: attendancePeriod.UpdatedBy,
	}, nil
}
