package logger

import (
	"github.com/hankeyyh/mxshop_user_srv/config"
	"go.uber.org/zap"
)

var defaultLogger *zap.Logger

func DefaultLogger() *zap.Logger {
	return defaultLogger
}

func InitDefaultLogger() {
	logConfig := config.Conf.Log
	var err error
	if logConfig.Level == "debug" {
		zap.NewProductionConfig()
		defaultLogger, err = zap.NewDevelopment()
	} else {
		defaultLogger, err = zap.NewProduction()
	}
	defaultLogger.WithOptions()
	if err != nil {
		panic(err)
	}
}

func init() {
	InitDefaultLogger()
}
