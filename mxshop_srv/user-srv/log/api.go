package log

import "context"

func Debug(ctx context.Context, msg string, fields ...Field) {
	logger := DefaultLogger()
	logger.Debug(msg, fields...)
}

func Info(ctx context.Context, msg string, fields ...Field) {
	logger := DefaultLogger()
	logger.Info(msg, fields...)
}

func Error(ctx context.Context, msg string, fields ...Field) {
	logger := DefaultLogger()
	logger.Error(msg, fields...)
}

func Warn(ctx context.Context, msg string, fields ...Field) {
	logger := DefaultLogger()
	logger.Warn(msg, fields...)
}
