package postgres

import (
	"context"

	"github.com/teguh522/payslip/cmd/internal/domain/attendance/entity"
	"gorm.io/gorm"
)

type AttendancePeriodImp struct {
	db *gorm.DB
}

func NewAttendancePeriodImp(db *gorm.DB) *AttendancePeriodImp {
	return &AttendancePeriodImp{
		db: db,
	}
}

func (repo *AttendancePeriodImp) CreateAttendancePeriod(ctx context.Context, attendancePeriod *entity.AttendancePeriod) error {
	if err := repo.db.WithContext(ctx).Save(attendancePeriod).Error; err != nil {
		return err
	}
	return nil
}
