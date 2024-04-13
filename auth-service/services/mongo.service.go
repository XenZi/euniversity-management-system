package services

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoService struct {
	cli *mongo.Client
}

func NewMongoService(ctx context.Context) (*MongoService, error) {
	uri := os.Getenv("AUTH_MONGO_URI")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	_, err = client.Database("auth").Collection("user").Indexes().CreateMany(context.Background(),
		[]mongo.IndexModel{
			{
				Keys:    bson.D{{Key: "email", Value: 1}},
				Options: options.Index().SetUnique(true),
			},
			{
				Keys:    bson.D{{Key: "personalIdentificationNumber", Value: 1}},
				Options: options.Index().SetUnique(true),
			},
		})
	if err != nil {
		log.Println("MongoService: Error while creating indexes: %v", err)
		return nil, err
	}
	return &MongoService{
		cli: client,
	}, nil
}

func (m MongoService) GetCLI() *mongo.Client {
	return m.cli
}
