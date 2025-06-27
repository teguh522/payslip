package postgres

import (
	"context"

	"github.com/teguh522/payslip/cmd/internal/domain/user/entity"
	"gorm.io/gorm"
)

type UserRepositoryImp struct {
	db *gorm.DB
}

func NewUserRepositoryImp(db *gorm.DB) *UserRepositoryImp {
	return &UserRepositoryImp{
		db: db,
	}
}
func (repo *UserRepositoryImp) CreateUser(ctx context.Context, user *entity.User) error {
	if err := repo.db.WithContext(ctx).Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepositoryImp) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	var user entity.User
	if err := repo.db.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return &user, nil
}
