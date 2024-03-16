package auth

import (
	"context"

	"github.com/kk7453603/user_RESTfull_API/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, username, password string) (*models.User, error)
}
