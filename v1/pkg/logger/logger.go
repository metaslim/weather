package logger

import (
	"github.com/sirupsen/logrus"
)

func NewLogrus(env string, level string) *logrus.Logger {
	log := logrus.New()
	log.WithFields(logrus.Fields{
		"environment": env,
	})

	var logLevel logrus.Level
	var err error
	{
		logLevel, err = logrus.ParseLevel(level)
		if err == nil {
			log.SetLevel(logLevel)
		} else {
			log.WithFields(logrus.Fields{
				"configLogLevel":  logLevel,
				"defaultLogLevel": log.GetLevel(),
			}).Warn("Unable to parse log level, using default")
		}
	}

	return log
}
