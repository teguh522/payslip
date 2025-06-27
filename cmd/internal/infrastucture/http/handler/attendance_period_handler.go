package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/teguh522/payslip/cmd/internal/application/attendance/command"
	"github.com/teguh522/payslip/cmd/internal/application/attendance/dto"
	"github.com/teguh522/payslip/cmd/internal/application/attendance/usecase"
)

type AttendancePeriodHandler struct {
	createAttendancePeriodUseCase *usecase.CreateAttendancePeriodUseCase
	validator                     *validator.Validate
}

func NewAttendancePeriodHandler(createAttendancePeriodUseCase *usecase.CreateAttendancePeriodUseCase) *AttendancePeriodHandler {
	return &AttendancePeriodHandler{
		createAttendancePeriodUseCase: createAttendancePeriodUseCase,
		validator:                     validator.New(),
	}
}

func (h *AttendancePeriodHandler) CreateAttendancePeriod(c *gin.Context) {
	var req dto.CreateAttendancePeriodRequest
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

	cmd := command.NewCreateAttendancePeriodCommand(req.StartDate, req.EndDate, req.CreatedBy, req.UpdatedBy)
	resp, err := h.createAttendancePeriodUseCase.Execute(c.Request.Context(), *cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	c.JSON(http.StatusCreated, resp)
}
