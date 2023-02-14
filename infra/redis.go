package infra

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisDatabase struct {
	client *redis.Client
}

func NewRedisDatabase() (*RedisDatabase, error) {
	client := new(RedisDatabase)

	client.client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return client, nil
}

func (r *RedisDatabase) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *RedisDatabase) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	val, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return r.client.Set(ctx, key, val, ttl).Err()
}

func (r *RedisDatabase) Close() error {
	return r.client.Close()
}

func (r *RedisDatabase) Ping(ctx context.Context) error {
	return r.client.Ping(ctx).Err()
}
