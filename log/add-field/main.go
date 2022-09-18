package main

import (
	"github.com/nite-coder/blackbear/pkg/log"
	"github.com/nite-coder/blackbear/pkg/log/handler/console"
)

func main() {
	logger := log.New()
	opts := console.ConsoleOptions{DisableColor: false}
	clog := console.New(opts)
	logger.AddHandler(clog, log.AllLevels...)
	log.SetLogger(logger)

	logger = log.Str("app_id", "blackbear").Logger()

	logger.Debug("Hello World")
	logger.Info("Hello World")
	logger.Warn("Hello World")
	logger.Error("Hello World")
}
