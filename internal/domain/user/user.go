package user

import "time"

type User interface {
	ID() string
	Username() string
	Password(string) bool
	Email() string
	Phone() string
	CreatedAt() time.Time
	UpdatedAt() time.Time
}
