package handler

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/kofalt/go-memoize"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/metaslim/weather/v1/pkg/di_container"
	"github.com/metaslim/weather/v1/pkg/response"
	"github.com/metaslim/weather/v1/pkg/weatheragent"
)

func TestWeather(t *testing.T) {
	testCases := []struct {
		desc                 string
		openWeatherResponse  response.WeatherResponse
		openWeatherError     error
		weatherStackResponse response.WeatherResponse
		weatherStackError    error
		expectedCode         int
		expectedBody         string
	}{
		{
			desc: "Return 200 and use reponse from Weather Response if both weather agents are fine",
			openWeatherResponse: response.WeatherResponse{
				WindSpeed:   10,
				Temperature: 20,
			},
			openWeatherError: nil,
			weatherStackResponse: response.WeatherResponse{
				WindSpeed:   30,
				Temperature: 40,
			},
			weatherStackError: nil,
			expectedCode:      http.StatusOK,
			expectedBody:      `{"wind_speed":10,"temperature_degrees":20}`,
		},
		{
			desc: "Return 200 and use reponse from Open Weather if Open Weather is working and Weather Stack is not working",
			openWeatherResponse: response.WeatherResponse{
				WindSpeed:   10,
				Temperature: 20,
			},
			openWeatherError:     nil,
			weatherStackResponse: response.WeatherResponse{},
			weatherStackError:    fmt.Errorf("error from Weather Stack"),
			expectedCode:         http.StatusOK,
			expectedBody:         `{"wind_speed":10,"temperature_degrees":20}`,
		},
		{
			desc:                "Return 200 and use reponse from Weather Stack if Open Weather is not working",
			openWeatherResponse: response.WeatherResponse{},
			openWeatherError:    fmt.Errorf("error from Open Weather"),
			weatherStackResponse: response.WeatherResponse{
				WindSpeed:   30,
				Temperature: 40,
			},
			weatherStackError: nil,
			expectedCode:      http.StatusOK,
			expectedBody:      `{"wind_speed":30,"temperature_degrees":40}`,
		},
		{
			desc:                 "Return 500 and use reponse from weatherStackResponse",
			openWeatherResponse:  response.WeatherResponse{},
			openWeatherError:     fmt.Errorf("error from Open Weather"),
			weatherStackResponse: response.WeatherResponse{},
			weatherStackError:    fmt.Errorf("error from Weather Stack"),
			expectedCode:         http.StatusInternalServerError,
			expectedBody:         `{"error_message":"Unable to get any weather data"}`,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			diContainer := di_container.DIContainer{}

			diContainer.Log = logrus.New()

			diContainer.WeatherAgents = &[]weatheragent.WeatherAgent{
				&weatheragent.WeatherAgentMock{
					GetDataFunc: func(in1 context.Context, in2 string) (response.WeatherResponse, error) {
						return tc.openWeatherResponse, tc.openWeatherError
					},
				},
				&weatheragent.WeatherAgentMock{
					GetDataFunc: func(in1 context.Context, in2 string) (response.WeatherResponse, error) {
						return tc.weatherStackResponse, tc.weatherStackError
					},
				},
			}

			diContainer.Cache = memoize.NewMemoizer(1*time.Second, 1*time.Second)
			diContainer.Cache.Storage.Flush()

			req, _ := http.NewRequest(http.MethodGet, "/weather", nil)
			params := req.URL.Query()
			params.Add("city", "melbourne")
			req.URL.RawQuery = params.Encode()
			req = req.WithContext(di_container.ContextWithDIC(req.Context(), &diContainer))

			recorder := httptest.NewRecorder()
			handler := http.HandlerFunc(Weather)

			handler.ServeHTTP(recorder, req)

			body, _ := ioutil.ReadAll(recorder.Body)

			assert.Equal(t, tc.expectedCode, recorder.Code)
			assert.Equal(t, tc.expectedBody, string(body))
		})
	}
}
