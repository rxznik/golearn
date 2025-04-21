package client

import (
	"net/http"

	"github.com/rxznik/golearn/weather-cli/internal/config"
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
