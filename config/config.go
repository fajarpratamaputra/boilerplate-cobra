package config

import (
	"github.com/sahalazain/go-common/config"
)

var DefaultConfig = map[string]interface{}{
	"REDIS_HOST":     "127.0.0.1:6379",
	"REDIS_PASSWORD": "",
	"REDIS_DB":       8,
}

var Config config.Getter
var Url string

func Load() error {
	cfgClient, err := config.Load(DefaultConfig, Url)
	if err != nil {
		return err
	}

	Config = cfgClient

	return nil
}
