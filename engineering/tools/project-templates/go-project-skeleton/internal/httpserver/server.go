package httpserver

import (
	"context"
	"net/http"

	"go-project-skeleton/internal/config"
)

type Server struct {
	server *http.Server
}

func New(addr string, handler http.Handler, cfg config.HTTPServerConfig) *Server {
	return &Server{
		server: &http.Server{
			Addr:         addr,
			Handler:      handler,
			ReadTimeout:  cfg.ReadTimeout,
			WriteTimeout: cfg.WriteTimeout,
			IdleTimeout:  cfg.IdleTimeout,
		},
	}
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
