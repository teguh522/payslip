package entity

import (
	"time"

	"github.com/google/uuid"
	employee "github.com/teguh522/payslip/cmd/internal/domain/employee/entity"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Username     string    `gorm:"unique;not null"`
	PasswordHash string    `gorm:"not null"`
	Role         string    `gorm:"not null"`
	CreatedBy    string    `gorm:"not null"`
	UpdatedBy    string    `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`

	Employee employee.Employee `gorm:"foreignKey:UserID"`
}

func NewUser(username, passwordHash, createdBy, UpdatedBy, role string) (*User, error) {
	return &User{
		Username:     username,
		PasswordHash: passwordHash,
		Role:         role,
		CreatedBy:    createdBy,
		UpdatedBy:    UpdatedBy,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, nil
}
