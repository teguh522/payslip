package postgres

import (
	"context"

	"github.com/teguh522/payslip/cmd/internal/domain/attendance/entity"
	"gorm.io/gorm"
)

type AttendanceRepositoryImp struct {
	db *gorm.DB
}

func NewAttendanceRepositoryImp(db *gorm.DB) *AttendanceRepositoryImp {
	return &AttendanceRepositoryImp{
		db: db,
	}
}
func (repo *AttendanceRepositoryImp) CreateAttendance(ctx context.Context, attendance *entity.Attendance) error {
	if err := repo.db.WithContext(ctx).Save(attendance).Error; err != nil {
		return err
	}
	return nil
}
