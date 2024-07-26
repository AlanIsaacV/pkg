package ehandler

import (
	"context"
	"os"
	"runtime/debug"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/AlanIsaacV/pkg/gcp"
)

func InitLog() {
	if Config().Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	zerolog.LevelFieldName = "severity"
	zerolog.TimestampFieldName = "timestamp"
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = stackMarshall

	logger := log.With().Caller().Stack().Logger()
	if gcp.Config().Service == "" {
		logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	log.Logger = logger
}

func stackMarshall(err error) interface{} {
	return string(debug.Stack())
}

func UpdateCtx(ctx context.Context, logCtx zerolog.Context) context.Context {
	return logCtx.Ctx(ctx).Logger().WithContext(ctx)
}
