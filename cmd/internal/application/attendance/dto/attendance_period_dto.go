package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateAttendancePeriodRequest struct {
	StartDate time.Time `json:"start_date" validate:"required"`
	EndDate   time.Time `json:"end_date" validate:"required"`
	CreatedBy string    `json:"created_by" validate:"required"`
	UpdatedBy string    `json:"updated_by" validate:"required"`
}

type CreateAttendancePeriodResponse struct {
	ID        uuid.UUID `json:"id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Status    string    `json:"status"`
	CreatedBy string    `json:"created_by"`
	UpdatedBy string    `json:"updated_by"`
}
