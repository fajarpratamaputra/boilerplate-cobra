package main

import (
	"context"
	"log"
	"top-ranking-worker/infra"
	"top-ranking-worker/lineup/summarizer"
	"top-ranking-worker/writer"
)

func calculate(ctx context.Context) {
	mongoDb, err := infra.NewMongoDatabase(ctx)
	if err != nil {
		log.Fatal(err)
	}

	wrt := writer.NewWriter()

	if err = calculatePerMenu(ctx, "fyp", wrt, mongoDb); err != nil {
		log.Fatal(err)
	}
}

func calculatePerMenu(ctx context.Context, menu string, wrt writer.Writer, mongoDb *infra.MongoDatabase) error {
	sm := summarizer.NewCalculatorSummarizer()

	results, err := sm.Summarize(ctx, menu, mongoDb)
	if err != nil {
		return err
	}

	if results == nil {
		return nil
	}

	key := "shorts:master:" + menu
	if err = wrt.Write(ctx, key, results); err != nil {
		return err
	}

	return nil
}
