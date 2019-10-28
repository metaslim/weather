package weatheragent

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/metaslim/weather/v1/pkg/response"
)

type weatherStack struct {
	httpClient *http.Client
	key        string
}

type WeatherStackWeather struct {
	Temperature int `json:"temperature"`
	WindSpeed   int `json:"wind_speed"`
}

//easyjson:json
type WeatherStackResponse struct {
	Current WeatherStackWeather `json:"current"`
}

func NewWeatherStack(key string, httpClient *http.Client) weatherStack {
	return weatherStack{
		key:        key,
		httpClient: httpClient,
	}
}

func (ws weatherStack) GetData(ctx context.Context, city string) (response.WeatherResponse, error) {
	req, _ := http.NewRequest("GET", "http://api.weatherstack.com/current", nil)
	req = req.WithContext(ctx)

	q := req.URL.Query()
	q.Add("access_key", ws.key)
	q.Add("query", city)
	q.Add("unit", "m")
	req.URL.RawQuery = q.Encode()

	res, err := ws.httpClient.Do(req)
	if err != nil {
		return response.WeatherResponse{}, err
	}
	defer res.Body.Close()

	apiResponse := &WeatherStackResponse{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return response.WeatherResponse{}, err
	}

	_ = apiResponse.UnmarshalJSON(body)

	return response.WeatherResponse{
		WindSpeed:   apiResponse.Current.WindSpeed,
		Temperature: apiResponse.Current.Temperature,
	}, nil
}
