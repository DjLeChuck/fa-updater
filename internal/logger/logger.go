package logger

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Init() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}

	log.Logger = zerolog.New(output).With().Timestamp().Logger()
}

func Info(msg string) {
	log.Info().Msg(msg)
}

func Infof(format string, v ...interface{}) {
	log.Info().Msgf(format, v...)
}

func Fatal(err error, msg string) {
	log.Fatal().Err(err).Msg(msg)
}
func Fatalf(err error, format string, v ...interface{}) {
	log.Fatal().Err(err).Msgf(format, v...)
}

func Error(err error, msg string) {
	log.Error().Err(err).Msg(msg)
}

func Errorf(err error, format string, v ...interface{}) {
	log.Error().Err(err).Msgf(format, v...)
}
