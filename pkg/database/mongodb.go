package database

import (
	"context"
	"fmt"

	"github.com/zazhiladhf/newbie-ecommerce/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo(cfg config.MongoDB) (*mongo.Database, error) {
	uri := fmt.Sprintf("%s://%s:%s@%s:%s", cfg.Driver, cfg.Username, cfg.Password, cfg.Host, cfg.Port)
	opts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	db := client.Database(cfg.Name)

	return db, nil
}
