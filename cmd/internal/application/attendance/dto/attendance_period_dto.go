package dto

import (
	"github.com/google/uuid"
	"github.com/teguh522/payslip/cmd/internal/pkg/helper"
)

type CreateAttendancePeriodRequest struct {
	StartDate helper.DateOnly `json:"start_date" validate:"required"`
	EndDate   helper.DateOnly `json:"end_date" validate:"required"`
	CreatedBy string          `json:"created_by" validate:"required"`
	UpdatedBy string          `json:"updated_by" validate:"required"`
}

type CreateAttendancePeriodResponse struct {
	ID        uuid.UUID       `json:"id"`
	StartDate helper.DateOnly `json:"start_date"`
	EndDate   helper.DateOnly `json:"end_date"`
	Status    string          `json:"status"`
	CreatedBy string          `json:"created_by"`
	UpdatedBy string          `json:"updated_by"`
}
