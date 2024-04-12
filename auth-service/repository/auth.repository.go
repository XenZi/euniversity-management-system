package repository

import "go.mongodb.org/mongo-driver/mongo"

type AuthRepository struct {
	cli *mongo.Client
}

func NewAuthRepository(cli *mongo.Client) (*AuthRepository, error) {
	return &AuthRepository{
		cli: cli,
	}, nil
}
