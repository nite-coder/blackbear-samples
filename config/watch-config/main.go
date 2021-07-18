package main

import (
	"time"

	"github.com/nite-coder/blackbear/pkg/config"
	"github.com/nite-coder/blackbear/pkg/config/provider/file"
	"github.com/nite-coder/blackbear/pkg/log"
	"github.com/nite-coder/blackbear/pkg/log/handler/console"
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

	fileProvder.OnChangedEvent = func(oldContent string, newContent string) {
		err := InitLogger()
		if err != nil {
			panic(err)
		}
	}

	config.AddProvider(fileProvder)

	err = InitLogger()
	if err != nil {
		panic(err)
	}

	for {
		log.Info("Hello")
		log.Debug("Debug")

		time.Sleep(2 * time.Second)
	}
}

type LogItem struct {
	Name     string
	Type     string
	MinLevel string `mapstructure:"min_level"`
}

func InitLogger() error {
	logger := log.New()

	logItems := []LogItem{}
	err := config.Scan("log", &logItems)
	if err != nil {
		return err
	}

	for _, target := range logItems {
		switch target.Type {
		case "console":
			clog := console.New()
			levels := log.GetLevelsFromMinLevel(target.MinLevel)
			logger.AddHandler(clog, levels...)
		case "gelf":
		}

	}

	log.SetLogger(logger)

	return nil
}
