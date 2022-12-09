package logger

import (
	"context"

	"go.uber.org/zap/zapcore"
)

var staticLogger *Logger

// InitStaticLogger inits static logger singleton
func InitStaticLogger(serviceName, environment string) error {
	logger, err := New(serviceName, environment)
	if err != nil {
		return err
	}

	staticLogger = logger
	return nil
}

// Info logs at info level
func Info(ctx context.Context, message string, fields ...zapcore.Field) {
	staticLogger.Info(ctx, message, fields...)
}

// Error logs at error level
func Error(ctx context.Context, err error, fields ...zapcore.Field) {
	staticLogger.Error(ctx, err, fields...)
}
