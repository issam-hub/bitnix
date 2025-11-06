package mongo

import (
	"user-service/config"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func OpenDB(cfg config.Config) (*mongo.Client, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(cfg.DB.URI))
	if err != nil {
		return nil, err
	}
	return client, nil
}
