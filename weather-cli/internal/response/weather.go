package response

type WeatherOK struct {
	Latitude       float64 `json:"-"`
	Longitude      float64 `json:"-"`
	Elevation      float64 `json:"-"`
	GenerationTime float64 `json:"-"`
	UTCOffset      int     `json:"-"`
	TimeZone       string  `json:"-"`
	TimeZoneAbbr   string  `json:"-"`

	Hourly      *HourlyData      `json:"hourly"`
	HourlyUnits *HourlyUnitsData `json:"-"`
}

type HourlyData struct {
	Time []string  `json:"time"`
	Temp []float64 `json:"temperature_2m"`
}

type HourlyUnitsData struct {
	Temp2m string `json:"-"`
}
