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
		resp    *response.GeoResponse
		mockErr error
	}{
		{
			name: "Success",
			resp: &response.GeoResponse{
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
			resp: &response.GeoResponse{
				Error: &response.ResponseError{
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
	defer logger.Sync()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			t.Parallel()

			geoClientMock := mocks.NewGeoClient(t)
			geoClientMock.
				On("GetGeoData", logger, "test").
				Return(tc.resp, tc.mockErr).
				Once()

			coordinates, err := geo.GetCityCoordinates(logger, geoClientMock, "test")

			if tc.resp == nil {
				assert.Equal(t, tc.mockErr, err)
				return
			}

			if tc.resp.Error != nil {
				errorString := tc.resp.Error.Error + ": " + tc.resp.Error.Reason
				assert.Equal(t, errorString, err.Error())
			}

			if tc.resp.OK == nil && coordinates != nil {
				t.Errorf("expected coordinates to be nil, got %v", coordinates)
			}

			if tc.mockErr != nil || tc.resp.Error != nil {
				return
			}

			assert.NotNil(t, coordinates)
			assert.Len(t, coordinates, 2)

			trueCoordinates := []float64{
				tc.resp.OK.Results[0].Latitude,
				tc.resp.OK.Results[0].Longitude,
			}

			assert.Equal(t, trueCoordinates, coordinates)
		})
	}
}
