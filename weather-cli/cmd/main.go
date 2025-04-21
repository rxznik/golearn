package main

import (
	"fmt"
	"os"

	"github.com/rxznik/golearn/weather-cli/internal/action"
	"github.com/rxznik/golearn/weather-cli/internal/client"
	"github.com/rxznik/golearn/weather-cli/internal/config"

	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func main() {

	cfg := config.MustLoad()
	logger := mustSetupLogger(cfg.Env)
	defer logger.Sync()

	client := client.NewClient(&cfg.Client)
	action := action.NewAction(client, logger)
	app := &cli.App{
		Name:   "weather-cli",
		Usage:  "Получает погоду по названию города (e.g. weather-cli Москва)",
		Action: action.Pipeline,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "loud",
				Aliases: []string{"l"},
				Value:   "false",
				Usage:   "показать логи (true, false)",
			},
		},
		Args:      true,
		ArgsUsage: "city name",
	}

	logger.Debug("starting app")

	if err := app.Run(os.Args); err != nil {
		if cfg.Env == "dev" {
			logger.Fatal("failed to run app", zap.Error(err))
		}
		fmt.Printf("\x1b[31m%s\x1b[0m\n", err.Error())
		os.Exit(1)
	}
}

func mustSetupLogger(envType string) *zap.Logger {
	switch envType {
	case "dev":
		return zap.Must(zap.NewDevelopment())
	case "prod":
		return zap.Must(zap.NewProduction())
	default:
		zap.L().Fatal("unknown env type", zap.String("env", envType))
		return nil
	}
}
