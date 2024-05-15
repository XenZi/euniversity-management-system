package services

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoService struct {
	cli *mongo.Client
}

func NewMongoService(ctx context.Context) (*MongoService, error) {
	uri := os.Getenv("DORM_MONGO_URI")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	return &MongoService{
		cli: client,
	}, nil
}

func (m MongoService) GetCLI() *mongo.Client {
	return m.cli
}
