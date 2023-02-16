package main

import (
	"context"
	"log"
	"top-ranking-worker/lineup"
	"top-ranking-worker/writer"
)

func calculate(ctx context.Context) {
	wrt := writer.NewWriter()
	calculator := lineup.NewCalculator()

	results, err := calculator.Calculate(ctx, nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	res := lineup.ToLineup(results)

	if err = wrt.Write(ctx, res); err != nil {
		log.Fatal(err)
	}
}
