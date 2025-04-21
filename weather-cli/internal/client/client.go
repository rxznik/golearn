package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/rxznik/golearn/weather-cli/internal/config"
	"github.com/rxznik/golearn/weather-cli/internal/response"

	"go.uber.org/zap"
)

type Client struct {
	HTTP       *http.Client
	GeoCfg     *config.ConfigGeo
	WeatherCfg *config.ConfigWeather
}

func NewClient(cfg *config.ConfigClient) *Client {
	httpClient := &http.Client{Timeout: cfg.Timeout}

	return &Client{
		HTTP:       httpClient,
		GeoCfg:     &cfg.Geo,
		WeatherCfg: &cfg.Weather,
	}
}

func (c *Client) GetGeoData(logger *zap.Logger, city string) (*response.Response, error) {
	url_ := fmt.Sprintf(
		"%s/search?name=%s&count=1&language=%s&format=json",
		c.GeoCfg.URL,
		url.QueryEscape(city),
		c.GeoCfg.Language,
	)

	logger.Debug("got url", zap.String("url", url_))

	resp, err := c.HTTP.Get(url_)
	// resp, err := c.HTTP.Get(url)
	if err != nil {
		logger.Error("failed to get geo data", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()

	// // remove after fix
	// // -----------------
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("\x1b[33m%s\x1b[0m\n", string(body))
	// // -----------------

	switch resp.StatusCode {

	case http.StatusOK:
		var geoBody response.GeoOK
		if err := json.NewDecoder(resp.Body).Decode(&geoBody); err != nil {
			logger.Error("failed to decode geo data", zap.Error(err))
			return nil, err
		}

		return &response.Response{OK: geoBody}, nil

	default:
		var errBody response.ResponseError
		if err := json.NewDecoder(resp.Body).Decode(&errBody); err != nil {
			logger.Error("failed to decode error body", zap.Error(err))
			return nil, err
		}

		return &response.Response{Error: errBody}, nil
	}
}

func (c *Client) GetWeatherData(logger *zap.Logger, lat, lon float64) (*response.Response, error) {
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

		return &response.Response{OK: weatherBody}, nil

	default:
		var errBody response.ResponseError
		if err := json.NewDecoder(resp.Body).Decode(&errBody); err != nil {
			logger.Error("failed to decode error body", zap.Error(err))
			return nil, err
		}

		return &response.Response{Error: errBody}, nil
	}
}
