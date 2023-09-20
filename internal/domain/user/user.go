package user

import "time"

type User interface {
	ID() string
	Username() string
	Password() string
	Email() string
	Phone() string
	CreatedAt() time.Time
	UpdatedAt() time.Time

	SetPassword(string)
}

func ValidateUser(user User) error {
	return nil
}
