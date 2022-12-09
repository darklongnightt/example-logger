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
	fmt.Println("Static Logger Example")
	fmt.Println("=============================================")

	if err := logger.InitStaticLogger("licenseengine", "uat"); err != nil {
		log.Fatal("error init static logger:", err)
	}

	logger.Info(context.Background(), "some info", zap.String("url", "https://sample.com"), zap.Duration("api resp", 3*time.Millisecond), zap.Int("attempt", 3))
	logger.Error(context.Background(), errors.New("some error"), zap.String("url", "https://sample.com"), zap.Duration("api resp", 3*time.Millisecond), zap.Int("attempt", 3))
}
