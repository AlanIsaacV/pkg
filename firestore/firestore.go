package firestore

import (
	"context"
	"os"
	"sync"

	"cloud.google.com/go/firestore"
	"github.com/goccy/go-json"
	"github.com/rs/zerolog/log"
	"google.golang.org/api/option"
)

type Credentials struct {
	ProjectID string `json:"project_id"`
}

func getProjectFromCred(credRaw []byte) string {
	var cred Credentials
	if err := json.Unmarshal(credRaw, &cred); err != nil {
		log.Error().Err(err).Msg("Error marshaling firestore credentials")
	}
	return cred.ProjectID
}

var (
	client     *firestore.Client
	clientOnce sync.Once
)

func Client() *firestore.Client {
	clientOnce.Do(
		func() {
			var err error
			var cred []byte
			var conf = Config()

			if conf.Credentials != "" {
				cred = []byte(conf.Credentials)
			} else if conf.CredentialsFile != "" {
				cred, err = os.ReadFile(conf.CredentialsFile)
				if err != nil {
					log.Error().Err(err).Msg("Error reading firestore credentials file")
				}
			}

			if conf.ProjectID == "" {
				conf.ProjectID = getProjectFromCred(cred)
			}

			client, err = firestore.NewClient(
				context.Background(), conf.ProjectID, option.WithCredentialsJSON(cred),
			)

			if err != nil {
				log.Fatal().Err(err).Msg("Error creating firestore client")
			}
		},
	)
	return client
}
