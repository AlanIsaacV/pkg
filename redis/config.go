package redis

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

type Settings struct {
	Addr   string `default:"localhost:6379"`
	Db     int    `default:"0"`
	Prefix string
}

var (
	config     Settings
	configOnce sync.Once
)

func Config() Settings {
	configOnce.Do(
		func() {
			if err := envconfig.Process("REDIS", &config); err != nil {
				log.Fatal().Err(err).Msg("Error loading redis config")
			}
		},
	)
	return config
}
