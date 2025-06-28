package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/teguh522/payslip/cmd/internal/application/attendance/command"
	"github.com/teguh522/payslip/cmd/internal/application/attendance/dto"
	"github.com/teguh522/payslip/cmd/internal/application/attendance/usecase"
)

type AttendanceHandler struct {
	createAttendanceUsecase *usecase.AttendanceUseCase
	validator               *validator.Validate
}

func NewAttendanceHandler(createAttendanceUseCase *usecase.AttendanceUseCase) *AttendanceHandler {
	return &AttendanceHandler{
		createAttendanceUsecase: createAttendanceUseCase,
		validator:               validator.New(),
	}
}

func (h *AttendanceHandler) CreateAttendance(c *gin.Context) {
	var req dto.CreateAttendanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validator.Struct(req); err != nil {
		validationErrors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors[err.Field()] = fmt.Sprintf("Field '%s' is not valid (tag: %s)", err.Field(), err.Tag())
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "details": validationErrors})
		return
	}

	if req.Date.Weekday() == time.Saturday || req.Date.Weekday() == time.Sunday {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Attendance cannot be checked out on weekends"})
		return
	}

	cmd := command.NewAttendanceCommand(req.Date, req.CheckIn, req.CreatedBy, req.UpdatedBy, req.EmployeeID, req.PeriodID)
	resp, err := h.createAttendanceUsecase.Execute(c.Request.Context(), *cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	c.JSON(http.StatusCreated, resp)
}

func (h *AttendanceHandler) CreateAttendanceCheckOut(c *gin.Context) {
	var req dto.AttendanceCheckOutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.validator.Struct(req); err != nil {
		validationErrors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors[err.Field()] = fmt.Sprintf("Field '%s' is not valid (tag: %s)", err.Field(), err.Tag())
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "details": validationErrors})
		return
	}
	cmd := command.NewAttendanceCheckOutCommand(req.Date, req.Checkout, req.UpdatedBy, req.EmployeeID, req.PeriodID)
	resp, err := h.createAttendanceUsecase.ExecuteCheckOut(c.Request.Context(), *cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
