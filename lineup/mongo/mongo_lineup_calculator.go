package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"top-ranking-worker/lineup/domain"
	"top-ranking-worker/lineup/infra"
)

type Calculator struct {
	Database *infra.MongoDatabase
}

func (lc *Calculator) Calculate(ctx context.Context, contents []domain.Content, interactions []domain.Interaction) (*map[int]float64, error) {
	coll := lc.Database.GetCollection("lineup", "contents")

	cur, err := coll.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		// do something with result....
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return nil, nil
}
