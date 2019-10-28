package config

import (
	"github.com/kelseyhightower/envconfig"
)

const EnvVarsPrefix string = ""
const AppName string = "weather"

// WeatherConfig contains all the configuration for a particular runtime environment
type WeatherConfig struct {
	Env                     string `env:"API_ENV" default:"development"`
	Port                    int    `env:"API_PORT" default:"8080"`
	LogLevel                string `env:"API_LOG_LEVEL" default:"debug"`
	OpenWeatherKey          string `env:"API_OPEN_WEATHER_KEY" default:"fee8ba54caa15f294e33013a5756981a"`
	WeatherStackKey         string `env:"API_WEATHER_STACK_KEY" default:"1f7c9848bb5082b82e6f3501cf7eeb5c"`
	CacheDurationSecond     int    `env:"API_CACHE_DURATION_SECOND" default:"3"`
	CachePurgeSecond        int    `env:"API_CACHE_PURGE_SECOND" default:"60"`
	HTTPClientTimeoutSecond int    `env:"API_HTTP_TIMEOUT_SECOND" default:"3"`
}

func NewConfig() (*WeatherConfig, error) {
	var cfg WeatherConfig

	err := envconfig.Process(EnvVarsPrefix, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func StringWithDefault(val string, def string) string {
	if val == "" {
		return def
	}
	return val
}
