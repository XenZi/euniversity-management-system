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

const (
	fs    = "food-service"
	cards = "cards"
)

func NewMongoService(ctx context.Context) (*MongoService, error) {
	uri := os.Getenv("FOOD_MONGO_URI")
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
