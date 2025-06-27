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
	loginUserUseCase  *usecase.LoginUserUseCase
}

func NewUserHandler(createUserUseCase *usecase.CreateUserUseCase, loginUserUseCase *usecase.LoginUserUseCase) *UserHandler {
	return &UserHandler{
		createUserUseCase: createUserUseCase,
		validator:         validator.New(),
		loginUserUseCase:  loginUserUseCase,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest
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

	cmd := command.NewCreateUserCommand(req.Username, req.Password, req.Role, req.CreatedBy, req.UpdatedBy)
	resp, err := h.createUserUseCase.Execute(c.Request.Context(), *cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	c.JSON(http.StatusCreated, resp)
}
func (h *UserHandler) LoginUser(c *gin.Context) {
	var req dto.LoginUserRequest
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

	cmd := command.NewLoginUserCommand(req.Username, req.Password)
	resp, err := h.loginUserUseCase.Execute(c.Request.Context(), *cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
