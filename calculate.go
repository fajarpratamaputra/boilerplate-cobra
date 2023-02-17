package main

import (
	"context"
	"log"
	"top-ranking-worker/infra"
	"top-ranking-worker/lineup"
	"top-ranking-worker/writer"
)

func calculate(ctx context.Context) {
	mongoDb, err := infra.NewMongoDatabase(ctx)
	if err != nil {
		log.Fatal(err)
	}

	results, err := lineup.Summarize(ctx, mongoDb)
	if err != nil {
		log.Fatal(err)
	}

	if results == nil {
		return
	}

	wrt := writer.NewWriter()

	key := "shorts:master:top"
	if err = wrt.Write(ctx, key, results); err != nil {
		log.Fatal(err)
	}
}
