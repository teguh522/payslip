package router

import (
	"github.com/gin-gonic/gin"
	"github.com/teguh522/payslip/cmd/internal/infrastucture/http/handler"
)

func NewRouter(userHandler *handler.UserHandler) *gin.Engine {
	r := gin.Default()

	// Grouping routes for /users
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", userHandler.CreateUser)
	}

	return r
}
