package response

//easyjson:json
type WeatherResponse struct {
	WindSpeed   int `json:"wind_speed"`
	Temperature int `json:"temperature_degrees"`
}

//easyjson:json
type ErrorResponse struct {
	Message string `json:"error_message"`
}
