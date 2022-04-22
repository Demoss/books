package repository

import (
	"JMIND/internal/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const col = "BlockInfo"

type EthMongo struct {
	*mongo.Database
}

func NewEthMongo(db *mongo.Database) *EthMongo {
	return &EthMongo{Database: db}
}

func (r *EthMongo) SaveBlock(ctx context.Context, number domain.BlockInfo) error {
	_, err := r.Database.Collection(col).InsertOne(ctx, &number)
	if err != nil {
		return err
	}
	return nil
}

func (r *EthMongo) GetBlock(ctx context.Context, blockNumber string) (domain.BlockInfo, error) {
	filter := bson.M{"_id": blockNumber}

	var block domain.BlockInfo

	err := r.Database.Collection(col).FindOne(ctx, filter).Decode(&block)
	if err != nil {
		return domain.BlockInfo{}, err
	}
	return block, nil
}
