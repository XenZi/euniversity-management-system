package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoService struct {
	cli *mongo.Client
}

const (
	h1  = "healthcare"
	rec = "records"
)

func NewMongoService(ctx context.Context) (*MongoService, error) {
	uri := os.Getenv("HEALTHCARE_MONGO_URI")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	_, err = client.Database(h1).Collection(rec).Indexes().CreateOne(context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "patientID", Value: 1}},
			Options: options.Index().SetUnique(true),
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
