package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"top-ranking-worker/infra"
	"top-ranking-worker/lineup/domain"
	"top-ranking-worker/lineup/domain/mongo"
)

type Calculator struct {
	Database *infra.MongoDatabase
}

func (lc *Calculator) calculateThings(ctx context.Context, scale float64, collectionName string) ([]mongo.InteractionScoreSum, error) {
	coll := lc.Database.GetCollection("interactions", collectionName)

	filter := bson.A{
		bson.D{{"$sortByCount", "$contentid"}},
		bson.D{
			{"$set",
				bson.D{
					{"total_score",
						bson.D{
							{"$multiply",
								bson.A{
									"$count",
									scale,
								},
							},
						},
					},
				},
			},
		},
	}

	curr, err := coll.Aggregate(ctx, filter)
	if err != nil {
		return nil, err
	}

	var results []mongo.InteractionScoreSum
	if err = curr.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, err
}

func (lc *Calculator) Calculate(ctx context.Context, contents []domain.Content, interactions []domain.Interaction) (*map[int]float64, error) {
	var lineup = make(map[int]float64)
	results, err := lc.calculateThings(ctx, 2.0, "likes")
	if err != nil {
		return nil, err
	}

	for _, result := range results {
		lineup[result.ID] = result.TotalScore
	}

	return &lineup, nil
}
