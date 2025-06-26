package entity

import (
	"time"

	"github.com/google/uuid"
)

type Overtime struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Date        time.Time
	Hours       float64
	Description string
	CreatedBy   string    `gorm:"not null"`
	UpdatedBy   string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	EmployeeID  uuid.UUID `gorm:"type:uuid"`
	PeriodID    uuid.UUID `gorm:"type:uuid"`
}
