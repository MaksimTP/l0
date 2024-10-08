package logger

import (
	"github.com/rs/zerolog/log"
)

func InitLogger() {
	log.Info().Msg("Logging has been started")
}
