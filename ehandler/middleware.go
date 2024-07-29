package ehandler

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/AlanIsaacV/pkg/gcp"
)

func LoggerHandler(c *fiber.Ctx) error {
	gcpConfig := gcp.Config()

	if gcpConfig.Service == "" {
		c.SetUserContext(log.Logger.WithContext(c.UserContext()))
		return c.Next()
	}

	l := log.With().Dict(
		"httpRequest", zerolog.Dict().
			Str("requestMethod", c.Method()).
			Str("requestUrl", c.OriginalURL()).
			Str("userAgent", c.Get("user-agent")),
	)
	if body := c.BodyRaw(); len(body) > 0 {
		l.Bytes("requestBody", body)
	}

	if trace := c.Get("x-cloud-trace-context"); trace != "" {
		trace = fmt.Sprintf(
			"projects/%s/traces/%s", gcpConfig.Project,
			strings.Split(trace, "/")[0],
		)
		if trace != "" {
			l.Str("logging.googleapis.com/trace", trace)
		}
	}
	c.SetUserContext(UpdateCtx(c.UserContext(), l))
	return c.Next()
}
