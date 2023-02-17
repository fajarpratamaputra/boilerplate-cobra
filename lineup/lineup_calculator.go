package lineup

import (
	"context"
	"top-ranking-worker/infra"
	"top-ranking-worker/lineup/domain"
	"top-ranking-worker/lineup/mongo"
)

type Calculator interface {
	Calculate(ctx context.Context, name string, filter map[string]interface{}) (*domain.Lineup, error)
}

func Summarize(ctx context.Context, mongoDb *infra.MongoDatabase) (*domain.Lineup, error) {
	filter := map[string]interface{}{
		"service": "hot",
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
