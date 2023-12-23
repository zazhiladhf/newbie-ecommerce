package order

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	db *mongo.Database
}

func NewMongoRepository(db *mongo.Database) mongoRepository {
	return mongoRepository{
		db: db,
	}
}

// CreateOrder implements OrderRepository.
func (r mongoRepository) CreateOrder(ctx context.Context, payload Order) (order Order, err error) {
	result, err := r.db.Collection("orders").InsertOne(ctx, payload)
	if err != nil {
		return order, err
	}

	err = r.db.Collection("orders").FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&order)
	if err != nil {
		return order, err
	}

	return order, nil
}
