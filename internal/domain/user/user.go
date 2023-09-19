package user

import "time"

type User interface {
	ID() string
	Username() string
	Email() string
	Phone() string
	CreatedAt() time.Time
	ValidPassword(string) bool
}
