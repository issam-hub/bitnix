package mongodb

import (
	"context"
	"user-service/internal/domain/apperrors"
	"user-service/internal/domain/entities"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoUserRepository struct {
	Collection *mongo.Collection
}

func NewMongoUserRepository(client *mongo.Client) *MongoUserRepository {
	return &MongoUserRepository{
		Collection: client.Database("user-service").Collection("user"),
	}
}

func (mur MongoUserRepository) Create(ctx context.Context, user entities.User) error {
	dbUser := toDBUser(&user)

	_, err := mur.Collection.InsertOne(ctx, dbUser)
	if err != nil {
		return err
	}

	return nil
}

func (mur MongoUserRepository) GetByUsername(ctx context.Context, username string) (*entities.User, error) {
	target := bson.D{{Key: "username", Value: username}}
	var dbUser User
	err := mur.Collection.FindOne(ctx, target).Decode(&target)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, apperrors.ErrUserNotFound
		} else {
			return nil, err
		}
	}
	user := fromDBUser(&dbUser)
	return user, nil
}
