package weatheragent

import (
	"context"

	"github.com/metaslim/weather/v1/pkg/response"
)

type WeatherAgent interface {
	GetData(context.Context, string) (response.WeatherResponse, error)
}
