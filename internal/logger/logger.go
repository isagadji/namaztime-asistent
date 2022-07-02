package logger

import (
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	logger zerolog.Logger
}

func Init() *Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	l := &Logger{logger: zerolog.New(os.Stderr).With().Timestamp().Logger()}

	return l
}

func (l *Logger) Info(msg string) {
	l.logger.Info().Msg(msg)
}
