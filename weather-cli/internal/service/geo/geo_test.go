package geo_test

import (
	"errors"
	"testing"
	"github.com/rxznik/golearn/weather-cli/internal/response"
	"github.com/rxznik/golearn/weather-cli/internal/service/geo"
	"github.com/rxznik/golearn/weather-cli/internal/service/geo/mocks"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestCityCoordinates(t *testing.T) {
	testCases := []struct {
		name    string
		resp    *response.Response
		mockErr error
	}{
		{
			name: "Success",
			resp: &response.Response{
				OK: &response.GeoOK{
					Results: []response.GeoData{
						{
							Latitude:  1.0,
							Longitude: 2.0,
						},
					},
				},
			},
			mockErr: nil,
		},
		{
			name: "Error response from api",
			resp: &response.Response{
				Error: response.ResponseError{
					Error:  "Error code 1488",
					Reason: "Some reason",
				},
			},
			mockErr: nil,
		},
		{
			name:    "Error from mock",
			resp:    nil,
			mockErr: errors.New("mock error"),
		},
	}

	logger := zap.NewNop()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			geoClientMock := mocks.NewGeoClient(t)
			geoClientMock.
				On("GetGeoData", logger, "test").
				Return(tc.resp, tc.mockErr)

			coordinates, err := geo.GetCityCoordinates(logger, geoClientMock, "test")

			assert.Equal(t, tc.mockErr, err)
			if tc.resp.Error.Error != "" {
				errorString := tc.resp.Error.Error + ": " + tc.resp.Error.Reason
				assert.Equal(t, errorString, err.Error())
			}

			if tc.resp.OK == nil && coordinates != nil {
				t.Errorf("expected coordinates to be nil, got %v", coordinates)
			}

			if tc.mockErr != nil && tc.resp.Error.Error != "" {
				return
			}

			trueCoordinates := []float64{
				tc.resp.OK.(response.GeoOK).Results[0].Latitude,
				tc.resp.OK.(response.GeoOK).Results[0].Longitude,
			}

			assert.Equal(t, trueCoordinates, coordinates)
		})
	}
}
