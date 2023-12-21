package auth

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisRepository struct {
	db *redis.Client
}

func NewRedisRepository(db *redis.Client) RedisRepository {
	return RedisRepository{
		db: db,
	}
}

func (r RedisRepository) Set(ctx context.Context, email string, token string, lifeTime int) (err error) {
	expiration := time.Duration(lifeTime) * time.Hour
	err = r.db.Set(ctx, email, token, expiration).Err()
	if err != nil {
		log.Println("error when try to set data to redis with message :", err.Error())
		return
	}

	return
}

func (r RedisRepository) Get(ctx context.Context, email string) (token string, err error) {
	token, err = r.db.Get(ctx, email).Result()
	if err != nil {
		return "", nil
	}

	return token, nil
}
