package main

import (
	"github.com/nite-coder/blackbear/pkg/log"
	"github.com/nite-coder/blackbear/pkg/log/handler/console"
)

func main() {
	logger := log.New()
	// logger.DisableTimeField = true

	opts := console.ConsoleOptions{DisableColor: true}
	clog := console.New(opts)
	logger.AddHandler(clog, log.AllLevels...)
	log.SetLogger(logger)

	log.Debug("Hello World")
}
