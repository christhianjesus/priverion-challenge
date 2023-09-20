package user

import (
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user User) error
	GetByUsername(ctx context.Context, username string) (User, error)
	Get(ctx context.Context, userID string) (User, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, userID string) error
}
