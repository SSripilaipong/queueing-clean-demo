package rest

import (
	"context"
	"github.com/rs/cors"
	"net/http"
	"queueing-clean-demo/base"
	"queueing-clean-demo/rest/deps"
	"time"
)

type Server struct {
	deps         *deps.RestDeps
	port         string
	corsHandler  *cors.Cors
	readTimeout  time.Duration
	writeTimeout time.Duration
	server       *http.Server
}

func NewServer(deps *deps.RestDeps, port string) base.IServer {
	return &Server{
		deps:         deps,
		port:         port,
		corsHandler:  newCorsHandler(),
		readTimeout:  180 * time.Second,
		writeTimeout: 180 * time.Second,
	}
}

func (s *Server) Start() {
	go s.serve()
}

func (s *Server) Stop() error {
	err := s.server.Shutdown(context.Background())
	return err
}

func (s *Server) serve() {
	s.server = newHttpServer(s.port, getApiRouter(s.deps), s.corsHandler, s.readTimeout, s.writeTimeout)

	err := s.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
