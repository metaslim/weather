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

func TestOpenWeatherGetData(t *testing.T) {
	testCases := []struct {
		desc             string
		code             int
		body             string
		expectedResponse response.WeatherResponse
		expectedError    error
	}{
		{
			desc: "JSON object created from Open Weather Payload",
			code: http.StatusOK,
			body: `{"coord":{"lon":144.96,"lat":-37.81},"weather":[{"id":803,"main":"Clouds","description":"broken clouds","icon":"04n"}],"base":"stations","main":{"temp":17.07,"pressure":1019,"humidity":22,"temp_min":13.33,"temp_max":20},"visibility":10000,"wind":{"speed":8.7,"deg":360},"clouds":{"all":55},"dt":1572299569,"sys":{"type":1,"id":9554,"country":"AU","sunrise":1572290245,"sunset":1572339014},"timezone":39600,"id":2158177,"name":"Melbourne","cod":200}`,
			expectedResponse: response.WeatherResponse{
				WindSpeed:   8,
				Temperature: 17,
			},
			expectedError: nil,
		},
		{
			desc:             "Empty JSON object created from Open Weather Payload",
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

			openWeather := NewOpenWeather("random-key", httpClient)
			reponse, err := openWeather.GetData(context.Background(), "melbourne")

			assert.Equal(t, tc.expectedResponse, reponse)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
