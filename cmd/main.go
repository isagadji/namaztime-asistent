package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"marusya/internal/marusya"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog"

	"golang.org/x/sync/errgroup"
)

const (
	serviceName = "marusya"
)

var ErrStopped = errors.New("stopped")

func main() {
	gr, _ := errgroup.WithContext(context.Background())

	logger := httplog.NewLogger(serviceName, httplog.Options{
		LogLevel: "debug",
		JSON:     true,
	})

	marusyaService := marusya.NewService()
	transport := marusya.NewTransport(marusyaService, logger)

	gr.Go(func() error {
		router := chi.NewRouter()

		router.Use(httplog.RequestLogger(logger))
		router.Use(middleware.Heartbeat("/ping"))

		router.Use(cors.Handler(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"DNT", "Authorization", "Origin", "X-Requested-With", "X-Host", "X-Request-Id", "Timing-Allow-Origin", "Content-Type", "Accept", "Content-Range", "Range", "Keep-Alive", "User-Agent", "If-Modified-Since", "Cache-Control", "Content-Type"},
			AllowCredentials: false,
		}))

		router.Mount("/", transport.Handler())

		return http.ListenAndServe(":3000", router)
	})

	if err := gr.Wait(); err != nil && !errors.Is(err, ErrStopped) {
		logger.Error().Msg(fmt.Sprintf("unexpected error: %w", err))
	}

	logger.Info().Msg("service gracefully stopped")
}
