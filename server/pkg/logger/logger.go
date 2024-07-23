package logger

import (
	"os"

	"github.com/haguirrear/coffeeassistant/server/pkg/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

var Logger = CreateLogger()

func CreateLogger() zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	if !config.Conf.IsProd {
		logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	return logger
}
