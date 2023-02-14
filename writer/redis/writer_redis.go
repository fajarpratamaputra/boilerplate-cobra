package redis

import (
	"context"
	"time"
	"top-ranking-worker/infra"
)

type Writer struct {
	Client *infra.RedisDatabase
}

func NewWriter(c *infra.RedisDatabase) *Writer {
	return &Writer{
		Client: c,
	}
}

func (w *Writer) Write(ctx context.Context, lineup []int) error {
	ttl := 24 * time.Hour

	return w.Client.Set(ctx, "shorts:master:top", "", ttl)
}
