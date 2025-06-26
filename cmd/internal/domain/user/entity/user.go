package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Username     string    `gorm:"unique;not null"`
	PasswordHash string    `gorm:"not null"`
	Role         string    `gorm:"not null"`
	CreatedBy    string    `gorm:"not null"`
	UpdatedBy    string    `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

func NewUser(username, passwordHash, createdBy, UpdatedBy, role string) (*User, error) {
	return &User{
		ID:           uuid.New(),
		Username:     username,
		PasswordHash: passwordHash,
		Role:         role,
		CreatedBy:    createdBy,
		UpdatedBy:    UpdatedBy,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, nil
}
