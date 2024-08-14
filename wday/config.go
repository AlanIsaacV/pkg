package wday

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

type Settings struct {
	Countries   []string
	MarketOpen  int `default:"0" envconfig:"MARKET_OPEN"`
	MarketClose int `default:"0" envconfig:"MARKET_CLOSE"`
}

var (
	config     Settings
	configOnce sync.Once
)

func Config() Settings {
	configOnce.Do(
		func() {
			if err := envconfig.Process("WDAY", &config); err != nil {
				log.Fatal().Err(err).Msg("Error loading wday config")
			}
		},
	)
	return config
}
