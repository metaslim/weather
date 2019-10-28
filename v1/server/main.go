package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/kofalt/go-memoize"
	"github.com/metaslim/weather/v1/pkg/config"
	"github.com/metaslim/weather/v1/pkg/di_container"
	"github.com/metaslim/weather/v1/pkg/handler"
	"github.com/metaslim/weather/v1/pkg/logger"
	"github.com/metaslim/weather/v1/pkg/weatheragent"
	"github.com/sirupsen/logrus"
)

func main() {
	dic, err := di_container.NewDIContainer()
	if err != nil {
		panic(err)
	}

	initConfig(dic)
	initLogger(dic)
	initCache(dic, dic.Config)
	initWeatherAgents(dic, dic.Config)

	router := chi.NewRouter()

	router.Use(di_container.DependencyInjectionMiddleware(dic))

	router.HandleFunc("/v1/weather", handler.Weather)

	dic.Log.Infof("Starting Server at :%d", dic.Config.Port)
	dic.Log.Fatal(http.ListenAndServe(":"+strconv.Itoa(dic.Config.Port), router))
}

func initConfig(dic *di_container.DIContainer) {
	cfg, err := config.NewConfig()
	if err != nil {
		logrus.WithError(err).Fatal("Unable to create config", err)
	}
	dic.Config = cfg
}

func initLogger(dic *di_container.DIContainer) {
	dic.Log = logger.NewLogrus(dic.Config.Env, dic.Config.LogLevel)
}

func initCache(dic *di_container.DIContainer, cfg *config.WeatherConfig) {
	// Cache expensive calls in memory for x seconds, purging old entries every x seconds.
	cacheTime := time.Duration(cfg.CacheDurationSecond) * time.Second
	purgeTime := time.Duration(cfg.CachePurgeSecond) * time.Second
	dic.Cache = memoize.NewMemoizer(cacheTime, purgeTime)
}

func initWeatherAgents(dic *di_container.DIContainer, cfg *config.WeatherConfig) {
	httpClient := &http.Client{
		Timeout: time.Duration(cfg.HTTPClientTimeoutSecond) * time.Second,
	}

	dic.WeatherAgents = &[]weatheragent.WeatherAgent{
		weatheragent.NewOpenWeather(cfg.OpenWeatherKey, httpClient),
		weatheragent.NewWeatherStack(cfg.WeatherStackKey, httpClient),
	}
}
