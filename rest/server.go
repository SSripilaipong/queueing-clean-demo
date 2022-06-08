package rest

import (
	"context"
	"net/http"
	"queueing-clean-demo/rest/deps"
	"time"
)

type Server struct {
	Deps   *deps.RestDeps
	server *http.Server
}

func (s *Server) Start(port string) {
	corsHandler := newCorsHandler()

	readTimeout := 180 * time.Second
	writeTimeout := 180 * time.Second

	s.server = newHttpServer(port, getApiRouter(s.Deps), corsHandler, readTimeout, writeTimeout)
	go s.serve()
}

func (s *Server) Stop() error {
	err := s.server.Shutdown(context.Background())
	return err
}

func (s *Server) serve() {
	err := s.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
