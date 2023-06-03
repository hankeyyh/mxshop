package log

func Info(msg string, fields ...Field) {
	DefaultLogger().Info(msg, fields...)
}

func Debug(msg string, fields ...Field) {
	DefaultLogger().Debug(msg, fields...)
}

func Warn(msg string, fields ...Field) {
	DefaultLogger().Warn(msg, fields...)
}

func Error(msg string, fields ...Field) {
	DefaultLogger().Error(msg, fields...)
}
