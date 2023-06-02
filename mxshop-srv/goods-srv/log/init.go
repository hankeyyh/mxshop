package log

import "go.uber.org/zap"

var (
	defaultLogger *zap.Logger
)

func DefaultLogger() *zap.Logger {
	return defaultLogger
}

func Init() error {
	conf := zap.NewDevelopmentConfig()
	var err error
	defaultLogger, err = conf.Build()
	if err != nil {
		return err
	}
	return nil
}
