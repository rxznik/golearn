package response

type ResponseError struct {
	Error  string `json:"error" required:"true"`
	Reason string `json:"reason" required:"true"`
}

type Response struct {
	OK    any
	Error ResponseError
}

type GeoResponse struct {
	OK    *GeoOK
	Error *ResponseError
}

type WeatherResponse struct {
	OK    *WeatherOK
	Error *ResponseError
}
