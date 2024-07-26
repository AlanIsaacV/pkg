package ehandler

import (
	"errors"
	"io"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type ErrorMsg struct {
	Message    string   `json:"message"`
	StatusCode int      `json:"status_code"`
	Data       []string `json:"data"`
}

func EHandler(c *fiber.Ctx, err error) error {
	var msg ErrorMsg
	var e *fiber.Error

	code := 500
	if errors.As(err, &e) {
		msg = ErrorMsg{StatusCode: e.Code, Message: e.Message}
		if e.Code < 400 {
			code = e.Code
		} else if e.Code < 500 {
			code = 400
		}
	} else {
		msg = ErrorMsg{StatusCode: code, Message: err.Error()}
	}

	l := log.Ctx(c.UserContext()).With().Err(err).Logger()
	if code >= 500 {
		l.Error().Send()
	} else {
		l.Warn().Send()
	}

	return c.Status(code).JSON(msg)
}

func Cleanup(resources ...io.Closer) func() error {
	return func() error {
		for _, r := range resources {
			if err := r.Close(); err != nil {
				log.Error().Err(err).Send()
			}
		}
		return nil
	}
}
