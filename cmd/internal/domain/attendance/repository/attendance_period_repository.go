package repository

import (
	"context"

	"github.com/teguh522/payslip/cmd/internal/domain/attendance/entity"
)

type AttendancePeriodRepository interface {
	CreateAttendancePeriod(ctx context.Context, attendancePeriod *entity.AttendancePeriod) error
}
