package di_container

import (
	"github.com/kofalt/go-memoize"
	"github.com/metaslim/weather/v1/pkg/config"
	"github.com/sirupsen/logrus"
)

type DIContainer struct {
	Cache  *memoize.Memoizer
	Config *config.WeatherConfig
	Log    *logrus.Logger
}

func NewDIContainer(options ...func(*DIContainer) error) (*DIContainer, error) {
	var dic DIContainer

	return &dic, nil
}
