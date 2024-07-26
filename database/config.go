package database

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

type Settings struct {
	User string `required:"true"`
	Pass string `required:"true"`
	Host string `required:"true"`
	Name string `required:"true"`
}

var (
	config     Settings
	configOnce sync.Once
)

func Config() Settings {
	configOnce.Do(
		func() {
			if err := envconfig.Process("DB", &config); err != nil {
				log.Fatal().Err(err).Msg("Error loading database config")
			}
		},
	)
	return config
}
