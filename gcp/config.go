package gcp

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

type Settings struct {
	// https://cloud.google.com/run/docs/container-contract#env-vars
	Service string `envconfig:"K_SERVICE"`

	// https://cloud.google.com/run/docs/container-contract#metadata-server
	Project string `envconfig:"PROJECT_ID"`
	Region  string
	Email   string

	// https://cloud.google.com/run/docs/authenticating/service-to-service#use_the_metadata_server
	Audience string
}

var (
	metaServer = "http://metadata.google.internal/computeMetadata/v1/"
	config     Settings
	configOnce sync.Once
)

func Config() Settings {
	configOnce.Do(
		func() {
			if err := envconfig.Process("GCP", &config); err != nil {
				log.Fatal().Err(err).Msg("Error loading gcp config")
			}

			if config.Service != "" {
				if config.Project == "" {
					config.Project = getProjectId()
				}
				if config.Region == "" {
					config.Region = getRegion()
				}
				if config.Email == "" {
					config.Email = getEmail()
				}
			} else if config.Region == "" {
				config.Region = "us-east1"
			}
		},
	)
	return config
}
