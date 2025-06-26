package entity

import (
	"time"

	"github.com/google/uuid"
	overtime "github.com/teguh522/payslip/cmd/internal/domain/overtime/entity"
	payRoll "github.com/teguh522/payslip/cmd/internal/domain/payroll/entity"
	reimbursement "github.com/teguh522/payslip/cmd/internal/domain/reimbursement/entity"
)

type AttendancePeriod struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	StartDate      time.Time
	EndDate        time.Time
	Status         string                        // "open" or "closed"
	CreatedBy      string                        `gorm:"not null"`
	UpdatedBy      string                        `gorm:"not null"`
	CreatedAt      time.Time                     `gorm:"autoCreateTime"`
	UpdatedAt      time.Time                     `gorm:"autoUpdateTime"`
	Attendances    []Attendance                  `gorm:"foreignKey:PeriodID"`
	Overtimes      []overtime.Overtime           `gorm:"foreignKey:PeriodID"`
	Reimbursements []reimbursement.Reimbursement `gorm:"foreignKey:PeriodID"`
	Payrolls       []payRoll.Payroll             `gorm:"foreignKey:PeriodID"`
}
