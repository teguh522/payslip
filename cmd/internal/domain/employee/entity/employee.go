package entity

import (
	"time"

	"github.com/google/uuid"
	attendance "github.com/teguh522/payslip/cmd/internal/domain/attendance/entity"
	overtime "github.com/teguh522/payslip/cmd/internal/domain/overtime/entity"
	payRoll "github.com/teguh522/payslip/cmd/internal/domain/payroll/entity"
	paySlip "github.com/teguh522/payslip/cmd/internal/domain/payslip/entity"
	reimbursement "github.com/teguh522/payslip/cmd/internal/domain/reimbursement/entity"
)

type Employee struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string
	Salary    float64
	UserID    uuid.UUID
	CreatedBy string    `gorm:"not null"`
	UpdatedBy string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Attendances    []attendance.Attendance       `gorm:"foreignKey:EmployeeID"`
	Overtimes      []overtime.Overtime           `gorm:"foreignKey:EmployeeID"`
	Reimbursements []reimbursement.Reimbursement `gorm:"foreignKey:EmployeeID"`
	Payrolls       []payRoll.Payroll             `gorm:"foreignKey:EmployeeID"`
	Payslips       []paySlip.Payslip             `gorm:"foreignKey:EmployeeID"`
}
