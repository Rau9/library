package logger

import (
	"github.com/Rau9/library/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewProduction is a zap NewProduction wrapper with custom config
func NewProduction(options ...zap.Option) (*zap.Logger, error) {
	return cfg().Build(options...)
}

func cfg() zap.Config {
	initialFields := map[string]interface{}{
		"service": config.GetString("SERVICE_NAME"),
		"env":     config.GetString("ENVIRONMENT"),
	}
	logOutputPaths := []string{"stdout"}
	if config.GetString("ENVIRONMENT") == "production" {
		// In production, log as well to a logfile that can be picked up by a sidecar container
		logOutputPaths = append(logOutputPaths, config.GetString("LOG_FILE"))
	}
	sampling := &zap.SamplingConfig{
		Initial:    100,
		Thereafter: 100,
	}
	encoderConfig := &zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "@timestamp",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     func(string, zapcore.PrimitiveArrayEncoder) {},
	}
	return zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       config.GetString("ENVIRONMENT") != "production",
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          sampling,
		Encoding:          "json",
		EncoderConfig:     *encoderConfig,
		OutputPaths:       logOutputPaths,
		ErrorOutputPaths:  logOutputPaths,
		InitialFields:     initialFields, // NOTE: fields to add atop of the root logger
	}
}
