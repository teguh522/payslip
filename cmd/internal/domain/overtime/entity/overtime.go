package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/teguh522/payslip/cmd/internal/pkg/helper"
)

type Overtime struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Date        helper.DateOnly
	Hours       float64
	Description string
	CreatedBy   string    `gorm:"not null"`
	UpdatedBy   string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	EmployeeID  uuid.UUID `gorm:"type:uuid"`
	PeriodID    uuid.UUID `gorm:"type:uuid"`
}

func NewOvertime(date helper.DateOnly, hours float64, description, createdBy, updatedBy, employeeID, periodID string) (*Overtime, error) {
	if hours <= 0 {
		return nil, errors.New("hours must be greater than zero")
	}

	return &Overtime{
		Date:        date,
		Hours:       hours,
		Description: description,
		CreatedBy:   createdBy,
		UpdatedBy:   updatedBy,
		EmployeeID:  uuid.MustParse(employeeID),
		PeriodID:    uuid.MustParse(periodID),
	}, nil
}
