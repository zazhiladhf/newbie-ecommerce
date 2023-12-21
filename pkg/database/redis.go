package database

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/zazhiladhf/newbie-ecommerce/config"
)

func ConnectRedis(ctx context.Context, cfg config.Redis) (client *redis.Client, err error) {
	address := cfg.Host + ":" + cfg.Port
	client = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: cfg.Password,
	})

	err = client.Ping(ctx).Err()
	if err != nil {
		log.Println("cannot connect to redis")
		return
	}
	return
}
