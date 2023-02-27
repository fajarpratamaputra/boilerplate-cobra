package randomizer

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
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

	rand.Shuffle(len(keys), func(i, j int) {
		keys[i], keys[j] = keys[j], keys[i]
	})
	// print all keys
	var slice []map[string]interface{}
	for _, key := range keys {
		val, err := clientRedis.Get(ctx, key)
		if err != nil {
			fmt.Println(err)
		} else {
			var content domain.ContentRandom
			err := json.Unmarshal([]byte(val), &content)
			if err != nil {
				fmt.Println(err)
			} else {
				var contentType interface{}
				arr := strings.Split(key, ":")
				if menu == "news" {
					contentType = content.ContentType
				} else {
					contentType = arr[4]
				}

				jsonObj := map[string]interface{}{
					"content_id":   content.ContentId,
					"service":      arr[2],
					"content_type": contentType,
				}
				slice = append(slice, jsonObj)
			}
		}
	}

	return slice, err
}
