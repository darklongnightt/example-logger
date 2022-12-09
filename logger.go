package logger

import (
	"context"

	"github.com/darklongnightt/example-logger/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	contextKey     = "context"
	serviceNameKey = "service"
	environmentKey = "environment"
	requestIDKey   = "requestID"
)

// Field alias to zapcore.Field for shorter declarations
type Field zapcore.Field

// Logger contains logging methods and configs set during init
type Logger struct {
	serviceName string
	environment string
	logger      *zap.Logger
}

// New creates and return new logger instance and setup common log configs
func New(serviceName, environment string) (*Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	decoratedLogger := logger.With(zap.String(serviceNameKey, serviceName)).With(zap.String(environmentKey, environment))

	return &Logger{
		serviceName: serviceName,
		environment: environment,
		logger:      decoratedLogger,
	}, nil
}

// Info logs at info level
func (l *Logger) Info(ctx context.Context, message string, fields ...zapcore.Field) {
	l.withRequestID(ctx).Info(message, fields...)

	// Do other things
}

// Error logs at error level
func (l *Logger) Error(ctx context.Context, err error, fields ...zapcore.Field) {
	l.withRequestID(ctx).Error(err.Error(), fields...)

	// Do other things
}

func (l *Logger) withRequestID(ctx context.Context) *zap.Logger {
	return l.logger.With(zap.String(requestIDKey, utils.GetCorrelationID(ctx)))
}
