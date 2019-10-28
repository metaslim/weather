package config

import (
	"github.com/kelseyhightower/envconfig"
)

const EnvVarsPrefix string = ""
const AppName string = "weather"

// WeatherConfig contains all the configuration for a particular runtime environment
type WeatherConfig struct {
	Env                 string `default:"development"`
	Port                int    `default:"8080"`
	LogLevel            string `default:"debug"`
	WeatherStackKey     string `default:"1f7c9848bb5082b82e6f3501cf7eeb5c"`
	OpenWeatherKey      string `default:"fee8ba54caa15f294e33013a5756981a"`
	CacheDurationSecond int    `default:"3"`
	CachePurgeSecond    int    `default:"60"`
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
