package router

import (
	"github.com/gin-gonic/gin"
	"github.com/teguh522/payslip/cmd/internal/container"
	"github.com/teguh522/payslip/cmd/internal/pkg/config"
)

func NewRouter(userHandler *container.Handlers, cfg *config.Config) *gin.Engine {
	if cfg.App.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", userHandler.UserHandler.CreateUser)
	}

	return r
}
