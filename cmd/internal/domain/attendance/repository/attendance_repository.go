package repository

import (
	"context"

	"github.com/teguh522/payslip/cmd/internal/domain/attendance/entity"
)

type AttendanceRepository interface {
	CreateAttendance(ctx context.Context, attendance *entity.Attendance) error
	CreateAttendanceCheckOut(ctx context.Context, attendance *entity.Attendance) error
}
