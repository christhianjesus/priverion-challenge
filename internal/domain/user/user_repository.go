package user

import (
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user User) error
	GetByUsername(ctx context.Context, username string) (User, error)
	Get(ctx context.Context, userID string) (User, error)
	UpdateUser(ctx context.Context, user User) error
	DeleteUser(ctx context.Context, userID string) error
}
