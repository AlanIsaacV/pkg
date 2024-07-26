package gcp

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func FetchMetadata(endpoint string) (string, []error) {
	agent := fiber.Get(metaServer + endpoint)
	agent.Set("Metadata-Flavor", "Google")

	_, body, err := agent.String()
	if err != nil {
		log.Error().Errs("Request", err).Str("endpoint", endpoint).Send()
	}
	return body, err
}

func IdentityToken() string {
	body, _ := FetchMetadata("instance/service-accounts/default/identity?audience=" + Config().Audience)
	return body

}

func getProjectId() string {
	body, _ := FetchMetadata("project/project-id")
	return body
}

func getRegion() string {
	// Returns: `projects/PROJECT-NUMBER/regions/REGION`
	body, _ := FetchMetadata("instance/region")
	parts := strings.Split(body, "/")
	return parts[len(parts)-1]
}

func getEmail() string {
	body, _ := FetchMetadata("instance/service-accounts/default/email")
	return body
}
