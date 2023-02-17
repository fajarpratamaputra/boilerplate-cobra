package summarizer

import (
	"context"
	"top-ranking-worker/infra"
	"top-ranking-worker/lineup/calculator/mongo"
	"top-ranking-worker/lineup/domain"
)

type Summarizer interface {
	Summarize(ctx context.Context, mongoDb *infra.MongoDatabase) (*domain.Lineup, error)
}

type CalculatorSummarizer struct {
}

func NewCalculatorSummarizer() *CalculatorSummarizer {
	return &CalculatorSummarizer{}
}

func (cs *CalculatorSummarizer) Summarize(ctx context.Context, menu string, mongoDb *infra.MongoDatabase) (*domain.Lineup, error) {
	var filter map[string]interface{}
	if menu != "fyp" {
		filter = map[string]interface{}{
			"service": menu,
		}
	}

	lr := mongo.NewCalculator(mongoDb)

	likesResults, err := lr.Calculate(ctx, "likes", filter)
	if err != nil {
		return nil, err
	}

	viewResults, err := lr.Calculate(ctx, "views", filter)
	if err != nil {
		return nil, err
	}

	if likesResults == nil {
		return viewResults, nil
	}

	for i, likeResult := range *likesResults {
		_, isExist := (*viewResults)[i]
		if isExist {
			(*viewResults)[i].Score += likeResult.Score
			continue
		}

		(*viewResults)[i] = likeResult
	}

	return viewResults, nil
}
