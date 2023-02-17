package writer

import (
	"context"
	"log"
	"top-ranking-worker/infra"
	"top-ranking-worker/writer/redis"
)

type Writer interface {
	Write(ctx context.Context, key string, lineup interface{}) error
}

func NewWriter() Writer {
	c, err := infra.NewRedisDatabase()
	if err != nil {
		log.Fatal(err)
	}

	return redis.NewWriter(c)
}
