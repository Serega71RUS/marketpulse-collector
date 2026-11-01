package logger

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Level         string `default:"error" envconfig:"LOGGER_LEVEL"`
	PrettyConsole bool   `default:"false" envconfig:"LOGGER_PRETTY_CONSOLE"`
}

func Init(c Config) {
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if level, err := zerolog.ParseLevel(c.Level); err == nil && level != zerolog.NoLevel {
		zerolog.SetGlobalLevel(level)
	}

	var output io.Writer = os.Stderr
	if c.PrettyConsole {
		output = zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: "15:04:05",
		}
	}

	log.Logger = zerolog.New(output).With().Timestamp().Logger()

	log.Info().Msg("logger initialized")
}
