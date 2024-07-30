package ehandler

import (
	"context"
	"os"
	"runtime/debug"

	"github.com/rs/zerolog"
	_ "github.com/rs/zerolog/hlog"
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

func AddStr(ctx context.Context, key string, value string) {
	l := zerolog.Ctx(ctx)
	l.UpdateContext(func(c zerolog.Context) zerolog.Context { return c.Str(key, value) })
}
func AddBool(ctx context.Context, key string, value bool) {
	l := zerolog.Ctx(ctx)
	l.UpdateContext(func(c zerolog.Context) zerolog.Context { return c.Bool(key, value) })
}
func AddInt(ctx context.Context, key string, value int) {
	l := zerolog.Ctx(ctx)
	l.UpdateContext(func(c zerolog.Context) zerolog.Context { return c.Int(key, value) })
}
func AddInt8(ctx context.Context, key string, value int8) {
	l := zerolog.Ctx(ctx)
	l.UpdateContext(func(c zerolog.Context) zerolog.Context { return c.Int8(key, value) })
}
func AddInt16(ctx context.Context, key string, value int16) {
	l := zerolog.Ctx(ctx)
	l.UpdateContext(func(c zerolog.Context) zerolog.Context { return c.Int16(key, value) })
}
func AddInt32(ctx context.Context, key string, value int32) {
	l := zerolog.Ctx(ctx)
	l.UpdateContext(func(c zerolog.Context) zerolog.Context { return c.Int32(key, value) })
}
func AddInt64(ctx context.Context, key string, value int64) {
	l := zerolog.Ctx(ctx)
	l.UpdateContext(func(c zerolog.Context) zerolog.Context { return c.Int64(key, value) })
}
func AddUint(ctx context.Context, key string, value uint) {
	l := zerolog.Ctx(ctx)
	l.UpdateContext(func(c zerolog.Context) zerolog.Context { return c.Uint(key, value) })
}
func AddUint8(ctx context.Context, key string, value uint8) {
	l := zerolog.Ctx(ctx)
	l.UpdateContext(func(c zerolog.Context) zerolog.Context { return c.Uint8(key, value) })
}
func AddUint16(ctx context.Context, key string, value uint16) {
	l := zerolog.Ctx(ctx)
	l.UpdateContext(func(c zerolog.Context) zerolog.Context { return c.Uint16(key, value) })
}
func AddUint32(ctx context.Context, key string, value uint32) {
	l := zerolog.Ctx(ctx)
	l.UpdateContext(func(c zerolog.Context) zerolog.Context { return c.Uint32(key, value) })
}
func AddUint64(ctx context.Context, key string, value uint64) {
	l := zerolog.Ctx(ctx)
	l.UpdateContext(func(c zerolog.Context) zerolog.Context { return c.Uint64(key, value) })
}
func AddFloat32(ctx context.Context, key string, value float32) {
	l := zerolog.Ctx(ctx)
	l.UpdateContext(func(c zerolog.Context) zerolog.Context { return c.Float32(key, value) })
}
func AddFloat64(ctx context.Context, key string, value float64) {
	l := zerolog.Ctx(ctx)
	l.UpdateContext(func(c zerolog.Context) zerolog.Context { return c.Float64(key, value) })
}
