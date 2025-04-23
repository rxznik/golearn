package action

import (
	"errors"
	"fmt"
	"time"

	"github.com/rxznik/golearn/weather-cli/internal/client"
	"github.com/rxznik/golearn/weather-cli/internal/service/geo"
	"github.com/rxznik/golearn/weather-cli/internal/service/weather"

	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

type Action struct {
	client *client.Client
	logger *zap.Logger
}

func NewAction(client *client.Client, logger *zap.Logger) *Action {
	return &Action{client: client, logger: logger}
}

const timeLayout = "2006-01-02T15:04"

func (ac *Action) Pipeline(ctx *cli.Context) error {

	if !ctx.Bool("loud") {
		ac.logger = ac.logger.WithOptions(zap.IncreaseLevel(zap.FatalLevel))
	}
	city := ctx.Args().Get(0)
	if city == "" {
		return errors.New("expected 1 args (city name) but got 0")
	}

	ac.logger.Info("got city", zap.String("city", city))

	now := time.Now().Local()
	ac.logger.Info("got current time", zap.Time("time", now))

	coordinates, err := geo.GetCityCoordinates(ac.logger, ac.client, city)
	if err != nil {
		ac.logger.Error("failed to get city coordinates", zap.Error(err))
		return err
	}

	lat, lon := coordinates[0], coordinates[1]
	ac.logger.Info("got coordinates", zap.Float64("lat", lat), zap.Float64("lon", lon))

	times, temps, err := weather.GetTodayWeather(ac.logger, ac.client, lat, lon)
	ac.logger.Info("got today weather", zap.Any("times", times), zap.Any("temps", temps))

	if err != nil {
		ac.logger.Error("failed to get today weather", zap.Error(err))
		return err
	}

	fmt.Println("time\ttemp °С\n---------------")

	var currentTemp float64

	for i, timeStr := range times {
		time, err := time.Parse(timeLayout, timeStr)
		if err != nil {
			return err
		}

		localTime := time.Local()

		if localTime.Hour() < now.Hour() {
			continue
		}

		if localTime.Hour() == now.Hour() {
			currentTemp = temps[i]
		}

		timePretty := fmt.Sprintf("\x1b[36m%s\x1b[0m", localTime.Format("15:04"))
		var tempPretty string
		if temps[i] > currentTemp {
			tempPretty = fmt.Sprintf("\x1b[31m%.2f°C\x1b[0m", temps[i])
		} else {
			tempPretty = fmt.Sprintf("\x1b[33m%.2f°C\x1b[0m", temps[i])
		}
		fmt.Printf("%s\t%s\n", timePretty, tempPretty)
	}

	return nil
}
