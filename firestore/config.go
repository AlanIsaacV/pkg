package firestore

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

type Settings struct {
	CredentialsFile string `envconfig:"GOOGLE_APPLICATION_CREDENTIALS"`
	Credentials     string
	ProjectID       string
}

var (
	config     Settings
	configOnce sync.Once
)

func Config() Settings {
	configOnce.Do(
		func() {
			if err := envconfig.Process("FIRESTORE", &config); err != nil {
				log.Fatal().Err(err).Msg("Error loading firestore config")
			}
		},
	)
	return config
}
