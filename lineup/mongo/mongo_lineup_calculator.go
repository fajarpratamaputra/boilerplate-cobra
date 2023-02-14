package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	md "go.mongodb.org/mongo-driver/mongo"
	"top-ranking-worker/infra"
	"top-ranking-worker/lineup/domain"
	"top-ranking-worker/lineup/domain/mongo"
)

type Calculator struct {
	Database *infra.MongoDatabase
}

func (lc *Calculator) openCursor(ctx context.Context, collectionName string) (*md.Cursor, error) {
	coll := lc.Database.GetCollection("interactions", collectionName)

	return coll.Find(ctx, bson.D{})
}

func (lc *Calculator) calculateThings(ctx context.Context, curr *md.Cursor) (*map[int]float64, error) {
	var lineup = make(map[int]float64)

	for curr.Next(ctx) {
		var result mongo.InteractionModel
		if err := curr.Decode(&result); err != nil {
			return nil, err
		}

		var score float64

		switch result.Action {
		case "view":
			score += domain.ViewScale
			break
		case "like":
			score += domain.LoveScale
			break
		case "comment":
			score += domain.CommentScale
			break
		case "share":
			score += domain.ShareScale
			break
		}

		lineup[result.ContentID] += score

	}

	return &lineup, nil
}

func (lc *Calculator) Calculate(ctx context.Context, contents []domain.Content, interactions []domain.Interaction) (*map[int]float64, error) {
	curr, err := lc.openCursor(ctx, "likes")
	if err != nil {
		return nil, err
	}

	return lc.calculateThings(ctx, curr)
}
