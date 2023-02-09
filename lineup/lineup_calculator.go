package lineup

import (
	"context"
	"top-ranking-worker/infra"
	"top-ranking-worker/lineup/basic"
	"top-ranking-worker/lineup/domain"
	"top-ranking-worker/lineup/mongo"
)

type Calculator interface {
	Calculate(ctx context.Context, contents []domain.Content, interactions []domain.Interaction) (*map[int]float64, error)
}

func newCalculator() Calculator {
	return &basic.Calculator{}
}

func newMongoCalculator() Calculator {
	mongoDb, err := infra.NewMongoDatabase()
	if err != nil {
		return nil
	}

	return &mongo.Calculator{Database: mongoDb}
}
