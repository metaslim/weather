package weatheragent

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/metaslim/weather/v1/pkg/response"
	"github.com/stretchr/testify/assert"
)

func TestWeatherStackGetData(t *testing.T) {
	testCases := []struct {
		desc             string
		code             int
		body             string
		expectedResponse response.WeatherResponse
		expectedError    error
	}{
		{
			desc: "JSON object created from Weather Stack Payload",
			code: http.StatusOK,
			body: `{"request":{"type":"City","query":"Melbourne, Australia","language":"en","unit":"m"},"location":{"name":"Melbourne","country":"Australia","region":"Victoria","lat":"-37.817","lon":"144.967","timezone_id":"Australia\/Melbourne","localtime":"2019-10-29 08:27","localtime_epoch":1572337620,"utc_offset":"11.0"},"current":{"observation_time":"09:27 PM","temperature":17,"weather_code":113,"weather_icons":["https:\/\/assets.weatherstack.com\/images\/wsymbols01_png_64\/wsymbol_0001_sunny.png"],"weather_descriptions":["Sunny"],"wind_speed":28,"wind_degree":360,"wind_dir":"N","pressure":1019,"precip":0,"humidity":25,"cloudcover":0,"feelslike":17,"uv_index":5,"visibility":10,"is_day":"yes"}}`,
			expectedResponse: response.WeatherResponse{
				WindSpeed:   28,
				Temperature: 17,
			},
			expectedError: nil,
		},
		{
			desc:             "Empty JSON object created from Weather Stack Payload",
			code:             http.StatusInternalServerError,
			body:             ``,
			expectedResponse: response.WeatherResponse{},
			expectedError:    fmt.Errorf("CODE: [%d] body: [%s]", http.StatusInternalServerError, ""),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {

			httpClient := &ClientMock{
				DoFunc: func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: tc.code,
						Body:       ioutil.NopCloser(bytes.NewBufferString(tc.body)),
					}, nil
				},
			}

			weatherStack := NewWeatherStack("random-key", httpClient)
			reponse, err := weatherStack.GetData(context.Background(), "melbourne")

			assert.Equal(t, tc.expectedResponse, reponse)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
