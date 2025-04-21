package weather_test

import (
	"errors"
	"testing"

	"github.com/rxznik/golearn/weather-cli/internal/response"
	"github.com/rxznik/golearn/weather-cli/internal/service/weather"
	"github.com/rxznik/golearn/weather-cli/internal/service/weather/mocks"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestGetTodayWeather(t *testing.T) {
	testCases := []struct {
		name    string
		resp    *response.WeatherResponse
		mockErr error
	}{
		{
			name: "Success",
			resp: &response.WeatherResponse{
				OK: &response.WeatherOK{
					Hourly: &response.HourlyData{
						Time: []string{"1", "2", "3"},
						Temp: []float64{1, 2, 3},
					},
				},
			},
			mockErr: nil,
		},
		{
			name: "Error response from api",
			resp: &response.WeatherResponse{
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

			weatherClientMock := mocks.NewWeatherClient(t)
			weatherClientMock.
				On("GetWeatherData", logger, 1.0, 2.0).
				Return(tc.resp, tc.mockErr).
				Once()

			resTime, resTemp, err := weather.GetTodayWeather(logger, weatherClientMock, 1.0, 2.0)

			if tc.resp == nil {
				assert.Equal(t, tc.mockErr, err)
				return
			}

			if tc.resp.Error != nil {
				errorString := tc.resp.Error.Error + ": " + tc.resp.Error.Reason
				t.Logf("\x1b[33mexpected error %s, got %s\x1b[0m", errorString, err.Error())
				assert.Equal(t, errorString, err.Error())
			}

			if tc.resp.OK == nil && (resTime != nil || resTemp != nil) {
				t.Errorf("expected resTime and resTemp to be nil, got %v & %v", resTime, resTemp)
			}

			if tc.mockErr != nil || tc.resp.Error != nil {
				return
			}

			assert.NotNil(t, resTime)
			assert.NotNil(t, resTemp)
			assert.Equal(t, tc.resp.OK.Hourly.Time, resTime)
			assert.Equal(t, tc.resp.OK.Hourly.Temp, resTemp)
		})

	}
}
