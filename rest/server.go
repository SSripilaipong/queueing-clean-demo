package rest

import (
	"context"
	"fmt"
	"github.com/rs/cors"
	"net/http"
	"queueing-clean-demo/base"
	"queueing-clean-demo/rest/deps"
	"time"
)

type Server struct {
	depsFactory  func() deps.IRestDeps
	port         string
	corsHandler  *cors.Cors
	readTimeout  time.Duration
	writeTimeout time.Duration
	server       *http.Server
}

func NewServer(depsFactory func() deps.IRestDeps, port string) base.IServer {
	return &Server{
		depsFactory:  depsFactory,
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
	fmt.Println("rest started")

	restDeps := s.depsFactory()
	defer restDeps.Destroy()

	s.server = newHttpServer(s.port, getApiRouter(restDeps), s.corsHandler, s.readTimeout, s.writeTimeout)

	err := s.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(err)
	}

	fmt.Println("rest exited")
}
