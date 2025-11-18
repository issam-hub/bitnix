package mongodb

import (
	"context"
	"user-service/internal/domain/entities"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoUserRepository struct {
	Client *mongo.Client
}

func NewMongoUserRepository(client *mongo.Client) *MongoUserRepository {
	return &MongoUserRepository{
		Client: client,
	}
}

func (mur MongoUserRepository) Create(ctx context.Context, user entities.User) error {
	dbUser := toDBUser(&user)

	_, err := mur.Client.Database("user-service").Collection("user").InsertOne(ctx, dbUser)
	if err != nil {
		return err
	}

	return nil
}
