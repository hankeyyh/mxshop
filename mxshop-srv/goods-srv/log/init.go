package log

import (
	"github.com/hankeyyh/mxshop/mxshop-srv/goods-srv/config"
	"go.uber.org/zap"
)

var (
	defaultLogger *zap.Logger
)

func DefaultLogger() *zap.Logger {
	return defaultLogger
}

func Init() error {
	level := config.DefaultConfig.Log.Level
	var conf zap.Config
	if level == "debug" {
		conf = zap.NewDevelopmentConfig()
	} else {
		conf = zap.NewProductionConfig()
	}
	conf.OutputPaths = []string{
		"stderr",
		config.DefaultConfig.Log.FilePath,
	}
	var err error
	defaultLogger, err = conf.Build()
	if err != nil {
		return err
	}
	return nil
}

func init() {
	Init()
}
