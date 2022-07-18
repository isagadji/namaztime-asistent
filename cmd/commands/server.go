package commands

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"marusya/internal/aladhan"
	"marusya/internal/namaztime"

	"github.com/alecthomas/kong"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog"
	"golang.org/x/sync/errgroup"
)

var errStopped = errors.New("stopped")

type Server struct {
	VKAppID   string `kong:"required,name=vk-app-id,env=VK_APP_ID"`
	MarusyaID string `kong:"required,name=namaztime-id,env=MARUSYA_ID"`
	HTTPAddr  string `kong:"required,name=http-addr,env=HTTP_ADDR,group='HTTP server'"`
	LogLevel  string `kong:"optional,name=log-level,env=LOG_LEVEL,default=info"`

	Aladhan       aladhan.Flags `kong:"embed"`
	PostgresFlags PostgresFlags `kong:"embed"`
}

func (s *Server) Run(kVars kong.Vars) error {
	serviceName := kVars["serviceName"]
	logger := httplog.NewLogger(serviceName, httplog.Options{
		LogLevel: s.LogLevel,
		JSON:     true,
	})
	db, err := s.PostgresFlags.Init()
	if err != nil {
		return err
	}
	storage := namaztime.NewStorage(db)
	aladhanService := s.Aladhan.Init()
	marusyaService := namaztime.NewService(aladhanService, storage, logger)
	transport := namaztime.NewTransport(marusyaService, logger)

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

	if err := gr.Wait(); err != nil && !errors.Is(err, errStopped) {
		logger.Error().Msg(fmt.Sprintf("unexpected error: %v", err))
	}

	logger.Info().Msg("service gracefully stopped")
	return nil
}
