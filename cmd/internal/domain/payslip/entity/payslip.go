package entity

import (
	"time"

	"github.com/google/uuid"
)

type Payslip struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	FileURL    string
	CreatedBy  string    `gorm:"not null"`
	UpdatedBy  string    `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
	EmployeeID uuid.UUID `gorm:"type:uuid"`
	PayrollID  uuid.UUID `gorm:"type:uuid"`
}
