package application

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/flexer2006/lms_gameLife/http/server"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Width  int
	Height int
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
		fmt.Printf("Ошибка настройки логгера: %v\n", err)
		return 1
	}

	shutDownFunc, err := server.Run(ctx, logger, a.Cfg.Height, a.Cfg.Width)
	if err != nil {
		logger.Error("Ошибка запуска сервера", zap.Error(err))
		return 1
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	<-c

	if err := shutDownFunc(ctx); err != nil {
		logger.Error("Ошибка завершения работы сервера", zap.Error(err))
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

	return logger, nil
}
