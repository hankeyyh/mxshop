package initialize

import (
	"go.uber.org/zap"
	"mxshop-api/user-web/global"
)

func InitLogger() {
	logCong := global.ServerConfig.Log
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
	logger, err := logConfig.Build()
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(logger)
	zap.S().Info("InitLogger Suc")
}
