package cloudtasks

import (
	"fmt"
	"sync"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

type Settings struct {
	Url   string
	Name  string `envconfig:"TASK_QUEUE"`
	Delay int64

	Queue string `envconfig:"TASK_PATH"`
	// ProjectId string `envconfig:"PROJECT_ID"`
}

var (
	config     Settings
	configOnce sync.Once
)

func Config() Settings {
	configOnce.Do(
		func() {
			if err := envconfig.Process("TASK", &config); err != nil {
				log.Fatal().Err(err).Msg("Error loading tasks config")
			}

			if config.Queue == "" {
				config.Queue = fmt.Sprintf(
					"projects/%s/locations/%s/queues/%s",
					gcpConfig.Project, gcpConfig.Region, config.Name,
				)
			}
		},
	)
	return config
}
