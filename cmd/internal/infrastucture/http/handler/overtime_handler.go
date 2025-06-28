package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/teguh522/payslip/cmd/internal/application/overtime/command"
	"github.com/teguh522/payslip/cmd/internal/application/overtime/dto"
	"github.com/teguh522/payslip/cmd/internal/application/overtime/usecase"
)

type OvertimeHandler struct {
	createAttendanceUsecase *usecase.OvertimeUseCase
	validator               *validator.Validate
}

func NewOvertimeHandler(createAttendanceUseCase *usecase.OvertimeUseCase) *OvertimeHandler {
	return &OvertimeHandler{
		createAttendanceUsecase: createAttendanceUseCase,
		validator:               validator.New(),
	}
}
func (h *OvertimeHandler) CreateOvertime(c *gin.Context) {
	var req dto.CreateOvertimeRequest
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

	if req.Hours > 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "overtime max 3 jam"})
		return
	}

	cmd := command.NewOvertimeCommand(req.EmployeeID, req.PeriodID, req.Date, req.Hours, req.Description, req.CreatedBy, req.UpdatedBy)
	resp, err := h.createAttendanceUsecase.Execute(c.Request.Context(), *cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp)
}
