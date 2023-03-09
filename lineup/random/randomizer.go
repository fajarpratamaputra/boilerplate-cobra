package randomizer

import (
	"context"
	"fmt"
	"strings"
	"top-ranking-worker/infra"
	"top-ranking-worker/lineup/domain"
)

type Randomizer interface {
	Randomizer(ctx context.Context, mongoDb *infra.MongoDatabase) (*domain.Lineup, error)
}

type RandomizerStruct struct {
}

func NewRandomizer() *RandomizerStruct {
	return &RandomizerStruct{}
}

func (cs *RandomizerStruct) Randomizer(ctx context.Context, menu string) ([]map[string]interface{}, error) {
	clientRedis, _ := infra.NewRedisDatabase()
	var keystring string
	// Use the client to get a key from Redis
	if menu != "fyp" {
		keystring = fmt.Sprintf("shorts:master:%s:*", menu)
	} else {
		keystring = fmt.Sprintf("shorts:master:*")
	}

	keys, err := clientRedis.Keys(ctx, keystring)
	if err != nil {
		panic(err)
	}

	var slice []map[string]interface{}
	for _, key := range keys {
		value := strings.Split(key, ":")

		service := value[2]
		contentType := value[4]
		contentId := value[5]

		if service == "news" || contentType == "clip" {
			continue
		}

		jsonObj := map[string]interface{}{
			"content_id":   contentId,
			"service":      service,
			"content_type": contentType,
		}
		slice = append(slice, jsonObj)
	}

	return slice, err
}
