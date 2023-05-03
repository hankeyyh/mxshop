package log

import (
	"go.uber.org/zap"
	"mxshop-api/user-web/config"
)

var defaultLogger *zap.Logger

func DefaultLogger() *zap.Logger {
	return defaultLogger
}

func Init() {
	logCong := config.DefaultConfig().Log
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
