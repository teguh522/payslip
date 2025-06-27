package usecase

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/teguh522/payslip/cmd/internal/application/user/command"
	"github.com/teguh522/payslip/cmd/internal/application/user/dto"
	"github.com/teguh522/payslip/cmd/internal/domain/user/repository"
	"github.com/teguh522/payslip/cmd/internal/pkg/config"
	"github.com/teguh522/payslip/cmd/internal/pkg/security"
)

type LoginUserUseCase struct {
	userRepo     repository.UserRepository
	jwtSecretKey []byte
}

func NewLoginUserUseCase(userRepo repository.UserRepository, cfg *config.Config) *LoginUserUseCase {
	return &LoginUserUseCase{userRepo: userRepo,
		jwtSecretKey: []byte(cfg.App.JWTSecret),
	}
}

func (uc *LoginUserUseCase) Execute(ctx context.Context, cmd command.LoginUserCommand) (*dto.LoginUserResponse, error) {
	user, err := uc.userRepo.FindByUsername(ctx, cmd.Username)
	if err != nil {
		return nil, err
	}

	err = security.VerifyPassword(user.PasswordHash, cmd.Password)
	if err != nil {
		return nil, err
	}

	claims := jwt.MapClaims{
		"userId":   user.ID.String(),
		"userName": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(uc.jwtSecretKey)
	if err != nil {
		return nil, err
	}

	return &dto.LoginUserResponse{
		AccessToken: tokenString,
	}, nil
}
