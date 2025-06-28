package dto

import (
	"github.com/google/uuid"
	"github.com/teguh522/payslip/cmd/internal/pkg/helper"
)

type CreateAttendanceRequest struct {
	Date       helper.DateOnly `json:"date" validate:"required"`
	CheckIn    string          `json:"check_in" validate:"required"`
	CreatedBy  string          `json:"created_by" validate:"required"`
	UpdatedBy  string          `json:"updated_by" validate:"required"`
	EmployeeID uuid.UUID       `json:"employee_id" validate:"required"`
	PeriodID   uuid.UUID       `json:"period_id" validate:"required"`
}
type CreateAttendanceResponse struct {
	ID     string `json:"id"`
	Date   string `json:"date"`
	Status string `json:"status"`
}

type AttendanceCheckOutRequest struct {
	Date       helper.DateOnly `json:"date" validate:"required"`
	Checkout   string          `json:"check_out" validate:"required"`
	UpdatedBy  string          `json:"updated_by" validate:"required"`
	EmployeeID uuid.UUID       `json:"employee_id" validate:"required"`
	PeriodID   uuid.UUID       `json:"period_id" validate:"required"`
}
type AttendanceCheckOutResponse struct {
	Status string `json:"status"`
}
