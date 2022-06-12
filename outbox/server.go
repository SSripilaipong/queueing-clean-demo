package outbox

import (
	"context"
	"fmt"
	"queueing-clean-demo/base"
	d "queueing-clean-demo/outbox/deps"
)

type server struct {
	topicName   string
	ctx         context.Context
	cancel      context.CancelFunc
	exited      chan struct{}
	depsFactory func() d.IOutboxDeps
}

func NewServer(depsFactory func() d.IOutboxDeps, topicName string) base.IServer {
	ctx, cancel := context.WithCancel(context.Background())

	return &server{
		topicName:   topicName,
		ctx:         ctx,
		cancel:      cancel,
		exited:      make(chan struct{}),
		depsFactory: depsFactory,
	}
}

func (s *server) Start() {
	go s.serve()
}

func (s *server) Stop() error {
	s.cancel()
	<-s.exited
	return nil
}

func (s *server) serve() {
	fmt.Println("outbox started")

	deps := s.depsFactory()
	defer deps.Destroy()

	watcherLoop(s.ctx, deps.Stream(), func(data map[string]any) {
		events := extractLatestEvents(data)
		for _, event := range events {
			deps.Broker().Publish(s.topicName, event)
			fmt.Printf("published: %#v\n", event)
		}
	})

	s.exited <- struct{}{}
	fmt.Println("outbox exited")
}
