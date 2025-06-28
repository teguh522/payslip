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
	result := repo.db.WithContext(ctx).Where("employee_id = ? AND period_id = ? AND date = ?", attendance.EmployeeID, attendance.PeriodID, attendance.Date).First(&entity.Attendance{})
	if result.Error == gorm.ErrRecordNotFound {
		if err := repo.db.WithContext(ctx).Save(attendance).Error; err != nil {
			return err
		}
	}
	return nil
}
func (repo *AttendanceRepositoryImp) CreateAttendanceCheckOut(ctx context.Context, attendance *entity.Attendance) error {
	updates := map[string]interface{}{
		"check_out":  attendance.CheckOut,
		"updated_by": attendance.UpdatedBy,
	}
	tx := repo.db.WithContext(ctx).Model(&entity.Attendance{}).
		Where("employee_id = ? AND period_id = ? AND date=?", attendance.EmployeeID, attendance.PeriodID, attendance.Date).
		Updates(updates)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
