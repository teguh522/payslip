package usecase

import (
	"context"

	"github.com/teguh522/payslip/cmd/internal/application/user/command"
	"github.com/teguh522/payslip/cmd/internal/application/user/dto"
	"github.com/teguh522/payslip/cmd/internal/domain/user/entity"
	"github.com/teguh522/payslip/cmd/internal/domain/user/repository"
	"github.com/teguh522/payslip/cmd/internal/pkg/security"
)

type CreateUserUseCase struct {
	userRepo repository.UserRepository
}

func NewCreateUserUseCase(userRepo repository.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepo: userRepo,
	}
}
func (uc *CreateUserUseCase) Execute(ctx context.Context, cmd command.CreateUserCommand) (*dto.CreateUserResponse, error) {
	hashedPassword, err := security.HashPassword(cmd.Password)
	if err != nil {
		return nil, err
	}
	newUser, err := entity.NewUser(cmd.Username, hashedPassword, cmd.CreatedBy, cmd.UpdatedBy, cmd.Role)
	if err != nil {
		return nil, err
	}
	err = uc.userRepo.CreateUser(ctx, newUser)
	if err != nil {
		return nil, err
	}
	return &dto.CreateUserResponse{
		ID:        newUser.ID,
		Username:  newUser.Username,
		Role:      newUser.Role,
		CreatedBy: newUser.CreatedBy,
		UpdatedBy: newUser.UpdatedBy,
		CreatedAt: newUser.CreatedAt.String(),
		UpdatedAt: newUser.UpdatedAt.String(),
	}, nil
}
