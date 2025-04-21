package geo

import (
	"errors"
	"github.com/rxznik/golearn/weather-cli/internal/response"

	"go.uber.org/zap"
)

type GeoClient interface {
	GetGeoData(logger *zap.Logger, city string) (*response.Response, error)
}

func GetCityCoordinates(logger *zap.Logger, client GeoClient, city string) ([]float64, error) {
	resp, err := client.GetGeoData(logger, city)
	if err != nil {
		if resp.Error.Error != "" && resp.Error.Reason != "" {
			errorMsg := resp.Error.Error + ": " + resp.Error.Reason
			logger.Error(errorMsg, zap.Error(err))
			return nil, errors.New(errorMsg)
		}

		return nil, err
	}

	logger.Debug("got geo data from api", zap.Any("data", resp))

	geoResults, ok := resp.OK.(response.GeoOK)
	if !ok {
		errorMsg := "invalid geo results"
		logger.Error(errorMsg)
		return nil, errors.New(errorMsg)
	}

	if len(geoResults.Results) == 0 {
		errorMsg := "no geo results found"
		logger.Error(errorMsg)
		return nil, errors.New(errorMsg)
	}

	geoData := geoResults.Results[0]

	return []float64{geoData.Latitude, geoData.Longitude}, nil
}
