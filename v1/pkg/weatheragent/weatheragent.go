package weatheragent

import (
	"context"
	"net/http"

	"github.com/metaslim/weather/v1/pkg/response"
)

type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

type WeatherAgent interface {
	GetData(context.Context, string) (response.WeatherResponse, error)
}
