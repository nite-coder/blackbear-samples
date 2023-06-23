package main

import (
	"os"

	"github.com/nite-coder/blackbear/pkg/log"
	"github.com/nite-coder/blackbear/pkg/log/handler/text"
)

func main() {

	// json handler
	log.Debug().Msg("Hello World")
	// {"time":"2023-06-23T06:17:43Z","level":"DEBUG","msg":"Hello World"}

	// text handler
	opts := log.HandlerOptions{
		Level:       log.DebugLevel,
		DisableTime: true,
	}
	logger := log.New(text.New(os.Stderr, &opts))
	log.SetDefault(logger)
	log.Debug().Msg("Hello World")
	// 06:17:43.991 DEBUG  Hello World
}
