package log

import (
	"github.com/hankeyyh/mxshop-srv/user-srv/config"
	"go.uber.org/zap"
)

var defaultLogger *zap.Logger

func DefaultLogger() *zap.Logger {
	return defaultLogger
}

func InitDefaultLogger() {
	logConfig := config.Conf.Log
	var err error
	var defaultConf zap.Config
	if logConfig.Level == "debug" {
		defaultConf = zap.NewProductionConfig()
	} else {
		defaultConf = zap.NewDevelopmentConfig()
	}
	defaultConf.OutputPaths = []string{
		"stderr",
		logConfig.FilePath,
	}
	defaultLogger, err = defaultConf.Build()
	if err != nil {
		panic(err)
	}
}

func init() {
	InitDefaultLogger()
}
