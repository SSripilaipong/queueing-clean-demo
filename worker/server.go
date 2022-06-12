package worker

import (
	"context"
	"queueing-clean-demo/base"
)

type server struct {
	ctx    context.Context
	cancel context.CancelFunc
	exited chan struct{}
}

func NewServer() base.IServer {
	ctx, cancel := context.WithCancel(context.Background())

	return &server{
		ctx:    ctx,
		cancel: cancel,
		exited: make(chan struct{}),
	}
}

func (s *server) Start() {
	go RunWorker(s.ctx)
}

func (s *server) Stop() error {
	s.cancel()
	<-s.exited
	return nil
}
