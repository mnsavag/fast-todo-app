package log

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func NewLogger(conf Config) (*logrus.Logger, error) {
	if err := conf.Validate(); err != nil {
		return nil, fmt.Errorf("cant create logger: %w", err)
	}

	logger := logrus.New()

	switch conf.LogEnv {
	case envLocal:
		logger.SetLevel(logrus.DebugLevel)
	case envProd:
		logger.SetLevel(logrus.InfoLevel)
	}
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.JSONFormatter{})
	return logger, nil
}
