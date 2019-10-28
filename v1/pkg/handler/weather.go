package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/metaslim/weather/v1/pkg/di_container"
	"github.com/metaslim/weather/v1/pkg/response"
	"github.com/metaslim/weather/v1/pkg/weatheragent"
)

func Weather(w http.ResponseWriter, r *http.Request) {
	cache := di_container.DIC(r.Context()).Cache

	city := r.URL.Query()["city"]
	if len(city) < 1 || city[0] == "" {
		RespondJSON(w, response.ErrorResponse{
			Message: "Missing city",
		}, http.StatusInternalServerError)
	}

	reponse, err, _ := cache.Memoize(city[0], func() (interface{}, error) {
		result, err := getData(r.Context(), city[0])
		return result, err
	})

	if err != nil {
		RespondJSON(w, response.ErrorResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
	}
	RespondJSON(w, reponse.(response.WeatherResponse), http.StatusOK)

}

func getData(ctx context.Context, city string) (response.WeatherResponse, error) {
	agents := []weatheragent.WeatherAgent{
		weatheragent.OpenWeather{},
		weatheragent.WeatherStack{},
	}

	for _, agent := range agents {
		response, err := agent.GetData(ctx, city)
		if err == nil {
			return response, nil
		}
	}

	return response.WeatherResponse{}, fmt.Errorf("Unable to get any weather data")
}
