package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/flexer2006/lms_gameLife/http/server/handler"
	"github.com/flexer2006/lms_gameLife/internal/service"
	"go.uber.org/zap"
)

// createHandler creates a new HTTP handler with routing and middleware
func createHandler(_ context.Context, logger *zap.Logger, lifeService *service.GameService) (http.Handler, error) {
	muxHandler, err := handlers.New(lifeService, logger)
	if err != nil {
		return nil, fmt.Errorf("ошибка инициализации обработчика: %w", err)
	}
	// Middleware для обработчиков
	muxHandler = handlers.Decorate(muxHandler, loggingMiddleware(logger))

	return muxHandler, nil
}

// Run starts HTTP server with specified parameters
func Run(ctx context.Context, logger *zap.Logger, height, width int) (func(context.Context) error, error) {

	lifeService := service.NewGameService(height, width)

	muxHandler, err := createHandler(ctx, logger, lifeService)
	if err != nil {
		return nil, fmt.Errorf("ошибка создания маршрутизатора: %w", err)
	}

	srv := &http.Server{
		Addr:    ":8081",
		Handler: muxHandler,
	}

	go func() {

		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("ошибка ListenAndServe", zap.Error(err))
		}
	}()

	return func(ctx context.Context) error {
		return srv.Shutdown(ctx)
	}, nil
}

// loggingMiddleware creates middleware for logging HTTP requests
func loggingMiddleware(logger *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			next.ServeHTTP(w, r)

			duration := time.Since(start)
			logger.Info("HTTP запрос",
				zap.String("метод", r.Method),
				zap.String("путь", r.URL.Path),
				zap.Duration("длительность", duration),
			)
		})
	}
}
