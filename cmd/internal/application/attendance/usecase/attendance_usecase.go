package usecase

import (
	"context"
	"errors"

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
	attendance, err := entity.NewAttendance(cmd.Date, cmd.CheckIn, cmd.CreatedBy,
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
		Status: "checked in successfully",
	}, nil
}

func (uc *AttendanceUseCase) ExecuteCheckOut(ctx context.Context, cmd command.AttendanceCheckOutCommand) (*dto.AttendanceCheckOutResponse, error) {
	attendance, err := entity.NewAttendanceCheckOut(cmd.Date, cmd.Checkout, cmd.UpdatedBy, cmd.EmployeeID.String(), cmd.PeriodID.String())
	if err != nil {
		return nil, err
	}

	err = uc.attendanceRepo.CreateAttendanceCheckOut(ctx, attendance)
	if err != nil {
		return nil, errors.New("failed to check out attendance: pastikan sudah melakukan check-in sebelumnya")
	}

	return &dto.AttendanceCheckOutResponse{
		Status: "checked out successfully",
	}, nil
}
