package weatheragent

//go:generate easyjson -output_filename weatheragent_easyjson.go .
//go:generate moq -out mock.go . WeatherAgent
