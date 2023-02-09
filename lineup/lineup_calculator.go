package lineup

import (
	"top-ranking-worker/lineup/basic"
	"top-ranking-worker/lineup/domain"
)

type Calculator interface {
	Calculate(contents []domain.Content, interactions []domain.Interaction) (*map[int]float64, error)
}

func newCalculator() Calculator {
	return &basic.Calculator{}
}
