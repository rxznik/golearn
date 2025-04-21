package response

type GeoOK struct {
	Results []GeoData `json:"results"`
}

type GeoData struct {
	ID          int      `json:"-"`
	Name        string   `json:"-"`
	Latitude    float64  `json:"latitude"`
	Longitude   float64  `json:"longitude"`
	Elevation   float64  `json:"-"`
	FeatureCode string   `json:"-"`
	CountryCode string   `json:"-"`
	Admin1ID    int      `json:"-"`
	Admin2ID    int      `json:"-"`
	Admin3ID    int      `json:"-"`
	Admin4ID    int      `json:"-"`
	TimeZone    string   `json:"-"`
	Population  int      `json:"-"`
	Postcodes   []string `json:"-"`
	CountryID   int      `json:"-"`
	Admin1      string   `json:"-"`
	Admin2      string   `json:"-"`
	Admin3      string   `json:"-"`
	Admin4      string   `json:"-"`
}
