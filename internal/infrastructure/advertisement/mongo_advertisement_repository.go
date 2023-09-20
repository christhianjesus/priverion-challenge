package advertisement

import (
	"context"
	"errors"

	"github.com/christhianjesus/priverion-challenge/internal/domain/advertisement"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewMongoAdvertisementRepository(db *mongo.Database) advertisement.AdvertisementRepository {
	return &mongoAdvertisementRepository{db.Collection("advertisement")}
}

type mongoAdvertisementRepository struct {
	db *mongo.Collection
}

func (r *mongoAdvertisementRepository) GetAll(ctx context.Context) ([]advertisement.Advertisement, error) {
	cursor, err := r.db.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var advertisements []advertisement.Advertisement
	for cursor.Next(ctx) {
		var advertisement mongoAdvertisement
		if err := cursor.Decode(&advertisement.e); err != nil {
			return nil, err
		}
		advertisements = append(advertisements, &advertisement)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return advertisements, nil
}

func (r *mongoAdvertisementRepository) Create(ctx context.Context, advertisement advertisement.Advertisement) error {
	mongoAdvertisement := advertisement.(*mongoAdvertisement)
	_, err := r.db.InsertOne(ctx, mongoAdvertisement.e)

	return err
}

func (r *mongoAdvertisementRepository) GetOne(ctx context.Context, advertisementID string) (advertisement.Advertisement, error) {
	var advertisement mongoAdvertisement
	mongoAdvertisementID, _ := primitive.ObjectIDFromHex(advertisementID)

	err := r.db.FindOne(ctx, bson.M{"_id": mongoAdvertisementID}).Decode(&advertisement.e)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("Not found")
		}

		return nil, err
	}

	return &advertisement, nil
}

func (r *mongoAdvertisementRepository) Update(ctx context.Context, advertisement advertisement.Advertisement) error {
	mongoAdvertisement := advertisement.(*mongoAdvertisement)
	_, err := r.db.UpdateOne(ctx, bson.M{"_id": mongoAdvertisement.e.ID}, mongoAdvertisement.e)
	if err == mongo.ErrNoDocuments {
		return errors.New("Not found")
	}

	return err
}

func (r *mongoAdvertisementRepository) Delete(ctx context.Context, advertisementID string) error {
	mongoUserID, _ := primitive.ObjectIDFromHex(advertisementID)
	_, err := r.db.DeleteOne(ctx, bson.M{"_id": mongoUserID})
	if err == mongo.ErrNoDocuments {
		return errors.New("Not found")
	}

	return err
}
