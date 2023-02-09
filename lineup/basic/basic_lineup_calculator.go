package basic

import (
	"context"
	"top-ranking-worker/lineup/domain"
)

type Calculator struct {
}

func newCalculator() *Calculator {
	return &Calculator{}
}

func (lc *Calculator) Calculate(ctx context.Context, contents []domain.Content, interactions []domain.Interaction) (*map[int]float64, error) {
	lineupRanking := map[int]float64{}

	for _, interaction := range interactions {
		var score float64

		switch interaction.Action {
		case "view":
			score += domain.ViewScale
			break
		case "love":
			score += domain.LoveScale
			break
		case "comment":
			score += domain.CommentScale
			break
		case "share":
			score += domain.ShareScale
			break
		}

		lineupRanking[interaction.ContentId] += score
	}

	return &lineupRanking, nil
}
