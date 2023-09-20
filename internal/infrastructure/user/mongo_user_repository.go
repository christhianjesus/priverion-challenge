package user

import (
	"context"
	"errors"

	"github.com/christhianjesus/priverion-challenge/internal/domain/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewMongoUserRepository(db *mongo.Database) user.UserRepository {
	return &mongoUserRepository{db.Collection("user")}
}

type mongoUserRepository struct {
	db *mongo.Collection
}

func (r *mongoUserRepository) Create(ctx context.Context, user user.User) error {
	mongoUser := user.(*mongoUser)
	_, err := r.db.InsertOne(ctx, mongoUser.e)

	return err
}

func (r *mongoUserRepository) GetByUsername(ctx context.Context, username string) (user.User, error) {
	var user mongoUser

	err := r.db.FindOne(ctx, bson.M{"username": username}).Decode(&user.e)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("Not found")
		}

		return nil, err
	}

	return &user, nil
}

func (r *mongoUserRepository) Get(ctx context.Context, userID string) (user.User, error) {
	var user mongoUser
	mongoUserID, _ := primitive.ObjectIDFromHex(userID)

	err := r.db.FindOne(ctx, bson.M{"_id": mongoUserID}).Decode(&user.e)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("Not found")
		}

		return nil, err
	}

	return &user, nil
}

func (r *mongoUserRepository) Update(ctx context.Context, user user.User) error {
	mongoUser := user.(*mongoUser)
	_, err := r.db.UpdateOne(ctx, bson.M{"_id": mongoUser.e.ID}, bson.M{"$set": mongoUser.e})
	if err == mongo.ErrNoDocuments {
		return errors.New("Not found")
	}

	return err
}

func (r *mongoUserRepository) Delete(ctx context.Context, userID string) error {
	mongoUserID, _ := primitive.ObjectIDFromHex(userID)
	_, err := r.db.DeleteOne(ctx, bson.M{"_id": mongoUserID})
	if err == mongo.ErrNoDocuments {
		return errors.New("Not found")
	}

	return err
}
