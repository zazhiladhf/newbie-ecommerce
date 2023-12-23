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

// GetOrderByExternalId implements OrderRepository.
func (r mongoRepository) GetOrderByExternalId(ctx context.Context, externalId string) (order Order, err error) {
	err = r.db.Collection(orderCollection).FindOne(ctx, bson.M{"external_id": externalId}).Decode(&order)
	if err != nil {
		return order, err
	}

	return order, nil
}

// UpdateOrderStatus implements OrderRepository.
func (r mongoRepository) UpdateOrderStatus(ctx context.Context, order Order) (err error) {
	update := bson.M{
		"$set": bson.M{
			"status":  order.Status,
			"invoice": order.Invoice,
		},
	}

	_, err = r.db.Collection(orderCollection).UpdateOne(ctx, bson.M{"external_id": order.ExternalId}, update)
	if err != nil {
		return err
	}

	return nil
}

// GetOrdersByMerchant implements OrderRepository.
func (r mongoRepository) GetOrdersByMerchant(ctx context.Context, limit int, page int, merchantId int) (orders []Order, totalPage int, err error) {
	skip := (page - 1) * limit

	opts := options.Find()
	opts.SetLimit(int64(limit))
	opts.SetSkip(int64(skip))

	filter := bson.M{
		"product.merchant_id": merchantId,
	}

	cursor, err := r.db.Collection(orderCollection).Find(ctx, filter, opts)
	if err != nil {
		return nil, totalPage, err
	}

	if err = cursor.All(ctx, &orders); err != nil {
		return nil, totalPage, err
	}

	count, err := r.db.Collection(orderCollection).CountDocuments(ctx, filter)
	if err != nil {
		return nil, totalPage, err
	}

	totalPage = int(math.Ceil(float64(count) / float64(limit)))

	return orders, totalPage, nil
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

	cursor, err := r.db.Collection(orderCollection).Find(ctx, filter, opts)
	if err != nil {
		return nil, totalPage, err
	}

	if err = cursor.All(ctx, &orders); err != nil {
		return nil, totalPage, err
	}

	count, err := r.db.Collection(orderCollection).CountDocuments(ctx, filter)
	if err != nil {
		return nil, totalPage, err
	}

	totalPage = int(math.Ceil(float64(count) / float64(limit)))

	return orders, totalPage, nil
}

// CreateOrder implements OrderRepository.
func (r mongoRepository) CreateOrder(ctx context.Context, payload Order) (order Order, err error) {
	result, err := r.db.Collection(orderCollection).InsertOne(ctx, payload)
	if err != nil {
		return order, err
	}

	err = r.db.Collection(orderCollection).FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&order)
	if err != nil {
		return order, err
	}

	return order, nil
}
