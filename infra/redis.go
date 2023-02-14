package infra

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisDatabase struct {
	client *redis.Client
}

func NewRedisDatabase() (*RedisDatabase, error) {
	client := new(RedisDatabase)

	client.client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return client, nil
}

func (r *RedisDatabase) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *RedisDatabase) Set(ctx context.Context, key string, value string, ttl time.Duration) error {
	return r.client.Set(ctx, key, value, ttl).Err()
}

func (r *RedisDatabase) Close() error {
	return r.client.Close()
}

func (r *RedisDatabase) Ping(ctx context.Context) error {
	return r.client.Ping(ctx).Err()
}
