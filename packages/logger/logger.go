package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// SilentError logs an error without stack trace
func SilentError(err error) zap.Field {
	// you could even flip the encoding depending on the environment here if you want
	return zap.Field{Key: "error", Type: zapcore.StringType, String: err.Error()}
}

// New creates a new zap.Logger with the given log level
func New(level string) *zap.Logger {
	var (
		logLevel  zapcore.Level
		zapConfig zap.Config
	)

	switch level {
	case "DEBUG":
		logLevel = zapcore.DebugLevel
	case "INFO":
		logLevel = zapcore.InfoLevel
	case "WARN", "WARNING":
		logLevel = zapcore.WarnLevel
	case "ERROR":
		logLevel = zapcore.ErrorLevel
	default:
		logLevel = zapcore.InfoLevel
	}

	if level == "DEBUG" {
		zapConfig = zap.NewDevelopmentConfig()
	} else {
		zapConfig = zap.NewProductionConfig()
	}

	zapConfig.Level.SetLevel(logLevel)

	log, _ := zapConfig.Build(zap.AddStacktrace(zapcore.PanicLevel))
	return log
}
