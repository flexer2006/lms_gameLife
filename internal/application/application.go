package application

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/flexer2006/lms_gameLife/http/server"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
}

type Application struct {
	Cfg Config
}

func New(config Config) *Application {
	return &Application{
		Cfg: config,
	}
}

func (a *Application) Run(ctx context.Context) int {
	logger, err := setupLogger()
	if err != nil {
		log.Printf("Logger setup error: %v\n", err)
		return 1
	}

	if logger == nil {
		log.Println("Logger is nil")
		return 1
	}

	shutDownFunc, err := server.Run(ctx, logger)
	if err != nil {
		logger.Error("Server startup error", zap.Error(err))
		return 1
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	<-c

	if err := shutDownFunc(ctx); err != nil {
		logger.Error("Server termination error", zap.Error(err))
		return 1
	}

	return 0
}

func setupLogger() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	if logger == nil {
		return nil, fmt.Errorf("logger is nil")
	}

	return logger, nil
}
