package weatheragent

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/metaslim/weather/v1/pkg/response"
)

type weatherStack struct {
	key string
}

type Weather struct {
	Temperature int `json:"temperature"`
	WindSpeed   int `json:"wind_speed"`
}

//easyjson:json
type Response struct {
	Current Weather `json:"current"`
}

func NewWeatherStack(key string) weatherStack {
	return weatherStack{
		key: key,
	}
}

func (ws weatherStack) GetData(ctx context.Context, city string) (response.WeatherResponse, error) {
	httpClient := http.Client{}

	req, _ := http.NewRequest("GET", "http://api.weatherstack.com/current", nil)
	q := req.URL.Query()
	q.Add("access_key", ws.key)
	q.Add("query", city)
	req.URL.RawQuery = q.Encode()

	res, err := httpClient.Do(req)
	if err != nil {
		return response.WeatherResponse{}, err
	}
	defer res.Body.Close()

	apiResponse := &Response{}
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
