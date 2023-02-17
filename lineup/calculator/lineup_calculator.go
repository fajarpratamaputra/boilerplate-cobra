package calculator

import (
	"context"
	"top-ranking-worker/lineup/domain"
)

type Calculator interface {
	Calculate(ctx context.Context, name string, filter map[string]interface{}) (*domain.Lineup, error)
}
