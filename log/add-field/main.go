package main

import (
	"github.com/nite-coder/blackbear/pkg/log"
)

func main() {
	logger := log.With().Str("app_id", "blackbear").Logger()

	logger.Debug().Msg("Hello World")
	logger.Info().Msg("Hello World")
	logger.Warn().Msg("Hello World")
	logger.Error().Msg("Hello World")
}
