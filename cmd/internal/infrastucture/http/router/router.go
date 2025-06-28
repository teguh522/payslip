package router

import (
	"github.com/gin-gonic/gin"
	"github.com/teguh522/payslip/cmd/internal/container"
	"github.com/teguh522/payslip/cmd/internal/pkg/config"
)

func NewRouter(handler *container.Handlers, cfg *config.Config, middlewares *container.Middlewares) *gin.Engine {
	if cfg.App.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	userRoutes := r.Group("/login")
	{
		userRoutes.POST("/", handler.UserHandler.LoginUser)
	}
	authenticatedRoutes := r.Group("/users")
	{
		authenticatedRoutes.Use(middlewares.AuthMiddleware.Authenticate())
		authenticatedRoutes.POST("/", handler.UserHandler.CreateUser)
	}
	attendancePeriodRoutes := r.Group("/attendance")
	{
		attendancePeriodRoutes.Use(middlewares.AuthMiddleware.Authenticate())
		attendancePeriodRoutes.Use(middlewares.AuthMiddleware.RoleUserMiddleware())
		attendancePeriodRoutes.POST("/checkin", handler.AttendanceHandler.CreateAttendance)
		attendancePeriodRoutes.POST("/checkout", handler.AttendanceHandler.CreateAttendanceCheckOut)
		attendancePeriodRoutes.Use(middlewares.AuthMiddleware.RoleAdminMiddleware())
		attendancePeriodRoutes.POST("/periods", handler.AttendancePeriodHandler.CreateAttendancePeriod)
	}

	return r
}
