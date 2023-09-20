package advertisement

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type encapsulatedMongoAdvertisement struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	UserID    string             `bson:"user_id,omitempty" json:"user_id"`
	Title     string             `bson:"title,omitempty" json:"title"`
	Body      string             `bson:"body,omitempty" json:"body"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty" json:"updated_at"`
}

type mongoAdvertisement struct {
	e encapsulatedMongoAdvertisement
}

func (a *mongoAdvertisement) ID() string {
	return a.e.ID.Hex()
}

func (a *mongoAdvertisement) UserID() string {
	return a.e.UserID
}

func (a *mongoAdvertisement) Title() string {
	return a.e.Title
}

func (a *mongoAdvertisement) Body() string {
	return a.e.Body
}

func (a *mongoAdvertisement) CreatedAt() time.Time {
	return a.e.CreatedAt
}

func (a *mongoAdvertisement) UpdatedAt() time.Time {
	return a.e.UpdatedAt
}
