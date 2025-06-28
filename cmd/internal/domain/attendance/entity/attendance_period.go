package entity

import (
	"time"

	"github.com/google/uuid"
	overtime "github.com/teguh522/payslip/cmd/internal/domain/overtime/entity"
	payRoll "github.com/teguh522/payslip/cmd/internal/domain/payroll/entity"
	reimbursement "github.com/teguh522/payslip/cmd/internal/domain/reimbursement/entity"
	"github.com/teguh522/payslip/cmd/internal/pkg/helper"
)

type AttendancePeriod struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	StartDate      helper.DateOnly
	EndDate        helper.DateOnly
	Status         string
	CreatedBy      string                        `gorm:"not null"`
	UpdatedBy      string                        `gorm:"not null"`
	CreatedAt      time.Time                     `gorm:"autoCreateTime"`
	UpdatedAt      time.Time                     `gorm:"autoUpdateTime"`
	Attendances    []Attendance                  `gorm:"foreignKey:PeriodID"`
	Overtimes      []overtime.Overtime           `gorm:"foreignKey:PeriodID"`
	Reimbursements []reimbursement.Reimbursement `gorm:"foreignKey:PeriodID"`
	Payrolls       []payRoll.Payroll             `gorm:"foreignKey:PeriodID"`
}

func NewAttendancePeriod(startDate, endDate helper.DateOnly, createdBy, updatedBy string) (*AttendancePeriod, error) {
	return &AttendancePeriod{
		StartDate: startDate,
		EndDate:   endDate,
		Status:    "open",
		CreatedBy: createdBy,
		UpdatedBy: updatedBy,
	}, nil
}
