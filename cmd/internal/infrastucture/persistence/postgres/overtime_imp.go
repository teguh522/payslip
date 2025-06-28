package postgres

import (
	"context"

	"github.com/teguh522/payslip/cmd/internal/domain/overtime/entity"
	"gorm.io/gorm"
)

type OvertimeRepoImp struct {
	db *gorm.DB
}

func NewOvertimeRepoImp(db *gorm.DB) *OvertimeRepoImp {
	return &OvertimeRepoImp{
		db: db,
	}
}

func (repo *OvertimeRepoImp) CreateOvertime(ctx context.Context, overtime *entity.Overtime) error {
	result := repo.db.WithContext(ctx).Where("employee_id = ? AND period_id = ? AND date = ?", overtime.EmployeeID, overtime.PeriodID, overtime.Date).First(&entity.Overtime{})
	if result.Error == gorm.ErrRecordNotFound {
		if err := repo.db.WithContext(ctx).Save(overtime).Error; err != nil {
			return err
		}
	}
	return nil
}
