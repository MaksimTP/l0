package logger

import (
	"github.com/rs/zerolog/log"
)

func InitLogger() {
	// log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("Logging has been started")
}
