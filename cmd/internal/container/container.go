package container

import (
	userUseCase "github.com/teguh522/payslip/cmd/internal/application/user/usecase"
	userHandler "github.com/teguh522/payslip/cmd/internal/infrastucture/http/handler"
	userRepo "github.com/teguh522/payslip/cmd/internal/infrastucture/persistence/postgres"
	"gorm.io/gorm"
)

type Repositories struct {
	UserRepository *userRepo.UserRepositoryImp
}

type UseCases struct {
	CreateUserUseCase *userUseCase.CreateUserUseCase
}

type Handlers struct {
	UserHandler *userHandler.UserHandler
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		UserRepository: userRepo.NewUserRepositoryImp(db),
	}
}

func NewUseCases(repos *Repositories) *UseCases {
	return &UseCases{
		CreateUserUseCase: userUseCase.NewCreateUserUseCase(repos.UserRepository),
	}
}

func NewHandlers(useCases *UseCases) *Handlers {
	return &Handlers{
		UserHandler: userHandler.NewUserHandler(useCases.CreateUserUseCase),
	}
}
