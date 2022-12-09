package main

import (
	"context"
	"errors"
	logger "example-logger"
	"log"
	"time"

	"go.uber.org/zap"
)

func main() {
	logger, err := logger.New("licenseengine", "uat")
	if err != nil {
		log.Fatal("error init logger:", err)
	}

	logger.Info(context.Background(), "some info", zap.String("url", "https://sample.com"), zap.Duration("api resp", 3*time.Millisecond), zap.Int("attempt", 3))

	logger.Error(context.Background(), errors.New("some error"), zap.String("url", "https://sample.com"), zap.Duration("api resp", 3*time.Millisecond), zap.Int("attempt", 3))
}
