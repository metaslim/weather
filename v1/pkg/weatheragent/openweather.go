package weatheragent

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/metaslim/weather/v1/pkg/response"
)

type openWeather struct {
	httpClient *http.Client
	key        string
}

type OpenWeatherWind struct {
	WindSpeed float64 `json:"speed"`
}

type OpenWeatherMain struct {
	Temperature float64 `json:"temp"`
}

//easyjson:json
type OpenWeatherResponse struct {
	Main OpenWeatherMain `json:"main"`
	Wind OpenWeatherWind `json:"wind"`
}

func NewOpenWeather(key string, httpClient *http.Client) openWeather {
	return openWeather{
		key:        key,
		httpClient: httpClient,
	}
}

func (ow openWeather) GetData(ctx context.Context, city string) (response.WeatherResponse, error) {
	req, _ := http.NewRequest("GET", "http://api.openweathermap.org/data/2.5/weather", nil)
	req = req.WithContext(ctx)

	q := req.URL.Query()
	q.Add("appid", ow.key)
	q.Add("q", city)
	q.Add("units", "metric")
	req.URL.RawQuery = q.Encode()

	res, err := ow.httpClient.Do(req)
	if err != nil {
		return response.WeatherResponse{}, err
	}
	defer res.Body.Close()

	apiResponse := &OpenWeatherResponse{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return response.WeatherResponse{}, err
	}

	_ = apiResponse.UnmarshalJSON(body)

	return response.WeatherResponse{
		WindSpeed:   int(apiResponse.Wind.WindSpeed),
		Temperature: int(apiResponse.Main.Temperature),
	}, nil
}
