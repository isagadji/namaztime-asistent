package commands

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"marusya/internal/marusya"

	"github.com/alecthomas/kong"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog"
	"golang.org/x/sync/errgroup"
)

var ErrStopped = errors.New("stopped")

type Server struct {
	AppId    string `kong:"required,name=app-id,env=APP_ID"`
	HTTPAddr string `kong:"required,name=http-addr,env=HTTP_ADDR"`
	Aladhan  struct {
		ApiHost string `kong:"required,name=aladhan-api-host,env=ALADHAN_API_HOST"`
		ApiPath string `kong:"required,name=aladhan-api-path,env=ALADHAN_API_PATH"`
	} `embed:""`
}

func (s *Server) Run(kVars kong.Vars) error {
	serviceName := kVars["serviceName"]
	logger := httplog.NewLogger(serviceName, httplog.Options{
		LogLevel: "debug",
		JSON:     true,
	})

	marusyaService := marusya.NewService()
	transport := marusya.NewTransport(marusyaService, logger)

	gr, _ := errgroup.WithContext(context.Background())
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

		return http.ListenAndServe(s.HTTPAddr, router)
	})

	if err := gr.Wait(); err != nil && !errors.Is(err, ErrStopped) {
		logger.Error().Msg(fmt.Sprintf("unexpected error: %w", err))
	}

	logger.Info().Msg("service gracefully stopped")
	return nil
}
