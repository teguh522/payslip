package usecase

import (
	"context"

	"github.com/teguh522/payslip/cmd/internal/application/attendance/command"
	"github.com/teguh522/payslip/cmd/internal/application/attendance/dto"
	"github.com/teguh522/payslip/cmd/internal/domain/attendance/entity"
	"github.com/teguh522/payslip/cmd/internal/domain/attendance/repository"
)

type AttendanceUseCase struct {
	attendanceRepo repository.AttendanceRepository
}

func NewAttendanceUseCase(attendanceRepo repository.AttendanceRepository) *AttendanceUseCase {
	return &AttendanceUseCase{
		attendanceRepo: attendanceRepo,
	}
}

func (uc *AttendanceUseCase) Execute(ctx context.Context, cmd command.AttendanceCommand) (*dto.CreateAttendanceResponse, error) {
	attendance, err := entity.NewAttendance(cmd.Date, cmd.CheckIn, cmd.CheckOut, cmd.CreatedBy,
		cmd.UpdatedBy, cmd.EmployeeID.String(), cmd.PeriodID.String())
	if err != nil {
		return nil, err
	}

	err = uc.attendanceRepo.CreateAttendance(ctx, attendance)
	if err != nil {
		return nil, err
	}

	return &dto.CreateAttendanceResponse{
		ID:     attendance.ID.String(),
		Date:   attendance.Date.Format("2006-01-02"),
		Status: "success",
	}, nil
}
