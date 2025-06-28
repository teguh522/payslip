package container

import (
	attendanceUseCase "github.com/teguh522/payslip/cmd/internal/application/attendance/usecase"
	overtimeUseCase "github.com/teguh522/payslip/cmd/internal/application/overtime/usecase"
	userUseCase "github.com/teguh522/payslip/cmd/internal/application/user/usecase"
	"github.com/teguh522/payslip/cmd/internal/infrastucture/http/handler"
	"github.com/teguh522/payslip/cmd/internal/infrastucture/http/middleware"
	repo "github.com/teguh522/payslip/cmd/internal/infrastucture/persistence/postgres"
	"github.com/teguh522/payslip/cmd/internal/pkg/config"
	"gorm.io/gorm"
)

type Repositories struct {
	UserRepository             *repo.UserRepositoryImp
	AttendancePeriodRepository *repo.AttendancePeriodImp
	AttendanceRepository       *repo.AttendanceRepositoryImp
	OvertimeRepository         *repo.OvertimeRepoImp
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		UserRepository:             repo.NewUserRepositoryImp(db),
		AttendancePeriodRepository: repo.NewAttendancePeriodImp(db),
		AttendanceRepository:       repo.NewAttendanceRepositoryImp(db),
		OvertimeRepository:         repo.NewOvertimeRepoImp(db),
	}
}

type UseCases struct {
	CreateUserUseCase             *userUseCase.CreateUserUseCase
	LoginUserUseCase              *userUseCase.LoginUserUseCase
	CreateAttendancePeriodUseCase *attendanceUseCase.CreateAttendancePeriodUseCase
	CreateAttendanceUseCase       *attendanceUseCase.AttendanceUseCase
	CreateAttendance              *attendanceUseCase.AttendanceUseCase
	OvertimeUseCase               *overtimeUseCase.OvertimeUseCase
}

func NewUseCases(repos *Repositories, cfg *config.Config) *UseCases {
	return &UseCases{
		CreateUserUseCase:             userUseCase.NewCreateUserUseCase(repos.UserRepository),
		LoginUserUseCase:              userUseCase.NewLoginUserUseCase(repos.UserRepository, cfg),
		CreateAttendancePeriodUseCase: attendanceUseCase.NewCreateAttendancePeriodUseCase(repos.AttendancePeriodRepository),
		CreateAttendance:              attendanceUseCase.NewAttendanceUseCase(repos.AttendanceRepository),
		OvertimeUseCase:               overtimeUseCase.NewOvertimeUseCase(repos.OvertimeRepository),
	}
}

type Handlers struct {
	UserHandler             *handler.UserHandler
	AttendancePeriodHandler *handler.AttendancePeriodHandler
	AttendanceHandler       *handler.AttendanceHandler
	OvertimeHandler         *handler.OvertimeHandler
}

func NewHandlers(useCases *UseCases) *Handlers {
	return &Handlers{
		UserHandler:             handler.NewUserHandler(useCases.CreateUserUseCase, useCases.LoginUserUseCase),
		AttendancePeriodHandler: handler.NewAttendancePeriodHandler(useCases.CreateAttendancePeriodUseCase),
		AttendanceHandler:       handler.NewAttendanceHandler(useCases.CreateAttendance),
		OvertimeHandler:         handler.NewOvertimeHandler(useCases.OvertimeUseCase),
	}
}

type Middlewares struct {
	AuthMiddleware *middleware.AuthMiddleware
}

func NewMiddlewares(cfg *config.Config) *Middlewares {
	return &Middlewares{
		AuthMiddleware: middleware.NewAuthMiddleware(cfg),
	}
}
