package order

import (
	"context"
	"math"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepository struct {
	db *mongo.Database
}

func NewMongoRepository(db *mongo.Database) mongoRepository {
	return mongoRepository{
		db: db,
	}
}

// GetOrderHistories implements OrderRepository.
func (r mongoRepository) GetOrderHistories(ctx context.Context, limit int, page int, userId int) (orders []Order, totalPage int, err error) {
	skip := (page - 1) * limit

	opts := options.Find()
	opts.SetLimit(int64(limit))
	opts.SetSkip(int64(skip))

	filter := bson.M{
		"user_id": userId,
	}

	cursor, err := r.db.Collection("orders").Find(ctx, filter, opts)
	if err != nil {
		return nil, totalPage, err
	}

	if err = cursor.All(ctx, &orders); err != nil {
		return nil, totalPage, err
	}

	count, err := r.db.Collection("orders").CountDocuments(ctx, filter)
	if err != nil {
		return nil, totalPage, err
	}

	totalPage = int(math.Ceil(float64(count) / float64(limit)))

	return orders, totalPage, nil
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
