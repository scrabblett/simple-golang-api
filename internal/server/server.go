package server

import (
	"context"
	"net/http"
	"simple-golang-api/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         ":8080",
			Handler:      handler,
			ReadTimeout:  cfg.Server.Timeout,
			WriteTimeout: cfg.Server.Timeout,
			IdleTimeout:  cfg.Server.IdleTimeout,
		},
	}
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
