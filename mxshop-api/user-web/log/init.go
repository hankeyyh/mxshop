package log

import (
	"github.com/hankeyyh/mxshop/mxshop-api/user-web/config"
	"go.uber.org/zap"
)

var defaultLogger *zap.Logger

func DefaultLogger() *zap.Logger {
	return defaultLogger
}

func Init() {
	logCong := config.DefaultConfig.Log
	var logConfig zap.Config
	if logCong.Level == "debug" {
		logConfig = zap.NewDevelopmentConfig()
	} else {
		logConfig = zap.NewProductionConfig()
	}
	logConfig.OutputPaths = []string{
		"stderr",
		logCong.FilePath,
	}
	var err error
	defaultLogger, err = logConfig.Build()
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(defaultLogger)
}
