package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	md "go.mongodb.org/mongo-driver/mongo"
	"top-ranking-worker/config"
	"top-ranking-worker/infra"
	"top-ranking-worker/lineup/domain"
	"top-ranking-worker/lineup/domain/mongo"
)

// Calculator is a struct that contains the database
type Calculator struct {
	mongoDatabase *infra.MongoDatabase
}

func NewCalculator(infra *infra.MongoDatabase) *Calculator {
	return &Calculator{mongoDatabase: infra}
}

// openCursor opens a cursor to the collection
func (lc *Calculator) openCursor(ctx context.Context, collectionName string, filter map[string]interface{}) (*md.Cursor, error) {
	coll := lc.mongoDatabase.GetCollection(config.Config.GetString("MONGO_DB"), collectionName)

	p := make(bson.M)
	for k, v := range filter {
		p[k] = v
	}

	return coll.Find(ctx, p)
}

// calculateFromMongo calculates the score for each content
func (lc *Calculator) calculateFromMongo(ctx context.Context, curr *md.Cursor) (*domain.Lineup, error) {
	l := make(domain.Lineup)
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
		case "likes":
			score += domain.LoveScale
			break
		case "comments":
			score += domain.CommentScale
			break
		case "share":
			score += domain.ShareScale
			break
		}

		content := &domain.Content{
			Service:     result.Service,
			ContentType: result.ContentType,
			Score:       score,
		}

		l[result.ContentID] = content
	}

	if len(l) == 0 {
		return nil, nil
	}

	return &l, nil
}

// Calculate calculates the score for each content
func (lc *Calculator) Calculate(ctx context.Context, name string, filter map[string]interface{}) (*domain.Lineup, error) {
	curr, err := lc.openCursor(ctx, name, filter)
	if err != nil {
		return nil, err
	}

	results, err := lc.calculateFromMongo(ctx, curr)
	if err != nil {
		return nil, err
	}

	return results, nil
}
