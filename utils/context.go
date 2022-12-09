package utils

import (
	"context"

	"github.com/google/uuid"
)

const (
	reqIDKey = "reqIdCtxKeyType"
)

// WithCorrelationID binds correlation id to context variable
func WithCorrelationID(ctx context.Context, reqID string) context.Context {
	return context.WithValue(ctx, reqIDKey, reqID)
}

// GetCorrelationID get correlation id from context variable or generates a new one if not exists
func GetCorrelationID(ctx context.Context) string {
	correlationId, ok := ctx.Value(reqIDKey).(string)
	if !ok {
		return uuid.New().String()
	}
	return correlationId
}
