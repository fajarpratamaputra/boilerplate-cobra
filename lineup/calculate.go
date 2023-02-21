package lineup

import (
	"context"
	"fmt"
	"top-ranking-worker/infra"
	"top-ranking-worker/lineup/summarizer"
	"top-ranking-worker/writer"
)

func Calculate(ctx context.Context, wrt writer.Writer, mongoDb *infra.MongoDatabase, menu string) error {
	return calculatePerMenu(ctx, menu, wrt, mongoDb)
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

	key := fmt.Sprintf("shorts:lineup:%s:top", menu)
	if err = wrt.Write(ctx, key, results); err != nil {
		return err
	}

	return nil
}
