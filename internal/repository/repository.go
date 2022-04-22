package repository

import (
	"JMIND/internal/domain"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type Eth interface {
	SaveBlock(ctx context.Context, number domain.BlockInfo) error
	GetBlock(ctx context.Context, blockNumber string) (domain.BlockInfo, error)
}

type Repository struct {
	Eth
}

func NewRepository(database *mongo.Database) *Repository {
	return &Repository{Eth: NewEthMongo(database)}
}
