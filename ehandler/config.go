package ehandler

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

type Settings struct {
	Debug bool
}

var (
	config     Settings
	configOnce sync.Once
)

func Config() Settings {
	configOnce.Do(
		func() {
			if err := envconfig.Process("EHANDLER", &config); err != nil {
				log.Fatal().Err(err).Msg("Error loading ehandler config")
			}
		},
	)
	return config
}
