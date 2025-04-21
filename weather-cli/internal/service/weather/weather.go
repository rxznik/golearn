package weather

import (
	"errors"
	"github.com/rxznik/golearn/weather-cli/internal/response"

	"go.uber.org/zap"
)

type WeatherClient interface {
	GetWeatherData(logger *zap.Logger, lat, lon float64) (*response.Response, error)
}

func GetTodayWeather(logger *zap.Logger, client WeatherClient, lat, lon float64) ([]string, []float64, error) {
	resp, err := client.GetWeatherData(logger, lat, lon)
	if err != nil {
		if resp.Error.Error != "" && resp.Error.Reason != "" {
			errorMsg := resp.Error.Error + ": " + resp.Error.Reason
			logger.Error(errorMsg, zap.Error(err))
			return nil, nil, errors.New(errorMsg)
		}
		logger.Error("failed to get weather data", zap.Error(err))
		return nil, nil, err
	}
	logger.Debug("got weather data from api", zap.Any("data", resp))

	weatherData, ok := resp.OK.(response.WeatherOK)
	if !ok {
		errorMsg := "invalid weather data"
		logger.Error(errorMsg)
		return nil, nil, errors.New(errorMsg)
	}

	return weatherData.Hourly.Time, weatherData.Hourly.Temp, nil
}
