package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	md "go.mongodb.org/mongo-driver/mongo"
	"top-ranking-worker/infra"
	"top-ranking-worker/lineup/domain"
	"top-ranking-worker/lineup/domain/mongo"
)

// Calculator is a struct that contains the database
type Calculator struct {
	Database *infra.MongoDatabase
}

// openCursor opens a cursor to the collection
func (lc *Calculator) openCursor(ctx context.Context, collectionName string) (*md.Cursor, error) {
	coll := lc.Database.GetCollection("interactions", collectionName)

	return coll.Find(ctx, bson.D{})
}

// calculateThings calculates the score for each content
func (lc *Calculator) calculateThings(ctx context.Context, curr *md.Cursor) (domain.LineupMap, error) {
	l := make(domain.LineupMap)

	for curr.Next(ctx) {
		var result mongo.InteractionModel
		if err := curr.Decode(&result); err != nil {
			return nil, err
		}

		var score float64

		switch result.Action {
		case "views":
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

		content := &domain.LineupContent{
			Service:     result.Service,
			ContentType: result.ContentType,
			Score:       score,
		}

		l[result.ContentID] = content
	}

	return l, nil
}

// calculateScore is a helper function to calculate the score for each content
func (lc *Calculator) calculateScore(ctx context.Context, collectionName string) (domain.LineupMap, error) {
	curr, err := lc.openCursor(ctx, collectionName)
	if err != nil {
		return nil, err
	}

	results, err := lc.calculateThings(ctx, curr)
	if err != nil {
		return nil, err
	}

	return results, nil
}

// Calculate calculates the score for each content
func (lc *Calculator) Calculate(ctx context.Context, contents []domain.Content, interactions []domain.LineupContent) (domain.LineupMap, error) {
	likesResults, err := lc.calculateScore(ctx, "likes")
	if err != nil {
		return nil, err
	}

	viewResults, err := lc.calculateScore(ctx, "views")
	if err != nil {
		return nil, err
	}

	for i, likeResult := range likesResults {
		_, isExist := viewResults[i]
		if isExist {
			viewResults[i].Score += likeResult.Score
			continue
		}

		viewResults[i] = likeResult
	}

	return viewResults, nil
}
