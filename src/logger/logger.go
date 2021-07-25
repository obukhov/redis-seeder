package logger

import (
	"github.com/rs/zerolog"
	"os"
)

// NewConsoleLogger creates logger with given minimal logLevel and friendly cli formatter
func NewConsoleLogger(logLevel string) zerolog.Logger {

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	parsedLevel, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		logger.Fatal().Msgf("Cannot parse log level: %s", logLevel)
	}

	return logger.Level(parsedLevel)
}
