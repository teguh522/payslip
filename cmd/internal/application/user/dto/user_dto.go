package dto

import "github.com/google/uuid"

type CreateUserRequest struct {
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required,min=8"`
	Role      string `json:"role" validate:"required,oneof=admin user"`
	CreatedBy string `json:"created_by" validate:"required"`
	UpdatedBy string `json:"updated_by" validate:"required"`
}

type CreateUserResponse struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	CreatedBy string    `json:"created_by"`
	UpdatedBy string    `json:"updated_by"`
}

type LoginUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginUserResponse struct {
	AccessToken string `json:"access_token"`
}
