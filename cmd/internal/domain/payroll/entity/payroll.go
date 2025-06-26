package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/teguh522/payslip/cmd/internal/domain/payslip/entity"
)

type Payroll struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	BaseSalary     float64
	OvertimeAmount float64
	Reimbursement  float64
	TotalSalary    float64
	GeneratedAt    time.Time
	CreatedBy      string         `gorm:"not null"`
	UpdatedBy      string         `gorm:"not null"`
	CreatedAt      time.Time      `gorm:"autoCreateTime"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime"`
	EmployeeID     uuid.UUID      `gorm:"type:uuid"`
	PeriodID       uuid.UUID      `gorm:"type:uuid"`
	Payslip        entity.Payslip `gorm:"foreignKey:PayrollID"`
}
