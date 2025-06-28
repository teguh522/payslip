package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/teguh522/payslip/cmd/internal/pkg/helper"
)

type Attendance struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Date       helper.DateOnly
	CheckIn    string
	CheckOut   string
	CreatedBy  string    `gorm:"not null"`
	UpdatedBy  string    `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
	EmployeeID uuid.UUID `gorm:"type:uuid"`
	PeriodID   uuid.UUID `gorm:"type:uuid"`
}

func NewAttendance(date helper.DateOnly, checkin, checkout, createdby, updatedby, employeeid, periodid string) (*Attendance, error) {
	return &Attendance{
		Date:       date,
		CheckIn:    checkin,
		CheckOut:   checkout,
		CreatedBy:  createdby,
		UpdatedBy:  updatedby,
		EmployeeID: uuid.MustParse(employeeid),
		PeriodID:   uuid.MustParse(periodid),
	}, nil
}
