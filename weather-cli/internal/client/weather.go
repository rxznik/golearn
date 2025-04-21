package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rxznik/golearn/weather-cli/internal/response"
	"go.uber.org/zap"
)

func (c *Client) GetWeatherData(logger *zap.Logger, lat, lon float64) (*response.WeatherResponse, error) {
	url_ := fmt.Sprintf(
		"%s/forecast?latitude=%f&longitude=%f&hourly=%s&forecast_days=%d",
		c.WeatherCfg.URL,
		lat,
		lon,
		c.WeatherCfg.Hourly,
		c.WeatherCfg.ForecastDays,
	)

	logger.Debug("got url", zap.String("url", url_))

	resp, err := c.HTTP.Get(url_)
	if err != nil {
		logger.Error("failed to get weather data", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {

	case http.StatusOK:
		var weatherBody response.WeatherOK
		if err := json.NewDecoder(resp.Body).Decode(&weatherBody); err != nil {
			logger.Error("failed to decode weather data", zap.Error(err))
			return nil, err
		}

		return &response.WeatherResponse{OK: &weatherBody}, nil

	default:
		var errBody response.ResponseError
		if err := json.NewDecoder(resp.Body).Decode(&errBody); err != nil {
			logger.Error("failed to decode error body", zap.Error(err))
			return nil, err
		}

		return &response.WeatherResponse{Error: &errBody}, nil
	}
}
