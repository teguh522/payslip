package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/teguh522/payslip/cmd/internal/application/user/command"
	"github.com/teguh522/payslip/cmd/internal/application/user/dto"
	"github.com/teguh522/payslip/cmd/internal/application/user/usecase"
)

type UserHandler struct {
	createUserUseCase *usecase.CreateUserUseCase
	validator         *validator.Validate
}

func NewUserHandler(createUserUseCase *usecase.CreateUserUseCase) *UserHandler {
	return &UserHandler{
		createUserUseCase: createUserUseCase,
		validator:         validator.New(),
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) { // Menggunakan *gin.Context
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil { // Menggunakan ShouldBindJSON untuk parsing dan validasi
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validator.Struct(req); err != nil {
		// Tangani error validasi dari validator
		validationErrors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors[err.Field()] = fmt.Sprintf("Field '%s' is not valid (tag: %s)", err.Field(), err.Tag())
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "details": validationErrors})
		return
	}

	// Buat command dari request DTO
	cmd := command.NewCreateUserCommand(req.Username, req.Password, req.Role, req.CreatedBy, req.UpdatedBy)
	// Panggil use case
	resp, err := h.createUserUseCase.Execute(c.Request.Context(), *cmd) // Dereference cmd to pass as value
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	c.JSON(http.StatusCreated, resp) // Menggunakan c.JSON untuk mengirim respons JSON
}
