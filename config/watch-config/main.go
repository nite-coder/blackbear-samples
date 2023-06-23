package main

import (
	"os"
	"time"

	"github.com/nite-coder/blackbear/pkg/config"
	"github.com/nite-coder/blackbear/pkg/config/provider/file"
	"github.com/nite-coder/blackbear/pkg/log"
	"github.com/nite-coder/blackbear/pkg/log/handler/text"
)

func main() {
	fileProvder := file.New()

	err := fileProvder.Load()
	if err != nil {
		panic(err)
	}

	err = fileProvder.WatchConfig()
	if err != nil {
		panic(err)
	}

	config.OnChangedEvent(func() error {
		log.Info().Msg("file changed")
		err := InitLogger()
		if err != nil {
			return err
		}
		return nil
	})

	config.AddProvider(fileProvder)

	err = InitLogger()
	if err != nil {
		panic(err)
	}

	for {
		log.Info().Msg("Hello")
		log.Debug().Msg("Debug")

		time.Sleep(2 * time.Second)
	}
}

type LogItem struct {
	Name     string
	Type     string
	MinLevel string `mapstructure:"min_level"`
}

func InitLogger() error {
	logItems := []LogItem{}
	err := config.Scan("log", &logItems)
	if err != nil {
		return err
	}

	for _, target := range logItems {
		switch target.Type {
		case "text":
			opts := log.HandlerOptions{
				Level:       log.NewLevel(target.MinLevel),
				DisableTime: true,
			}
			logger := log.New(text.New(os.Stderr, &opts))
			log.SetDefault(logger)
		case "json":
		}

	}

	return nil
}
