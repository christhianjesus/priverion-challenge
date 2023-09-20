package user

import (
	"context"
)

type UserService interface {
	CreateUser(ctx context.Context, user User) error
	AuthUser(ctx context.Context, username, password string) (User, error)
	GetUser(ctx context.Context, userID string) (User, error)
	UpdateUser(ctx context.Context, user User) error
	DeleteUser(ctx context.Context, userID string) error
}
