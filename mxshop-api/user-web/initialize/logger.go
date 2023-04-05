package initialize

import (
	"go.uber.org/zap"
	"mxshop-api/user-web/config"
)

func InitLogger() {
	conf := config.Conf.Log
	var logConfig zap.Config
	if conf.Level == "debug" {
		logConfig = zap.NewDevelopmentConfig()
	} else {
		logConfig = zap.NewProductionConfig()
	}
	logConfig.OutputPaths = []string{
		"stderr",
		conf.FilePath,
	}
	logger, err := logConfig.Build()
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(logger)
}
