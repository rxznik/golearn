package geo

import (
	"errors"

	"github.com/rxznik/golearn/weather-cli/internal/response"

	"go.uber.org/zap"
)

type GeoClient interface {
	GetGeoData(logger *zap.Logger, city string) (*response.GeoResponse, error)
}

func GetCityCoordinates(logger *zap.Logger, client GeoClient, city string) ([]float64, error) {
	resp, err := client.GetGeoData(logger, city)
	if err != nil {
		logger.Error("failed to get geo data", zap.Error(err))
		return nil, err
	}

	if resp.Error != nil {
		errorMsg := resp.Error.Error + ": " + resp.Error.Reason
		err := errors.New(errorMsg)
		logger.Error("failed to get geo data", zap.Error(err))
		return nil, err
	}

	logger.Debug("got geo data from api", zap.Any("data", resp))

	if resp.OK == nil {
		errorMsg := "invalid geo data"
		logger.Error(errorMsg)
		return nil, errors.New(errorMsg)
	}

	geoResults := resp.OK

	if len(geoResults.Results) == 0 {
		errorMsg := "no geo results found"
		logger.Error(errorMsg)
		return nil, errors.New(errorMsg)
	}

	geoData := geoResults.Results[0]

	return []float64{geoData.Latitude, geoData.Longitude}, nil
}
