package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	logger "github.com/darklongnightt/example-logger"
	"go.uber.org/zap"
)

func main() {
	fmt.Println("=============================================")
	fmt.Println("Logger Example")
	fmt.Println("=============================================")

	logs, err := logger.New("licenseengine", "uat")
	if err != nil {
		log.Fatal("error init logger:", err)
	}

	logs.Info(context.Background(), "some info", zap.String("url", "https://sample.com"), zap.Duration("api resp", 3*time.Millisecond), zap.Int("attempt", 3))
	logs.Error(context.Background(), errors.New("some error"), zap.String("url", "https://sample.com"), zap.Duration("api resp", 3*time.Millisecond), zap.Int("attempt", 3))
}
