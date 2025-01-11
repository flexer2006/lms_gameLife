package server

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/flexer2006/lms_gameLife/http/server/handler"
	"github.com/flexer2006/lms_gameLife/internal/service"
	"go.uber.org/zap"
)

func createHandler(_ context.Context, logger *zap.Logger, lifeService *service.GameService) (http.Handler, error) {
	muxHandler, err := handlers.New(lifeService, logger)
	if err != nil {
		logger.Error("handler initialization error", zap.Error(err))
		return nil, err
	}

	muxHandler = handlers.Decorate(muxHandler, loggingMiddleware(logger))

	return muxHandler, nil
}

func Run(ctx context.Context, logger *zap.Logger) (func(context.Context) error, error) {
	lifeService := service.NewGameService()

	muxHandler, err := createHandler(ctx, logger, lifeService)
	if err != nil {
		logger.Error("router creation error", zap.Error(err))
		return nil, err
	}

	srv := &http.Server{
		Addr:    ":8081",
		Handler: muxHandler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("ListenAndServe error", zap.Error(err))
		}
	}()

	return func(ctx context.Context) error {
		return srv.Shutdown(ctx)
	}, nil
}

func loggingMiddleware(logger *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			next.ServeHTTP(w, r)

			duration := time.Since(start)
			logger.Info("HTTP request",
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
				zap.Duration("duration", duration),
			)
		})
	}
}
