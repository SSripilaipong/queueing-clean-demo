package worker

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"queueing-clean-demo/base"
	d "queueing-clean-demo/worker/deps"
)

type server struct {
	topicName   string
	depsFactory func() d.IWorkerDeps
	ctx         context.Context
	cancel      context.CancelFunc
	exited      chan struct{}
}

func NewServer(depsFactory func() d.IWorkerDeps, topicName string) base.IServer {
	ctx, cancel := context.WithCancel(context.Background())

	return &server{
		topicName:   topicName,
		depsFactory: depsFactory,
		ctx:         ctx,
		cancel:      cancel,
		exited:      make(chan struct{}),
	}
}

func (s *server) Start() {
	go s.serve()
}

func (s *server) serve() {
	deps := s.depsFactory()
	defer deps.Destroy()

	deps.Broker().Subscribe(s.ctx, s.topicName, func(delivery amqp.Delivery) {
		var msg message
		if err := json.Unmarshal(delivery.Body, &msg); err != nil {
			panic(err)
		}

		messageRoute(msg, deps)

		if err := delivery.Ack(false); err != nil {
			panic(err)
		}
	})
}

func (s *server) Stop() error {
	s.cancel()
	<-s.exited
	return nil
}
