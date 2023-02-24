package lineup

import (
	"context"
	"fmt"
	"time"
	"top-ranking-worker/infra"
	randomizer "top-ranking-worker/lineup/random"
)

func Random(ctx context.Context, menu string) error {
	return randomPerMenu(ctx, menu)
}

func randomPerMenu(ctx context.Context, menu string) error {
	sm := randomizer.NewRandomizer()

	slice, err := sm.Randomizer(ctx, menu)
	if err != nil {
		return err
	}
	clientRedis, _ := infra.NewRedisDatabase()
	key := fmt.Sprintf("shorts:lineup:%s:random", menu)
	// Marshal the slice to JSON
	err = clientRedis.Set(ctx, key, slice, 24*time.Hour)
	if err != nil {
		panic(err)
	}
	return nil
}
