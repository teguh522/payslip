package dto

import "github.com/teguh522/payslip/cmd/internal/pkg/helper"

type CreateOvertimeRequest struct {
	EmployeeID  string          `json:"employee_id" validate:"required"`
	PeriodID    string          `json:"period_id" validate:"required"`
	Date        helper.DateOnly `json:"date" validate:"required"`
	Hours       float64         `json:"hours" validate:"required,gte=0"`
	Description string          `json:"description" validate:"required"`
	CreatedBy   string          `json:"created_by" validate:"required"`
	UpdatedBy   string          `json:"updated_by" validate:"required"`
}

type CreateOvertimeResponse struct {
	Status string `json:"status"`
}
