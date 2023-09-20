package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EncapsulatedMongoUser struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Username  string             `bson:"username,omitempty" json:"username"`
	Password  string             `bson:"password,omitempty" json:"password"`
	Email     string             `bson:"email,omitempty" json:"email"`
	Phone     string             `bson:"phone,omitempty" json:"phone"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty" json:"updated_at"`
}

type mongoUser struct {
	e EncapsulatedMongoUser
}

func (u *mongoUser) ID() string {
	return u.e.ID.Hex()
}

func (u *mongoUser) Username() string {
	return u.e.Username
}

func (u *mongoUser) Password() string {
	return u.e.Password
}

func (u *mongoUser) Email() string {
	return u.e.Email
}

func (u *mongoUser) Phone() string {
	return u.e.Phone
}

func (u *mongoUser) CreatedAt() time.Time {
	return u.e.CreatedAt
}

func (u *mongoUser) UpdatedAt() time.Time {
	return u.e.UpdatedAt
}

func (u *mongoUser) SetPassword(password string) {
	u.e.Password = password
}
