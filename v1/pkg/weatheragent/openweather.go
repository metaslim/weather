package weatheragent

import (
	"context"

	owm "github.com/briandowns/openweathermap"
	"github.com/metaslim/weather/v1/pkg/di_container"
	"github.com/metaslim/weather/v1/pkg/response"
)

type OpenWeather struct{}

func (c OpenWeather) GetData(ctx context.Context, city string) (response.WeatherResponse, error) {
	cfg := di_container.DIC(ctx).Config
	log := di_container.DIC(ctx).Log

	w, err := owm.NewCurrent("C", "EN", cfg.OpenWeatherKey)
	if err != nil {
		log.Error(err)
		return response.WeatherResponse{}, err
	}
	w.CurrentByName(city)

	return response.WeatherResponse{
		WindSpeed:   int(w.Wind.Speed),
		Temperature: int(w.Main.Temp),
	}, nil
}
