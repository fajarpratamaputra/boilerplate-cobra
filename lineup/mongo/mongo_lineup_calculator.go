package mongo

import (
	"top-ranking-worker/lineup/domain"
)

type Calculator struct {
}

func (lc *Calculator) Calculate(contents []domain.Content, interactions []domain.Interaction) *map[int]float64 {
	return nil
}
