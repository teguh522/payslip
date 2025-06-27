package repository

import (
	"context"

	"github.com/teguh522/payslip/cmd/internal/domain/user/entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) error
	FindByUsername(ctx context.Context, username string) (*entity.User, error)
}
