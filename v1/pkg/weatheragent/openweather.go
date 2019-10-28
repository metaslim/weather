package weatheragent

import (
	"context"

	"github.com/briandowns/openweathermap"
	"github.com/metaslim/weather/v1/pkg/response"
)

type openWeather struct {
	key string
}

func NewOpenWeather(key string) openWeather {
	return openWeather{
		key: key,
	}
}

func (ow openWeather) GetData(ctx context.Context, city string) (response.WeatherResponse, error) {
	w, err := openweathermap.NewCurrent("C", "EN", ow.key)
	if err != nil {
		return response.WeatherResponse{}, err
	}
	w.CurrentByName(city)

	return response.WeatherResponse{
		WindSpeed:   int(w.Wind.Speed),
		Temperature: int(w.Main.Temp),
	}, nil
}
