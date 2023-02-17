package redis

import (
	"context"
	"time"
	"top-ranking-worker/infra"
	"top-ranking-worker/lineup/domain"
	redisDomain "top-ranking-worker/lineup/domain/redis"
)

type Writer struct {
	Client *infra.RedisDatabase
}

func NewWriter(c *infra.RedisDatabase) *Writer {
	return &Writer{
		Client: c,
	}
}

func convertLineupToPayload(lineup interface{}) []*redisDomain.LineupPayload {
	l := lineup.(*domain.Lineup)

	var domainLineup []*redisDomain.LineupPayload

	for i, content := range *l {
		domainLineup = append(domainLineup, &redisDomain.LineupPayload{
			ContentId:   i,
			Service:     content.Service,
			ContentType: content.ContentType,
		})
	}

	return domainLineup
}

func (w *Writer) Write(ctx context.Context, key string, lineup interface{}) error {
	ttl := 24 * time.Hour

	r := convertLineupToPayload(lineup)

	return w.Client.Set(ctx, key, r, ttl)
}
