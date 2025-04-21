package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/rxznik/golearn/weather-cli/internal/response"
	"go.uber.org/zap"
)

func (c *Client) GetGeoData(logger *zap.Logger, city string) (*response.GeoResponse, error) {
	url_ := fmt.Sprintf(
		"%s/search?name=%s&count=1&language=%s&format=json",
		c.GeoCfg.URL,
		url.QueryEscape(city),
		c.GeoCfg.Language,
	)

	logger.Debug("got url", zap.String("url", url_))

	resp, err := c.HTTP.Get(url_)

	if err != nil {
		logger.Error("failed to get geo data", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {

	case http.StatusOK:
		var geoBody response.GeoOK
		if err := json.NewDecoder(resp.Body).Decode(&geoBody); err != nil {
			logger.Error("failed to decode geo data", zap.Error(err))
			return nil, err
		}

		return &response.GeoResponse{OK: &geoBody}, nil

	default:
		var errBody response.ResponseError
		if err := json.NewDecoder(resp.Body).Decode(&errBody); err != nil {
			logger.Error("failed to decode error body", zap.Error(err))
			return nil, err
		}

		return &response.GeoResponse{Error: &errBody}, nil
	}
}
